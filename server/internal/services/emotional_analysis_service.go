package services

import (
	"context"
	"log"
	"os"
	"regexp"
	"server/internal/database"
	"server/models"
	"strconv"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func CreateEmotionalAnalysis(postText string) (int, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GOOGLE_AI_KEY")))
	if err != nil {
		log.Println(err)
	}

	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(
		ctx,
		genai.Text("Analyze this text. write the percentage of positivity it evokes from 0 to 100. Write only a number between 0 and 100 and percent char. No more text is needed.\nHere is the text of the post:\n\n"+postText+"\n\n"),
	)
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`\d{2,3}%`)
	var text string
	switch p := resp.Candidates[0].Content.Parts[0].(type) {
	case genai.Text:
		text = string(p)
	case genai.Blob:
		text = string(p.Data)
	default:
		log.Println("Unknown type")
	}
	match := re.Find([]byte(text))

	if match == nil {
		return 0, nil
	}

	percentage, err := strconv.Atoi(string(match[:len(match)-1]))
	if err != nil {
		return 0, err
	}

	var icon string
	if percentage <= 25 {
		icon = "☹️"
	} else if percentage <= 50 {
		icon = "😐"
	} else if percentage <= 75 {
		icon = "🙂"
	} else {
		icon = "😊"
	}
	emotionalAnalysis := models.EmotionalAnalysis{
		EmotionalState: percentage,
		EmotionalIcon:  icon,
	}

	dbService := database.Instance()
	id, err := dbService.AddEmotionalAnalysis(emotionalAnalysis)
	if err != nil {
		return 0, err
	}

	return id, nil
}

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"server/internal/database"
	"server/models"
	"strings"

	"github.com/go-chi/render"
	twitterscraper "github.com/n0madic/twitter-scraper"
)

func AddAuthorFunc(w http.ResponseWriter, r *http.Request) {
	var body struct {
		AuthorLink string `json:"author_link"`
		GroupId    int    `json:"group_id"`
	}

	if err := render.Decode(r, &body); err != nil {
		http.Error(w, "Cannot decode: "+err.Error(), http.StatusBadRequest)
		return
	}

	var author models.Author
	if strings.Contains(body.AuthorLink, "x.com") {
		var err error
		author, err = addXAuthor(body.AuthorLink, body.GroupId)
		if err != nil {
			http.Error(w, "Cannot add author: "+err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println("Author added successfully")
	} else {
		http.Error(w, "Invalid author link", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]interface{}{
		"message": "Author added successfully",
		"author": author,
	}
	json.NewEncoder(w).Encode(response)
}

func addXAuthor(link string, groupId int) (models.Author, error) {
	regularExpression := regexp.MustCompile(`^https://x.com/[a-zA-Z0-9]+$`)
	if !regularExpression.MatchString(link) {
		return models.Author{}, fmt.Errorf("Invalid link")
	}

	userName := strings.Split(link, "/")[3]

	twitterUsername := os.Getenv("TWITTER_USERNAME")
	twitterPassword := os.Getenv("TWITTER_PASSWORD")
	scraper := twitterscraper.New()
    err := scraper.Login(twitterUsername, twitterPassword)
	defer scraper.Logout()
	if err != nil {
		return models.Author{}, err
	}

    profile, err := scraper.GetProfile(userName)
	if err != nil {
		return models.Author{}, err
	}

	author := models.Author{
		Name:                  profile.Name,
		Username:              profile.Username,
		SocialMedia:           "X",
		AuthorImage:           strings.ReplaceAll(profile.Avatar, "normal", "400x400"),
		AuthorBackgroundImage: profile.Banner,
		GroupId:               groupId,
	}

	dbService := database.Instance()
	authorId, err := dbService.AddAuthor(author)
	if err != nil {
		return models.Author{}, err
	}

	author.AuthorId = authorId

	return author, nil
}

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	_ "os"
	"regexp"
	"server/dto"
	"server/internal/database"
	"server/internal/services"
	"server/models"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
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
		"author":  author,
	}
	json.NewEncoder(w).Encode(response)
}

func addXAuthor(link string, groupId int) (models.Author, error) {
	regularExpression := regexp.MustCompile(`^https://x.com/[a-zA-Z0-9_]+$`)
	if !regularExpression.MatchString(link) {
		return models.Author{}, fmt.Errorf("Invalid link")
	}

	userName := strings.Split(link, "/")[3]
	client := &http.Client{Timeout: 10 * time.Second}

	// Fetch user details
	userDetailsURL := "https://twitter154.p.rapidapi.com/user/details?username=" + userName
	req, err := http.NewRequest("GET", userDetailsURL, nil)
	if err != nil {
		return models.Author{}, err
	}
	req.Header.Add("x-rapidapi-key", "4780938895msh64888601af59894p1ad53ejsn1a0ae27d979f")
	req.Header.Add("x-rapidapi-host", "twitter154.p.rapidapi.com")

	res, err := client.Do(req)
	if err != nil {
		return models.Author{}, err
	}
	defer res.Body.Close()

	var body struct {
		Name                  string `json:"name"`
		UserName              string `json:"username"`
		AuthorImage           string `json:"profile_pic_url"`
		AuthorBackgroundImage string `json:"profile_banner_url"`
	}

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return models.Author{}, err
	}

	author := models.Author{
		Name:                  body.Name,
		Username:              body.UserName,
		SocialMedia:           "X",
		AuthorImage:           strings.ReplaceAll(body.AuthorImage, "normal", "400x400"),
		AuthorBackgroundImage: body.AuthorBackgroundImage,
		GroupId:               groupId,
	}

	dbService := database.Instance()
	authorId, err := dbService.AddAuthor(author)
	if err != nil {
		return models.Author{}, err
	}

	// Fetch user tweets
	tweetsURL := "https://twitter154.p.rapidapi.com/user/tweets?username=" + userName + "&limit=10&include_replies=false"
	req, err = http.NewRequest("GET", tweetsURL, nil)
	if err != nil {
		return models.Author{}, err
	}
	req.Header.Add("x-rapidapi-key", "4780938895msh64888601af59894p1ad53ejsn1a0ae27d979f")
	req.Header.Add("x-rapidapi-host", "twitter154.p.rapidapi.com")

	res, err = client.Do(req)
	if err != nil {
		return models.Author{}, err
	}
	defer res.Body.Close()

	var resultTwits struct {
		Results []struct {
			Text         string   `json:"text"`
			CreationDate string   `json:"creation_date"`
			Media        []string `json:"media_url"`
			Video        []struct {
				URL string `json:"url"`
			} `json:"video_url"`
		} `json:"results"`
	}

	if err := json.NewDecoder(res.Body).Decode(&resultTwits); err != nil {
		return models.Author{}, fmt.Errorf("error decoding JSON: %v", err)
	}

	fmt.Printf("Decoded tweets: %+v\n", resultTwits)
	iterator := 0
	for _, tweet := range resultTwits.Results {
		if iterator >= 10 {
			break
		}
		tweetTime, err := time.Parse(time.RubyDate, tweet.CreationDate)
		if err != nil {
			return models.Author{}, fmt.Errorf("error parsing date: %v", err)
		}

		fmt.Printf("Processing tweet: %s\n", tweet.Text)

		emotionalAnalysisId, err := services.CreateEmotionalAnalysis(tweet.Text)
		if err != nil {
			return models.Author{}, fmt.Errorf("error creating emotional analysis: %v", err)
		}

		postId, err := dbService.AddPost(models.Post{
			AuthorId:            authorId,
			TextContent:         tweet.Text,
			Date:                tweetTime,
			CountOfLikes:        0,
			EmotionalAnalysisId: emotionalAnalysisId,
		})
		if err != nil {
			return models.Author{}, fmt.Errorf("error adding post: %v", err)
		}

		for _, photo := range tweet.Media {
			fmt.Printf("Processing photo: %s\n", photo)
			if _, err := dbService.AddPhoto(models.XPhoto{URL: photo, PostId: postId}); err != nil {
				return models.Author{}, fmt.Errorf("error adding photo: %v", err)
			}
		}

		for _, video := range tweet.Video {
			fmt.Printf("Processing video: %s\n", video.URL)
			if _, err := dbService.AddVideo(models.XVideo{URL: video.URL, PostId: postId}); err != nil {
				return models.Author{}, fmt.Errorf("error adding video: %v", err)
			}
		}
		iterator++
	}

	author.AuthorId = authorId
	return author, nil
}

func GetAuthorsFunc(w http.ResponseWriter, r *http.Request) {
	var body struct {
		GroupId int `json:"group_id"`
	}

	if err := render.Decode(r, &body); err != nil {
		http.Error(w, "Cannot decode: "+err.Error(), http.StatusBadRequest)
		return
	}

	dbService := database.Instance()
	authors, err := dbService.GetAuthorsByGroupId(body.GroupId)
	if err != nil {
		http.Error(w, "Cannot retrieve authors: "+err.Error(), http.StatusInternalServerError)
		return
	}
	group, err := dbService.GetGroupById(body.GroupId)
	if err != nil {
		http.Error(w, "Cannot retrieve group: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	groupAnalysis, err := dbService.GetGroupEmotionalAnalysis(body.GroupId)
	if err != nil {
		http.Error(w, "Cannot retrieve group emotional analysis: "+err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]interface{}{
		"group": dto.Group{
			Group:             *group,
			EmotionalAnalysis: *groupAnalysis,
		},
		"authors": authors,
	}
	json.NewEncoder(w).Encode(response)
}

func GetAuthorByIdFunc(w http.ResponseWriter, r *http.Request) {
	authorId := chi.URLParam(r, "id")
	if authorId == "" {
		http.Error(w, "Author ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(authorId)
	if err != nil {
		http.Error(w, "Invalid author ID", http.StatusBadRequest)
		return
	}

	dbService := database.Instance()
	author, err := dbService.GetAuthorById(id)
	if err != nil {
		http.Error(w, "Cannot retrieve author: "+err.Error(), http.StatusInternalServerError)
		return
	}

	posts, err := dbService.GetPostsByAuthorId(id)
	if err != nil {
		http.Error(w, "Cannot retrieve posts: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var resultPosts []dto.Post
	for _, post := range posts {
		photos, err := dbService.GetPhotosByPostId(post.PostId)
		if err != nil {
			http.Error(w, "Cannot retrieve photos: "+err.Error(), http.StatusInternalServerError)
			return
		}

		videos, err := dbService.GetVideosByPostId(post.PostId)
		if err != nil {
			http.Error(w, "Cannot retrieve videos: "+err.Error(), http.StatusInternalServerError)
			return
		}

		comments, err := dbService.GetPostComments(post.PostId)
		if err != nil {
			http.Error(w, "Cannot retrieve comments: "+err.Error(), http.StatusInternalServerError)
			return
		}
		resultComments := []dto.Comment{}
		for _, comment := range comments {
			user, err := dbService.GetUserById(comment.UserId)
			if err != nil {
				http.Error(w, "Cannot retrieve user: "+err.Error(), http.StatusInternalServerError)
				return
			}
			resultComments = append(resultComments, dto.Comment{
				Text:     comment.Text,
				UserName: user.Username,
				Date:     comment.Date.Format("2006-01-02 15:04:05"),
			})
		}
		emotionalAnalysis, err := dbService.GetEmotionalAnalysisById(post.EmotionalAnalysisId)
		if err != nil || emotionalAnalysis == nil {
			http.Error(w, "Cannot retrieve emotional analysis: "+err.Error(), http.StatusInternalServerError)
			return
		}
		resultPosts = append(resultPosts, dto.Post{
			Post:     post,
			Photos:   photos,
			Videos:   videos,
			Comments: resultComments,
			EmotionalAnalysis: dto.EmotionalAnalysis{
				EmotionalState: emotionalAnalysis.EmotionalState,
				EmotionalIcon:  emotionalAnalysis.EmotionalIcon,
			},
		})
	}

	authorsEmotionalAnalysis, err := dbService.GetAuthorEmotionalAnalysis(id)
	if err != nil || authorsEmotionalAnalysis == nil {
		http.Error(w, "Cannot retrieve author emotional analysis: "+err.Error(), http.StatusInternalServerError)
		return
	}

	authorDTO := dto.Author{
		Author:            *author,
		EmotionalAnalysis: *authorsEmotionalAnalysis,
		Posts:             resultPosts,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(authorDTO)
}

func AddCommentFunc(w http.ResponseWriter, r *http.Request) {
	var body struct {
		PostId int    `json:"post_id"`
		UserId int    `json:"user_id"`
		Text   string `json:"text"`
	}

	if err := render.Decode(r, &body); err != nil {
		http.Error(w, "Cannot decode: "+err.Error(), http.StatusBadRequest)
		return
	}

	comment := models.Comment{
		PostId: body.PostId,
		UserId: body.UserId,
		Text:   body.Text,
		Date: time.Now(),
	}

	dbService := database.Instance()
	commentId, err := dbService.AddComment(comment)
	if err != nil {
		http.Error(w, "Cannot add comment: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]interface{}{
		"message":    "Comment added successfully",
		"comment_id": commentId,
	}
	json.NewEncoder(w).Encode(response)
}

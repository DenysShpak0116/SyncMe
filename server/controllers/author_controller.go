package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"server/dto"
	"server/internal/database"
	"server/models"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
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
		"author":  author,
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

	for tweet := range scraper.GetTweets(context.Background(), userName, 50) {
		if tweet.Error != nil {
			log.Println(tweet.Error)
		}
		postId, err := dbService.AddPost(models.Post{
			AuthorId:            authorId,
			TextContent:         tweet.Text,
			Date:                tweet.TimeParsed,
			CountOfLikes:        0,
			EmotionalAnalysisId: 1,
		})
		if err != nil {
			return models.Author{}, err
		}

		for _, photo := range tweet.Photos {
			_, err = dbService.AddPhoto(models.XPhoto{
				URL:    photo.URL,
				PostId: postId,
			})
			if err != nil {
				return models.Author{}, err
			}
		}

		for _, video := range tweet.Videos {
			_, err = dbService.AddVideo(models.XVideo{
				URL:    video.URL,
				PostId: postId,
			})
			if err != nil {
				return models.Author{}, err
			}
		}
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
	response := map[string]interface{}{
		"group":   group,
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

		resultPosts = append(resultPosts, dto.Post{
			Post:   post,
			Photos: photos,
			Videos: videos,
		})
	}

	authorDTO := dto.Author{
		Author: *author,
		Posts:  resultPosts,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(authorDTO)
}

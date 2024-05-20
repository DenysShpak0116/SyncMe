package models

import "time"

type Post struct {
	postId       int
	textContent  string
	imgContent   string
	videoContent string
	date         time.Time
	countOfLikes int
}
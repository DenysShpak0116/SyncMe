package models

import "time"

type Post struct {
	PostId              int       `db:"PostId"`
	TextContent         string    `db:"TextContent"`
	ImgContent          string    `db:"ImgContent"`
	VideoContent        string    `db:"VideoContent"`
	Date                time.Time `db:"Date"`
	CountOfLikes        int       `db:"CountOfLikes"`
	AuthorId            int       `db:"AuthorId"`
	EmotionalAnalysisId int       `db:"EmotionalAnalysisId"`
}

package models

import (
	_ "database/sql"
	"time"
)

type Post struct {
	PostId              int       `db:"PostId" json:"postId"`
	TextContent         string    `db:"TextContent" json:"textContent"`
	Date                time.Time `db:"Date" json:"date,omitempty"`
	CountOfLikes        int       `db:"CountOfLikes" json:"countOfLikes"`
	AuthorId            int       `db:"AuthorId" json:"authorId"`
	EmotionalAnalysisId int       `db:"EmotionalAnalysisId" json:"emotionalAnalysisId"`
}

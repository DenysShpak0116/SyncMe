package models

import "time"

type Comment struct {
	CommentId int       `db:"CommentId"`
	Text      string    `db:"Text"`
	Date      time.Time `db:"Date"`
	UserId    int       `db:"UserId"`
	PostId    int       `db:"PostId"`
}

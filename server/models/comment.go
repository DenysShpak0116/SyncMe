package models

import "time"

type Comment struct {
	commentId int
	text string
	date time.Time
}
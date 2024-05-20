package models

import "time"

type Message struct {
	messsageId int
	text       string
	sentAt     time.Time
	userFromId int
	userToId   int
}
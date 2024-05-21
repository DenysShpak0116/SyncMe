package models

import "time"

type Message struct {
	MessageId  int       `db:"MessageId"`
	Text       string    `db:"Text"`
	SentAt     time.Time `db:"SentAt"`
	UserFromId int       `db:"UserFromId"`
	UserToId   int       `db:"UserToId"`
}

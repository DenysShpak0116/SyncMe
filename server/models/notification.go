package models

import "time"

type Notification struct {
	notificationId int
	postText       string
	postImg        string
	date           time.Time
}
package models

import "time"

type Notification struct {
	NotificationId int       `db:"NotificationId"`
	Text           string    `db:"Text"`
	Date           time.Time `db:"Date"`
}

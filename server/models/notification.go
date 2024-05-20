package models

import "time"

type Notification struct {
	NotificationId int    `db:"NotificationId"`
	PostText       string `db:"PostText"`
	PostImg        string `db:"PostImg"`
	Date           time.Time `db:"Date"`
}
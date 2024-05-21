package models

type UserNotification struct {
	UserNotificationId int `db:"UserNotificationId"`
	NotificationId     int `db:"NotificationId"`
	UserId             int `db:"UserId"`
	IsRead             bool `db:"IsRead"`
}
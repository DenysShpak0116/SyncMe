package database

import (
	"context"
	"server/models"
	"time"
)

func (s *service) AddNotification(notification models.Notification) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	res, err := s.db.ExecContext(
		ctx, 
		"INSERT INTO notification (text, date) VALUES (?, ?)", 
		notification.Text, 
		notification.Date,
	)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *service) AddUserNotification(userId, notificationId int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	res, err := s.db.ExecContext(
		ctx, 
		"INSERT INTO usernotification (userid, notificationid) VALUES (?, ?)", 
		userId, 
		notificationId,
	)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *service) GetUserNotifications(userId int) ([]models.Notification, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	rows, err := s.db.QueryContext(
		ctx, 
		"SELECT n.NotificationId, n.text, n.date FROM notification n JOIN usernotification un ON n.NotificationId = un.notificationid WHERE un.userid = ?", 
		userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []models.Notification
	for rows.Next() {
		var notification models.Notification
		err := rows.Scan(&notification.NotificationId, &notification.Text, &notification.Date)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}

	return notifications, nil
}
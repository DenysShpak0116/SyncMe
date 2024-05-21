package models

type UserGroup struct {
	UserGroupId int `db:"UserGroupId"`
	GroupId     int `db:"GroupId"`
	UserId      int `db:"UserId"`
}
package dto

import (
	"server/models"
)

type Post struct {
	models.Post
	Photos   []models.XPhoto `json:"photos"`
	Videos   []models.XVideo `json:"videos"`
	Comments []Comment       `json:"comments"`
}

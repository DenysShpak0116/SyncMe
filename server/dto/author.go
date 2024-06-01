package dto

import (
	"server/models"
)

type Author struct {
	models.Author
	Posts []Post `json:"posts"`
}
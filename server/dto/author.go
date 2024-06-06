package dto

import (
	"server/models"
)

type Author struct {
	models.Author
	EmotionalAnalysis EmotionalAnalysis `json:"emotionalAnalysis"`
	Posts             []Post            `json:"posts"`
}

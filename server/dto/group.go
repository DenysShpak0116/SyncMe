package dto

import "server/models"

type Group struct {
	models.Group
	EmotionalAnalysis EmotionalAnalysis `json:"emotional_analysis"`
}
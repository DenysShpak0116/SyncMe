package models

type EmotionalAnalysis struct {
	EmotionalAnalysisId int    `db:"EmotionalAnalysisId"`
	EmotionalState      int    `db:"EmotionalState"`
	EmotionalIcon       string `db:"EmotionalIcon"`
}
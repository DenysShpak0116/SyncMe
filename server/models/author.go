package models

type Author struct {
	AuthorId              int    `db:"AuthorId"`
	Name                  string `db:"Name"`
	SocialMedia           string `db:"SocialMedia"`
	AuthorImage           string `db:"AuthorImage"`
	AuthorBackgroundImage string `db:"AuthorBackgroundImage"`
	GroupId               int    `db:"GroupId"`
	EmotionalAnalysisId   int    `db:"EmotionalAnalysisId"`
}

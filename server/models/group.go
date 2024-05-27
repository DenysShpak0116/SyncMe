package models

type Group struct {
	GroupId              int    `db:"GroupId"`
	Name                 string `db:"Name"`
	Description          string `db:"Decsription"`
	GroupImage           string `db:"GroupImage"`
	GroupBackgroundImage string `db:"GroupBackgroundImage"`
	EmotionalAnalysisId  int    `db:"EmotionalAnalysisId"`
}

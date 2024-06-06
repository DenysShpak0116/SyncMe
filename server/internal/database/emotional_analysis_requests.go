package database

import (
	"context"
	"server/dto"
	"server/models"
	"time"
)

func (s *service) AddEmotionalAnalysis(emotionalAnalysis models.EmotionalAnalysis) (int, error) {
	query := `INSERT INTO emotionalanalysis (emotionalstate, emotionalIcon) VALUES (?, ?)`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	result, err := s.db.ExecContext(
		ctx,
		query, 
		emotionalAnalysis.EmotionalState, 
		emotionalAnalysis.EmotionalIcon,
	)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int(id), nil
}

func (s *service) GetEmotionalAnalysisById(id int) (*models.EmotionalAnalysis, error) {
	query := `SELECT emotionalanalysisid, emotionalstate, emotionalIcon FROM emotionalanalysis WHERE id = ?`
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := s.db.QueryRowContext(ctx, query, id)
	emotionalAnalysis := models.EmotionalAnalysis{}
	err := row.Scan(
		&emotionalAnalysis.EmotionalAnalysisId,
		&emotionalAnalysis.EmotionalState,
		&emotionalAnalysis.EmotionalIcon,
	)
	if err != nil {
		return nil, err
	}

	return &emotionalAnalysis, nil
}

func (s *service) GetAuthorEmotionalAnalysis(authorId int) (*dto.EmotionalAnalysis, error) {
	query := `SELECT AVG(emotionalstate) FROM emotionalanalysis 
			  INNER JOIN post ON emotionalanalysis.emotionalanalysisid = post.emotionalanalysisid 
			  WHERE post.authorid = ?`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var averageEmotionalState float64
	err := s.db.QueryRowContext(ctx, query, authorId).Scan(&averageEmotionalState)
	if err != nil {
		return nil, err
	}

	var icon string
	if averageEmotionalState <= 25 {
		icon = "☹️"
	} else if averageEmotionalState <= 50 {
		icon = "😐"
	} else if averageEmotionalState <= 75 {
		icon = "🙂"
	} else {
		icon = "😊"
	}
	emotionalAnalysis := &dto.EmotionalAnalysis{
		EmotionalState: int(averageEmotionalState),
		EmotionalIcon:  icon,
	}

	return emotionalAnalysis, nil
}

func (s *service) GetGroupEmotionalAnalysis(groupId int) (*dto.EmotionalAnalysis, error) {
	query := `SELECT AVG(emotionalstate) FROM emotionalanalysis 
			  INNER JOIN post ON emotionalanalysis.emotionalanalysisid = post.emotionalanalysisid 
			  INNER JOIN author ON post.authorid = author.authorid
			  WHERE author.groupid = ?`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var averageEmotionalState float64
	err := s.db.QueryRowContext(ctx, query, groupId).Scan(&averageEmotionalState)
	if err != nil {
		return nil, err
	}

	var icon string
	if averageEmotionalState <= 25 {
		icon = "☹️"
	} else if averageEmotionalState <= 50 {
		icon = "😐"
	} else if averageEmotionalState <= 75 {
		icon = "🙂"
	} else {
		icon = "😊"
	}
	emotionalAnalysis := &dto.EmotionalAnalysis{
		EmotionalState: int(averageEmotionalState),
		EmotionalIcon:  icon,
	}

	return emotionalAnalysis, nil
}
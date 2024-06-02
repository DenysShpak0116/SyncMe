package controllers

import (
	"net/http"
	"server/internal/database"
	"server/models"
	"time"

	"github.com/go-chi/render"
)

func AddMessageFunc(w http.ResponseWriter, r *http.Request) {
	var body struct {
		MessageText string    `json:"message_text"`
		SentAt      time.Time `json:"sent_at"`
		UserFromId  int       `json:"user_from_id"`
		UserToId    int       `json:"user_to_id"`
	}
	if err := render.Decode(r, &body); err != nil {
		http.Error(w, "Cannot decode: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Add message to database
	dbService := database.Instance()
	message := models.Message{
		Text:       body.MessageText,
		SentAt:     body.SentAt,
		UserFromId: body.UserFromId,
		UserToId:   body.UserToId,
	}
	messageId, err := dbService.AddMessage(message)
	if err != nil {
		http.Error(w, "Cannot add message: "+err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"message_id": messageId,
	})
}

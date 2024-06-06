package controllers

import (
	"encoding/json"
	"net/http"
	"server/internal/database"
	"server/models"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
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

	response := map[string]interface{}{
		"message_id": messageId,
	}

	json.NewEncoder(w).Encode(response)
}

func GetMessageFunc(w http.ResponseWriter, r *http.Request) {
	var body struct {
		DisscusserId  int `json:"disscusser_id"`
		CurrentUserId int `json:"current_user_id"`
	}
	if err := render.Decode(r, &body); err != nil {
		http.Error(w, "Cannot decode: "+err.Error(), http.StatusBadRequest)
		return
	}

	dbService := database.Instance()

	chat, err := dbService.GetChat(body.DisscusserId, body.CurrentUserId)
	if err != nil {
		http.Error(w, "Cannot get chat: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"chat": chat,
	}

	json.NewEncoder(w).Encode(response)
}

func DeleteMessageFunc(w http.ResponseWriter, r *http.Request) {
	var body struct {
		MessageId int `json:"message_id"`
	}
	if err := render.Decode(r, &body); err != nil {
		http.Error(w, "Cannot decode: "+err.Error(), http.StatusBadRequest)
		return
	}

	dbService := database.Instance()
	err := dbService.DeleteMessage(body.MessageId)
	if err != nil {
		http.Error(w, "Cannot delete message: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message_id": body.MessageId,
	}

	json.NewEncoder(w).Encode(response)
}

func GetChatsFunc(w http.ResponseWriter, r *http.Request) {
	UserId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	dbService := database.Instance()
	chats, err := dbService.GetUserChats(UserId)
	if err != nil {
		http.Error(w, "Cannot get chats: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"chats": chats,
	}

	json.NewEncoder(w).Encode(response)
}

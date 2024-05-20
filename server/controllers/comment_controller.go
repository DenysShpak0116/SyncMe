package controllers

import (
    "encoding/json"
    "net/http"
    "time"

    "github.com/jmoiron/sqlx"
    "server/models"
)

var DB *sqlx.DB

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var comment models.Comment
    err := json.NewDecoder(r.Body).Decode(&comment)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    comment.Date = time.Now()

    query := `INSERT INTO comments (text, date) VALUES (:text, :date)`
    res, err := DB.NamedExec(query, comment)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    id, err := res.LastInsertId()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    comment.CommentId = int(id)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(comment)
}

package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dieg0code/twitter-clone/db"
	"github.com/dieg0code/twitter-clone/models"
)

// SaveTweet is the function to save a tweet in the database
func SaveTweet(w http.ResponseWriter, r *http.Request) {
	var tweet models.Tweet
	err := json.NewDecoder(r.Body).Decode(&tweet)
	if err != nil {
		http.Error(w, "Error al leer el tweet "+err.Error(), 400)
		return
	}

	register := models.CreateTweet{
		UserId:    UserID,
		Message:   tweet.Message,
		CreatedAt: time.Now(),
	}

	_, status, err := db.InsertTweet(register)
	if err != nil {
		http.Error(w, "Error al insertar el tweet "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "No se puedo insertar el tweet", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

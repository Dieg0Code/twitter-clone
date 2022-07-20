package routers

import (
	"net/http"

	"github.com/dieg0code/twitter-clone/db"
)

func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar un ID", http.StatusBadRequest)
		return
	}

	err := db.DeleteTweetDB(ID, UserID)
	if err != nil {
		http.Error(w, "Error al borrar el tweet"+err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

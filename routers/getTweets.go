package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dieg0code/twitter-clone/db"
)

func GetTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El id es requerido", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "La pagina es requerida", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "La pagina debe ser un numero", http.StatusBadRequest)
		return
	}

	pag := int64(page)
	res, correct := db.ReadTweets(ID, pag)
	if !correct {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

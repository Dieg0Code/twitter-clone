package routers

import (
	"encoding/json"
	"net/http"

	"github.com/dieg0code/twitter-clone/db"
	"github.com/dieg0code/twitter-clone/models"
)

// ModifyProfile modifies a user's profile
func ModifyProfile(w http.ResponseWriter, r *http.Request) {
	var usr models.User

	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(), 400)
	}

	var status bool
	status, err = db.ModifyRegister(usr, UserID)
	if err != nil {
		http.Error(w, "Error al modificar el registro"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha podido modificar el registro", 400)
	}

	w.WriteHeader(http.StatusCreated)
}

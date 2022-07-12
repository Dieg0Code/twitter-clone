package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dieg0code/twitter-clone/db"
	"github.com/dieg0code/twitter-clone/jwt"
	"github.com/dieg0code/twitter-clone/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var usr models.User

	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		http.Error(w, "Usuario y/o contraseña invalidos "+err.Error(), 400)
		return
	}
	if len(usr.Email) == 0 {
		http.Error(w, "El email es requerido", 400)
		return
	}
	document, exist := db.LoginTry(usr.Email, usr.Password)
	if !exist {
		http.Error(w, "Usuario y/o contraseña invalidos", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "Error al generar el token "+err.Error(), 400)
		return
	}

	res := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)

	//Almacenar el token en las cookies
	expirationTime := time.Now().Add(time.Hour * 24) // expire in 24 hours
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}

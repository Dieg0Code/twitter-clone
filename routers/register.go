package routers

import (
	"encoding/json"
	"net/http"

	"github.com/dieg0code/twitter-clone/db"
	"github.com/dieg0code/twitter-clone/models"
)

//Register is the function to register a new user into the database
func Register(w http.ResponseWriter, r *http.Request) {

	var usr models.User
	err := json.NewDecoder(r.Body).Decode(&usr) //El body de un http request es un stream por lo que solo se puede leer una vez, despues se destruye
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(usr.Email) == 0 {
		http.Error(w, "El email es requerido", 400)
		return
	}

	if len(usr.Password) < 6 {
		http.Error(w, "La contraseña debe tener al menos 6 caracteres", 400)
		return
	}

	_, finded, _ := db.CheckIfUserExists(usr.Email)
	if finded == true {
		http.Error(w, "El usuario ya existe", 400)
		return
	}

	_, status, err := db.InserRegister(usr)
	if err != nil {
		http.Error(w, "Error al registrar el usuario "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro de usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

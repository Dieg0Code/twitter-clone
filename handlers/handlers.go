package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/dieg0code/twitter-clone/middlew"
	"github.com/dieg0code/twitter-clone/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Handler set the port and start to listen and handle requests
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/get-profile", middlew.CheckDB(middlew.ValidateJWT(routers.GetProfile))).Methods("GET")
	router.HandleFunc("/update-profile", middlew.CheckDB(middlew.ValidateJWT(routers.ModifyProfile))).Methods("PUT")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}

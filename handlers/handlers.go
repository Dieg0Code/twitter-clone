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
	router.HandleFunc("/getProfile", middlew.CheckDB(middlew.ValidateJWT(routers.GetProfile))).Methods("GET")
	router.HandleFunc("/updateProfile", middlew.CheckDB(middlew.ValidateJWT(routers.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.CheckDB(middlew.ValidateJWT(routers.SaveTweet))).Methods("POST")
	router.HandleFunc("/readTweets", middlew.CheckDB(middlew.ValidateJWT(routers.GetTweets))).Methods("GET")
	router.HandleFunc("/deleteTweet", middlew.CheckDB(middlew.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}

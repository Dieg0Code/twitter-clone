package main

import (
	"log"

	"github.com/dieg0code/twitter-clone/db"
	"github.com/dieg0code/twitter-clone/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("[main] Error al conectar con la base de datos")
		return
	}
	handlers.Handlers()
}

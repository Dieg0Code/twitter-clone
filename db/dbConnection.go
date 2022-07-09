package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN is a global variable that stores our MongoDB connection
var MongoCN = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://dieg0code:<password>@cluster0.yhvu9.mongodb.net/?retryWrites=true&w=majority")

/*
	Un contexto (context) es un espacio en memoria
	donde yo puedo ir compartiendo cosas, puedo ejecutar una intrucción
	de mongo y dentro de un contexto, decirle que esa ejecución no puede
	tardar mas de 5 seg, con esto evitamos que la aplicación se cuelgue.

	Los contextos nos sirven para comunicar información entre ejecución y ejecución, y además nos permite
	setear una seríe de valores como un timeout. Una vez que generamos un contexto y le decimos que tiene
	que ejecutar algo dentro de ese contexto y ese contexto tiene una ejecución máxima de 5 segundos, si llega
	a haber un cuelgue en la base de datos no me va a colgar la api
*/

// DB connection
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions) // context.TODO() es decir sin restricciones de tiempo
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Printf("Conexión a la base de datos establecida")
	return client
}

// CheckConnection ping to MongoDB
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}

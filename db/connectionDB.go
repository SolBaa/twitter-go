package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCT >mongo ct conecta la base de datps
var MongoCT = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://solbaa:Homero1994@cluster0.vw9qa.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

// ConnectDB es la funcion para conectar la base de datos
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Successful connection to DB")
	return client
}

// CheckConnection es la funcion que permite que checquear la conexion
func CheckConnection() int {
	err := MongoCT.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1

}

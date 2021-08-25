package main

import (
	"log"

	"github.com/SolBaa/twitter-go/db"
	"github.com/SolBaa/twitter-go/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Connection to DB failed!")
		return
	}
	handlers.Handlers()
}

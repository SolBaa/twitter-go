package handlers

import (
	"log"
	"net/http"
	"os"

	middlew "github.com/SolBaa/twitter-go/middlewares"
	"github.com/SolBaa/twitter-go/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Handlers seteo mi puerto y pongo a escuchar mi servidor
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")
	// router.HandleFunc("/profile", middlew.CheckDB(middlew.ValidoJWT(routers.Profile))).Methods("GET")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

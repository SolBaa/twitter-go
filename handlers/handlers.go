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
	router.HandleFunc("/profile", middlew.CheckDB(middlew.ValidateJWT(routers.Profile))).Methods("GET")
	router.HandleFunc("/updateProfile", middlew.CheckDB(middlew.ValidateJWT(routers.UpdateProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.CheckDB(middlew.ValidateJWT(routers.SendTweet))).Methods("POST")
	router.HandleFunc("/tweet", middlew.CheckDB(middlew.ValidateJWT(routers.GetTweets))).Methods("GET")
	router.HandleFunc("/tweet", middlew.CheckDB(middlew.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/avatar", middlew.CheckDB(middlew.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/avatar", middlew.CheckDB(routers.GetAvatar)).Methods("GET")
	router.HandleFunc("/banner", middlew.CheckDB(middlew.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/banner", middlew.CheckDB(routers.GetBanner)).Methods("GET")

	router.HandleFunc("/relation", middlew.CheckDB(middlew.ValidateJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/relation", middlew.CheckDB(middlew.ValidateJWT(routers.BajaRelacion))).Methods("DELETE")
	router.HandleFunc("/relation", middlew.CheckDB(middlew.ValidateJWT(routers.CheckRelation))).Methods("GET")

	router.HandleFunc("/users", middlew.CheckDB(middlew.ValidateJWT(routers.UsersList))).Methods("GET")
	router.HandleFunc("/leoTweetsSeguidores", middlew.CheckDB(middlew.ValidateJWT(routers.GetFollowersTweets))).Methods("GET")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

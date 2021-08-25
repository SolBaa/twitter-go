package routers

import (
	"encoding/json"
	"net/http"

	"github.com/SolBaa/twitter-go/db"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) == 0 {
		http.Error(w, "DEbe enviar el parametro ID", http.StatusBadRequest)
		return
	}
	profile, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(w, "Ocurrio un error al buscar el registro"+err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}

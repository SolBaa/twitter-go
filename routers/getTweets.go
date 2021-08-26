package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/SolBaa/twitter-go/db"
)

func GetTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) == 0 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return

	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parametro page", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "No se pudo obtener la pagina ", http.StatusBadRequest)
		return
	}
	pag := int64(page)

	resp, ok := db.GetTweets(ID, pag)
	if !ok {
		http.Error(w, "Erro al leer los tweets", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

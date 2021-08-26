package routers

import (
	"encoding/json"
	"net/http"

	"github.com/SolBaa/twitter-go/db"
	"github.com/SolBaa/twitter-go/models"
)

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(), http.StatusInternalServerError)
		return
	}
	status := false
	status, err = db.ChangeRegister(t, IDUser)
	if err != nil {
		http.Error(w, "Ocurrio un error al modifical el registro. Intente nuevamente"+err.Error(), http.StatusInternalServerError)
		return
	}
	if !status {
		http.Error(w, "No se modifico el register"+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

}

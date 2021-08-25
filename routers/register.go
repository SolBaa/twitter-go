package routers

import (
	"encoding/json"
	"net/http"

	"github.com/SolBaa/twitter-go/db"

	"github.com/SolBaa/twitter-go/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, err.Error()+" Error en los datos recibidos", 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El Email del usuario es requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "La contraseña debe contener al menos 6 caracteres", 400)
		return
	}

	_, find, _ := db.CheckIfUserExist(t.Email)
	if find {
		http.Error(w, "Ya existe un usuario con ese Email", 400)
		return
	}

	_, status, err := db.InsertRegister(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario"+err.Error(), 400)
		return
	}

	if status {
		http.Error(w, "No se ha logrado insertar el registro de usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

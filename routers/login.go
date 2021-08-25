package routers

import (
	"encoding/json"
	"net/http"
	"time"

	bd "github.com/SolBaa/twitter-go/db"
	"github.com/SolBaa/twitter-go/jwt"
	"github.com/SolBaa/twitter-go/models"
)

// Login realiza el login
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Usuario o contraseña invalidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido "+err.Error(), 400)
		return
	}
	document, exists := bd.CheckLogin(t.Email, t.Password)

	if exists == false {
		http.Error(w, "Usuario y/o contraseña invalidos "+err.Error(), 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(document)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar generar el Token correspondiente "+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}

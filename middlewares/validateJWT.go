package middlew

import (
	"net/http"

	"github.com/SolBaa/twitter-go/routers"
)

//ValidateJWT permite validar el JWT que nos viene en la peticion
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el token!"+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}

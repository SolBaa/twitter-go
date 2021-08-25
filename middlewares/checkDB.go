package middlew

import (
	"net/http"

	"github.com/SolBaa/twitter-go/db"
)

// CheckDB
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Connection lost with the DB", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}

package jwt

import (
	"time"

	"github.com/SolBaa/twitter-go/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// GeneroJWT genera el encriptado con JWT e
func GeneroJWT(t models.User) (string, error) {
	miClave := []byte("Golang-Solcha")

	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Name,
		"apellidos":        t.LastName,
		"fecha_nacimiento": t.BirthDate,
		"biografia":        t.Biography,
		"ubicacion ":       t.Location,
		"sitioweb":         t.WebSite,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}

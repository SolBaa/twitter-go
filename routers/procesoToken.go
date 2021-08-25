package routers

import (
	"errors"
	"strings"

	"github.com/SolBaa/twitter-go/db"
	"github.com/SolBaa/twitter-go/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// Email es el email devielto por el p=modelo que se usara en toos los endpoints
var Email string

// IDUser es el id deuelto dentro del modelo que se usara en todos los ennpoints
var IDUser string

// ProcesoToken para extraer sus valores
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("Golang-Solcha")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := db.CheckIfUserExist(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, encontrado, IDUser, nil

	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("Token invalido")
	}
	return claims, false, string(""), err
}

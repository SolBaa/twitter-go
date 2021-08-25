package jwt

import (
	"time"

	"github.com/SolBaa/twitter-go/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateJWT genera el encriptado con JWT
func GenerateJWT(t models.User) (string, error) {
	miClave := []byte("Golang-Solcha")

	payload := jwt.MapClaims{
		"email":      t.Email,
		"name":       t.Name,
		"last_name":  t.LastName,
		"birth_date": t.BirthDate,
		"biography":  t.Biography,
		"location ":  t.Location,
		"web_site":   t.WebSite,
		"_id":        t.ID.Hex(),
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}

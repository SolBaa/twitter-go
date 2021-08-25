package db

import "golang.org/x/crypto/bcrypt"

// HashPassword funcion para encriptar la contrase;a
func HashPassword(pass string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}

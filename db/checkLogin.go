package db

import (
	"github.com/SolBaa/twitter-go/models"
	"golang.org/x/crypto/bcrypt"
)

func CheckLogin(email string, password string) (models.User, bool) {
	usu, find, _ := CheckIfUserExist(email)
	if find {
		return usu, false
	}
	passWord := []byte(password)
	passwordDB := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passWord)
	if err != nil {
		return usu, false
	}
	return usu, true

}

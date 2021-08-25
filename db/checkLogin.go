package db

import (
	"github.com/SolBaa/twitter-go/models"
	"golang.org/x/crypto/bcrypt"
)

func CheckLogin(email string, password string) (models.User, bool) {
	user, find, _ := CheckIfUserExist(email)
	if !find {
		return user, false
	}
	passWord := []byte(password)
	passwordDB := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passWord)
	if err != nil {
		return user, false
	}
	return user, true

}

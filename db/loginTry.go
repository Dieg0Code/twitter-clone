package db

import (
	"github.com/dieg0code/twitter-clone/models"
	"golang.org/x/crypto/bcrypt"
)

// LoginTry veriies email and password and returns the user if it exists
func LoginTry(email string, password string) (models.User, bool) {
	usr, finded, _ := CheckIfUserExists(email)
	if !finded {
		return usr, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usr.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usr, false
	}
	return usr, true
}

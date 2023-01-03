package handler

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password []byte, cost int) (string, error) {

	pass, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	return string(pass), err
}

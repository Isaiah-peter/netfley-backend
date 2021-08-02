package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("fail to hash password: %v", err)
	}

	return string(hashedPassword), nil

}

func CheckPassword(password string, heshedpassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(heshedpassword), []byte(password))
}

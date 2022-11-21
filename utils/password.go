package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(p string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hashedPass), nil
}

func CheckPassword(inputPassword, dbPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(inputPassword))
}

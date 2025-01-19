package helper

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

func CheckHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ValidatePassword(password string) (bool, error) {

	if !regexp.MustCompile(`[a-zA-Z]`).MatchString(password) {
		return false, fmt.Errorf("password must contain at least one letter")
	}

	if !regexp.MustCompile(`\d`).MatchString(password) {
		return false, fmt.Errorf("password must contain at least one digit")
	}

	if !regexp.MustCompile(`[!@#$%^&*]`).MatchString(password) {
		return false, fmt.Errorf("password must contain at least one special character %s", "(!@#$%^&*)")
	}

	return true, nil
}

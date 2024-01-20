package utilities

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Encrypts password
func EncryptPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, "failed to encrypt password")
	}

	return string(passwordHash), nil
}

func ValidatePassword(password string, passwordHash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid password")
	}

	return nil
}
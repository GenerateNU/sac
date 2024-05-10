package auth

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/GenerateNU/sac/backend/constants"
)

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	if !hasDigit(password) {
		return fmt.Errorf("password must contain at least one digit")
	}

	if !hasSpecialChar(password) {
		return fmt.Errorf("password must contain at least one special character")
	}

	return nil
}

func hasDigit(str string) bool {
	return regexp.MustCompile(`[0-9]`).MatchString(str)
}

func hasSpecialChar(str string) bool {
	for _, c := range constants.SPECIAL_CHARACTERS {
		if strings.Contains(str, string(c)) {
			return true
		}
	}

	return false
}

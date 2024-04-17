package auth

import (
	"regexp"
	"strings"

	"github.com/GenerateNU/sac/backend/src/constants"
	"github.com/GenerateNU/sac/backend/src/errors"
)

func ValidatePassword(password string) *errors.Error {
	if len(password) < 8 {
		return &errors.InvalidPasswordNotLongEnough
	}

	if !hasDigit(password) {
		return &errors.InvalidPasswordNoDigit
	}

	if !hasSpecialChar(password) {
		return &errors.InvalidPasswordNoSpecialCharacter
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

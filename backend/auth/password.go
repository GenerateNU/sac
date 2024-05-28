package auth

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/GenerateNU/sac/backend/constants"
)

func ValidatePassword(password string) error {
	var errs []string

	if len(password) < 8 {
		errs = append(errs, "must be at least 8 characters long")
	}

	if len(password) > 128 { // see https://github.com/OWASP/ASVS/issues/756
		errs = append(errs, "must be at most 128 characters long")
	}

	if !hasDigit(password) {
		errs = append(errs, "must contain at least one digit")
	}

	if !hasSpecialChar(password) {
		errs = append(errs, fmt.Sprintf("must contain at least one special character from: [%v]", string(constants.SPECIAL_CHARACTERS)))
	}

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, ", "))
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

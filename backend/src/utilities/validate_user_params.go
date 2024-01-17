package utilities

import (
	"errors"
	"strconv"
)

// Ensures the string is a proper NUID.
func ValidateNUID(nuid string) error {
	_, err := strconv.Atoi(nuid)
	if err != nil || len(nuid) != 9 {
		return errors.New("invalid nuid")
	}
	return nil
}

// Ensures the string is a proper email.
func ValidateEmail(email string) error {
	// TODO: email validation - id like to use https://github.com/mcnijman/go-emailaddress for this
	return nil
}

// Ensures the string is a proper password.
func ValidatePassword(password string) error {
	// TODO: password validation - what are rules (>8 chars, 2 numbers, etc.) and also what do we want
	// to use for that (raw code or is there a nice library
	return nil
}

// Ensures the string is a proper college.
func ValidateCollege(college string) error {
	colleges := []string{"CAMD", "DMSB", "KCCS", "CE", "BCHS", "SL", "CPS", "CS", "CSSH"}
	invalidCollege := true
	for i := 0; i < len(colleges); i++ {
		if string(college) == colleges[i] {
			invalidCollege = false
		}
	}
	if invalidCollege {
		return errors.New("college is invalid")
	}
	return nil
}

// Ensures the string is a proper year.
func ValidateYear(year uint) error {
	years := []uint{1, 2, 3, 4, 5, 6}
	invalidYear := true
	for i := 0; i < len(years); i++ {
		if year == years[i] {
			invalidYear = false
		}
	}
	if invalidYear {
		return errors.New("year is invalid")
	}
	return nil
}

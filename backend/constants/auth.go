package constants

import "time"

const (
	ACCESS_TOKEN_EXPIRY  time.Duration = time.Hour * 12
	REFRESH_TOKEN_EXPIRY time.Duration = time.Hour * 24 * 30
	OTP_LENGTH           int           = 6
	OTP_EXPIRY           time.Duration = time.Minute * 5
)

var SPECIAL_CHARACTERS = []rune{' ', '!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '@', '[', '\\', ']', '^', '_', '`', '{', '|', '}', '~'} // see https://owasp.org/www-community/password-special-characters

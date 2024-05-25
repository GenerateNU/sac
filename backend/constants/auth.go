package constants

import "time"

const (
	ACCESS_TOKEN_EXPIRY  time.Duration = time.Minute * 24 * 30 // temporary TODO: change to 60 minutes
	REFRESH_TOKEN_EXPIRY time.Duration = time.Minute * 24 * 30
	CSRF_TOKEN_LENGTH    int           = 32
)

var SPECIAL_CHARACTERS = []rune{' ', '!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '@', '[', '\\', ']', '^', '_', '`', '{', '|', '}', '~'} // see https://owasp.org/www-community/password-special-characters

package constants

import "strconv"

const (
	DEFAULT_LIMIT int = 10
	DEFAULT_PAGE  int = 1
)

// TODO vault
var (
	DEFAULT_LIMIT_STRING string = strconv.Itoa(DEFAULT_LIMIT)
	DEFAULT_PAGE_STRING  string = strconv.Itoa(DEFAULT_PAGE)
)

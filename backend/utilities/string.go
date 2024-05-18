package utilities

import (
	"strings"
	"unicode"
)

func ToSnakeCase(s string) string {
	var builder strings.Builder
	builder.Grow(len(s) * 2) // Preallocate space for efficiency

	for i, c := range s {
		if i > 0 && unicode.IsUpper(c) {
			builder.WriteRune('_')
		}
		builder.WriteRune(unicode.ToLower(c))
	}

	return builder.String()
}

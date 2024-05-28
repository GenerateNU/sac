package utilities

import (
	"fmt"
	"os"
)

func Exit(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(0)
}

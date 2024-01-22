package commands

import (
	"os"
	"path/filepath"
)

var ROOT_DIR, _ = os.Getwd()
var FRONTEND_DIR = filepath.Join(ROOT_DIR, "/frontend")
var BACKEND_DIR = filepath.Join(ROOT_DIR, "/backend/src")

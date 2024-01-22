package commands

import (
	"path/filepath"

	"github.com/GenerateNU/sac/cli/utils"
)

var ROOT_DIR, _ = utils.GetRootDir()
var FRONTEND_DIR = filepath.Join(ROOT_DIR, "/frontend")
var BACKEND_DIR = filepath.Join(ROOT_DIR, "/backend/src")

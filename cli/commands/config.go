package commands

import (
	"path/filepath"

	"github.com/GenerateNU/sac/cli/utils"
)

var (
	ROOT_DIR, _     = utils.GetRootDir()
	FRONTEND_DIR    = filepath.Join(ROOT_DIR, "/frontend")
	BACKEND_DIR     = filepath.Join(ROOT_DIR, "/backend")
	BACKEND_SRC_DIR = filepath.Join(BACKEND_DIR, "/src")
	CONFIG          = filepath.Join(ROOT_DIR, "/config")
	MIGRATION_FILE  = filepath.Join(BACKEND_SRC_DIR, "/migrations/data.sql")
)

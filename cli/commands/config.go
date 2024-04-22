package commands

import (
	"path/filepath"

	"github.com/GenerateNU/sac/cli/utils"
)

var (
	ROOT_DIR, _  = utils.GetRootDir()
	FRONTEND_DIR = filepath.Join(ROOT_DIR, "/frontend")
	BACKEND_DIR  = filepath.Join(ROOT_DIR, "/backend")
	CONFIG       = filepath.Join(ROOT_DIR, "/config")
	MOCK_FILE    = filepath.Join(BACKEND_DIR, "/mock/data.sql")
)

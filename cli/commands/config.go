package commands

import (
	"path/filepath"

	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/cli/utils"
)

var ROOT_DIR, _ = utils.GetRootDir()
var FRONTEND_DIR = filepath.Join(ROOT_DIR, "/frontend")
var BACKEND_DIR = filepath.Join(ROOT_DIR, "/backend/src")
var CONFIG, _ = config.GetConfiguration(filepath.Join(ROOT_DIR, "/config"))
var MIGRATION_FILE = filepath.Join(BACKEND_DIR, "/migrations/data.sql")

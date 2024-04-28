package helpers

import (
	"path/filepath"

	"github.com/GenerateNU/sac/cli/utils"
)

var (
	ROOT_DIR, _   = utils.GetRootDir()
	FRONTEND_DIR  = filepath.Join(ROOT_DIR, "/frontend")
	MOBILE_DIR    = filepath.Join(FRONTEND_DIR, "/mobile")
	DASHBOARD_DIR = filepath.Join(FRONTEND_DIR, "/dashboard")
	LIB_DIR       = filepath.Join(FRONTEND_DIR, "/lib")
	WEB_DIR       = filepath.Join(FRONTEND_DIR, "/web")
	BACKEND_DIR   = filepath.Join(ROOT_DIR, "/backend")
	CONFIG        = filepath.Join(ROOT_DIR, "/config")
	MIGRATIONS    = filepath.Join(BACKEND_DIR, "/migrations")
	MOCK_FILE     = filepath.Join(BACKEND_DIR, "/mock/data.sql")
)

var (
	FRONTEND_RUN_TARGETS = []string{"web", "mobile", "dashboard"}
	FRONTEND_CI_TARGETS  = []string{"web", "mobile", "dashboard", "lib"}
)

var TARGET_DIRS = map[string]string{
	"web":       WEB_DIR,
	"mobile":    MOBILE_DIR,
	"dashboard": DASHBOARD_DIR,
	"lib":       LIB_DIR,
}

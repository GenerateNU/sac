package helpers

import (
	"fmt"
	"os/exec"
)

func RunBackend() error {
	err := InitDB()
	if err != nil {
		return fmt.Errorf("error initializing database: %w", err)
	}

	err = UpRedis()
	if err != nil {
		return fmt.Errorf("error initializing redis: %w", err)
	}

	err = Execute(exec.Command("templ", "generate", "&&", "go", "run", "main.go"), BACKEND_DIR)
	if err != nil {
		return fmt.Errorf("error running backend: %w", err)
	}

	return nil
}

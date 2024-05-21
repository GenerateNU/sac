package helpers

import (
	"fmt"
	"os/exec"
)

func UpRedis() error {
	err := Execute(exec.Command("docker-compose", "up", "-d"), BACKEND_DIR)
	if err != nil {
		return fmt.Errorf("error initializing redis: %w", err)
	}

	return nil
}

func DownRedis() error {
	err := Execute(exec.Command("docker-compose", "down"), BACKEND_DIR)
	if err != nil {
		return fmt.Errorf("error stopping redis: %w", err)
	}

	return nil
}

func RestartRedis() error {
	err := Execute(exec.Command("docker-compose", "restart"), BACKEND_DIR)
	if err != nil {
		return fmt.Errorf("error restarting redis: %w", err)
	}

	return nil
}

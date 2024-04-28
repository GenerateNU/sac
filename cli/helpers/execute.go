package helpers

import (
	"fmt"
	"os"
	"os/exec"
)

func Execute(cmd *exec.Cmd, dir string) error {
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error starting %s: %w", dir, err)
	}

	return nil
}

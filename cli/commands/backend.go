package commands

import (
	"fmt"
	"os"
	"os/exec"

	_ "github.com/lib/pq"
	"github.com/urfave/cli/v2"
)

func BackendCommand() *cli.Command {
	command := &cli.Command{
		Name:    "backend",
		Usage:   "Starts the backend server",
		Aliases: []string{"be"},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "use-dev-dot-env",
				Usage:   "Use the .env file in the backend directory",
				Aliases: []string{"d"},
			},
		},
		Action: func(c *cli.Context) error {
			err := RunBackend()
			if err != nil {
				return cli.Exit(err.Error(), 1)
			}

			return nil
		},
	}

	return command
}

func RunBackend() error {
	cmd := exec.Command("go", "run", "main.go")
	cmd.Dir = BACKEND_SRC_DIR
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error starting backend: %w", err)
	}
	return nil
}

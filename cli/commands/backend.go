package commands

import (
	"fmt"
	"os"
	"os/exec"

	_ "github.com/lib/pq"
	"github.com/urfave/cli/v2"
)

var backendCmd = &cli.Command{
	Name:  "backend",
	Usage: "Starts the backend server in development mode",
	Action: func(c *cli.Context) error {
		cmd := exec.Command("go", "run", "backend/main.go", "--use-dev-dot-env")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("error starting backend: %w", err)
		}
		return nil
	},
}

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
			RunBackend()

			return nil
		},
	}

	return command
}

func RunBackend() error {
	cmd := exec.Command("go", "run", "main.go")
	cmd.Dir = BACKEND_DIR
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error starting backend: %w", err)
	}
	return nil
}

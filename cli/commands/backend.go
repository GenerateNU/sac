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
		Name:     "be",
		Usage:    "Run the backend",
		Category: "Development",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "d",
				Usage: "Use the .env.dev file in the backend directory",
			},
		},
		Action: func(c *cli.Context) error {
			err := RunBackend(c.Bool("d"))
			if err != nil {
				return cli.Exit(err.Error(), 1)
			}

			return nil
		},
	}

	return command
}

func RunBackend(useDevDotEnv bool) error {
	cmd := exec.Command("go", "run", "main.go")

	if useDevDotEnv {
		cmd.Args = append(cmd.Args, "--use-dev-dot-env")
	}

	cmd.Dir = BACKEND_SRC_DIR
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error starting backend: %w", err)
	}
	return nil
}

package commands

import (
	"fmt"
	"os"
	"os/exec"

	_ "github.com/lib/pq"
	"github.com/urfave/cli/v2"
)

func RunBackendCommand() *cli.Command {
	command := cli.Command{
		Name:     "be",
		Usage:    "Run the backend",
		Category: "Development",
		Action: func(c *cli.Context) error {
			if c.Args().Len() > 0 {
				return cli.Exit("Invalid arguments", 1)
			}

			err := RunBE()
			if err != nil {
				return cli.Exit(err.Error(), 1)
			}

			return nil
		},
	}

	return &command
}

func RunBE() error {
	goCmd := exec.Command("go", "run", "main.go")
	goCmd.Dir = BACKEND_DIR

	goCmd.Stdout = os.Stdout
	goCmd.Stderr = os.Stderr

	fmt.Println("Running backend")

	err := goCmd.Run()
	if err != nil {
		return fmt.Errorf("error running backend: %w", err)
	}

	return nil
}

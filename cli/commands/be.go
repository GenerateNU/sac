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
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "use-dev-dot-env",
				Usage:   "Use the development .env file",
				Value:   false,
				Aliases: []string{"d"},
			},
		},
		Action: func(c *cli.Context) error {
			if c.Args().Len() > 0 {
				return cli.Exit("Invalid arguments", 1)
			}

			useDevDotEnv := c.Bool("use-dev-dot-env")

			err := RunBE(useDevDotEnv)
			if err != nil {
				return cli.Exit(err.Error(), 1)
			}

			return nil
		},
	}

	return &command
}

func RunBE(useDevDotEnv bool) error {
	var goCmd *exec.Cmd

	if useDevDotEnv {
		goCmd = exec.Command("go", "run", "main.go", "--use-dev-dot-env")
	} else {
		goCmd = exec.Command("go", "run", "main.go")
	}

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

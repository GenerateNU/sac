package commands

import (
	"fmt"
	"os"
	"os/exec"
	"slices"

	"github.com/urfave/cli/v2"
)

func FormatCommand() *cli.Command {
	command := cli.Command{
		Name:     "format",
		Aliases:  []string{"f"},
		Usage:    "Runs formatting tools",
		Category: "CI",
		Subcommands: []*cli.Command{
			{
				Name:    "frontend",
				Usage:   "Format the frontend",
				Aliases: []string{"f"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "target",
						Aliases: []string{"t"},
						Value:   "mobile",
						Usage:   "Format a specific frontend type (web or mobile)",
					},
				},
				Action: func(c *cli.Context) error {
					if c.Args().Len() > 0 {
						return cli.Exit("Invalid arguments", 1)
					}

					target := c.String("target")
					if !slices.Contains(FRONTEND_CI_TARGETS, target) {
						return cli.Exit("Invalid frontend type: must be 'web', 'mobile', 'dashboard', or 'lib'", 1)
					}

					err := FormatFrontend(target)
					if err != nil {
						return cli.Exit(err.Error(), 1)
					}

					return nil
				},
			},
			{
				Name:    "backend",
				Usage:   "Format the backend",
				Aliases: []string{"b"},
				Action: func(c *cli.Context) error {
					if c.Args().Len() > 0 {
						return cli.Exit("Invalid arguments", 1)
					}

					err := FormatBackend()
					if err != nil {
						return cli.Exit(err.Error(), 1)
					}

					return nil
				},
			},
		},
	}

	return &command
}

func FormatFrontend(target string) error {
	dir, ok := TARGET_DIRS[target]
	if !ok {
		return cli.Exit("Invalid frontend type", 1)
	}

	cmd := exec.Command("yarn", "run", "format")

	cmd.Dir = dir

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func FormatBackend() error {
	fmt.Println("Formatting backend")

	cmd := exec.Command("gofumpt", "-l", "-w", ".")
	cmd.Dir = BACKEND_DIR

	err := cmd.Run()
	if err != nil {
		return cli.Exit("Failed to format backend", 1)
	}

	fmt.Println("Backend formatted")
	return nil
}

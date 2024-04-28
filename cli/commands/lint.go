package commands

import (
	"fmt"
	"os"
	"os/exec"
	"slices"

	"github.com/urfave/cli/v2"
)

func LintCommand() *cli.Command {
	command := cli.Command{
		Name:     "lint",
		Aliases:  []string{"l"},
		Usage:    "Runs linting tools",
		Category: "CI",
		Subcommands: []*cli.Command{
			{
				Name:    "frontend",
				Usage:   "Lint the frontend",
				Aliases: []string{"f"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "target",
						Aliases: []string{"t"},
						Value:   "mobile",
						Usage:   "Lint a specific frontend type (web or mobile)",
					},
					&cli.BoolFlag{
						Name:    "fix",
						Aliases: []string{"f"},
						Usage:   "Fix linting errors",
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

					err := LintFrontend(target)
					if err != nil {
						return cli.Exit(err.Error(), 1)
					}

					return nil
				},
			},
			{
				Name:    "backend",
				Usage:   "Lint the backend",
				Aliases: []string{"b"},
				Action: func(c *cli.Context) error {
					if c.Args().Len() > 0 {
						return cli.Exit("Invalid arguments", 1)
					}

					err := LintBackend()
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

func LintFrontend(target string) error {
	dir, ok := TARGET_DIRS[target]
	if !ok {
		return cli.Exit("Invalid frontend type", 1)
	}

	cmd := exec.Command("yarn", "run", "lint")

	cmd.Dir = dir

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func LintBackend() error {
	fmt.Println("Linting backend")

	cmd := exec.Command("golangci-lint", "run", "--fix")
	cmd.Dir = BACKEND_DIR

	err := cmd.Run()
	if err != nil {
		return cli.Exit("Failed to lint backend", 1)
	}

	fmt.Println("Backend linted")

	return nil
}

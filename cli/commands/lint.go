package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func LintCommand() *cli.Command {
    command := cli.Command{
        Name:    "lint",
        Aliases: []string{"l"},
        Usage:   "Runs linting tools",
		Category: "CI",
		Subcommands: []*cli.Command{
			{
				Name:  "frontend",
				Usage: "Lint the frontend",
				Aliases: []string{"f"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "target",
						Aliases: []string{"t"},
						Value:   "mobile",
						Usage:   "Lint a specific frontend type (web or mobile)",
					},
					&cli.BoolFlag{
						Name: "fix",
						Aliases: []string{"f"},
						Usage: "Fix linting errors",
					},
				},
				Action: func(c *cli.Context) error {
					if c.Args().Len() > 0 {
						return cli.Exit("Invalid arguments", 1)
					}

					target := c.String("target")
					if target != "web" && target != "mobile" {
						return cli.Exit("Invalid frontend type: must be 'web' or 'mobile'", 1)
					}

					fix := c.Bool("fix")

					err := LintFrontend(target, fix)
					if err != nil {
						return cli.Exit(err.Error(), 1)
					}

					return nil
				},
			},
			{
				Name:  "backend",
				Usage: "Lint the backend",
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

func LintFrontend(target string, fix bool) error {
	switch target {
	case "web":
		return LintWeb(fix)
	case "mobile":
		return LintMobile(fix)
	default:
		return LintMobile(fix)
	}
}

func LintBackend() error {
	fmt.Println("Linting backend")

	cmd := exec.Command("go", "vet", "./...")
	cmd.Dir = BACKEND_DIR

	err := cmd.Run()
	if err != nil {
		return cli.Exit("Failed to lint backend", 1)
	}

	fmt.Println("Backend linted")

	return nil
}

func LintWeb(fix bool) error {
	return nil
}

func LintMobile(fix bool) error {
	var mobileCmd *exec.Cmd
	if fix {
		mobileCmd = exec.Command("yarn", "run", "lint", "--fix")
	} else {
		mobileCmd = exec.Command("yarn", "run", "lint")
	}
	mobileCmd.Dir = FRONTEND_DIR + "/sac-mobile"

	mobileCmd.Stdout = os.Stdout
	mobileCmd.Stderr = os.Stderr
	mobileCmd.Stdin = os.Stdin 

	if err := mobileCmd.Run(); err != nil {
		return err
	}

	return nil
}
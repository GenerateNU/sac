package commands

import (
	"fmt"
	"os"
	"os/exec"

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
					if target != "web" && target != "mobile" {
						return cli.Exit("Invalid frontend type: must be 'web' or 'mobile'", 1)
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
	switch target {
	case "web":
		return FormatWeb()
	case "mobile":
		return FormatMobile()
	default:
		return FormatMobile()
	}
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

func FormatWeb() error {
	return nil
}

func FormatMobile() error {
	mobileCmd := exec.Command("yarn", "run", "format")
	mobileCmd.Dir = FRONTEND_DIR + "/sac-mobile"

	mobileCmd.Stdout = os.Stdout
	mobileCmd.Stderr = os.Stderr
	mobileCmd.Stdin = os.Stdin

	if err := mobileCmd.Run(); err != nil {
		return err
	}

	return nil
}

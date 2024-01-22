package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func LintCommand() *cli.Command {
	command := cli.Command{
		Name:  "lint",
		Usage: "Runs linter",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "frontend",
				Aliases: []string{"f"},
				Usage:   "Runs frontend linter",
				Value:   "",
			},
			&cli.BoolFlag{
				Name:    "backend",
				Aliases: []string{"b"},
				Usage:   "Runs backend linter",
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("lint", c.String("frontend"), c.Bool("backend"))
			return nil
		},
	}

	return &command
}

package commands

import (
	"fmt"
	"os/exec"
	"sync"

	"github.com/urfave/cli/v2"
)

func FormatCommand() *cli.Command {
	command := cli.Command{
		Name:  "format",
		Usage: "Runs formatting tools",
		Aliases: []string{"f"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "frontend",
				Aliases: []string{"f"},
				Value:   "",
				Usage:   "Formats a specific frontend folder",
			},
			&cli.BoolFlag{
				Name:    "backend",
				Aliases: []string{"b"},
				Usage:   "Formats the backend",
			},
		},
		Action: func(c *cli.Context) error {
			if c.Args().Len() > 0 {
				return cli.Exit("Invalid arguments", 1)
			}

			if c.String("frontend") == "" && !c.Bool("backend") {
				return cli.Exit("Must specify frontend or backend", 1)
			}

			folder := c.String("frontend")
			runFrontend := folder != ""
			runBackend := c.Bool("backend")

			Format(folder, runFrontend, runBackend)

			return nil
		},
	}

	return &command
}

func Format(folder string, runFrontend bool, runBackend bool) error {
	var wg sync.WaitGroup

	// Start the backend if specified
	if runBackend {
		wg.Add(1)
		go func() {
			defer wg.Done()
			BackendFormat()
		}()
	}

	// Start the frontend if specified
	if runFrontend {
		wg.Add(1)
		go func() {
			defer wg.Done()
			FrontendFormat(folder)
		}()
	}

	wg.Wait()

	return nil
}

func BackendFormat() error {
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

func FrontendFormat(folder string) error {
	fmt.Println("UNIMPLEMENTED")
	return nil
}

package commands

import (
	"fmt"
	"os/exec"
	"sync"

	"github.com/urfave/cli/v2"
)

func LintCommand() *cli.Command {
	command := cli.Command{
		Name:    "lint",
		Aliases: []string{"l"},
		Usage:   "Runs linting tools",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "frontend",
				Aliases: []string{"f"},
				Value:   "",
				Usage:   "Lint a specific frontend folder",
			},
			&cli.BoolFlag{
				Name:    "backend",
				Aliases: []string{"b"},
				Usage:   "Lint the backend",
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

			err := Lint(folder, runFrontend, runBackend)
			if err != nil {
				return cli.Exit(err.Error(), 1)
			}

			return nil
		},
	}

	return &command
}

func Lint(folder string, runFrontend bool, runBackend bool) error {
	var wg sync.WaitGroup
	errChan := make(chan error, 1)

	// Start the backend if specified
	if runBackend {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := BackendLint()
			if err != nil {
				errChan <- err
			}
		}()
	}

	// Start the frontend if specified
	if runFrontend {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := FrontendLint(folder)
			if err != nil {
				errChan <- err
			}
		}()
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}

func BackendLint() error {
	fmt.Println("Linting backend")

	cmd := exec.Command("golangci-lint", "run")
	cmd.Dir = BACKEND_DIR

	err := cmd.Run()
	if err != nil {
		return cli.Exit("Failed to lint backend", 1)
	}

	fmt.Println("Backend linted")

	return nil
}

func FrontendLint(folder string) error {
	fmt.Println("UNIMPLEMENTED")
	return nil
}

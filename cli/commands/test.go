package commands

import (
	"fmt"
	"os/exec"
	"sync"

	"github.com/urfave/cli/v2"
)

func TestCommand(backendDir string, frontendDir string) *cli.Command {
	command := cli.Command{
		Name:  "test",
		Usage: "Runs tests",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "frontend",
				Aliases: []string{"f"},
				Value:   "",
				Usage:   "Runs frontend tests",
			},
			&cli.BoolFlag{
				Name:    "backend",
				Aliases: []string{"b"},
				Usage:   "Runs backend tests",
			},
		},
		Action: func(c *cli.Context) error {
			if c.Args().Len() > 0 {
				return cli.Exit("Invalid arguments", 1)
			}

			if c.String("frontend") == "" && !c.Bool("backend") {
				return cli.Exit("Must specify frontend or backend", 1)
			}

			fmt.Println("Frontend", c.String("frontend"))

			folder := c.String("frontend")
			runFrontend := folder != ""
			runBackend := c.Bool("backend")

			Test(backendDir, frontendDir, folder, runFrontend, runBackend)

			return nil
		},
	}

	return &command
}

func Test(backendDir string, frontendDir string, folder string, runFrontend bool, runBackend bool) error {
	var wg sync.WaitGroup

	// Start the backend if specified
	if runBackend {
		wg.Add(1)
		go func() {
			defer wg.Done()
			BackendTest(backendDir)
		}()
	}

	// Start the frontend if specified
	if runFrontend {
		wg.Add(1)
		go func() {
			defer wg.Done()
			FrontendTest(frontendDir, folder)
		}()
	}

	wg.Wait()

	return nil
}

func BackendTest(backendDir string) error {
	cmd := exec.Command("go", "test", "./...")
	cmd.Dir = backendDir + "/.."

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return cli.Exit("Failed to run backend tests", 1)
	}

	fmt.Println(string(out))

	return nil
}

func FrontendTest(frontendDir string, folder string) error {
	cmd := exec.Command("yarn", "run", "test")
	cmd.Dir = frontendDir

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return cli.Exit("Failed to run frontend tests", 1)
	}

	fmt.Println(string(out))

	return nil
}

package commands

import (
	"fmt"
	"os/exec"
	"sync"

	"github.com/urfave/cli/v2"
)

func TestCommand() *cli.Command {
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

			folder := c.String("frontend")
			runFrontend := folder != ""
			runBackend := c.Bool("backend")

			Test(folder, runFrontend, runBackend)

			return nil
		},
	}

	return &command
}

func Test(folder string, runFrontend bool, runBackend bool) error {
	var wg sync.WaitGroup

	// Start the backend if specified
	if runBackend {
		wg.Add(1)
		go func() {
			defer wg.Done()
			BackendTest()
		}()
	}

	// Start the frontend if specified
	if runFrontend {
		wg.Add(1)
		go func() {
			defer wg.Done()
			FrontendTest(folder)
		}()
	}

	wg.Wait()

	return nil
}

func BackendTest() error {
	cmd := exec.Command("go", "test", "./...")
	cmd.Dir = BACKEND_DIR + "/.."

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return cli.Exit("Failed to run backend tests", 1)
	}

	fmt.Println(string(out))

	cmd = exec.Command("./scripts/clean_old_test_dbs.sh")
	cmd.Dir = ROOT_DIR

	out, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return cli.Exit("Failed to clean old test databases", 1)
	}

	return nil
}

func FrontendTest(folder string) error {
	cmd := exec.Command("yarn", "run", "test")
	cmd.Dir = FRONTEND_DIR + "/" + folder

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return cli.Exit("Failed to run frontend tests", 1)
	}

	fmt.Println(string(out))

	return nil
}

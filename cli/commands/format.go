package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func FormatCommand() *cli.Command {
	command := cli.Command{
		Name:  "format",
		Usage: "Runs formatter",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "frontend",
				Aliases: []string{"f"},
				Usage:   "Runs frontend formatter",
				Value:   "",
			},
			&cli.BoolFlag{
				Name:    "backend",
				Aliases: []string{"b"},
				Usage:   "Runs backend formatter",
			},
		},
		Action: Format,
	}

	return &command
}

func Format(c *cli.Context) error {
	if c.Args().Len() > 0 {
		return cli.Exit("Invalid arguments", 1)
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return cli.Exit("Error getting current directory", 1)
	}

	frontendDir := filepath.Join(currentDir, "frontend/")
	backendDir := filepath.Join(currentDir, "github.com/GenerateNU/sac/backend/")
	list, err := os.ReadDir(frontendDir)
	if err != nil {
		return cli.Exit("Error reading frontend directory", 1)
	}

	if !c.IsSet("frontend") && !c.IsSet("backend") {
		formatBackend(backendDir)
		for _, f := range list {
			if f.IsDir() {
				formatFrontend(frontendDir, f.Name())
			}
		}
	}

	if c.IsSet("frontend") && c.IsSet("backend") {
		formatFrontend(frontendDir, c.String("frontend"))
		formatBackend(backendDir)
	}

	if c.IsSet("frontend") {
		formatFrontend(frontendDir, c.String("frontend"))
	}
	if c.IsSet("backend") {
		formatBackend(backendDir)
	}

	return nil
}

func formatFrontend(frontendDir string, folder string) {
	cmd := exec.Command("yarn", "format")
	cmd.Dir = filepath.Join(frontendDir, folder)

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error formatting frontend, run yarn format in frontend folder")
	}

	fmt.Println("frontend", cmd.Dir) // remove
}

func formatBackend(backendDir string) {
	cmd := exec.Command("go", "fmt", "./...")
	cmd.Dir = filepath.Join(backendDir)

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error formatting backend, run go fmt ./... in backend folder")
	}

	fmt.Println("backend", cmd.Dir) // remove
}

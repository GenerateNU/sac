package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func SwaggerCommand() *cli.Command {
	command := cli.Command{
		Name:    "swagger",
		Aliases: []string{"swag"},
		Usage:   "Runs `swag init` to update Swagger documentation for the backend API",
		Action: func(c *cli.Context) error {
			cmd := exec.Command("swag", "init")
			cmd.Dir = BACKEND_SRC_DIR
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				return fmt.Errorf("error running swag init: %w", err)
			}
			return nil
		},
	}

	return &command
}

func Swagger() error {
	cmd := exec.Command("swag", "init")
	cmd.Dir = BACKEND_SRC_DIR

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return cli.Exit("Failed to generate swagger.json", 1)
	}

	fmt.Println(string(out))
	return nil
}

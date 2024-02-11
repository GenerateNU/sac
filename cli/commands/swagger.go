package commands

import (
	"fmt"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func SwaggerCommand() *cli.Command {
	command := cli.Command{
		Name:    "swagger",
		Aliases: []string{"swag"},
		Usage:   "Updates the swagger documentation",
		Action: func(c *cli.Context) error {
			if c.Args().Len() > 0 {
				return cli.Exit("Invalid arguments", 1)
			}

			err := Swagger()
			if err != nil {
				return cli.Exit(err.Error(), 1)
			}
			return nil
		},
	}

	return &command
}

func Swagger() error {
	cmd := exec.Command("swag", "init")
	cmd.Dir = BACKEND_DIR

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return cli.Exit("Failed to generate swagger.json", 1)
	}

	fmt.Println(string(out))
	return nil
}

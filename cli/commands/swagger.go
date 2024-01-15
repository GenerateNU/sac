package commands

import (
	"fmt"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func SwaggerCommand(backendDir string) *cli.Command {
	command := cli.Command{
		Name:  "swagger",
		Usage: "Updates the swagger documentation",
		Action: func(c *cli.Context) error {
			if c.Args().Len() > 0 {
				return cli.Exit("Invalid arguments", 1)
			}

			Swagger(backendDir)
			return nil
		},
	}

	return &command
}


func Swagger(backendDir string) error {
	cmd := exec.Command("swag", "init")
	cmd.Dir = backendDir 

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return cli.Exit("Failed to generate swagger.json", 1)
	}

	fmt.Println(string(out))
	return nil
}
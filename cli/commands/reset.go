package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func ResetDBCommand() *cli.Command {
	command := cli.Command{
		Name:  "resetdb",
		Usage: "Resets the database",
		Action: func(c *cli.Context) error {
			fmt.Println("resetdb")
			return nil
		},
	}

	return &command
}
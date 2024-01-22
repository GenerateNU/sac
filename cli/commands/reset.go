package commands

import (
	"fmt"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func ResetDBCommand() *cli.Command {
	command := cli.Command{
		Name:  "reset",
		Usage: "Resets the database",
		Action: func(c *cli.Context) error {
			if c.Args().Len() > 0 {
				return cli.Exit("Invalid arguments", 1)
			}

			ResetDB()
			return nil
		},
	}

	return &command
}

func ResetDB() error {
	fmt.Println("Resetting database")

	DropDB()

	cmd := exec.Command("sleep", "1")
	cmd.Dir = BACKEND_DIR 

	err := cmd.Run()
	if err != nil {
		return cli.Exit("Error running sleep", 1)
	}

	Migrate()

	fmt.Println("Done resetting database")
	
	return nil
}
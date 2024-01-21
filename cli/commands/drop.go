package commands

import (
	"fmt"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func DropDBCommand() *cli.Command {
	command := cli.Command{
		Name:  "drop",
		Usage: "Drops the database",
		Action: func(c *cli.Context) error {
			if c.Args().Len() > 0 {
				return cli.Exit("Invalid arguments", 1)
			}

			DropDB()
			return nil
		},
	}

	return &command
}

func DropDB() error {
	fmt.Println("Droping database")

	cmd := exec.Command("../../scripts/drop_db.sh")
	cmd.Dir = BACKEND_DIR

	output, err := cmd.CombinedOutput()
	if err != nil {
		return cli.Exit("Error running drop_db.sh", 1)
	}

	fmt.Println(string(output))

	fmt.Println("Done droping database")

	return nil
}
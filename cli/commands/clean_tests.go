package commands

import (
	"fmt"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func ClearDBCommand() *cli.Command {
	command := cli.Command{
		Name:  "clean",
		Usage: "Remove databases used for testing",
		Action: func(c *cli.Context) error {
			if c.Args().Len() > 0 {
				return cli.Exit("Invalid arguments", 1)
			}

			ClearDB()
			return nil
		},
	}

	return &command
}

func ClearDB() error { 

	fmt.Println("Clearing databases")

	cmd := exec.Command("./scripts/clean_old_test_dbs.sh")
	cmd.Dir = ROOT_DIR

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return cli.Exit("Failed to clean old test databases", 1)
	}

	fmt.Println("Databases cleared")
	
	return nil
}

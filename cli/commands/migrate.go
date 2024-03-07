package commands

import (
	"fmt"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func MigrateCommand() *cli.Command {
	command := cli.Command{
		Name:     "migrate",
		Aliases:  []string{"m"},
		Usage:    "Migrate the database, creating tables and relationships",
		Category: "Database Operations",
		Action: func(c *cli.Context) error {
			if c.Args().Len() > 0 {
				return cli.Exit("Invalid arguments", 1)
			}

			err := Migrate()
			if err != nil {
				return cli.Exit(err.Error(), 1)
			}
			return nil
		},
	}

	return &command
}

func Migrate() error {
	fmt.Println("Migrating database")

	goCmd := exec.Command("go", "run", "main.go", "--only-migrate")
	goCmd.Dir = BACKEND_SRC_DIR

	output, err := goCmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running main.go:", err)
	}

	fmt.Println(string(output))

	fmt.Println("Done migrating database")

	return nil
}

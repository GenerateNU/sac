package commands

import (
	"fmt"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func MigrateCommand() *cli.Command {
	command := cli.Command{
		Name:  "migrate",
		Usage: "Migrate the database",
		Action: func(c *cli.Context) error {
			if c.Args().Len() > 0 {
				return cli.Exit("Invalid arguments", 1)
			}

			Migrate()
			return nil
		},
	}

	return &command
}

func Migrate() error {
	fmt.Println("Migrating database")

	goCmd := exec.Command("go", "run", "main.go", "--only-migrate")
	goCmd.Dir = BACKEND_DIR

	output, err := goCmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running main.go:", err)
	}

	fmt.Println(string(output))

	fmt.Println("Inserting data into database")

	scriptCmd := exec.Command("./scripts/insert_db.sh")
	scriptCmd.Dir = ROOT_DIR

	output, err = scriptCmd.CombinedOutput()
	if err != nil {
		return cli.Exit("Error running insert_db.sh", 1)
	}

	fmt.Println(string(output))
	fmt.Println("Done migrating database")

	return nil
}

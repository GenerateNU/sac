package commands

import (
	"fmt"
	"os/exec"
	"sync"

	"github.com/urfave/cli/v2"
)

func MigrateCommand(backendDir string) *cli.Command {
	command := cli.Command{
		Name:  "migrate",
		Usage: "Migrate the database",
		Action: func(c *cli.Context) error {
			if c.Args().Len() > 0 {
				return cli.Exit("Invalid arguments", 1)
			}

			Migrate(backendDir)
			return nil
		},
	}

	return &command
}

func Migrate(backendDir string) error {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		cmd := exec.Command("go", "run", "main.go", "--only-migrate")
		cmd.Dir = backendDir

		fmt.Println("Running main.go")
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error running main.go:", err)
		}
	}()

	cmd := exec.Command("../../scripts/insert_db.sh")
	cmd.Dir = backendDir

	output, err := cmd.CombinedOutput()
	if err != nil {
		return cli.Exit("Error running insert_db.sh", 1)
	}

	fmt.Println(string(output))

	wg.Wait()
	fmt.Println("Insert_db.sh completed")
	
	return nil
}

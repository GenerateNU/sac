package commands

import (
	"fmt"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func ResetCommand() *cli.Command {
	command := cli.Command{
		Name:     "reset",
		Aliases:  []string{"r"},
		Usage:    "Resets the database, dropping all tables, clearing data, and re-running migrations",
		Category: "Database Operations",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "data",
				Usage: "Reset only data, not the entire database, will re-run migrations",
			},
		},
		Action: func(c *cli.Context) error {
			if c.Args().Len() > 0 {
				return cli.Exit("Invalid arguments", 1)
			}

			if c.Bool("data") {
				err := ResetData()
				if err != nil {
					return cli.Exit(err.Error(), 1)
				}
			} else {
				err := ResetMigration()
				if err != nil {
					return cli.Exit(err.Error(), 1)
				}
			}

			return nil
		},
	}

	return &command
}

func ResetData() error {
	fmt.Println("Clearing database")

	err := DropData()
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

	cmd := exec.Command("sleep", "1")
	cmd.Dir = BACKEND_SRC_DIR

	err = cmd.Run()
	if err != nil {
		return cli.Exit("Error running sleep", 1)
	}

	err = Migrate()
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

	cmd = exec.Command("sleep", "1")
	cmd.Dir = BACKEND_SRC_DIR

	err = cmd.Run()
	if err != nil {
		return cli.Exit("Error running sleep", 1)
	}

	err = InsertDB()
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

	fmt.Println("Data reset successfully")

	return nil
}

func ResetMigration() error {
	fmt.Println("Resetting migration")

	err := DropDB()
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

	cmd := exec.Command("sleep", "1")
	cmd.Dir = BACKEND_SRC_DIR

	err = cmd.Run()
	if err != nil {
		return cli.Exit("Error running sleep", 1)
	}

	err = Migrate()
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}

	fmt.Println("Migration reset successfully")

	return nil
}

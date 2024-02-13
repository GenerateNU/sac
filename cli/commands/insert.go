package commands

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"

	"github.com/lib/pq"
	"github.com/urfave/cli/v2"
)

func InsertCommand() *cli.Command {
	command := cli.Command{
		Name:     "insert",
		Category: "Database Operations",
		Aliases:  []string{"i"},
		Usage:    "Inserts mock data into the database",
		Action: func(c *cli.Context) error {
			if c.Args().Len() > 0 {
				return cli.Exit("Invalid arguments", 1)
			}

			err := InsertDB()
			if err != nil {
				return cli.Exit(err.Error(), 1)
			}

			return nil
		},
	}

	return &command
}

func InsertDB() error {
	db, err := sql.Open("postgres", CONFIG.Database.WithDb())
	if err != nil {
		return err
	}

	defer db.Close()

	var exists bool

	err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' LIMIT 1);").Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		fmt.Println("Database does not exist or has no tables. Running database migration.")

		migrateCmd := exec.Command("go", "run", "main.go", "--only-migrate")

		migrateCmd.Dir = BACKEND_DIR

		if err := migrateCmd.Run(); err != nil {
			return fmt.Errorf("error running migration: %w", err)
		}
	} else {
		fmt.Println("Database exists with tables.")
	}

	migrationSQL, err := os.ReadFile(MIGRATION_FILE)
	if err != nil {
		return fmt.Errorf("error reading migration file: %w", err)
	}

	_, err = db.Exec(string(migrationSQL))
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			fmt.Println("PostgreSQL Error:")
			fmt.Println("Code:", pqErr.Code)
			fmt.Println("Message:", pqErr.Message)
		} else {
			return fmt.Errorf("error executing migration: %w", err)
		}
	}

	fmt.Println("Data inserted successfully.")

	return nil
}

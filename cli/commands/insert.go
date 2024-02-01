package commands

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
	"github.com/urfave/cli/v2"
)

func InsertDBCommand() *cli.Command {
	command := cli.Command{
		Name:  "insert",
		Usage: "Inserts mock data into the database",
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

	insertCmd := exec.Command("psql", "-h", CONFIG.Database.Host, "-p", strconv.Itoa(int(CONFIG.Database.Port)), "-U", CONFIG.Database.Username, "-d", CONFIG.Database.DatabaseName, "-a", "-f", MIGRATION_FILE)

	var output bytes.Buffer
	insertCmd.Stdout = &output
	insertCmd.Stderr = &output

	if err := insertCmd.Run(); err != nil {
		return fmt.Errorf("error inserting data: %w", err)
	}

	if strings.Contains(output.String(), "ROLLBACK") {
		return errors.New("insertion failed, rolling back")
	}

	fmt.Println("Data inserted successfully.")

	return nil
}

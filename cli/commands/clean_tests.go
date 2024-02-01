package commands

import (
	"database/sql"
	"fmt"
	"os/user"
	"sync"

	_ "github.com/lib/pq"
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

			err := CleanTestDBs()
			if err != nil {
				return cli.Exit(err.Error(), 1)
			}

			return nil
		},
	}

	return &command
}

func CleanTestDBs() error {
	fmt.Println("Cleaning test databases")

	db, err := sql.Open("postgres", CONFIG.Database.WithDb())
	if err != nil {
		return err
	}

	defer db.Close()

	currentUser, err := user.Current()
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}

	rows, err := db.Query("SELECT datname FROM pg_database WHERE datistemplate = false AND datname != 'postgres' AND datname != $1 AND datname != $2", currentUser.Username, CONFIG.Database.DatabaseName)
	if err != nil {
		return err
	}

	defer rows.Close()

	var wg sync.WaitGroup

	for rows.Next() {
		var dbName string

		if err := rows.Scan(&dbName); err != nil {
			return err
		}

		wg.Add(1)

		go func(dbName string) {
			defer wg.Done()

			fmt.Printf("Dropping database %s\n", dbName)

			if _, err := db.Exec(fmt.Sprintf("DROP DATABASE %s", dbName)); err != nil {
				fmt.Printf("Failed to drop database %s: %v\n", dbName, err)
			}
		}(dbName)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	wg.Wait()

	return nil
}

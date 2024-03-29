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
		Name:     "clean",
		Category: "Database Operations",
		Aliases:  []string{"c"},
		Usage:    "Remove databases used for testing",
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

	query := "SELECT datname FROM pg_database WHERE datistemplate = false AND datname != 'postgres' AND datname != $1 AND datname != $2 AND datname LIKE 'sac_test_%';"
	rows, err := db.Query(query, currentUser.Username, CONFIG.Database.DatabaseName)
	if err != nil {
		return err
	}

	defer rows.Close()

	var wg sync.WaitGroup
	var dropped, failed int

	for rows.Next() {
		var dbName string

		if err := rows.Scan(&dbName); err != nil {
			return err
		}

		wg.Add(1)

		go func(dbName string) {
			defer wg.Done()

			fmt.Printf("Dropping database %s\n", dbName)

			_, err := db.Exec(fmt.Sprintf("DROP DATABASE %s", dbName))
			if err != nil {
				fmt.Printf("Failed to drop database %s: %v\n", dbName, err)
				failed++
			} else {
				dropped++
			}
		}(dbName)
	}

	wg.Wait()

	fmt.Printf("\nSummary:\n  - Databases dropped: %d\n  - Databases failed to drop: %d\n", dropped, failed)

	if failed > 0 {
		return fmt.Errorf("failed to drop %d database(s)", failed)
	}

	return nil
}

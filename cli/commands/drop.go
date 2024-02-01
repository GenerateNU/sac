package commands

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
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

			err := DropDB()
			if err != nil {
				return cli.Exit(err.Error(), 1)
			}

			return nil
		},
	}

	return &command
}

func DropDB() error {
	fmt.Println("Dropping database")

	db, err := sql.Open("postgres", CONFIG.Database.WithDb())
	if err != nil {
		return err
	}

	defer db.Close()

	var tableCount int

	err = db.QueryRow("SELECT COUNT(*) FROM pg_tables WHERE schemaname = 'public'").Scan(&tableCount)
	if err != nil {
		return fmt.Errorf("error checking tables: %w", err)
	}

	if tableCount == 0 {
		fmt.Println("No tables to drop. The database is empty.")
		return nil
	}

	fmt.Println("Generating DROP TABLE statements...")

	rows, err := db.Query("SELECT tablename FROM pg_tables WHERE schemaname = 'public'")
	if err != nil {
		return fmt.Errorf("error generating DROP TABLE statements: %w", err)
	}

	defer rows.Close()

	var wg sync.WaitGroup

	fmt.Println("Dropping tables...")

	for rows.Next() {
		var tablename string
		if err := rows.Scan(&tablename); err != nil {
			return fmt.Errorf("error reading table name: %w", err)
		}

		wg.Add(1)
		go func(table string) {
			defer wg.Done()
			dropStmt := fmt.Sprintf("DROP TABLE IF EXISTS \"%s\" CASCADE", table)
			if _, err := db.Exec(dropStmt); err != nil {
				fmt.Printf("Error dropping table %s: %v\n", table, err)
			} else {
				fmt.Printf("Dropped table %s\n", table)
			}
		}(tablename)
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("error in rows handling: %w", err)
	}

	wg.Wait()
	fmt.Println("All tables dropped successfully.")
	return nil
}

package commands

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/urfave/cli/v2"
)

var dbMutex sync.Mutex

func DropCommand() *cli.Command {
	command := cli.Command{
		Name:     "drop",
		Aliases:  []string{"d"},
		Usage:    "Drop data with a migration or drops the entire database",
		Category: "Database Operations",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "data",
				Usage: "Drop only data, not the entire database",
			},
		},
		Action: func(c *cli.Context) error {
			if c.Args().Len() > 0 {
				return cli.Exit("Invalid arguments", 1)
			}

			if c.Bool("data") {
				err := DropData()
				if err != nil {
					return cli.Exit(err.Error(), 1)
				}
			} else {
				err := DropDB()
				if err != nil {
					return cli.Exit(err.Error(), 1)
				}
			}

			return nil
		},
	}

	return &command
}

func DropData() error {
	fmt.Println("Clearing database")

	dbMutex.Lock()
	defer dbMutex.Unlock()

	db, err := sql.Open("postgres", CONFIG.Database.WithDb())
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("SELECT tablename FROM pg_tables WHERE schemaname = 'public'")
	if err != nil {
		return fmt.Errorf("error retrieving tables: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var tablename string
		if err := rows.Scan(&tablename); err != nil {
			return fmt.Errorf("error scanning table name: %w", err)
		}

		deleteStmt := fmt.Sprintf("DELETE FROM \"%s\"", tablename)
		_, err := db.Exec(deleteStmt)
		if err != nil {
			return fmt.Errorf("error deleting rows from table %s: %w", tablename, err)
		}
		fmt.Printf("Removed all rows from table %s\n", tablename)
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("error in rows handling: %w", err)
	}

	Migrate()

	fmt.Println("All rows removed successfully.")
	return nil
}

func DropDB() error {
	fmt.Println("Dropping database")

	dbMutex.Lock()
	defer dbMutex.Unlock()

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

	fmt.Println("Dropping tables...")

	for rows.Next() {
		var tablename string
		if err := rows.Scan(&tablename); err != nil {
			return fmt.Errorf("error reading table name: %w", err)
		}

		dropStmt := fmt.Sprintf("DROP TABLE IF EXISTS \"%s\" CASCADE", tablename)
		if _, err := db.Exec(dropStmt); err != nil {
			fmt.Printf("Error dropping table %s: %v\n", tablename, err)
		} else {
			fmt.Printf("Dropped table %s\n", tablename)
		}
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("error in rows handling: %w", err)
	}

	fmt.Println("All tables dropped successfully.")
	return nil
}

package commands

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"sync"

	"github.com/GenerateNU/sac/backend/config"
	_ "github.com/lib/pq"
	"github.com/urfave/cli/v2"
)

func DBCommand() *cli.Command {
	command := &cli.Command{
		Name:  "db",
		Usage: "Database management",
		Subcommands: []*cli.Command{
			{
				Name:  "init",
				Usage: "Initialize the database",
				Action: func(c *cli.Context) error {
					err := InitDB()
					if err != nil {
						return cli.Exit(err.Error(), 1)
					}

					return nil
				},
			},
			{
				Name:  "down",
				Usage: "Migrate down the database",
				Action: func(c *cli.Context) error {
					err := DownDB()
					if err != nil {
						return cli.Exit(err.Error(), 1)
					}

					return nil
				},
			},
			{
				Name:  "reset",
				Usage: "Reset the database",
				Action: func(c *cli.Context) error {
					err := DownDB()
					if err != nil {
						return cli.Exit(err.Error(), 1)
					}

					err = InitDB()
					if err != nil {
						return cli.Exit(err.Error(), 1)
					}

					return nil
				},
			},
			{
				Name:  "clean-tests",
				Usage: "Remove databases used for testing",
				Action: func(c *cli.Context) error {
					return CleanTestDBs(context.Background())
				},
			},
		},
	}

	return command
}

func InitDB() error {
	cmd := exec.Command("./scripts/init_db.sh")

	cmd.Dir = ROOT_DIR
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error initializing database: %w", err)
	}
	return nil
}

func DownDB() error {
	config, err := config.GetConfiguration(CONFIG, false)
	if err != nil {
		return err
	}

	// #nosec G204 -- MIGRATIONS is a constant file path configured in cli/commands/config.go and config.Database.PostgresConn() is configured by the config/local.yml
	cmd := exec.Command("migrate", "-path", MIGRATIONS, "-database", config.Database.PostgresConn(), "-verbose", "down", "-all")

	cmd.Dir = ROOT_DIR
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error running down migrations: %w", err)
	}
	return nil
}

func CleanTestDBs(ctx context.Context) error {
	fmt.Println("Cleaning test databases")

	config, err := config.GetConfiguration(CONFIG, false)
	if err != nil {
		return err
	}

	db, err := sql.Open("postgres", config.Database.WithDb())
	if err != nil {
		return err
	}

	defer db.Close()

	currentUser, err := user.Current()
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}

	rows, err := db.Query("SELECT datname FROM pg_database WHERE datistemplate = false AND datname != 'postgres' AND datname != $1 AND datname != $2 AND datname LIKE 'sac_test_%';", currentUser.Username, config.Database.DatabaseName)
	if err != nil {
		return err
	}

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

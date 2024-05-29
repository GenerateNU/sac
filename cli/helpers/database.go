package helpers

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"sync"

	"github.com/GenerateNU/sac/backend/config"
	"github.com/lib/pq"
)

func InitDB() error {
	err := Execute(exec.Command("./scripts/init_db.sh"), ROOT_DIR)
	if err != nil {
		return fmt.Errorf("error initializing database: %w", err)
	}

	return nil
}

func DownDB() error {
	config, err := config.GetConfiguration(CONFIG)
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

func CleanTestDBs() error {
	config, err := config.GetConfiguration(CONFIG)
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

func InsertDB() error {
	config, err := config.GetConfiguration(CONFIG)
	if err != nil {
		return err
	}

	db, err := sql.Open("postgres", config.Database.WithDb())
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

	migrationSQL, err := os.ReadFile(MOCK_FILE)
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

	return nil
}

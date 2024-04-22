package commands

import (
	"fmt"
	"os"
	"os/exec"

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

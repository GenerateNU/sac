package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"path/filepath"

	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/database"
	_ "github.com/GenerateNU/sac/backend/docs"
	"github.com/GenerateNU/sac/backend/integrations"
	"github.com/GenerateNU/sac/backend/integrations/email"
	"github.com/GenerateNU/sac/backend/integrations/file"
	"github.com/GenerateNU/sac/backend/search"
	"github.com/GenerateNU/sac/backend/server"
)

func CheckServerRunning(host string, port uint16) error {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}

func Exit(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(0)
}

func configureIntegrations(config *config.Settings) integrations.Integrations {
	integrations := integrations.Integrations{
		File:  file.NewAWSProvider(config.AWS),
		Email: email.NewResendClient(config.Resend, true),
	}
	return integrations
}

func main() {
	onlyMigrate := flag.Bool("only-migrate", false, "Specify if you want to only perform the database migration")
	configPath := flag.String("config", filepath.Join("..", "config"), "Specify the path to the config directory")
	useDevDotEnv := flag.Bool("use-dev-dot-env", true, "Specify if you want to use the .env.dev file")
	seedSearch := flag.Bool("seed-search", false, "Specify if you want to perform search seeding.")

	flag.Parse()

	config, err := config.GetConfiguration(*configPath, *useDevDotEnv)
	if err != nil {
		Exit("Error getting configuration: %s", err.Error())
	}

	err = CheckServerRunning(config.Application.Host, config.Application.Port)
	if err == nil {
		Exit("A server is already running on %s:%d.\n", config.Application.Host, config.Application.Port)
	}

	db, err := database.ConfigureDB(*config)
	if err != nil {
		Exit("Error migrating database: %s", err.Error())
	}

	if *onlyMigrate {
		return
	}

	err = database.ConnPooling(db)
	if err != nil {
		Exit("Error with connection pooling: %s", err.Error())
	}

	integrations := configureIntegrations(config)

	if *seedSearch {
		search.SeedClubs(db)
	}

	app := server.Init(db, integrations, *config)

	err = app.Listen(fmt.Sprintf("%s:%d", config.Application.Host, config.Application.Port))
	if err != nil {
		Exit("Error starting server: %s", err.Error())
	}
}

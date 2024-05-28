package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net"
	"path/filepath"

	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/database"
	"github.com/GenerateNU/sac/backend/database/store"
	_ "github.com/GenerateNU/sac/backend/docs"
	"github.com/GenerateNU/sac/backend/integrations"
	"github.com/GenerateNU/sac/backend/integrations/email"
	"github.com/GenerateNU/sac/backend/integrations/file"
	"github.com/GenerateNU/sac/backend/integrations/oauth"
	"github.com/GenerateNU/sac/backend/search"
	"github.com/GenerateNU/sac/backend/server"
	"github.com/GenerateNU/sac/backend/telemetry"
	"github.com/GenerateNU/sac/backend/utilities"
)

func main() {
	onlyMigrate := flag.Bool("only-migrate", false, "Specify if you want to only perform the database migration")
	seedSearch := flag.Bool("seed-search", false, "Specify if you want to seed the opensearch nodes.")
	configPath := flag.String("config", filepath.Join("..", "config"), "Specify the path to the config directory")
	useDevDotEnv := flag.Bool("use-dev-dot-env", true, "Specify if you want to use the .env.dev file")

	flag.Parse()

	config, err := config.GetConfiguration(*configPath)

	X := config.Database.Password.Expose()
	print(X)

	if err != nil {
		utilities.Exit("Error getting configuration: %s", err.Error())
	}

	err = checkServerRunning(config.Application.Host, config.Application.Port)
	if err == nil {
		utilities.Exit("A server is already running on %s:%d.\n", config.Application.Host, config.Application.Port)
	}

	db, err := database.ConfigureDB(*config)
	if err != nil {
		utilities.Exit("Error migrating database: %s", err.Error())
	}

	if *onlyMigrate {
		return
	}

	if *seedSearch {
		if err := search.SeedClubs(db); err != nil {
			return
		}

		if err := search.SeedEvents(db); err != nil {
			return
		}
	}

	err = database.ConnPooling(db)
	if err != nil {
		utilities.Exit("Error with connection pooling: %s", err.Error())
	}

	stores := store.ConfigureRedis(*config)

	integrations := configureIntegrations(config)

	tp := telemetry.InitTracer()

	slog.Info("appease linter since we aren't tracing anything yet", "tracer", telemetry.Tracer)

	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			slog.Error("error shutting down tracer", "error", err)
		}
	}()

	app := server.Init(db, stores, integrations, *config)

	err = app.Listen(fmt.Sprintf("%s:%d", config.Application.Host, config.Application.Port))
	if err != nil {
		utilities.Exit("Error starting server: %s", err.Error())
	}
}

func checkServerRunning(host string, port uint16) error {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}

func configureIntegrations(config *config.Settings) integrations.Integrations {
	googleClient := oauth.GoogleOAuthClient{
		OAuthConfig: config.GoogleSettings,
	}
	outlookClient := oauth.OutlookOAuthClient{
		OAuthConfig: config.OutlookSettings,
	}

	oauthResources := oauth.OAuthResources{
		Google:  googleClient,
		Outlook: outlookClient,
	}

	integrations := integrations.Integrations{
		File:  file.NewAWSProvider(config.AWS),
		Email: email.NewResendClient(config.Resend, true),
		OAuth: oauthResources,
	}

	return integrations
}

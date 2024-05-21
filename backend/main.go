package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
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
	search_test "github.com/GenerateNU/sac/backend/search/base"
	"github.com/GenerateNU/sac/backend/server"
	"github.com/GenerateNU/sac/backend/telemetry"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("sac")

func main() {
	onlyMigrate := flag.Bool("only-migrate", false, "Specify if you want to only perform the database migration")
	seedSearch := flag.Bool("seed-search", true, "Specify if you want to seed the opensearch nodes.")
	configPath := flag.String("config", filepath.Join("..", "config"), "Specify the path to the config directory")
	useDevDotEnv := flag.Bool("use-dev-dot-env", true, "Specify if you want to use the .env.dev file")

	flag.Parse()

	config, err := config.GetConfiguration(*configPath, *useDevDotEnv)
	if err != nil {
		exit("Error getting configuration: %s", err.Error())
	}

	err = checkServerRunning(config.Application.Host, config.Application.Port)
	if err == nil {
		exit("A server is already running on %s:%d.\n", config.Application.Host, config.Application.Port)
	}

	db, err := database.ConfigureDB(*config)
	if err != nil {
		exit("Error migrating database: %s", err.Error())
	}

	if *onlyMigrate {
		return
	}

	if *seedSearch {
		search.SeedClubs(db)
		search.SeedEvents(db)

		search_test.UpsertClubToIndex(db, "f15c4bb5-77a7-4efe-a47c-10f5e6a46b50")
		search_test.DeleteClubFromIndex(db, "f15c4bb5-77a7-4efe-a47c-10f5e6a46b50")

		search_test.UpsertEventToIndex(db, "be83bd22-638a-4000-aaee-9d2d2115fbd2")
		search_test.DeleteEventFromIndex(db, "be83bd22-638a-4000-aaee-9d2d2115fbd2")
	}

	err = database.ConnPooling(db)
	if err != nil {
		exit("Error with connection pooling: %s", err.Error())
	}

	integrations := configureIntegrations(config)

	tp := telemetry.InitTracer()

	slog.Info("appease linter since we aren't tracing anything yet", "tracer", tracer)

	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			slog.Error("error shutting down tracer", "error", err)
		}
	}()

	app := server.Init(db, integrations, *config)

	err = app.Listen(fmt.Sprintf("%s:%d", config.Application.Host, config.Application.Port))
	if err != nil {
		exit("Error starting server: %s", err.Error())
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

func exit(format string, a ...interface{}) {
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

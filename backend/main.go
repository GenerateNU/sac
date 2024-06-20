package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/GenerateNU/sac/backend/background"
	"github.com/GenerateNU/sac/backend/background/jobs"
	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/database"
	"github.com/GenerateNU/sac/backend/database/store"
	"github.com/GenerateNU/sac/backend/integrations"
	"github.com/GenerateNU/sac/backend/integrations/email"
	"github.com/GenerateNU/sac/backend/integrations/file"
	"github.com/GenerateNU/sac/backend/integrations/oauth/soth/sothic"
	"github.com/GenerateNU/sac/backend/search"
	"github.com/GenerateNU/sac/backend/server"
	"github.com/GenerateNU/sac/backend/telemetry"
	"github.com/GenerateNU/sac/backend/utilities"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	onlyMigrate, configPath := parseFlags()
	config, err := config.GetConfiguration(*configPath)
	if err != nil {
		utilities.Exit("Error getting configuration: %s", err.Error())
	}

	if checkServerRunning(config.Application.Host, config.Application.Port) == nil {
		utilities.Exit("A server is already running on %s:%d.\n", config.Application.Host, config.Application.Port)
	}

	ctx := context.Background()

	db, err := database.ConfigureDB(*config)
	if err != nil {
		utilities.Exit("Error migrating database: %s", err.Error())
	}

	if *onlyMigrate {
		return
	}

	if err := database.ConnPooling(db); err != nil {
		utilities.Exit("Error with connection pooling: %s", err.Error())
	}

	startBackgroundJobs(ctx, db, config)

	stores := store.ConfigureStores(config.RedisLimiter)
	sothic.Init(config.Session)
	integrations := configureIntegrations(&config.Integrations)

	tp := telemetry.InitTracer()
	defer shutdownTracer(tp)

	app := server.Init(db, stores, *integrations, *config)

	go startServer(app, config.Application.Host, config.Application.Port)

	waitForShutdown(app)
}

func parseFlags() (onlyMigrate *bool, configPath *string) {
	onlyMigrate = flag.Bool("only-migrate", false, "Specify if you want to only perform the database migration")
	configPath = flag.String("config", filepath.Join("..", "config", ".env.dev"), "Specify the path to the config file (.env)")
	flag.Parse()
	return
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

func startBackgroundJobs(ctx context.Context, db *gorm.DB, settings *config.Settings) {
	jobs := jobs.New(db, settings)
	background.Go(jobs.WelcomeSender(ctx))
	background.Go(jobs.ClubEmbeddings(ctx))
	background.Go(jobs.EventEmbeddings(ctx))
}

func configureIntegrations(config *config.Integrations) *integrations.Integrations {
	return &integrations.Integrations{
		File:  file.NewAWSProvider(config.AWS),
		Email: email.NewResendClient(config.Resend, true),
	}
}

func shutdownTracer(tp *sdktrace.TracerProvider) {
	if err := tp.Shutdown(context.Background()); err != nil {
		slog.Error("error shutting down tracer", "error", err)
	}
}

func startServer(app *fiber.App, host string, port uint16) {
	if err := app.Listen(fmt.Sprintf("%s:%d", host, port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func waitForShutdown(app *fiber.App) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("Shutting down server")
	if err := app.Shutdown(); err != nil {
		slog.Error("failed to shutdown server", "error", err)
	}
	slog.Info("Server shutdown")
}

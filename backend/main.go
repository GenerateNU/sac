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
	"github.com/GenerateNU/sac/backend/integrations/calendar"
	"github.com/GenerateNU/sac/backend/integrations/email"
	"github.com/GenerateNU/sac/backend/integrations/file"
	"github.com/GenerateNU/sac/backend/integrations/search"
	"github.com/GenerateNU/sac/backend/server"
	"github.com/GenerateNU/sac/backend/tests/api/mocks"
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

func configureIntegrations(config *config.Settings, connectToPinecone bool) integrations.Integrations {
	openAi := search.NewOpenAIClient(config.OpenAISettings)
	var pinecone search.SearchClientInterface

	if connectToPinecone {
		pinecone = search.NewPineconeClient(openAi, config.PineconeSettings)
	} else {
		pinecone = mocks.NewPineconeMockClient()
	}

	// Calendar integrations:
	googleClient := calendar.GoogleClient{
		OAuthClient: calendar.NewGoogleOAuthClient(config.GoogleClientSettings),
	}
	outlookClient := calendar.OutlookClient{Settings: config.OutlookClientSettings}
	calendarClient := calendar.CalendarClient{
		Google: googleClient,
		Outlook: outlookClient,
	}

	integrations := integrations.Integrations{
		File:   file.NewAWSProvider(config.AWS),
		AI:     openAi,
		Email:  email.NewResendClient(config.ResendSettings, true),
		Search: pinecone,
		Calendar: calendarClient,
	}
	return integrations
}

func main() {
	onlyMigrate := flag.Bool("only-migrate", false, "Specify if you want to only perform the database migration")
	onlySeedPinecone := flag.Bool("seed-pinecone", false, "Specify if want to only perform the pinecone database seeding")
	connectToPinecone := flag.Bool("connect-to-pinecone", false, "Connect to a real Pinecone instance instead of mock")
	configPath := flag.String("config", filepath.Join("..", "config"), "Specify the path to the config directory")
	useDevDotEnv := flag.Bool("use-dev-dot-env", true, "Specify if you want to use the .env.dev file")

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

	if *onlySeedPinecone {
		openAi := search.NewOpenAIClient(config.OpenAISettings)
		pinecone := search.NewPineconeClient(openAi, config.PineconeSettings)

		err := pinecone.Seed(db)
		if err != nil {
			Exit("Error seeding PineconeDB: %s\n", err.Error())
		}
		return
	}

	err = database.ConnPooling(db)
	if err != nil {
		Exit("Error with connection pooling: %s", err.Error())
	}

	integrations := configureIntegrations(config, *connectToPinecone)

	app := server.Init(db, integrations, *config)

	err = app.Listen(fmt.Sprintf("%s:%d", config.Application.Host, config.Application.Port))
	if err != nil {
		Exit("Error starting server: %s", err.Error())
	}
}

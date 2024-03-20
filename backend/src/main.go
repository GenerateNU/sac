package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"path/filepath"

	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/database"
	_ "github.com/GenerateNU/sac/backend/src/docs"
	"github.com/GenerateNU/sac/backend/src/search"
	"github.com/GenerateNU/sac/backend/src/server"
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

func main() {
	onlyMigrate := flag.Bool("only-migrate", false, "Specify if you want to only perform the database migration")
	onlySeedPinecone := flag.Bool("seed-pinecone", false, "Specify if want to only perform the pinecone database seeding")
	configPath := flag.String("config", filepath.Join("..", "..", "config"), "Specify the path to the config directory")
	useDevDotEnv := flag.Bool("use-dev-dot-env", false, "Specify if you want to use the .env.dev file")

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
			fmt.Printf("Error seeding PineconeDB: %s\n", err.Error())
			return
		}
		return
	}

	err = database.ConnPooling(db)
	if err != nil {
		Exit("Error with connection pooling: %s", err.Error())
	}

	openAi := search.NewOpenAIClient(config.OpenAISettings)
	pinecone := search.NewPineconeClient(openAi, config.PineconeSettings)

	app := server.Init(db, pinecone, *config)

	err = app.Listen(fmt.Sprintf("%s:%d", config.Application.Host, config.Application.Port))
	if err != nil {
		Exit("Error starting server: %s", err.Error())
	}
}

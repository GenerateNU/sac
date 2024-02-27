package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/database"
	_ "github.com/GenerateNU/sac/backend/src/docs"
	"github.com/GenerateNU/sac/backend/src/search"
	"github.com/GenerateNU/sac/backend/src/server"
)

// @title SAC API
// @version 1.0
// @description Backend Server for SAC App
// @contact.name	David Oduneye and Garrett Ladley
// @contact.email	generatesac@gmail.com
// @host 127.0.0.1:8080
// @BasePath /api/v1
func main() {
	onlyMigrate := flag.Bool("only-migrate", false, "Specify if you want to only perform the database migration")
	onlySeedPinecone := flag.Bool("seed-pinecone", false, "Specify if want to only perform the pinecone database seeding")
	configPath := flag.String("config", filepath.Join("..", "..", "config"), "Specify the path to the config directory")
	useDevDotEnv := flag.Bool("use-dev-dot-env", false, "Specify if you want to use the .env.dev file")

	flag.Parse()

	config, err := config.GetConfiguration(*configPath, *useDevDotEnv)
	if err != nil {
		panic(fmt.Sprintf("Error getting configuration: %s", err.Error()))
	}

	db, err := database.ConfigureDB(*config)
	if err != nil {
		panic(fmt.Sprintf("Error configuring database: %s", err.Error()))
	}

	openAi := search.NewOpenAIClient(config.OpenAISettings)
	pinecone := search.NewPineconeClient(openAi, config.PineconeSettings)

	if *onlyMigrate {
		return
	}

	if *onlySeedPinecone {
		err := pinecone.Seed(db)
		if err != nil {
			fmt.Printf("Error seeding PineconeDB: %s\n", err.Error())
			return
		}
		return
	}

	err = database.ConnPooling(db)
	if err != nil {
		panic(fmt.Sprintf("Error connecting to database: %s", err.Error()))
	}

	app := server.Init(db, pinecone, *config)

	err = app.Listen(fmt.Sprintf("%s:%d", config.Application.Host, config.Application.Port))
	if err != nil {
		panic(fmt.Sprintf("Error starting server: %s", err.Error()))
	}
}

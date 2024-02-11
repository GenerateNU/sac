package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/GenerateNU/sac/backend/src/config"
	"github.com/GenerateNU/sac/backend/src/database"
	"github.com/GenerateNU/sac/backend/src/server"
)

// @title SAC API
// @version 1.0
// @description Backend Server for SAC App

// @contact.name	David Oduneye and Garrett Ladley
// @contact.email	oduneye.d@northeastern.edu and ladley.g@northeastern.edu
// @host 127.0.0.1:8080
// @BasePath /api/v1
func main() {
	onlyMigrate := flag.Bool("only-migrate", false, "Specify if you want to only perform the database migration")
	configPath := flag.String("config", filepath.Join("..", "..", "config"), "Specify the path to the config directory")

	flag.Parse()

	config, err := config.GetConfiguration(*configPath)
	if err != nil {
		panic(fmt.Sprintf("Error getting configuration: %s", err.Error()))
	}

	db, err := database.ConfigureDB(*config)
	if err != nil {
		panic(fmt.Sprintf("Error configuring database: %s", err.Error()))
	}

	if *onlyMigrate {
		return
	}

	err = database.ConnPooling(db)
	if err != nil {
		panic(fmt.Sprintf("Error connecting to database: %s", err.Error()))
	}

	app := server.Init(db, *config)

	err = app.Listen(fmt.Sprintf("%s:%d", config.Application.Host, config.Application.Port))
	if err != nil {
		panic(fmt.Sprintf("Error starting server: %s", err.Error()))
	}
}

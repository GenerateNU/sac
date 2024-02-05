package main

import (
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
	config, err := config.GetConfiguration(filepath.Join("..", "..", "config"))
	if err != nil {
		panic(fmt.Sprintf("Error getting configuration: %s", err.Error()))
	}

	db, err := database.ConfigureDB(config)
	if err != nil {
		panic(fmt.Sprintf("Error configuring database: %s", err.Error()))
	}

	err = database.ConnPooling(db)
	if err != nil {
		panic(err)
	}

	app := server.Init(db, config)

	err = app.Listen(fmt.Sprintf("%s:%d", config.Application.Host, config.Application.Port))
	if err != nil {
		panic(err)
	}
}

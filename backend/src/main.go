package main

import (
	"backend/src/config"
	"backend/src/database"
	"backend/src/server"
	"fmt"
)

// @title SAC API
// @version 1.0
// @description Backend Server for SAC App

// @contact.name	David Oduneye and Garrett Ladley
// @contact.email	oduneye.d@northeastern.edu and ladley.g@northeastern.edu
// @host 127.0.0.1:8080
// @BasePath /api/v1
func main() {
	config, err := config.GetConfiguration("../../config")

	if err != nil {
		panic(err)
	}

	db, err := database.CreatePostgresConnection(config)

	if err != nil {
		panic(err)
	}

	app := server.Init(db)

	app.Listen(fmt.Sprintf("%s:%d", config.Application.Host, config.Application.Port))
}

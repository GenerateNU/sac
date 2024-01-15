package main

import (
	"backend/src/config"
	"backend/src/database"
	"backend/src/server"
	"flag"
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
	onlyMigrate := flag.Bool("only-migrate", false, "Specify if you want to only perform the database migration")

	flag.Parse()

	config, err := config.GetConfiguration("../../config")

	if err != nil {
		panic(err)
	}

	db, err := database.ConfigureDB(config)

	if err != nil {
		panic(err)
	}

	if *onlyMigrate {
		return
	}

	app := server.Init(db)

	app.Listen(fmt.Sprintf("%s:%d", config.Application.Host, config.Application.Port))
}

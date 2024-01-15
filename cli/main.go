package main

import (
	"log"
	"os"
	"path/filepath"

	"cli/commands"

	"github.com/urfave/cli/v2"
)


func main() {
	ROOT_DIR, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	BACKEND_DIR := filepath.Join(ROOT_DIR, "/backend/src")
	FRONTEND_DIR := filepath.Join(ROOT_DIR, "/frontend")

	app := &cli.App{
		Name:  "sac-cli",
		Usage: "CLI for SAC",
		Commands: []*cli.Command{
			commands.SwaggerCommand(BACKEND_DIR),
			// commands.StartCommand(BACKEND_DIR, FRONTEND_DIR), // Dont use
			commands.MigrateCommand(BACKEND_DIR),
			commands.TestCommand(BACKEND_DIR, FRONTEND_DIR),
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

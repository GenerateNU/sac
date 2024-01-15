package main

import (
	"log"
	"os"
	"path/filepath"

	"cli/commands"

	"github.com/urfave/cli/v2"
)

// Usage: mycli [command] [options]
// Commands:
// swagger 		  	   Generate swagger.json X
// start --frontend [name]   Start frontend and/or backend (frontend optional) run backend with ngrok instead of locally ( Run ngrok and set API_ENDPOINT env) X
// start --backend     Start backend X
// start X
// initdb              Initialize Docker for PostgreSQL [x]
// resetdb             Reset Docker for PostgreSQL
// test                Run tests
// test --frontend     Run frontend tests
// test --backend      Run backend tests
// lint                Run linter
// lint --frontend     Run frontend linter
// lint --backend      Run backend linter
// format              Run formatter
// format --frontend   Run frontend formatter
// format --backend    Run backend formatter

// Options:
// -h, --help          Show help message


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
			commands.StartCommand(BACKEND_DIR, FRONTEND_DIR), // Dont use
			commands.MigrateCommand(BACKEND_DIR),
			commands.TestCommand(BACKEND_DIR, FRONTEND_DIR),
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

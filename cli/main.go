package main

import (
	"fmt"
	"log"
	"os"

	"cli/commands"

	"github.com/urfave/cli/v2"
)

// Usage: mycli [command] [options]
// Commands:
// swagger 		  	   Generate swagger.json
// start --frontend    Start frontend and/or backend (frontend optional) run backend with ngrok instead of locally ( Run ngrok and set API_ENDPOINT env)
// initdb              Initialize Docker for PostgreSQL
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
	app := &cli.App{
		Name:  "sac-cli",
		Usage: "CLI for SAC",
		Commands: []*cli.Command{
			swaggerCommand(),
			{
				Name:  "start",
				Usage: "Start frontend and/or backend",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "ngrok",
						Aliases: []string{"n"},
						Value:   "",
						Usage:   "Run backend with ngrok instead of locally (requires ngrok URL)",
					},
					&cli.StringFlag{
						Name:    "frontend",
						Aliases: []string{"f"},
						Value:   "",
						Usage:   "Specify the frontend folder",
					},
				},
				Action: func(c *cli.Context) error {
					ngrok := c.Bool("ngrok")
					folder := c.String("frontend")

					fmt.Println("start", ngrok, folder)
					return nil
				},
			},
			{
				Name:  "initdb",
				Usage: "Initializes the database",
				Action: func(c *cli.Context) error {
					fmt.Println("initdb")
					return nil
				},
			},
			{
				Name:  "resetdb",
				Usage: "Resets the database",
				Action: func(c *cli.Context) error {
					fmt.Println("resetdb")
					return nil
				},
			},
			{
				Name:  "test",
				Usage: "Runs tests",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "frontend",
						Aliases: []string{"f"},
						Usage:   "Runs frontend tests",
					},
					&cli.BoolFlag{
						Name:    "backend",
						Aliases: []string{"b"},
						Usage:   "Runs backend tests",
					},
				},
				Action: func(c *cli.Context) error {
					fmt.Println("test", c.Bool("frontend"), c.Bool("backend"))
					return nil
				},
			},
			{
				Name:  "lint",
				Usage: "Runs linter",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "frontend",
						Aliases: []string{"f"},
						Usage:   "Runs frontend linter",
						Value:   "",
					},
					&cli.BoolFlag{
						Name:    "backend",
						Aliases: []string{"b"},
						Usage:   "Runs backend linter",
					},
				},
				Action: func(c *cli.Context) error {
					fmt.Println("lint", c.String("frontend"), c.Bool("backend"))
					return nil
				},
			},
			{
				Name:  "format",
				Usage: "Runs formatter",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "frontend",
						Aliases: []string{"f"},
						Usage:   "Runs frontend formatter",
						Value:   "",
					},
					&cli.BoolFlag{
						Name:    "backend",
						Aliases: []string{"b"},
						Usage:   "Runs backend formatter",
					},
				},
				Action: commands.Format,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func swaggerCommand() *cli.Command {
	command := cli.Command{
		Name:  "swagger",
		Usage: "Updates the swagger documentation",
		Action: commands.Swagger,
	}

	return &command
}

// func lintFrontend(currentDir string, folder string) {
// 	cmd := exec.Command("yarn", "lint") 
// 	cmd.Dir = filepath.Join(currentDir, "../frontend/", folder) 

// 	err := cmd.Run()
// 	if err != nil {
// 		fmt.Println("Error linting frontend, run yarn lint in frontend folder")
// 	}

// 	fmt.Println("frontend", cmd.Dir)
// }
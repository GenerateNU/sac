package main

import (
	"log"
	"os"

	"github.com/GenerateNU/sac/cli/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "sac",
		Usage: "CLI for the GenerateNU SAC",
		Commands: []*cli.Command{
			commands.SwaggerCommand(),
			commands.DBCommand(),
			commands.ClearDBCommand(),
			commands.InsertCommand(),
			commands.BackendCommand(),
			commands.FrontendCommand(),
			commands.TestCommand(), // TODO: frontend tests
			commands.FormatCommand(),
			commands.LintCommand(),
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

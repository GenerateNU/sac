package main

import (
	"log"
	"os"

	"github.com/GenerateNU/sac/cli/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "sac-cli",
		Usage: "CLI for the GenerateNU SAC",
		Commands: []*cli.Command{
			commands.SwaggerCommand(),
			commands.ClearDBCommand(),
			commands.MigrateCommand(),
			commands.ResetCommand(),
			commands.InsertCommand(),
			commands.DropCommand(),
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

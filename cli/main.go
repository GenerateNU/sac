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
		Usage: "CLI for SAC",
		Commands: []*cli.Command{
			commands.SwaggerCommand(),
			commands.ClearDBCommand(),
			commands.MigrateCommand(),
			commands.ResetDBCommand(),
			commands.DropDBCommand(),
			commands.TestCommand(), // TODO: frontend tests
			commands.FormatCommand(), // TODO: frontend format
			commands.LintCommand(), // TODO: frontend lint
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/GenerateNU/sac/cli/helpers"
	"github.com/spf13/cobra"
)

var swagCmd = &cobra.Command{
	Use:   "swagger",
	Short: "Run swagger init",
	Aliases: []string{"swag"},
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.Execute(exec.Command("swag", "init"), helpers.BACKEND_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(swagCmd)
}
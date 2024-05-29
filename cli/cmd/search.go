package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/GenerateNU/sac/cli/helpers"
	"github.com/spf13/cobra"
)

var seedSearchCmd = &cobra.Command{
	Use:   "seed-search",
	Short: "Seed Search",
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.Execute(exec.Command("go", "run", "main.go", "seed-search"), helpers.BACKEND_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(seedSearchCmd)
}

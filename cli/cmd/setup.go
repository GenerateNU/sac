package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Installs and sets up the project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not implemented")
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}

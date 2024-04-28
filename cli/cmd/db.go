package cmd

import (
	"fmt"
	"os"

	"github.com/GenerateNU/sac/cli/helpers"
	"github.com/spf13/cobra"
)

var dbCmd = &cobra.Command{
	Use:   "database",
	Aliases: []string{"db"},
	Short: "Database management commands",
}

func init() {
	rootCmd.AddCommand(dbCmd)
	dbCmd.AddCommand(dbInitCmd)
	dbCmd.AddCommand(dbDownCmd)
	dbCmd.AddCommand(dbResetCmd)
	dbCmd.AddCommand(dbInsertCmd)
}

var dbInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the database",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.InitDB()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var dbDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Migrate down the database",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.DownDB()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var dbResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the database",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.DownDB()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = helpers.InitDB()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var dbInsertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Insert mock data into the database",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.InsertDB()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
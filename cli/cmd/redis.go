package cmd

import (
	"fmt"
	"os"

	"github.com/GenerateNU/sac/cli/helpers"
	"github.com/spf13/cobra"
)

var redisCmd = &cobra.Command{
	Use:     "redis",
	Aliases: []string{"redis"},
	Short:   "Redis management commands",
}

func init() {
	rootCmd.AddCommand(redisCmd)
	redisCmd.AddCommand(redisInitCmd)
}

var redisInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize redis",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.InitRedis()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
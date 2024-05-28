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
	redisCmd.AddCommand(redisDownCmd)
	redisCmd.AddCommand(redisRestartCmd)
}

var redisInitCmd = &cobra.Command{
	Use:   "up",
	Short: "Start the Redis container(s)",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.UpRedis()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var redisDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Stop the Redis container(s)",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.DownRedis()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var redisRestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart the Redis container(s)",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.RestartRedis()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

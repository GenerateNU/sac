package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/GenerateNU/sac/cli/helpers"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Aliases: []string{"r"},
	Short: "Run the backend or frontend",
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.AddCommand(runFeCmd)
	runCmd.AddCommand(runBeCmd)

	runFeCmd.AddCommand(runWebCmd)
	runFeCmd.AddCommand(runMobileCmd)

	var platform string
	runMobileCmd.Flags().StringVarP(&platform, "platform", "p", "ios", "The platform to run the mobile frontend on")
	runFeCmd.AddCommand(runDashboardCmd)
}

var runBeCmd = &cobra.Command{
	Use:   "backend",
	Short: "Run the backend",
	Aliases: []string{"be"},
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.Execute(exec.Command("go", "run", "main.go"), helpers.BACKEND_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var runFeCmd = &cobra.Command{
	Use:   "frontend",
	Short: "Run the frontend",
	Aliases: []string{"fe"},
}

var runWebCmd = &cobra.Command{
	Use:   "web",
	Short: "Run the web frontend",
	Aliases: []string{"w"},
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.Execute(exec.Command("yarn", "run", "dev"), helpers.WEB_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var runMobileCmd = &cobra.Command{
	Use:   "mobile",
	Short: "Run the mobile frontend",
	Aliases: []string{"m"},
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		platform, _ := cmd.Flags().GetString("platform")

		platform = "run:" + platform

		err := helpers.Execute(exec.Command("npx", "expo", platform), helpers.MOBILE_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var runDashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Run the dashboard frontend",
	Aliases: []string{"d"},
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.Execute(exec.Command("yarn", "run", "dev"), helpers.DASHBOARD_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
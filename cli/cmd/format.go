package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/GenerateNU/sac/cli/helpers"
	"github.com/spf13/cobra"
)

var formatCmd = &cobra.Command{
	Use:     "format",
	Aliases: []string{"f"},
	Short:   "Formatting commands",
}

func init() {
	rootCmd.AddCommand(formatCmd)
	formatCmd.AddCommand(formatFrontendCmd)
	formatCmd.AddCommand(formatBackendCmd)
	formatCmd.AddCommand(formatCliCmd)

	formatFrontendCmd.AddCommand(formatWebCmd)
	formatFrontendCmd.AddCommand(formatMobileCmd)
	formatFrontendCmd.AddCommand(formatDashboardCmd)
	formatFrontendCmd.AddCommand(formatLibCmd)
}

var formatFrontendCmd = &cobra.Command{
	Use:     "frontend",
	Aliases: []string{"fe"},
	Short:   "Frontend formatting commands",
}

var formatWebCmd = &cobra.Command{
	Use:     "web",
	Aliases: []string{"w"},
	Short:   "Frontend web formatting commands",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.Execute(exec.Command("yarn", "run", "format"), helpers.WEB_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var formatMobileCmd = &cobra.Command{
	Use:     "mobile",
	Aliases: []string{"m"},
	Short:   "Formats the frontend mobile",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.Execute(exec.Command("yarn", "run", "format"), helpers.MOBILE_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var formatDashboardCmd = &cobra.Command{
	Use:     "dashboard",
	Aliases: []string{"d"},
	Short:   "Formats the frontend dashboard",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.Execute(exec.Command("yarn", "run", "format"), helpers.DASHBOARD_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var formatLibCmd = &cobra.Command{
	Use:     "lib",
	Aliases: []string{"l"},
	Short:   "Formats the frontend lib",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.Execute(exec.Command("yarn", "run", "format"), helpers.LIB_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var formatBackendCmd = &cobra.Command{
	Use:     "backend",
	Short:   "Formats the backend",
	Aliases: []string{"be"},
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := exec.LookPath("gofumpt")
		if err != nil {
			fmt.Println("gofumpt is not installed. Please run the following command to install it:")
			fmt.Println("go install mvdan.cc/gofumpt@latest")
			os.Exit(1)
		}

		err = helpers.Execute(exec.Command("templ", "generate", "&&", "gofumpt", "-l", "-w", "."), helpers.BACKEND_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var formatCliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Formats the cli",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := exec.LookPath("gofumpt")
		if err != nil {
			fmt.Println("gofumpt is not installed. Please run the following command to install it:")
			fmt.Println("go install mvdan.cc/gofumpt@latest")
			os.Exit(1)
		}

		err = helpers.Execute(exec.Command("templ", "generate", "&&", "gofumpt", "-l", "-w", "."), helpers.CLI_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

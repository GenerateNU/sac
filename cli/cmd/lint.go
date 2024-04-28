package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/GenerateNU/sac/cli/helpers"
	"github.com/spf13/cobra"
)

var lintCmd = &cobra.Command{
	Use:     "lint",
	Aliases: []string{"l"},
	Short:   "Linting commands",
}

func init() {
	rootCmd.AddCommand(lintCmd)
	lintCmd.AddCommand(lintFrontendCmd)
	lintCmd.AddCommand(lintBackendCmd)
	lintCmd.AddCommand(lintCliCmd)

	lintFrontendCmd.AddCommand(lintWebCmd)
	lintFrontendCmd.AddCommand(lintMobileCmd)
	lintFrontendCmd.AddCommand(lintDashboardCmd)
	lintFrontendCmd.AddCommand(lintLibCmd)
}

var lintFrontendCmd = &cobra.Command{
	Use:     "frontend",
	Aliases: []string{"fe"},
	Short:   "Frontend linting commands",
}

var lintWebCmd = &cobra.Command{
	Use:     "web",
	Short:   "Lints the frontend web",
	Aliases: []string{"w"},
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.Execute(exec.Command("yarn", "run", "lint", "--fix"), helpers.WEB_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var lintMobileCmd = &cobra.Command{
	Use:     "mobile",
	Short:   "Lints the frontend mobile",
	Aliases: []string{"m"},
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.Execute(exec.Command("yarn", "run", "lint", "--fix"), helpers.MOBILE_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var lintDashboardCmd = &cobra.Command{
	Use:     "dashboard",
	Short:   "Lints the frontend dashboard",
	Aliases: []string{"d"},
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.Execute(exec.Command("yarn", "run", "lint", "--fix"), helpers.DASHBOARD_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var lintLibCmd = &cobra.Command{
	Use:     "lib",
	Short:   "Lints the lib directory",
	Aliases: []string{"l"},
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.Execute(exec.Command("yarn", "run", "lint", "--fix"), helpers.LIB_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var lintBackendCmd = &cobra.Command{
	Use:     "backend",
	Short:   "Lints the backend",
	Aliases: []string{"be"},
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := exec.LookPath("golangci-lint")
		if err != nil {
			fmt.Println("golangci-lint is not installed. Please install it by running the following command:")
			fmt.Println("curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.57.2")
			os.Exit(1)
		}

		err = helpers.Execute(exec.Command("golangci-lint", "run", "--fix"), helpers.BACKEND_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var lintCliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Lints the CLI",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := exec.LookPath("golangci-lint")
		if err != nil {
			fmt.Println("golangci-lint is not installed. Please install it by running the following command:")
			fmt.Println("curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.57.2")
			os.Exit(1)
		}

		err = helpers.Execute(exec.Command("golangci-lint", "run", "--fix"), helpers.CLI_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

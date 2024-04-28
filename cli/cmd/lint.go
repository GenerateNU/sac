package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/GenerateNU/sac/cli/helpers"
	"github.com/spf13/cobra"
)

var lintCmd = &cobra.Command{
	Use:   "lint",
	Aliases: []string{"l"},
	Short: "Linting commands",
}

func init() {
	rootCmd.AddCommand(lintCmd)
	lintCmd.AddCommand(lintFrontendCmd)
	lintCmd.AddCommand(lintBackendCmd)

	lintFrontendCmd.AddCommand(lintWebCmd)
	lintFrontendCmd.AddCommand(lintMobileCmd)
	lintFrontendCmd.AddCommand(lintDashboardCmd)
	lintFrontendCmd.AddCommand(lintLibCmd)

	lintWebCmd.Flags().BoolP("fix", "f", false, "Fix linting errors")
	lintMobileCmd.Flags().BoolP("fix", "f", false, "Fix linting errors")
	lintDashboardCmd.Flags().BoolP("fix", "f", false, "Fix linting errors")
}

var lintFrontendCmd = &cobra.Command{
	Use:   "frontend",
	Aliases: []string{"fe"},
	Short: "Frontend linting commands",
}

var lintWebCmd = &cobra.Command{
	Use:   "web",
	Short: "Frontend web linting commands",
	Aliases: []string{"w"},
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fix, _ := cmd.Flags().GetBool("fix")

		var command string
		if fix {
			command = " --fix"
		}

		err := helpers.Execute(exec.Command("yarn", "run", "lint"+command), helpers.WEB_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var lintMobileCmd = &cobra.Command{
	Use:   "mobile",
	Short: "Frontend mobile linting commands",
	Aliases: []string{"m"},
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fix, _ := cmd.Flags().GetBool("fix")

		var command string
		if fix {
			command = " --fix"
		}

		err := helpers.Execute(exec.Command("yarn", "run", "lint"+command), helpers.MOBILE_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var lintDashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Frontend dashboard linting commands",
	Aliases: []string{"d"},
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fix, _ := cmd.Flags().GetBool("fix")

		var command string
		if fix {
			command = " --fix"
		}

		err := helpers.Execute(exec.Command("yarn", "run", "lint"+command), helpers.DASHBOARD_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var lintLibCmd = &cobra.Command{
	Use:   "lib",
	Short: "Frontend lib linting commands",
	Aliases: []string{"l"},
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fix, _ := cmd.Flags().GetBool("fix")

		var command string
		if fix {
			command = " --fix"
		}

		err := helpers.Execute(exec.Command("yarn", "run", "lint"+command), helpers.LIB_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var lintBackendCmd = &cobra.Command{
	Use:   "backend",
	Short: "Backend linting commands",
	Aliases: []string{"be"},
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: golangci-lint 
		err := helpers.Execute(exec.Command("go", "lint", "./..."), helpers.BACKEND_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

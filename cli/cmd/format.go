package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/GenerateNU/sac/cli/helpers"
	"github.com/spf13/cobra"
)

var formatCmd = &cobra.Command{
	Use:   "format",
	Aliases: []string{"f"},
	Short: "Executeing commands",
}

func init() {
	rootCmd.AddCommand(formatCmd)
	formatCmd.AddCommand(formatFrontendCmd)
	formatCmd.AddCommand(formatBackendCmd)

	formatFrontendCmd.AddCommand(formatWebCmd)
	formatFrontendCmd.AddCommand(formatMobileCmd)
	formatFrontendCmd.AddCommand(formatDashboardCmd)
	formatFrontendCmd.AddCommand(formatLibCmd)
}

var formatFrontendCmd = &cobra.Command{
	Use:   "frontend",
	Aliases: []string{"fe"},
	Short: "Frontend formatting commands",
}

var formatWebCmd = &cobra.Command{
	Use:   "web",
	Aliases: []string{"w"},
	Short: "Frontend web formatting commands",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.Execute(exec.Command("yarn", "run", "format"), helpers.WEB_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var formatMobileCmd = &cobra.Command{
	Use:   "mobile",
	Aliases: []string{"m"},
	Short: "Frontend mobile formatting commands",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.Execute(exec.Command("yarn", "run", "format"), helpers.MOBILE_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var formatDashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Aliases: []string{"d"},
	Short: "Frontend dashboard formatting commands",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.Execute(exec.Command("yarn", "run", "format"), helpers.DASHBOARD_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var formatLibCmd = &cobra.Command{
	Use:   "lib",
	Aliases: []string{"l"},
	Short: "Frontend lib formatting commands",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.Execute(exec.Command("yarn", "run", "format"), helpers.LIB_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var formatBackendCmd = &cobra.Command{
	Use:   "backend",
	Short: "Backend formatting commands",
	Aliases: []string{"be"},
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.Execute(exec.Command("go", "fmt", "./..."), helpers.BACKEND_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

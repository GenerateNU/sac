package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/GenerateNU/sac/cli/helpers"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Aliases: []string{"t"},
	Short: "Testing commands",
}

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.AddCommand(testFrontendCmd)
	testCmd.AddCommand(testBackendCmd)
	testCmd.AddCommand(testCleanDatabaseCmd)

	testBackendCmd.Flags().BoolP("clean", "c", false, "Clean the test cache after running tests")

	testFrontendCmd.AddCommand(testWebCmd)
	testFrontendCmd.AddCommand(testMobileCmd)
	testFrontendCmd.AddCommand(testDashboardCmd)
	testFrontendCmd.AddCommand(testLibCmd)
}

var testFrontendCmd = &cobra.Command{
	Use:   "frontend",
	Aliases: []string{"fe"},
	Short: "Frontend testing commands",
}

var testWebCmd = &cobra.Command{
	Use:   "web",
	Short: "Frontend web testing commands",
	Aliases: []string{"w"},
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.Execute(exec.Command("yarn", "run", "test"), helpers.WEB_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var testMobileCmd = &cobra.Command{
	Use:   "mobile",
	Short: "Frontend mobile testing commands",
	Aliases: []string{"m"},
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.Execute(exec.Command("yarn", "run", "test"), helpers.MOBILE_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var testDashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Frontend dashboard testing commands",
	Aliases: []string{"d"},
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.Execute(exec.Command("yarn", "run", "test"), helpers.DASHBOARD_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var testLibCmd = &cobra.Command{
	Use:   "lib",
	Short: "Frontend lib testing commands",
	Aliases: []string{"l"},
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.Execute(exec.Command("yarn", "run", "test"), helpers.LIB_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var testBackendCmd = &cobra.Command{
	Use:   "backend",
	Short: "Backend testing commands",
	Aliases: []string{"be"},
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		clear, _ := cmd.Flags().GetBool("clean")

		err := helpers.Execute(exec.Command("go", "test", "./..."), helpers.BACKEND_DIR)
		if err != nil {
			fmt.Println(err)
		}

		if clear {
			err := helpers.CleanTestDBs()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	},
}


var testCleanDatabaseCmd = &cobra.Command{
	Use:   "clean-db",
	Short: "Remove databases used for testing",
	Aliases: []string{"clean"},
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.CleanTestDBs()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
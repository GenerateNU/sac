package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/GenerateNU/sac/cli/helpers"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:     "run",
	Aliases: []string{"r"},
	Short:   "Run commands",
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.AddCommand(runFeCmd)
	runCmd.AddCommand(runBeCmd)
	runBeCmd.Flags().BoolP("migrate", "m", false, "Specify if you want to only perform the database migration")
	runBeCmd.Flags().BoolP("seed-pinecone", "s", false, "Specify if you want to only perform the pinecone database seeding")
	runBeCmd.Flags().BoolP("connect-to-pinecone", "c", false, "Connect to a real Pinecone instance instead of mock")

	runFeCmd.AddCommand(runWebCmd)
	runFeCmd.AddCommand(runMobileCmd)

	var platform string
	runMobileCmd.Flags().StringVarP(&platform, "platform", "p", "ios", "The platform to run the mobile frontend on")
	runFeCmd.AddCommand(runDashboardCmd)
}

var runBeCmd = &cobra.Command{
	Use:     "backend",
	Short:   "Run the backend",
	Aliases: []string{"be"},
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		migrate, _ := cmd.Flags().GetBool("migrate")
		seedPinecone, _ := cmd.Flags().GetBool("seed-pinecone")
		connectToPinecone, _ := cmd.Flags().GetBool("connect-to-pinecone")

		if migrate {
			err := helpers.Execute(exec.Command("go", "run", "main.go", "--only-migrate"), helpers.BACKEND_DIR)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		if seedPinecone {
			err := helpers.Execute(exec.Command("go", "run", "main.go", "--seed-pinecone"), helpers.BACKEND_DIR)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		if connectToPinecone {
			err := helpers.Execute(exec.Command("go", "run", "main.go", "--connect-to-pinecone"), helpers.BACKEND_DIR)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		if !migrate && !seedPinecone && !connectToPinecone {
			err := helpers.Execute(exec.Command("go", "run", "main.go"), helpers.BACKEND_DIR)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	},
}

var runFeCmd = &cobra.Command{
	Use:     "frontend",
	Short:   "Frontend run commands",
	Aliases: []string{"fe"},
}

var runWebCmd = &cobra.Command{
	Use:     "web",
	Short:   "Run the web frontend",
	Aliases: []string{"w"},
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.Execute(exec.Command("yarn", "run", "dev"), helpers.WEB_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var runMobileCmd = &cobra.Command{
	Use:     "mobile",
	Short:   "Run the mobile frontend",
	Aliases: []string{"m"},
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		platform, _ := cmd.Flags().GetString("platform")

		platform = "run:" + platform

		output := exec.Command("npx", "expo", platform)
		output.Dir = helpers.MOBILE_DIR
		output.Stdout = os.Stdout
		output.Stderr = os.Stderr
		output.Stdin = os.Stdin

		err := helpers.Execute(output, helpers.MOBILE_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var runDashboardCmd = &cobra.Command{
	Use:     "dashboard",
	Short:   "Run the dashboard frontend",
	Aliases: []string{"d"},
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.Execute(exec.Command("yarn", "run", "dev"), helpers.DASHBOARD_DIR)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

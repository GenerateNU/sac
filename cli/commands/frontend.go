package commands

import (
	"os"
	"os/exec"
	"slices"

	"github.com/urfave/cli/v2"
)

func FrontendCommand() *cli.Command {
	command := cli.Command{
		Name:     "fe",
		Usage:    "Run the frontend",
		Category: "Development",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "target",
				Aliases: []string{"t"},
				Value:   "mobile",
				Usage:   "Run a specific frontend type (web or mobile)",
			},
			&cli.StringFlag{
				Name:    "platform",
				Aliases: []string{"p"},
				Usage:   "Run a specific platform for mobile frontend",
				Value:   "ios",
			},
		},
		Action: func(c *cli.Context) error {
			if c.Args().Len() > 0 {
				return cli.Exit("Invalid arguments", 1)
			}

			if !slices.Contains(FRONTEND_RUN_TARGETS, c.String("target")) {
				return cli.Exit("Invalid frontend type: must be 'web', 'mobile', or 'dashboard'", 1)
			}

			err := RunFE(c.String("type"), c.String("platform"))
			if err != nil {
				return cli.Exit(err.Error(), 1)
			}

			return nil
		},
	}

	return &command
}

func RunFE(feType string, platform string) error {
	switch feType {
	case "mobile":
		return RunMobileFE(platform)
	case "web":
		return RunWebFE()
	case "dashboard":
		return RunDashboardFE()
	default:
		return RunMobileFE(platform)
	}
}

func RunMobileFE(platform string) error {
	mobileCmd := exec.Command("yarn", "run", platform)
	mobileCmd.Dir = MOBILE_DIR

	mobileCmd.Stdout = os.Stdout
	mobileCmd.Stderr = os.Stderr
	mobileCmd.Stdin = os.Stdin

	if err := mobileCmd.Run(); err != nil {
		return err
	}

	return nil
}

func RunWebFE() error {
	webCmd := exec.Command("yarn", "run")
	webCmd.Dir = WEB_DIR

	webCmd.Stdout = os.Stdout
	webCmd.Stderr = os.Stderr
	webCmd.Stdin = os.Stdin

	if err := webCmd.Run(); err != nil {
		return err
	}

	return nil
}

func RunDashboardFE() error {
	dashboardCmd := exec.Command("yarn", "run")
	dashboardCmd.Dir = DASHBOARD_DIR

	dashboardCmd.Stdout = os.Stdout
	dashboardCmd.Stderr = os.Stderr
	dashboardCmd.Stdin = os.Stdin

	if err := dashboardCmd.Run(); err != nil {
		return err
	}

	return nil
}

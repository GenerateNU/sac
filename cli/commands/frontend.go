package commands

import (
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func FrontendCommand() *cli.Command {
	command := cli.Command{
		Name:     "fe",
		Usage:    "Run the frontend",
		Category: "Development",
		Subcommands: []*cli.Command{
			{
				Name:  "mobile",
				Aliases: []string{"m"},
				Usage: "Run the mobile frontend",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "platform",
						Usage: "The platform to run the mobile frontend on",
						Value: "ios",
					},
				},
				Action: func(c *cli.Context) error {
					return RunFE("mobile", c.String("platform"))
				},
			},
			{
				Name:  "web",
				Aliases: []string{"w"},
				Usage: "Run the web frontend",
				Action: func(c *cli.Context) error {
					return RunFE("web", "")
				},
			},
			{
				Name:  "dashboard",
				Aliases: []string{"d"},
				Usage: "Run the dashboard frontend",
				Action: func(c *cli.Context) error {
					return RunFE("dashboard", "")
				},
			},
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
	mobileCmd.Dir = FRONTEND_DIR + "/mobile"

	mobileCmd.Stdout = os.Stdout
	mobileCmd.Stderr = os.Stderr
	mobileCmd.Stdin = os.Stdin

	if err := mobileCmd.Run(); err != nil {
		return err
	}

	return nil
}

func RunWebFE() error {
	webCmd := exec.Command("yarn", "run", "dev")
	webCmd.Dir = FRONTEND_DIR + "/web"

	webCmd.Stdout = os.Stdout
	webCmd.Stderr = os.Stderr
	webCmd.Stdin = os.Stdin

	if err := webCmd.Run(); err != nil {
		return err
	}

	return nil
}

func RunDashboardFE() error {
	dashboardCmd := exec.Command("yarn", "run", "dev")
	dashboardCmd.Dir = FRONTEND_DIR + "/dashboard"

	dashboardCmd.Stdout = os.Stdout
	dashboardCmd.Stderr = os.Stderr
	dashboardCmd.Stdin = os.Stdin

	if err := dashboardCmd.Run(); err != nil {
		return err
	}

	return nil
}

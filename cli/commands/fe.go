package commands

import (
	"os"
	"os/exec"

	_ "github.com/lib/pq"
	"github.com/urfave/cli/v2"
)

func RunFrontendCommand() *cli.Command {
	command := cli.Command{
		Name:  "fe",
		Usage: "Run the frontend",
		Category: "Development",
		Flags: []cli.Flag{
			&cli.StringFlag{
                Name:    "target",
                Aliases: []string{"t"},
                Value:   "mobile",
                Usage:   "Run a specific frontend type (web or mobile)",
            },
			&cli.StringFlag{
				Name:  "platform",
				Aliases: []string{"p"},
				Usage: "Run a specific platform for mobile frontend",
				Value: "ios",
			},
		},
		Action: func(c *cli.Context) error {
			if c.Args().Len() > 0 {
				return cli.Exit("Invalid arguments", 1)
			}

			target := c.String("target")
            if target != "web" && target != "mobile" {
                return cli.Exit("Invalid frontend type: must be 'web' or 'mobile'", 1)
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
	default:
		return RunMobileFE(platform)
	}
}

func RunMobileFE(platform string) error {
    mobileCmd := exec.Command("yarn", "run", platform)
    mobileCmd.Dir = FRONTEND_DIR + "/sac-mobile"

    mobileCmd.Stdout = os.Stdout
    mobileCmd.Stderr = os.Stderr
    mobileCmd.Stdin = os.Stdin 

    if err := mobileCmd.Run(); err != nil {
        return err
    }

    return nil
}

func RunWebFE() error {
	return nil
}
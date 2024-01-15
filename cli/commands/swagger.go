package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func Swagger(c *cli.Context) error {
	if c.Args().Len() > 0 {
		return cli.Exit("Invalid arguments", 1)
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return cli.Exit("Error getting current directory", 1)
	}

	backendDir := filepath.Join(currentDir, "../backend/src")

	cmd := exec.Command("swag", "init")
	cmd.Dir = backendDir

	err = cmd.Run()
	if err != nil {
		return cli.Exit("Error generating swagger docs", 1)
	}

	fmt.Println("Swagger docs generated")
	
	return nil
}
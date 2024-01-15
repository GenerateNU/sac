package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func InsertDB(c *cli.Context) error {
	if c.Args().Len() > 0 {
		return cli.Exit("Invalid arguments", 1)
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return cli.Exit("Error getting current directory", 1)
	}

	backendDir := filepath.Join(currentDir, "../backend/src")

	cmd := exec.Command("./init_db.sh")
	cmd.Dir = backendDir

	output, err := cmd.CombinedOutput()
	if err != nil {
		return cli.Exit("Error running init_db.sh", 1)
	}

	fmt.Println(string(output))
	
	return nil
}
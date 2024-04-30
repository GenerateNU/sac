package helpers

import (
	"fmt"
	"os/exec"
)

func InitRedis() error {
	err := Execute(exec.Command("./scripts/init_redis.sh"), ROOT_DIR)
	if err != nil {
		return fmt.Errorf("error initializing redis: %w", err)
	}

	return nil
}
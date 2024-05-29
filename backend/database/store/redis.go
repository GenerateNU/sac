package store

import (
	"bytes"
	"fmt"
	"log/slog"
	"os/exec"

	"github.com/GenerateNU/sac/backend/config"
)

type Stores struct {
	Limiter     LimiterInterface
	Blacklist   BlacklistInterface
	ActiveToken ActiveTokenInterface
}

func NewStores(limiter LimiterInterface, blacklist BlacklistInterface, activeToken ActiveTokenInterface) *Stores {
	return &Stores{
		Limiter:     limiter,
		Blacklist:   blacklist,
		ActiveToken: activeToken,
	}
}

func ConfigureRedis(settings config.Settings) *Stores {
	stores := NewStores(
		NewLimiter(NewRedisClient(settings.RedisLimiter.Username, settings.RedisLimiter.Host, settings.RedisLimiter.Port, settings.RedisLimiter.Password, settings.RedisLimiter.DB)),
		NewBlacklist(NewRedisClient(settings.RedisBlacklist.Username, settings.RedisBlacklist.Host, settings.RedisBlacklist.Port, settings.RedisBlacklist.Password, settings.RedisBlacklist.DB)),
		NewActiveToken(NewRedisClient(settings.RedisActiveTokens.Username, settings.RedisActiveTokens.Host, settings.RedisActiveTokens.Port, settings.RedisActiveTokens.Password, settings.RedisActiveTokens.DB)),
	)

	MustEstablishConn()

	return stores
}

// TODO: this will be a part of a larger function that will check all services
func MustEstablishConn() {
	isRunning, err := isDockerComposeRunning()
	if err != nil {
		panic(err)
	}

	if !*isRunning {
		if err := restartServices(); err != nil {
			panic(err)
		}
	}
}

func isDockerComposeRunning() (*bool, error) {
	cmd := exec.Command("docker-compose", "ps", "-q")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("error checking if docker-compose is running: %s", err.Error())
	}

	result := out.Len() > 0
	return &result, nil
}

func restartServices() error {
	cmd := exec.Command("docker-compose", "up", "-d")

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error restarting services: %s\nconsole output: %s", err.Error(), out.String())
	}

	slog.Info("Services restarted successfully.")
	return nil
}

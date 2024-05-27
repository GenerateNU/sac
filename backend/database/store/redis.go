package store

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"os/exec"

	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/utilities"
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
	var activeTokens *ActiveToken
	var blacklist *Blacklist
	var limiter *Limiter

	for _, service := range settings.Redis {
		client := NewRedisClient(service.Username, service.Host, service.Port, service.Password, service.DB)
		switch service.Username {
		case "redis_active_tokens":
			activeTokens = NewActiveToken(client, context.Background())
		case "redis_blacklist":
			blacklist = NewBlacklist(client, context.Background())
		case "redis_limiter":
			limiter = NewLimiter(client, context.Background())
		default:
			slog.Error("unknown redis service", service.Username, "skipping...")
		}
	}

	stores := NewStores(limiter, blacklist, activeTokens)

	MustEstablishConn()

	return stores
}

// TODO: this will be a part of a larger function that will check all services
func MustEstablishConn() {
	if !isDockerComposeRunning() {
		restartServices()
	}
}

func isDockerComposeRunning() bool {
	cmd := exec.Command("docker-compose", "ps", "-q")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		utilities.Exit(fmt.Sprintf("Error checking if Docker Compose is running: %s", err.Error()))
	}

	return out.Len() > 0
}

func restartServices() {
	cmd := exec.Command("docker-compose", "up", "-d")

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		slog.Error("error restarting services: ", err.Error(), "output: ", out.String(), "exiting...")
		return
	}

	slog.Info("Services restarted successfully.")
}

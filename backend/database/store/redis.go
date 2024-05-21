package store

import (
	"bytes"
	"fmt"
	"log/slog"
	"os/exec"
	"strings"

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
	var activeTokens *ActiveToken
	var blacklist *Blacklist
	var limiter *Limiter

	for _, service := range settings.Redis {
		client := NewRedisClient(service.Username, service.Host, service.Port, service.Password, service.DB)
		switch service.Username {
		case "redis_active_tokens":
			activeTokens = NewActiveToken(client)
		case "redis_blacklist":
			blacklist = NewBlacklist(client)
		case "redis_limiter":
			limiter = NewLimiter(client)
		default:
			slog.Error(fmt.Sprintf("unknown redis service: %s", service.Username))
		}
	}

	stores := NewStores(limiter, blacklist, activeTokens)

	EstablishConn(settings)

	return stores
}

func EstablishConn(settings config.Settings) {
	services := settings.Redis

	for _, service := range services {
		if !isContainerRunning(service.Username) {
			restartServices()
		}
	}
}

func isContainerRunning(name string) bool {
	sanitizedName := strings.ReplaceAll(name, " ", "")
	cmd := exec.Command("docker", "ps", "-f", fmt.Sprintf("name=%s", sanitizedName), "-q") //nolint:gosec
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		slog.Error(fmt.Sprintf("error checking if service %s is running: %v\noutput: %s\n", name, err, out.String()))
		return false
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
		slog.Error(fmt.Sprintf("error restarting services: %v\noutput: %s\n", err, out.String()))
	}
}

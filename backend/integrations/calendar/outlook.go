package calendar

import (
	"github.com/GenerateNU/sac/backend/config"
)

type OutlookClient struct {
	Settings *config.OutlookClientSettings
}

func (client OutlookClient) Authorize() string { return "" }

func (client OutlookClient) Token(grantType string, code *string) error { return nil }
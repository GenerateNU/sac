package config

import "fmt"

type ApplicationSettings struct {
	Port      uint16 `env:"PORT"`
	Host      string `env:"HOST"`
	BaseUrl   string `env:"BASE_URL"`
	PublicUrl string `env:"PUBLIC_URL"`
}

func (s *ApplicationSettings) ApplicationURL() string {
	var host string
	if s.Host == "127.0.0.1" {
		host = "localhost"
	} else {
		host = s.Host
	}
	return fmt.Sprintf("http://%s:%d", host, s.Port)
}

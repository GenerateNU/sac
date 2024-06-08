package config

import "fmt"

type ApplicationSettings struct {
	Port    uint16 `env:"PORT"`
	Host    string `env:"HOST"`
	BaseUrl string `env:"BASE_URL"`
}

func (s *ApplicationSettings) ApplicationURL() string {
	return fmt.Sprintf("http://%s:%d", s.Host, s.Port)
}

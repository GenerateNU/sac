package config

type ApplicationSettings struct {
	Port      uint16 `env:"PORT"`
	Host      string `env:"HOST"`
	PublicURL string `env:"PUBLIC_URL"`
}

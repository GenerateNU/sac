package config

type ApplicationSettings struct {
	Port    uint16 `env:"PORT"`
	Host    string `env:"HOST"`
	BaseUrl string `env:"BASE_URL"`
}

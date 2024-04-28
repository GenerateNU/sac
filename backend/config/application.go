package config

type ApplicationSettings struct {
	Port    uint16 `yaml:"port"`
	Host    string `yaml:"host"`
	BaseUrl string `yaml:"baseurl"`
}

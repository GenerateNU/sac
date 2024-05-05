package config

// import (
// 	"errors"
// 	"fmt"

// 	m "github.com/garrettladley/mattress"
// )

// type RedisSettings struct {
// 	Host     string            `yaml:"host"`
// 	Port     uint              `yaml:"port"`
// 	Password *m.Secret[string] `yaml:"password"`
// 	DB       int               `yaml:"db"`
// }

// func (int *intermediateRedisSettings) into() (*RedisSettings, error) {
// 	password, err := m.NewSecret(int.Password)
// 	if err != nil {
// 		return nil, errors.New("failed to create secret from password")
// 	}

// 	return &RedisSettings{
// 		Host:     int.Host,
// 		Port:     int.Port,
// 		Password: password,
// 		DB:       int.DB,
// 	}, nil
// }

// type intermediateRedisSettings struct {
// 	Host     string `yaml:"host"`
// 	Port     uint   `yaml:"port"`
// 	Password string `yaml:"password"`
// 	DB       int    `yaml:"db"`
// }

// func (s *RedisSettings) Conn() string {
// 	return fmt.Sprintf("host=%s port=%d password=%s db=%d",
// 		s.Host, s.Port, s.Password.Expose(), s.DB)
// }

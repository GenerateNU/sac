package config

import (
	"os"
	"github.com/joho/godotenv"
)

type AWSSettings struct {
	BUCKET_NAME string
	ID          string
	SECRET      string
}

func ConfigAWS() AWSSettings {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	return AWSSettings{
		BUCKET_NAME: os.Getenv("BUCKET_NAME"),
		ID:          os.Getenv("AWS_ID"),
		SECRET:      os.Getenv("AWS_SECRET")}
}
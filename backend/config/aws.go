package config

import (
	"fmt"

	m "github.com/garrettladley/mattress"
)

type AWSSettings struct {
	BucketName *m.Secret[string]
	Id         *m.Secret[string]
	Secret     *m.Secret[string]
	Region     *m.Secret[string]
}

type intermediateAWSSettings struct {
	BucketName string `env:"BUCKET_NAME"`
	Id         string `env:"ID"`
	Secret     string `env:"SECRET"`
	Region     string `env:"REGION"`
}

func (i *intermediateAWSSettings) into() (*AWSSettings, error) {
	bucketName, err := m.NewSecret(i.BucketName)
	if err != nil {
		return nil, fmt.Errorf("failed to create secret from bucket name: %s", err.Error())
	}

	id, err := m.NewSecret(i.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to create secret from ID: %s", err.Error())
	}

	secret, err := m.NewSecret(i.Secret)
	if err != nil {
		return nil, fmt.Errorf("failed to create secret from secret: %s", err.Error())
	}

	region, err := m.NewSecret(i.Region)
	if err != nil {
		return nil, fmt.Errorf("failed to create secret from region: %s", err.Error())
	}

	return &AWSSettings{
		BucketName: bucketName,
		Id:         id,
		Secret:     secret,
		Region:     region,
	}, nil
}

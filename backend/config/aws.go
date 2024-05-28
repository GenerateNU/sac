package config

import (
	"fmt"

	m "github.com/garrettladley/mattress"
)

type AWSSettings struct {
	BucketName *m.Secret[string]
	ID         *m.Secret[string]
	Secret     *m.Secret[string]
	Region     *m.Secret[string]
}

type intermediateAWSSettings struct {
	bucketName string `env:"BUCKET_NAME"`
	id         string `env:"ID"`
	secret     string `env:"SECRET"`
	region     string `env:"REGION"`
}

func (i *intermediateAWSSettings) into() (*AWSSettings, error) {
	bucketName, err := m.NewSecret(i.bucketName)
	if err != nil {
		return nil, fmt.Errorf("failed to create secret from bucket name: %s", err.Error())
	}

	id, err := m.NewSecret(i.id)
	if err != nil {
		return nil, fmt.Errorf("failed to create secret from ID: %s", err.Error())
	}

	secret, err := m.NewSecret(i.secret)
	if err != nil {
		return nil, fmt.Errorf("failed to create secret from secret: %s", err.Error())
	}

	region, err := m.NewSecret(i.region)
	if err != nil {
		return nil, fmt.Errorf("failed to create secret from region: %s", err.Error())
	}

	return &AWSSettings{
		BucketName: bucketName,
		ID:         id,
		Secret:     secret,
		Region:     region,
	}, nil
}

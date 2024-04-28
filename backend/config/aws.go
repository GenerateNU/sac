package config

import (
	"errors"
	"os"

	m "github.com/garrettladley/mattress"
)

type AWSSettings struct {
	BUCKET_NAME *m.Secret[string]
	ID          *m.Secret[string]
	SECRET      *m.Secret[string]
	REGION      *m.Secret[string]
}

func readAWSSettings() (*AWSSettings, error) {
	bucketName := os.Getenv("SAC_AWS_BUCKET_NAME")
	if bucketName == "" {
		return nil, errors.New("SAC_AWS_BUCKET_NAME is not set")
	}

	secretBucketName, err := m.NewSecret(bucketName)
	if err != nil {
		return nil, errors.New("failed to create secret from bucket name")
	}

	id := os.Getenv("SAC_AWS_ID")
	if id == "" {
		return nil, errors.New("SAC_AWS_ID is not set")
	}

	secretID, err := m.NewSecret(id)
	if err != nil {
		return nil, errors.New("failed to create secret from id")
	}

	secret := os.Getenv("SAC_AWS_SECRET")
	if secret == "" {
		return nil, errors.New("SAC_AWS_SECRET is not set")
	}

	secretSecret, err := m.NewSecret(secret)
	if err != nil {
		return nil, errors.New("failed to create secret from secret")
	}

	region := os.Getenv("SAC_AWS_REGION")
	if region == "" {
		return nil, errors.New("SAC_AWS_REGION is not set")
	}

	reigonSecret, err := m.NewSecret(region)
	if err != nil {
		return nil, errors.New("failed to create secret from region")
	}

	return &AWSSettings{
		BUCKET_NAME: secretBucketName,
		ID:          secretID,
		SECRET:      secretSecret,
		REGION:      reigonSecret,
	}, nil
}

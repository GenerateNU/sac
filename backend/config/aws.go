package config

import m "github.com/garrettladley/mattress"

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
		return nil, err
	}

	id, err := m.NewSecret(i.Id)
	if err != nil {
		return nil, err
	}

	secret, err := m.NewSecret(i.Secret)
	if err != nil {
		return nil, err
	}

	region, err := m.NewSecret(i.Region)
	if err != nil {
		return nil, err
	}

	return &AWSSettings{
		BucketName: bucketName,
		Id:         id,
		Secret:     secret,
		Region:     region,
	}, nil
}

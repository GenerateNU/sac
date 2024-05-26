package config

import (
	"os"
	"strconv"
	"time"

	"github.com/garrettladley/mattress"
	"github.com/joho/godotenv"
)

func GetConfiguration(path string) (*Settings, error) {
	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}

	applicationPort, err := strconv.Atoi(os.Getenv("SAC_APPLICATION_PORT"))
	if err != nil {
		return nil, err
	}

	applicationHost := os.Getenv("SAC_APPLICATION_HOST")

	applicationBaseUrl := os.Getenv("SAC_APPLICATION_BASE_URL")

	databaseUsername := os.Getenv("SAC_DB_USERNAME")

	databasePassword, err := mattress.NewSecret(os.Getenv("SAC_DB_PASSWORD"))
	if err != nil {
		return nil, err
	}

	databasePort, err := strconv.Atoi(os.Getenv("SAC_DB_PORT"))
	if err != nil {
		return nil, err
	}

	databaseHost := os.Getenv("SAC_DB_HOST")
	databaseName := os.Getenv("SAC_DB_NAME")

	databaseRequireSsl, err := strconv.ParseBool(os.Getenv("SAC_DB_REQUIRE_SSL"))
	if err != nil {
		return nil, err
	}

	superuserPassword, err := mattress.NewSecret(os.Getenv("SAC_SUDO_PASSWORD"))
	if err != nil {
		return nil, err
	}

	authAccessKey, err := mattress.NewSecret(os.Getenv("SAC_AUTH_ACCESS_KEY"))
	if err != nil {
		return nil, err
	}

	authRefreshKey, err := mattress.NewSecret(os.Getenv("SAC_AUTH_REFRESH_KEY"))
	if err != nil {
		return nil, err
	}

	awsBucketName, err := mattress.NewSecret(os.Getenv("SAC_AWS_BUCKET_NAME"))
	if err != nil {
		return nil, err
	}
	awsId, err := mattress.NewSecret(os.Getenv("SAC_AWS_ID"))
	if err != nil {
		return nil, err
	}

	awsSecret, err := mattress.NewSecret(os.Getenv("SAC_AWS_SECRET"))
	if err != nil {
		return nil, err
	}

	awsRegion, err := mattress.NewSecret(os.Getenv("SAC_AWS_REGION"))
	if err != nil {
		return nil, err
	}

	resendApiKey, err := mattress.NewSecret(os.Getenv("SAC_RESEND_API_KEY"))
	if err != nil {
		return nil, err
	}

	calendarMaxTerminationDate, err := time.Parse("01-02-2006", os.Getenv("SAC_CALENDAR_MAX_TERMINATION_DATE"))
	if err != nil {
		return nil, err
	}

	return &Settings{
		Application: ApplicationSettings{
			Port:    uint16(applicationPort),
			Host:    applicationHost,
			BaseUrl: applicationBaseUrl,
		},
		Database: DatabaseSettings{
			Username:     databaseUsername,
			Password:     databasePassword,
			Port:         uint(databasePort),
			Host:         databaseHost,
			DatabaseName: databaseName,
			RequireSSL:   databaseRequireSsl,
		},
		SuperUser: SuperUserSettings{
			Password: superuserPassword,
		},
		Auth: AuthSettings{
			AccessKey:  authAccessKey,
			RefreshKey: authRefreshKey,
		},
		AWS: AWSSettings{
			BucketName: awsBucketName,
			Id:         awsId,
			Secret:     awsSecret,
			Region:     awsRegion,
		},
		Resend: ResendSettings{
			APIKey: resendApiKey,
		},
		Calendar: CalendarSettings{
			MaxTerminationDate: calendarMaxTerminationDate,
		},
	}, nil
}

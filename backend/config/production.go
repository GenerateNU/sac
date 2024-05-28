package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	m "github.com/garrettladley/mattress"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type ProductionSettings struct {
	Database    ProductionDatabaseSettings    `yaml:"database"`
	Application ProductionApplicationSettings `yaml:"application"`
	Calendar    CalendarSettings              `yaml:"calendar"`
}

type ProductionDatabaseSettings struct {
	RequireSSL bool `yaml:"requiressl"`
}

type ProductionApplicationSettings struct {
	Port uint16 `yaml:"port"`
	Host string `yaml:"host"`
}

func readProd(v *viper.Viper) (*Settings, error) {
	var prodSettings ProductionSettings

	env := string(EnvironmentProduction)

	v.SetConfigName(env)

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read %s configuration: %w", env, err)
	}

	if err := v.Unmarshal(&prodSettings); err != nil {
		return nil, fmt.Errorf("failed to unmarshal configuration: %w", err)
	}

	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("failed to load .env: %w", err)
	}

	appPrefix := "APP_"
	applicationPrefix := fmt.Sprintf("%sAPPLICATION__", appPrefix)
	dbPrefix := fmt.Sprintf("%sDATABASE__", appPrefix)
	superUserPrefix := fmt.Sprintf("%sSUPERUSER__", appPrefix)
	authSecretPrefix := fmt.Sprintf("%sAUTHSECRET__", appPrefix)

	portStr := os.Getenv(fmt.Sprintf("%sPORT", appPrefix))
	portInt, err := strconv.ParseUint(portStr, 10, 16)
	if err != nil {
		return nil, fmt.Errorf("failed to parse port: %w", err)
	}

	dbPassword, err := m.NewSecret(os.Getenv(fmt.Sprintf("%sUSERNAME", dbPrefix)))
	if err != nil {
		return nil, errors.New("failed to create secret from database password")
	}

	superPassword, err := m.NewSecret(os.Getenv(fmt.Sprintf("%sPASSWORD", superUserPrefix)))
	if err != nil {
		return nil, errors.New("failed to create secret from super user password")
	}

	authAccessKey, err := m.NewSecret(os.Getenv(fmt.Sprintf("%sACCESS_TOKEN", authSecretPrefix)))
	if err != nil {
		return nil, errors.New("failed to create secret from access token")
	}

	authRefreshKey, err := m.NewSecret(os.Getenv(fmt.Sprintf("%sREFRESH_TOKEN", authSecretPrefix)))
	if err != nil {
		return nil, errors.New("failed to create secret from refresh token")
	}

	awsSettings, err := readAWSSettings()
	if err != nil {
		return nil, fmt.Errorf("failed to read AWS settings: %w", err)
	}

	resendSettings, err := readResendSettings()
	if err != nil {
		return nil, fmt.Errorf("failed to read Resend settings: %w", err)
	}

	googleSettings, err := readGoogleOAuthSettings()
	if err != nil {
		return nil, fmt.Errorf("failed to read Google OAuth settings: %w", err)
	}

	outlookSettings, err := readOutlookOAuthSettings()
	if err != nil {
		return nil, fmt.Errorf("failed to read Outlook OAuth settings: %w", err)
	}

	return &Settings{
		Application: ApplicationSettings{
			Port:    uint16(portInt),
			Host:    prodSettings.Application.Host,
			BaseUrl: os.Getenv(fmt.Sprintf("%sBASE_URL", applicationPrefix)),
		},
		Database: DatabaseSettings{
			Username:     os.Getenv(fmt.Sprintf("%sUSERNAME", dbPrefix)),
			Password:     dbPassword,
			Host:         os.Getenv(fmt.Sprintf("%sHOST", dbPrefix)),
			Port:         uint(portInt),
			DatabaseName: os.Getenv(fmt.Sprintf("%sDATABASE_NAME", dbPrefix)),
			RequireSSL:   prodSettings.Database.RequireSSL,
		},
		SuperUser: SuperUserSettings{
			Password: superPassword,
		},
		Auth: AuthSettings{
			AccessKey:  authAccessKey,
			RefreshKey: authRefreshKey,
		},
		Pinecone:        *pineconeSettings,
		OpenAI:          *openAISettings,
		AWS:             *awsSettings,
		Resend:          *resendSettings,
		Calendar:        prodSettings.Calendar,
		GoogleSettings:  *googleSettings,
		OutlookSettings: *outlookSettings,
	}, nil
}

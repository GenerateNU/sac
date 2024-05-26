package config

import (
	"time"

	m "github.com/garrettladley/mattress"
)

type Settings struct {
	Application ApplicationSettings `envPrefix:"SAC_APPLICATION_"`
	Database    DatabaseSettings    `envPrefix:"SAC_DB_"`
	SuperUser   SuperUserSettings   `envPrefix:"SAC_SUDO_"`
	Auth        AuthSettings        `envPrefix:"SAC_AUTH_"`
	AWS         AWSSettings         `envPrefix:"SAC_AWS_"`
	Resend      ResendSettings      `envPrefix:"SAC_RESEND_"`
	Calendar    CalendarSettings    `envPrefix:"SAC_CALENDAR_"`
}

type ApplicationSettings struct {
	Port    uint16 `env:"PORT" envDefault:"8080"`
	Host    string `env:"HOST" envDefault:"127.0.0.1"`
	BaseUrl string `env:"BASE_URL" envDefault:"http://127.0.0.1"`
}

type SuperUserSettings struct {
	Password *m.Secret[string] `env:"PASSWORD" envDefault:"Password#1"`
}

type AuthSettings struct {
	AccessKey  *m.Secret[string] `env:"ACCESS_KEY" envDefault:"g(r|##*?>Qp}h37+,T2"`
	RefreshKey *m.Secret[string] `env:"REFRESH_KEY" envDefault:"amk*2!gG}1i8D9RwJS$p"`
}

type AWSSettings struct {
	BucketName *m.Secret[string] `env:"BUCKET_NAME"`
	Id         *m.Secret[string] `env:"ID"`
	Secret     *m.Secret[string] `env:"SECRET"`
	Region     *m.Secret[string] `env:"REGION"`
}

type ResendSettings struct {
	APIKey *m.Secret[string] `env:"API_KEY"`
}

type CalendarSettings struct {
	MaxTerminationDate time.Time `env:"MAX_TERMINATION_DATE" envDefault:"12-31-2024"`
}

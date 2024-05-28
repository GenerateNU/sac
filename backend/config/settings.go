package config

type Settings struct {
	Application ApplicationSettings
	Database    DatabaseSettings
	SuperUser   SuperUserSettings
	Auth        AuthSettings
	AWS         AWSSettings
	Resend      ResendSettings
	Calendar    CalendarSettings
	Google      OAuthSettings
	Outlook     OAuthSettings
}

type IntermediateSettings struct {
	Application     ApplicationSettings           `envPrefix:"SAC_APPLICATION_"`
	Database        IntermediateDatabaseSettings  `envPrefix:"SAC_DB_"`
	SuperUser       IntermediateSuperUserSettings `envPrefix:"SAC_SUDO_"`
	Auth            IntermediateAuthSettings      `envPrefix:"SAC_AUTH_"`
	AWS             IntermediateAWSSettings       `envPrefix:"SAC_AWS_"`
	Resend          IntermediateResendSettings    `envPrefix:"SAC_RESEND_"`
	Calendar        IntermediateCalendarSettings  `envPrefix:"SAC_CALENDAR_"`
	GoogleSettings  GoogleOAuthSettings           `envPrefix:"SAC_GOOGLE_OAUTH"`
	OutlookSettings OutlookOAuthSettings          `envPrefix:"SAC_OUTLOOK_OAUTH"`
}

func (i *IntermediateSettings) Into() (*Settings, error) {
	database, err := i.Database.Into()
	if err != nil {
		return nil, nil
	}

	superUser, err := i.SuperUser.Into()
	if err != nil {
		return nil, nil
	}

	auth, err := i.Auth.Into()
	if err != nil {
		return nil, nil
	}

	aws, err := i.AWS.Into()
	if err != nil {
		return nil, nil
	}

	resend, err := i.Resend.Into()
	if err != nil {
		return nil, nil
	}

	calendar, err := i.Calendar.Into()
	if err != nil {
		return nil, nil
	}

	google, err := i.GoogleSettings.Into()
	if err != nil {
		return nil, nil
	}

	outlook, err := i.OutlookSettings.Into()
	if err != nil {
		return nil, nil
	}

	return &Settings{
		Application: i.Application,
		Database:    *database,
		SuperUser:   *superUser,
		Auth:        *auth,
		AWS:         *aws,
		Resend:      *resend,
		Calendar:    *calendar,
		Google:      *google,
		Outlook:     *outlook,
	}, nil
}

package config

type Settings struct {
	Application       ApplicationSettings
	Database          DatabaseSettings
	RedisActiveTokens RedisSettings
	RedisBlacklist    RedisSettings
	RedisLimiter      RedisSettings
	SuperUser         SuperUserSettings
	Auth              AuthSettings
	AWS               AWSSettings
	Resend            ResendSettings
	Calendar          CalendarSettings
	Google            OAuthSettings
	Outlook           OAuthSettings
}

type intermediateSettings struct {
	application       ApplicationSettings              `envPrefix:"SAC_APPLICATION_"`
	database          intermediateDatabaseSettings     `envPrefix:"SAC_DB_"`
	redisActiveTokens intermediateRedisSettings        `envPrefix:"SAC_REDIS_ACTIVE_TOKENS_"`
	redisBlacklist    intermediateRedisSettings        `envPrefix:"SAC_REDIS_BLACKLIST_"`
	redisLimiter      intermediateRedisSettings        `envPrefix:"SAC_REDIS_LIMITER_"`
	superUser         intermediateSuperUserSettings    `envPrefix:"SAC_SUDO_"`
	auth              intermediateAuthSettings         `envPrefix:"SAC_AUTH_"`
	aws               intermediateAWSSettings          `envPrefix:"SAC_AWS_"`
	resend            intermediateResendSettings       `envPrefix:"SAC_RESEND_"`
	calendar          intermediateCalendarSettings     `envPrefix:"SAC_CALENDAR_"`
	googleSettings    intermediateGoogleOAuthSettings  `envPrefix:"SAC_GOOGLE_OAUTH"`
	outlookSettings   intermdeiateOutlookOAuthSettings `envPrefix:"SAC_OUTLOOK_OAUTH"`
}

func (i *intermediateSettings) into() (*Settings, error) {
	database, err := i.database.into()
	if err != nil {
		return nil, err
	}

	redisActiveTokens, err := i.redisActiveTokens.into()
	if err != nil {
		return nil, err
	}

	redisBlacklist, err := i.redisBlacklist.into()
	if err != nil {
		return nil, err
	}

	redisLimiter, err := i.redisLimiter.into()
	if err != nil {
		return nil, err
	}

	superUser, err := i.superUser.into()
	if err != nil {
		return nil, err
	}

	auth, err := i.auth.into()
	if err != nil {
		return nil, err
	}

	aws, err := i.aws.into()
	if err != nil {
		return nil, err
	}

	resend, err := i.resend.into()
	if err != nil {
		return nil, err
	}

	calendar, err := i.calendar.into()
	if err != nil {
		return nil, err
	}

	google, err := i.googleSettings.into()
	if err != nil {
		return nil, err
	}

	outlook, err := i.outlookSettings.into()
	if err != nil {
		return nil, err
	}

	return &Settings{
		Application:       i.application,
		Database:          *database,
		RedisActiveTokens: *redisActiveTokens,
		RedisBlacklist:    *redisBlacklist,
		RedisLimiter:      *redisLimiter,
		SuperUser:         *superUser,
		Auth:              *auth,
		AWS:               *aws,
		Resend:            *resend,
		Calendar:          *calendar,
		Google:            *google,
		Outlook:           *outlook,
	}, nil
}

package config

type Settings struct {
	Application       ApplicationSettings
	Database          DatabaseSettings
	RedisActiveTokens RedisSettings
	RedisBlacklist    RedisSettings
	RedisLimiter      RedisSettings
	SuperUser         SuperUserSettings
	Auth              AuthSettings
	Calendar          CalendarSettings
	Integrations
}

type Integrations struct {
	GoogleOauth  OAuthSettings
	OutlookOauth OAuthSettings
	AWS          AWSSettings
	Resend       ResendSettings
	Search       SearchSettings
}

type intermediateSettings struct {
	Application       ApplicationSettings              `envPrefix:"SAC_APPLICATION_"`
	Database          intermediateDatabaseSettings     `envPrefix:"SAC_DB_"`
	RedisActiveTokens intermediateRedisSettings        `envPrefix:"SAC_REDIS_ACTIVE_TOKENS_"`
	RedisBlacklist    intermediateRedisSettings        `envPrefix:"SAC_REDIS_BLACKLIST_"`
	RedisLimiter      intermediateRedisSettings        `envPrefix:"SAC_REDIS_LIMITER_"`
	SuperUser         intermediateSuperUserSettings    `envPrefix:"SAC_SUDO_"`
	Auth              intermediateAuthSettings         `envPrefix:"SAC_AUTH_"`
	AWS               intermediateAWSSettings          `envPrefix:"SAC_AWS_"`
	Resend            intermediateResendSettings       `envPrefix:"SAC_RESEND_"`
	Calendar          intermediateCalendarSettings     `envPrefix:"SAC_CALENDAR_"`
	GoogleSettings    intermediateGoogleOAuthSettings  `envPrefix:"SAC_GOOGLE_OAUTH_"`
	OutlookSettings   intermdeiateOutlookOAuthSettings `envPrefix:"SAC_OUTLOOK_OAUTH_"`
	Search            SearchSettings                   `envPrefix:"SAC_SEARCH_"`
}

func (i *intermediateSettings) into() (*Settings, error) {
	database, err := i.Database.into()
	if err != nil {
		return nil, err
	}

	redisActiveTokens, err := i.RedisActiveTokens.into()
	if err != nil {
		return nil, err
	}

	redisBlacklist, err := i.RedisBlacklist.into()
	if err != nil {
		return nil, err
	}

	redisLimiter, err := i.RedisLimiter.into()
	if err != nil {
		return nil, err
	}

	superUser, err := i.SuperUser.into()
	if err != nil {
		return nil, err
	}

	auth, err := i.Auth.into()
	if err != nil {
		return nil, err
	}

	aws, err := i.AWS.into()
	if err != nil {
		return nil, err
	}

	resend, err := i.Resend.into()
	if err != nil {
		return nil, err
	}

	calendar, err := i.Calendar.into()
	if err != nil {
		return nil, err
	}

	google, err := i.GoogleSettings.into()
	if err != nil {
		return nil, err
	}

	outlook, err := i.OutlookSettings.into()
	if err != nil {
		return nil, err
	}

	return &Settings{
		Application:       i.Application,
		Database:          *database,
		RedisActiveTokens: *redisActiveTokens,
		RedisBlacklist:    *redisBlacklist,
		RedisLimiter:      *redisLimiter,
		SuperUser:         *superUser,
		Auth:              *auth,
		Calendar:          *calendar,
		Integrations: Integrations{
			GoogleOauth:  *google,
			OutlookOauth: *outlook,
			AWS:          *aws,
			Resend:       *resend,
			Search:       i.Search,
		},
	}, nil
}

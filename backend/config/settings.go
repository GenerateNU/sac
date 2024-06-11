package config

type Settings struct {
	Application  ApplicationSettings
	Database     DatabaseSettings
	DBCache      RedisSettings
	Session      SessionSettings
	RedisLimiter RedisSettings
	SuperUser    SuperUserSettings
	Calendar     CalendarSettings
	Integrations
}

type Integrations struct {
	Google   GoogleOAuthSettings
	Microsft MicrosoftOAuthSettings
	AWS      AWSSettings
	Resend   ResendSettings
	Search   SearchSettings
}

type intermediateSettings struct {
	Application  ApplicationSettings               `envPrefix:"SAC_APPLICATION_"`
	Database     intermediateDatabaseSettings      `envPrefix:"SAC_DB_"`
	DBCache      intermediateRedisSettings         `envPrefix:"SAC_REDIS_DB_CACHE_"`
	RedisSession intermediateRedisSettings         `envPrefix:"SAC_REDIS_SESSION_"`
	Session      intermediateSessionSettings       `envPrefix:"SAC_SESSION_"`
	RedisLimiter intermediateRedisSettings         `envPrefix:"SAC_REDIS_LIMITER_"`
	SuperUser    intermediateSuperUserSettings     `envPrefix:"SAC_SUDO_"`
	AWS          intermediateAWSSettings           `envPrefix:"SAC_AWS_"`
	Resend       intermediateResendSettings        `envPrefix:"SAC_RESEND_"`
	Calendar     intermediateCalendarSettings      `envPrefix:"SAC_CALENDAR_"`
	Google       intermediateGoogleOAuthSettings   `envPrefix:"SAC_GOOGLE_OAUTH_"`
	Microsft     intermediateMicrosoftOAuthSetting `envPrefix:"SAC_MICROSOFT_OAUTH_"`
	Search       SearchSettings                    `envPrefix:"SAC_SEARCH_"`
}

func (i *intermediateSettings) into() (*Settings, error) {
	database, err := i.Database.into()
	if err != nil {
		return nil, err
	}

	dbCache, err := i.DBCache.into()
	if err != nil {
		return nil, err
	}

	redisSession, err := i.RedisSession.into()
	if err != nil {
		return nil, err
	}

	session, err := i.Session.into(*redisSession)
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

	google, err := i.Google.into()
	if err != nil {
		return nil, err
	}

	microsoft, err := i.Microsft.into()
	if err != nil {
		return nil, err
	}

	return &Settings{
		Application:  i.Application,
		DBCache:      *dbCache,
		Database:     *database,
		Session:      *session,
		RedisLimiter: *redisLimiter,
		SuperUser:    *superUser,
		Calendar:     *calendar,
		Integrations: Integrations{
			Google:   *google,
			Microsft: *microsoft,
			AWS:      *aws,
			Resend:   *resend,
			Search:   i.Search,
		},
	}, nil
}

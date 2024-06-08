package config

type GoogleOAuthSettings struct {
	OAuthSettings
}

type intermediateGoogleOAuthSettings struct {
	intermediateOAuthSettings
}

func (i *intermediateGoogleOAuthSettings) into() (*GoogleOAuthSettings, error) {
	res, err := i.intermediateOAuthSettings.into()
	if err != nil {
		return nil, err
	}

	return &GoogleOAuthSettings{*res}, nil
}

package config

const (
	tenantID = "a8eec281-aaa3-4dae-ac9b-9a398b9215e7"
)

type MicrosoftOAuthSettings struct {
	OAuthSettings
	Tenant string
}

type intermediateMicrosoftOAuthSetting struct {
	intermediateOAuthSettings
}

func (i *intermediateMicrosoftOAuthSetting) into() (*MicrosoftOAuthSettings, error) {
	res, err := i.intermediateOAuthSettings.into()
	if err != nil {
		return nil, err
	}

	return &MicrosoftOAuthSettings{
		OAuthSettings: *res,
		Tenant:        tenantID,
	}, nil
}

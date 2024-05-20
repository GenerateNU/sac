package oauth

type GoogleExchangeCodeRequestBody struct {
	Code          string `json:"code" validate:"required"`
	ClientId      string `json:"client_id" validate:"required"`
	ClientSecret  string `json:"client_secret" validate:"required"`
	RedirectUri   string `json:"redirect_uri" validate:"required"`
	GrantType     string `json:"grant_type" validate:"required"`
}

type OAuthToken struct {
	AccessToken 	string 	`json:"access_token" validate:"required"`
	ExpiresIn   	int 	`json:"expires_in" validate:"required"`
	RefreshToken 	string 	`json:"refresh_token" validate:"required"`
	Scope			string  `json:"scope" validate:"required"`
	TokenType		string  `json:"token_type" validate:"required"`
}

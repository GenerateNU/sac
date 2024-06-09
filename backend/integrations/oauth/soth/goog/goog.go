package goog

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	go_json "github.com/goccy/go-json"

	"github.com/GenerateNU/sac/backend/integrations/oauth/soth"
	"github.com/GenerateNU/sac/backend/utilities"

	"golang.org/x/oauth2"

	m "github.com/garrettladley/mattress"
	goog "golang.org/x/oauth2/google"
)

const (
	endpointProfile string = "https://www.googleapis.com/oauth2/v2/userinfo"
)

var (
	endpoint      oauth2.Endpoint = goog.Endpoint
	defaultScopes []string        = []string{"email", "https://www.googleapis.com/auth/calendar.events", "https://www.googleapis.com/auth/calendar.readonly"}
)

// New creates a new Google provider, and sets up important connection details.
// You should always call `google.New` to get a new Provider. Never try to create
// one manually.
func New(clientKey *m.Secret[string], secret *m.Secret[string], callbackURL string, scopes ...string) *Provider {
	p := &Provider{
		ClientKey:    clientKey,
		Secret:       secret,
		CallbackURL:  callbackURL,
		providerName: "google",

		// We can get a refresh token from Google by this option.
		// See https://developers.google.com/identity/protocols/oauth2/openid-connect#access-type-param
		authCodeOptions: []oauth2.AuthCodeOption{
			oauth2.AccessTypeOffline,
		},
	}
	p.config = newConfig(p, scopes)
	return p
}

// Provider is the implementation of `oauth.Provider` for accessing Google.
type Provider struct {
	ClientKey       *m.Secret[string]
	Secret          *m.Secret[string]
	CallbackURL     string
	config          *oauth2.Config
	authCodeOptions []oauth2.AuthCodeOption
	providerName    string
}

// Name is the name used to retrieve this provider later.
func (p *Provider) Name() string {
	return p.providerName
}

// SetName is to update the name of the provider (needed in case of multiple providers of 1 type)
func (p *Provider) SetName(name string) {
	p.providerName = name
}

// Debug is a no-op for the goog package.
func (p *Provider) Debug(debug bool) {}

// BeginAuth asks Google for an authentication endpoint.
func (p *Provider) BeginAuth(state string) (soth.Session, error) {
	url := p.config.AuthCodeURL(state, p.authCodeOptions...)
	session := &Session{
		AuthURL: url,
	}
	return session, nil
}

type googleUser struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	FirstName string `json:"given_name"`
	LastName  string `json:"family_name"`
	Link      string `json:"link"`
	Picture   string `json:"picture"`
}

// FetchUser will go to Google and access basic information about the user.
func (p *Provider) FetchUser(session soth.Session) (soth.User, error) {
	sess := session.(*Session)
	user := soth.User{
		AccessToken:  sess.AccessToken,
		Provider:     p.Name(),
		RefreshToken: sess.RefreshToken,
		ExpiresAt:    sess.ExpiresAt,
		IDToken:      sess.IDToken,
	}

	if user.AccessToken == "" {
		// Data is not yet retrieved, since accessToken is still empty.
		return user, fmt.Errorf("%s cannot get user information without accessToken", p.providerName)
	}

	resp, err := utilities.Request(http.MethodGet, fmt.Sprintf("%s?access_token=%s", endpointProfile, url.QueryEscape(sess.AccessToken)), nil)
	if err != nil {
		return user, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return user, fmt.Errorf("%s responded with a %d trying to fetch user information", p.providerName, resp.StatusCode)
	}

	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return user, err
	}

	var u googleUser
	if err := go_json.Unmarshal(responseBytes, &u); err != nil {
		return user, err
	}

	// Extract the user data we got from Google into our oauth.User.
	user.Name = u.Name
	user.FirstName = u.FirstName
	user.LastName = u.LastName
	user.NickName = u.Name
	user.Email = u.Email
	user.AvatarURL = u.Picture
	user.UserID = u.ID
	// Google provides other useful fields such as 'hd'; get them from RawData
	if err := go_json.Unmarshal(responseBytes, &user.RawData); err != nil {
		return user, err
	}

	return user, nil
}

func newConfig(provider *Provider, scopes []string) *oauth2.Config {
	c := &oauth2.Config{
		ClientID:     provider.ClientKey.Expose(),
		ClientSecret: provider.Secret.Expose(),
		RedirectURL:  provider.CallbackURL,
		Endpoint:     endpoint,
		Scopes:       []string{},
	}

	if len(scopes) > 0 {
		c.Scopes = append(c.Scopes, scopes...)
	} else {
		c.Scopes = append(c.Scopes, defaultScopes...)
	}
	return c
}

// RefreshTokenAvailable refresh token is provided by auth provider or not
func (p *Provider) RefreshTokenAvailable() bool {
	return true
}

// RefreshToken get new access token based on the refresh token
func (p *Provider) RefreshToken(refreshToken string) (*oauth2.Token, error) {
	token := &oauth2.Token{RefreshToken: refreshToken}
	ts := p.config.TokenSource(soth.ContextForClient(utilities.Client()), token)
	newToken, err := ts.Token()
	if err != nil {
		return nil, err
	}
	return newToken, err
}

// SetPrompt sets the prompt values for the google OAuth call. Use this to
// force users to choose and account every time by passing "select_account",
// for example.
// See https://developers.google.com/identity/protocols/OpenIDConnect#authenticationuriparameters
func (p *Provider) SetPrompt(prompt ...string) {
	if len(prompt) == 0 {
		return
	}
	p.authCodeOptions = append(p.authCodeOptions, oauth2.SetAuthURLParam("prompt", strings.Join(prompt, " ")))
}

// SetHostedDomain sets the hd parameter for google OAuth call.
// Use this to force user to pick user from specific hosted domain.
// See https://developers.google.com/identity/protocols/oauth2/openid-connect#hd-param
func (p *Provider) SetHostedDomain(hd string) {
	if hd == "" {
		return
	}
	p.authCodeOptions = append(p.authCodeOptions, oauth2.SetAuthURLParam("hd", hd))
}

// SetLoginHint sets the login_hint parameter for the Google OAuth call.
// Use this to prompt the user to log in with a specific account.
// See https://developers.google.com/identity/protocols/oauth2/openid-connect#login-hint
func (p *Provider) SetLoginHint(loginHint string) {
	if loginHint == "" {
		return
	}
	p.authCodeOptions = append(p.authCodeOptions, oauth2.SetAuthURLParam("login_hint", loginHint))
}

// SetAccessType sets the access_type parameter for the Google OAuth call.
// If an access token is being requested, the client does not receive a refresh token unless a value of offline is specified.
// See https://developers.google.com/identity/protocols/oauth2/openid-connect#access-type-param
func (p *Provider) SetAccessType(at string) {
	if at == "" {
		return
	}
	p.authCodeOptions = append(p.authCodeOptions, oauth2.SetAuthURLParam("access_type", at))
}

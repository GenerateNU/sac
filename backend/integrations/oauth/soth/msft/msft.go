package msft

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"

	go_json "github.com/goccy/go-json"

	"github.com/GenerateNU/sac/backend/integrations/oauth/soth"
	"github.com/GenerateNU/sac/backend/utilities"

	"github.com/markbates/going/defaults"
	"golang.org/x/oauth2"

	m "github.com/garrettladley/mattress"
)

const (
	// #nosec G101
	authURLFmt string = "https://login.microsoftonline.com/%s/oauth2/v2.0/authorize"
	// #nosec G101
	tokenURLFmt     string = "https://login.microsoftonline.com/%s/oauth2/v2.0/token"
	endpointProfile string = "https://graph.microsoft.com/v1.0/me"
)

var defaultScopes = []string{"openid", "offline_access", "user.read", "calendars.readwrite", "email", "profile"}

// New creates a new microsoftonline Provider, and sets up important connection details.
// You should always call `msft.New` to get a new Provider. Never try to create
// one manually.
func New(clientKey *m.Secret[string], callbackURL string, tenant string, scopes ...string) *Provider {
	p := &Provider{
		ClientKey:    clientKey,
		CallbackURL:  callbackURL,
		ProviderName: "microsoftonline",
		tenant:       tenant,
	}

	p.config = newConfig(p, scopes)
	return p
}

// Provider is the implementation of `soth.Provider` for accessing microsoftonline.
type Provider struct {
	ClientKey    *m.Secret[string]
	CallbackURL  string
	config       *oauth2.Config
	ProviderName string
	tenant       string
}

// Name is the name used to retrieve this Provider later.
func (p *Provider) Name() string {
	return p.ProviderName
}

// SetName is to update the name of the Provider (needed in case of multiple Providers of 1 type)
func (p *Provider) SetName(name string) {
	p.ProviderName = name
}

// Debug is a no-op for the msft package.
func (p *Provider) Debug(debug bool) {}

// BeginAuth asks MicrosoftOnline for an authentication end-point.
func (p *Provider) BeginAuth(state string) (soth.Session, error) {
	return &Session{
		AuthURL: p.config.AuthCodeURL(state),
	}, nil
}

// FetchUser will go to MicrosoftOnline and access basic information about the user.
func (p *Provider) FetchUser(session soth.Session) (soth.User, error) {
	msSession := session.(*Session)
	user := soth.User{
		AccessToken: msSession.AccessToken,
		Provider:    p.Name(),
		ExpiresAt:   msSession.ExpiresAt,
	}

	if user.AccessToken == "" {
		return user, fmt.Errorf("%s cannot get user information without accessToken", p.ProviderName)
	}

	resp, err := utilities.Request(http.MethodGet, endpointProfile, nil, utilities.Authorization(msSession.AccessToken))
	if err != nil {
		return user, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return user, fmt.Errorf("%s responded with a %d trying to fetch user information", p.ProviderName, resp.StatusCode)
	}

	user.AccessToken = msSession.AccessToken

	err = userFromReader(resp.Body, &user)
	return user, err
}

// RefreshTokenAvailable refresh token is provided by auth Provider or not
// available for microsoft online as session size hit the limit of max cookie size
func (p *Provider) RefreshTokenAvailable() bool {
	return false
}

// RefreshToken get new access token based on the refresh token
func (p *Provider) RefreshToken(refreshToken string) (*oauth2.Token, error) {
	if refreshToken == "" {
		return nil, errors.New("no refresh token provided")
	}

	token := &oauth2.Token{RefreshToken: refreshToken}
	ts := p.config.TokenSource(soth.ContextForClient(utilities.Client()), token)
	newToken, err := ts.Token()
	if err != nil {
		return nil, err
	}
	return newToken, err
}

func newConfig(provider *Provider, scopes []string) *oauth2.Config {
	var (
		authURL  string
		tokenURL string
	)
	if provider.tenant == "" {
		authURL = fmt.Sprintf(authURLFmt, "common")
		tokenURL = fmt.Sprintf(tokenURLFmt, "common")
	} else {
		authURL = fmt.Sprintf(authURLFmt, provider.tenant)
		tokenURL = fmt.Sprintf(tokenURLFmt, provider.tenant)
	}

	c := &oauth2.Config{
		ClientID:    provider.ClientKey.Expose(),
		RedirectURL: provider.CallbackURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authURL,
			TokenURL: tokenURL,
		},
		Scopes: []string{},
	}

	c.Scopes = append(c.Scopes, scopes...)
	if len(scopes) == 0 {
		c.Scopes = append(c.Scopes, defaultScopes...)
	}

	return c
}

func userFromReader(r io.Reader, user *soth.User) error {
	buf := &bytes.Buffer{}
	tee := io.TeeReader(r, buf)

	u := struct {
		ID                string `json:"id"`
		Name              string `json:"displayName"`
		Email             string `json:"mail"`
		FirstName         string `json:"givenName"`
		LastName          string `json:"surname"`
		UserPrincipalName string `json:"userPrincipalName"`
	}{}

	if err := go_json.NewDecoder(tee).Decode(&u); err != nil {
		return err
	}

	raw := map[string]interface{}{}
	if err := go_json.NewDecoder(buf).Decode(&raw); err != nil {
		return err
	}

	user.UserID = u.ID
	user.Email = defaults.String(u.Email, u.UserPrincipalName)
	user.Name = u.Name
	user.NickName = u.Name
	user.FirstName = u.FirstName
	user.LastName = u.LastName
	user.RawData = raw

	return nil
}

package sacic

import (
	"bytes"
	"compress/gzip"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

	"github.com/GenerateNU/sac/backend/integrations/oauth"
	"github.com/GenerateNU/sac/backend/utilities"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	m "github.com/garrettladley/mattress"
)

// SessionName is the key used to access the session store.
const SessionName = "_sacic_session"

// Store can/should be set by applications using sacic. The default is a cookie store.
var (
	Store *session.Store
)

type key int

// ProviderParamKey can be used as a key in context when passing in a provider
const ProviderParamKey key = iota

// Init initializes the sacic package with the session secret and store.
// Must be called before using any of the other functions in this package.
func Init(sessionSecret *m.Secret[string], store fiber.Storage) {
	Store = session.New(session.Config{
		Storage:      store,
		KeyLookup:    fmt.Sprintf("header:%s", SessionName),
		KeyGenerator: func() string { return sessionSecret.Expose() },
	})
}

/*
BeginAuthHandler is a convenience handler for starting the authentication process.
It expects to be able to get the name of the provider from the query parameters
as either "provider" or ":provider".

BeginAuthHandler will redirect the user to the appropriate authentication end-point
for the requested provider.
*/
func BeginAuthHandler(c *fiber.Ctx) error {
	url, err := GetAuthURL(c)
	if err != nil {
		slog.Error("sacic: error getting auth URL", "error", err)
		return utilities.BadRequest(errors.New("error getting auth URL"))
	}

	return c.Redirect(url, http.StatusTemporaryRedirect)
}

// SetState sets the state string associated with the given request.
// If no state string is associated with the request, one will be generated.
// This state is sent to the provider and can be retrieved during the
// callback.
func SetState(c *fiber.Ctx) string {
	state := c.Query("state")
	if len(state) > 0 {
		return state
	}

	// If a state query param is not passed in, generate a random
	// base64-encoded nonce so that the state on the auth URL
	// is unguessable, preventing CSRF attacks, as described in
	//
	// https://auth0.com/docs/protocols/oauth2/oauth-state#keep-reading
	nonceBytes := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, nonceBytes)
	if err != nil {
		panic("sacic: source of randomness unavailable: " + err.Error())
	}
	return base64.URLEncoding.EncodeToString(nonceBytes)
}

// GetState gets the state returned by the provider during the callback.
// This is used to prevent CSRF attacks, see
// http://tools.ietf.org/html/rfc6749#section-10.12
func GetState(c *fiber.Ctx) string {
	if c.Method() == fiber.MethodPost {
		return c.FormValue("state")
	}
	return c.Query("state")
}

/*
GetAuthURL starts the authentication process with the requested provided.
It will return a URL that should be used to send users to.

It expects to be able to get the name of the provider from the query parameters
as either "provider" or ":provider".

I would recommend using the BeginAuthHandler instead of doing all of these steps
yourself, but that's entirely up to you.
*/
func GetAuthURL(c *fiber.Ctx) (string, error) {
	providerName, err := GetProviderName(c)
	if err != nil {
		return "", err
	}

	provider, err := oauth.GetProvider(providerName)
	if err != nil {
		return "", err
	}
	sess, err := provider.BeginAuth(SetState(c))
	if err != nil {
		return "", err
	}

	url, err := sess.GetAuthURL()
	if err != nil {
		return "", err
	}

	err = StoreInSession(providerName, sess.Marshal(), c)
	if err != nil {
		return "", err
	}

	return url, err
}

/*
CompleteUserAuth does what it says on the tin. It completes the authentication
process and fetches all the basic information about the user from the provider.

It expects to be able to get the name of the provider from the query parameters
as either "provider" or ":provider".
*/
func CompleteUserAuth(c *fiber.Ctx) (oauth.User, error) {
	providerName, err := GetProviderName(c)
	if err != nil {
		return oauth.User{}, err
	}

	provider, err := oauth.GetProvider(providerName)
	if err != nil {
		return oauth.User{}, err
	}

	value, err := GetFromSession(providerName, c)
	if err != nil {
		return oauth.User{}, err
	}
	defer Logout(c)
	sess, err := provider.UnmarshalSession(value)
	if err != nil {
		return oauth.User{}, err
	}

	err = validateState(c, sess)
	if err != nil {
		return oauth.User{}, err
	}

	user, err := provider.FetchUser(sess)
	if err == nil {
		// user can be found with existing session data
		return user, err
	}

	paramsMap := c.Queries()

	params := url.Values{}
	for key, value := range paramsMap {
		params.Add(key, value)
	}

	// get new token and retry fetch
	_, err = sess.Authorize(provider, params)
	if err != nil {
		return oauth.User{}, err
	}

	err = StoreInSession(providerName, sess.Marshal(), c)
	if err != nil {
		return oauth.User{}, err
	}

	outhu, err := provider.FetchUser(sess)
	return outhu, err
}

// validateState ensures that the state token param from the original
// AuthURL matches the one included in the current (callback) request.
func validateState(c *fiber.Ctx, sess oauth.Session) error {
	rawAuthURL, err := sess.GetAuthURL()
	if err != nil {
		return err
	}

	authURL, err := url.Parse(rawAuthURL)
	if err != nil {
		return err
	}

	reqState := GetState(c)

	originalState := authURL.Query().Get("state")
	if originalState != "" && (originalState != reqState) {
		return errors.New("state token mismatch")
	}
	return nil
}

// Logout invalidates a user session.
func Logout(c *fiber.Ctx) error {
	sess, err := Store.Get(c)
	if err != nil {
		return err
	}
	sess.Destroy()
	return nil
}

// GetProviderName is a function used to get the name of a provider
// for a given request. By default, this provider is fetched from
// the URL query string. If you provide it in a different way,
// assign your own function to this variable that returns the provider
// name for your request.
func GetProviderName(c *fiber.Ctx) (string, error) {
	// try to get it from the url param ":provider"
	if p := c.Params("provider"); p != "" {
		return p, nil
	}

	// try to get it from the context's value of "provider" key
	if p, ok := c.Locals("provider").(string); ok {
		return p, nil
	}

	// try to get it from the go-context's value of providerContextKey key
	if p, ok := c.Context().Value(ProviderParamKey).(string); ok {
		return p, nil
	}

	// As a fallback, loop over the used providers, if we already have a valid session for any provider (ie. user has already begun authentication with a provider), then return that provider name
	providers := oauth.GetProviders()
	sess, _ := Store.Get(c)
	for _, provider := range providers {
		p := provider.Name()
		value := sess.Get(p)
		if _, ok := value.(string); ok {
			return p, nil
		}
	}

	// if not found then return an empty string with the corresponding error
	return "", errors.New("you must select a provider")
}

// GetContextWithProvider returns a new request context containing the provider
func GetContextWithProvider(c *fiber.Ctx, provider string) *fiber.Ctx {
	c.Locals("provider", provider)
	return c
}

// StoreInSession stores a specified key/value pair in the session.
func StoreInSession(key string, value string, c *fiber.Ctx) error {
	sess, err := Store.Get(c)
	if err != nil {
		return err
	}

	if err := updateSessionValue(sess, key, value); err != nil {
		return err
	}

	sess.Save()
	return nil
}

// GetFromSession retrieves a previously-stored value from the session.
// If no value has previously been stored at the specified key, it will return an error.
func GetFromSession(key string, c *fiber.Ctx) (string, error) {
	sess, err := Store.Get(c)
	if err != nil {
		return "", errors.New("could not find a matching session for this request")
	}
	value, err := getSessionValue(sess, key)
	if err != nil {
		return "", errors.New("could not find a matching session for this request")
	}

	return value, nil
}

func getSessionValue(sess *session.Session, key string) (string, error) {
	value := sess.Get(key)
	if value == nil {
		return "", errors.New("could not find a matching session for this request")
	}

	rdata := strings.NewReader(value.(string))
	r, err := gzip.NewReader(rdata)
	if err != nil {
		return "", err
	}
	s, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(s), nil
}

func updateSessionValue(sess *session.Session, key, value string) error {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte(value)); err != nil {
		return err
	}
	if err := gz.Flush(); err != nil {
		return err
	}
	if err := gz.Close(); err != nil {
		return err
	}

	sess.Set(key, b.String())
	return nil
}

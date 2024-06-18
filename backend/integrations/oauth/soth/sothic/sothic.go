package sothic

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

	"github.com/GenerateNU/sac/backend/config"
	"github.com/GenerateNU/sac/backend/database/store"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/GenerateNU/sac/backend/integrations/oauth/crypt"
	"github.com/GenerateNU/sac/backend/integrations/oauth/soth"
)

type key int

const (
	SessionName string = "_sac_session"

	// ProviderParamKey can be used as a key in context when passing in a provider
	ProviderParamKey key = iota
)

var (
	// Session can/should be set by applications using sothic.
	SessionStore *session.Store

	encrypter func(string) (string, error)
	decrypter func(string) (string, error)
)

// MUST be called before using the package
func Init(sessionSettings config.SessionSettings) {
	config := session.Config{
		Storage:   store.NewRedisClient(sessionSettings.Redis),
		KeyLookup: fmt.Sprintf("header:%s", SessionName),
	}

	encrypter = func(s string) (string, error) {
		return crypt.Encrypt(s, sessionSettings.PassPhrase.Expose())
	}

	decrypter = func(s string) (string, error) {
		return crypt.Decrypt(s, sessionSettings.PassPhrase.Expose())
	}

	SessionStore = session.New(config)
}

/*
BeginAuthHandler is a convenience handler for starting the authentication process.

It expects to be able to get the name of the provider from the path parameter
":provider" or as set by SetProvider.

BeginAuthHandler will redirect the user to the appropriate authentication end-point
for the requested provider.

See https://github.com/markbates/goth/examples/main.go to see this in action.
*/
func BeginAuthHandler(c *fiber.Ctx) error {
	url, err := GetAuthURL(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	c.Set("redirect", url)
	return c.SendStatus(http.StatusOK)
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
		panic(fmt.Sprintf("sothic: source of randomness unavailable: %v", err.Error()))
	}
	return base64.URLEncoding.EncodeToString(nonceBytes)
}

// GetState gets the state returned by the provider during the callback.
// This is used to prevent CSRF attacks, see
// http://tools.ietf.org/html/rfc6749#section-10.12
func GetState(c *fiber.Ctx) string {
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
	if SessionStore == nil {
		return "", errors.New("session store unexpectedly nil")
	}

	providerName, err := GetProviderName(c)
	if err != nil {
		return "", err
	}

	provider, err := soth.GetProvider(providerName)
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

	if err := StoreInSession(providerName, sess.Marshal(), c); err != nil {
		return "", err
	}

	return url, err
}

/*
CompleteUserAuth does what it says on the tin. It completes the authentication
process and fetches all of the basic information about the user from the provider.

It expects to be able to get the name of the provider from the path parameter
":provider" or as set by SetProvider.

This method automatically ends the session. You can prevent this behavior by
passing in options. Please note that any options provided in addition to the
first will be ignored.

See https://github.com/markbates/goth/examples/main.go to see this in action.
*/
func CompleteUserAuth(c *fiber.Ctx) (soth.User, error) {
	if SessionStore == nil {
		return soth.User{}, errors.New("session store unexpectedly nil")
	}

	providerName, err := GetProviderName(c)
	if err != nil {
		return soth.User{}, err
	}

	provider, err := soth.GetProvider(providerName)
	if err != nil {
		return soth.User{}, err
	}

	value, err := GetFromSession(providerName, c)
	if err != nil {
		return soth.User{}, err
	}

	sess, err := provider.UnmarshalSession(value)
	if err != nil {
		return soth.User{}, err
	}

	err = validateState(c, sess)
	if err != nil {
		return soth.User{}, err
	}

	user, err := provider.FetchUser(sess)
	if err == nil {
		// user can be found with existing session data
		return user, err
	}

	// get new token and retry fetch
	_, err = sess.Authorize(provider, &Params{c: c})
	if err != nil {
		return soth.User{}, err
	}

	if err := StoreInSession(providerName, sess.Marshal(), c); err != nil {
		return soth.User{}, err
	}

	gu, err := provider.FetchUser(sess)

	return gu, err
}

// validateState ensures that the state token param from the original
// AuthURL matches the one included in the current (callback) request.
func validateState(c *fiber.Ctx, sess soth.Session) error {
	rawAuthURL, err := sess.GetAuthURL()
	if err != nil {
		return err
	}

	authURL, err := url.Parse(rawAuthURL)
	if err != nil {
		return err
	}

	originalState := authURL.Query().Get("state")

	if originalState != "" && (originalState != c.Query("state")) {
		return errors.New("state token mismatch")
	}
	return nil
}

// Logout invalidates a user session.
func Logout(c *fiber.Ctx) error {
	session, err := SessionStore.Get(c)
	if err != nil {
		return err
	}

	if err := session.Destroy(); err != nil {
		return err
	}

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

	// try to get it from the Fasthttp context's value of providerContextKey key
	if p := c.Locals(fmt.Sprint(ProviderParamKey)); p != "" {
		provider, ok := p.(string)
		if ok && provider != "" {
			return provider, nil
		}
	}

	// As a fallback, loop over the used providers, if we already have a valid session for any provider (ie. user has already begun authentication with a provider), then return that provider name
	providers := soth.GetProviders()
	session, err := SessionStore.Get(c)
	if err != nil {
		return "", err
	}

	for _, provider := range providers {
		p := provider.Name()
		value := session.Get(p)
		if _, ok := value.(string); ok {
			return p, nil
		}
	}

	// if not found then return an empty string with the corresponding error
	return "", errors.New("you must select a provider")
}

func SetProvider(c *fiber.Ctx, provider string) {
	c.Locals(fmt.Sprint(ProviderParamKey), provider)
}

// StoreInSession stores a specified key/value pair in the session.
func StoreInSession(key string, value string, c *fiber.Ctx) error {
	session, err := SessionStore.Get(c)
	if err != nil {
		slog.Error("error getting session", "error", err)
		return err
	}

	if err := updateSessionValue(session, key, value); err != nil {
		return err
	}

	// saved here
	return session.Save()
}

// GetFromSession retrieves a previously-stored value from the session.
// If no value has previously been stored at the specified key, it will return an error.
func GetFromSession(key string, c *fiber.Ctx) (string, error) {
	session, err := SessionStore.Get(c)
	if err != nil {
		return "", err
	}

	value, err := getSessionValue(session, key)
	if err != nil {
		slog.Error("error getting session value", "error", err)
		return "", errors.New("could not find a matching session for this request")
	}

	return value, nil
}

func getSessionValue(store *session.Session, key string) (string, error) {
	value := store.Get(key)
	if value == nil {
		return "", fmt.Errorf("key %s not found in session", key)
	}

	decrypted, err := decrypter(value.(string))
	if err != nil {
		return "", err
	}

	r, err := gzip.NewReader(strings.NewReader(decrypted))
	if err != nil {
		return "", err
	}
	s, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(s), nil
}

func updateSessionValue(session *session.Session, key, value string) error {
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
	encrypted, err := encrypter(b.String())
	if err != nil {
		return err
	}
	session.Set(key, encrypted)
	return nil
}

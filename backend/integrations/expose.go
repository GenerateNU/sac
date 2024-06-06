package integrations

import (
	"github.com/GenerateNU/sac/backend/integrations/email"
	"github.com/GenerateNU/sac/backend/integrations/file"
	"github.com/GenerateNU/sac/backend/integrations/oauth"
)

type Integrations struct {
	Email email.EmailClientInterface
	File  file.FileClientInterface
	OAuth oauth.OAuthProviders
}

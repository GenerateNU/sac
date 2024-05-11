package integrations

import (
	"github.com/GenerateNU/sac/backend/integrations/email"
	"github.com/GenerateNU/sac/backend/integrations/file"
)

type Integrations struct {
	Email email.EmailClientInterface
	File  file.FileClientInterface
}

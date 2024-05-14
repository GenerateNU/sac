package integrations

import (
	"github.com/GenerateNU/sac/backend/integrations/email"
	"github.com/GenerateNU/sac/backend/integrations/file"
	"github.com/GenerateNU/sac/backend/integrations/search"
)

type Integrations struct {
	Search search.SearchClientInterface
	Email  email.EmailClientInterface
	File   file.FileClientInterface
}

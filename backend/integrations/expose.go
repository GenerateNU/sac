package integrations

import (
	"github.com/GenerateNU/sac/backend/integrations/cache"
	"github.com/GenerateNU/sac/backend/integrations/email"
	"github.com/GenerateNU/sac/backend/integrations/file"
	"github.com/GenerateNU/sac/backend/integrations/search"
)

type Integrations struct {
	Search      search.SearchClientInterface
	AI          search.AIClientInterface
	Email       email.EmailClientInterface
	File        file.FileClientInterface
	Blacklist   cache.BlacklistInterface
	ActiveToken cache.ActiveTokenInterface
}

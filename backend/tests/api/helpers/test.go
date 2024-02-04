package helpers

import (
	"testing"

	"github.com/huandu/go-assert"
)

func InitTest(t *testing.T) ExistingAppAssert {
	assert := assert.New(t)
	app, err := spawnApp()

	assert.NilError(err)

	return ExistingAppAssert{
		App:    *app,
		Assert: assert,
	}
}

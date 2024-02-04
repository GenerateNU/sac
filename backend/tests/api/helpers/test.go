package helpers

import (
	"testing"

	"github.com/huandu/go-assert"
)

func InitTest(t *testing.T) (TestApp, *assert.A) {
	assert := assert.New(t)
	app, err := spawnApp()

	assert.NilError(err)

	return *app, assert
}

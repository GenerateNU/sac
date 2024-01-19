package tests

import (
	"backend/src/utilities"
	"testing"

	"github.com/huandu/go-assert"
)

func TestThatContainsWorks(t *testing.T) {
	assert := assert.New(t)

	slice := []string{"foo", "bar", "baz"}

	assert.Assert(utilities.Contains(slice, "foo"))
	assert.Assert(utilities.Contains(slice, "bar"))
	assert.Assert(utilities.Contains(slice, "baz"))
	assert.Assert(!utilities.Contains(slice, "qux"))
}

package tests

import (
	"testing"

	"github.com/GenerateNU/sac/backend/src/utilities"

	"github.com/huandu/go-assert"
)

func TestThatAllCasingPermutationsWorks(t *testing.T) {
	assert := assert.New(t)

	expectedPermutations := []string{"foo", "Foo", "fOo", "foO", "FOo", "FoO", "fOO", "FOO"}

	acutalPermutations := AllCasingPermutations("foo")

	for _, permutation := range expectedPermutations {
		assert.Assert(utilities.Contains(acutalPermutations, permutation))
	}
}

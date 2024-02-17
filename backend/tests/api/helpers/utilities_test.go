package helpers

import (
	"slices"
	"testing"

	"github.com/huandu/go-assert"
)

func TestThatAllCasingPermutationsWorks(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	expectedPermutations := []string{"foo", "Foo", "fOo", "foO", "FOo", "FoO", "fOO", "FOO"}

	acutalPermutations := AllCasingPermutations("foo")

	for _, permutation := range expectedPermutations {
		assert.Assert(slices.Contains(acutalPermutations, permutation))
	}
}

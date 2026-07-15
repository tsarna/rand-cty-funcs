package randcty

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Every name is under the rand:: namespace. `random()` in particular becomes
// `rand::float`, not `rand::random` — a leaf name should not repeat its namespace.
func TestNamesAreNamespaced(t *testing.T) {
	funcs := GetRandomFunctions()

	for name := range funcs {
		assert.True(t, strings.HasPrefix(name, "rand::"), "%s() is not under the rand:: namespace", name)
	}
	assert.Contains(t, funcs, "rand::float", "the [0,1) generator should be rand::float")
	assert.NotContains(t, funcs, "rand::random", "rand::random repeats the namespace; it is rand::float")
}

// Every function, and every parameter, carries a cty description. These functions have
// honest cty signatures (concrete parameter types, no variadics, describable returns),
// so they need no functy extern — the cty metadata is the whole of their documentation,
// and it must be complete.
func TestEverythingIsDescribed(t *testing.T) {
	for name, fn := range GetRandomFunctions() {
		assert.NotEmpty(t, fn.Description(), "%s() has no cty Description", name)

		for _, p := range fn.Params() {
			assert.NotEmpty(t, p.Description, "%s() parameter %q has no Description", name, p.Name)
		}
		assert.Nil(t, fn.VarParam(),
			"%s() has grown a VarParam; its cty signature can no longer be honest and it now "+
				"needs a functy extern", name)
	}
}

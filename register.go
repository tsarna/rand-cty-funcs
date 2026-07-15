package randcty

import "github.com/zclconf/go-cty/cty/function"

// GetRandomFunctions returns all random-number cty functions for registration
// in an HCL2 eval context.
//
// The names are namespaced under `rand::`. HCL parses `a::b(x)` natively as a single
// flat map key, so this is a naming choice, not a structural one; the leaf names drop
// the `rand` prefix the flat names carried and sort together. `random()` becomes
// `rand::float` — it returns a float in [0.0, 1.0), and a leaf name should not repeat
// its namespace.
func GetRandomFunctions() map[string]function.Function {
	return map[string]function.Function{
		"rand::float":   RandomFunc,
		"rand::int":     RandIntFunc,
		"rand::uniform": RandUniformFunc,
		"rand::gauss":   RandGaussFunc,
		"rand::choice":  RandChoiceFunc,
		"rand::sample":  RandSampleFunc,
		"rand::shuffle": RandShuffleFunc,
	}
}

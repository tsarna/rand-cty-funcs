package randcty

import (
	"fmt"
	"math/rand"

	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

// RandomFunc returns a random float in [0.0, 1.0).
var RandomFunc = function.New(&function.Spec{
	Description: "Returns a random float in [0.0, 1.0)",
	Params:      []function.Parameter{},
	Type:        function.StaticReturnType(cty.Number),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		return cty.NumberFloatVal(rand.Float64()), nil
	},
})

// RandIntFunc returns a random integer N such that a <= N <= b.
var RandIntFunc = function.New(&function.Spec{
	Description: "Returns a random integer N such that a <= N <= b",
	Params: []function.Parameter{
		{Name: "a", Type: cty.Number},
		{Name: "b", Type: cty.Number},
	},
	Type: function.StaticReturnType(cty.Number),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		a, _ := args[0].AsBigFloat().Int64()
		b, _ := args[1].AsBigFloat().Int64()
		if b < a {
			return cty.NilVal, fmt.Errorf("randint: b must be >= a")
		}
		return cty.NumberIntVal(rand.Int63n(b-a+1) + a), nil
	},
})

// RandUniformFunc returns a random float N such that a <= N <= b.
var RandUniformFunc = function.New(&function.Spec{
	Description: "Returns a random float N such that a <= N <= b",
	Params: []function.Parameter{
		{Name: "a", Type: cty.Number},
		{Name: "b", Type: cty.Number},
	},
	Type: function.StaticReturnType(cty.Number),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		a, _ := args[0].AsBigFloat().Float64()
		b, _ := args[1].AsBigFloat().Float64()
		if b < a {
			return cty.NilVal, fmt.Errorf("randuniform: b must be >= a")
		}
		return cty.NumberFloatVal(a + rand.Float64()*(b-a)), nil
	},
})

// RandGaussFunc returns a random float from a Gaussian distribution with the
// given mean (mu) and standard deviation (sigma).
var RandGaussFunc = function.New(&function.Spec{
	Description: "Returns a random float from a Gaussian distribution with the given mean and standard deviation",
	Params: []function.Parameter{
		{Name: "mu", Type: cty.Number},
		{Name: "sigma", Type: cty.Number},
	},
	Type: function.StaticReturnType(cty.Number),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		mu, _ := args[0].AsBigFloat().Float64()
		sigma, _ := args[1].AsBigFloat().Float64()
		return cty.NumberFloatVal(rand.NormFloat64()*sigma + mu), nil
	},
})

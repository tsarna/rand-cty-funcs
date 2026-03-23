package randcty

import "github.com/zclconf/go-cty/cty/function"

// GetRandomFunctions returns all random-number cty functions for registration
// in an HCL2 eval context.
func GetRandomFunctions() map[string]function.Function {
	return map[string]function.Function{
		"random":      RandomFunc,
		"randint":     RandIntFunc,
		"randuniform": RandUniformFunc,
		"randgauss":   RandGaussFunc,
		"randchoice":  RandChoiceFunc,
		"randsample":  RandSampleFunc,
		"randshuffle": RandShuffleFunc,
	}
}

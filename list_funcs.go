package randcty

import (
	"fmt"
	"math/rand"

	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

// RandChoiceFunc returns a random element from a list.
var RandChoiceFunc = function.New(&function.Spec{
	Description: "Returns a random element from a list",
	Params: []function.Parameter{
		{Name: "list", Type: cty.List(cty.DynamicPseudoType)},
	},
	Type: func(args []cty.Value) (cty.Type, error) {
		listTy := args[0].Type()
		if listTy == cty.DynamicPseudoType {
			return cty.DynamicPseudoType, nil
		}
		return listTy.ElementType(), nil
	},
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		var items []cty.Value
		for it := args[0].ElementIterator(); it.Next(); {
			_, v := it.Element()
			items = append(items, v)
		}
		if len(items) == 0 {
			return cty.NilVal, fmt.Errorf("randchoice: list must not be empty")
		}
		return items[rand.Intn(len(items))], nil
	},
})

// RandSampleFunc returns k unique random elements from a list (without replacement).
var RandSampleFunc = function.New(&function.Spec{
	Description: "Returns k unique random elements from a list",
	Params: []function.Parameter{
		{Name: "list", Type: cty.List(cty.DynamicPseudoType)},
		{Name: "k", Type: cty.Number},
	},
	Type: func(args []cty.Value) (cty.Type, error) {
		listTy := args[0].Type()
		if listTy == cty.DynamicPseudoType {
			return cty.DynamicPseudoType, nil
		}
		return listTy, nil
	},
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		var items []cty.Value
		for it := args[0].ElementIterator(); it.Next(); {
			_, v := it.Element()
			items = append(items, v)
		}
		k64, _ := args[1].AsBigFloat().Int64()
		k := int(k64)
		if k < 0 {
			return cty.NilVal, fmt.Errorf("randsample: k must be non-negative")
		}
		if k > len(items) {
			return cty.NilVal, fmt.Errorf("randsample: k (%d) exceeds list length (%d)", k, len(items))
		}
		if k == 0 {
			return cty.ListValEmpty(retType.ElementType()), nil
		}
		perm := rand.Perm(len(items))
		result := make([]cty.Value, k)
		for i := 0; i < k; i++ {
			result[i] = items[perm[i]]
		}
		return cty.ListVal(result), nil
	},
})

// RandShuffleFunc returns a shuffled copy of a list.
var RandShuffleFunc = function.New(&function.Spec{
	Description: "Returns a shuffled copy of a list",
	Params: []function.Parameter{
		{Name: "list", Type: cty.List(cty.DynamicPseudoType)},
	},
	Type: func(args []cty.Value) (cty.Type, error) {
		return args[0].Type(), nil
	},
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		var items []cty.Value
		for it := args[0].ElementIterator(); it.Next(); {
			_, v := it.Element()
			items = append(items, v)
		}
		if len(items) == 0 {
			return cty.ListValEmpty(retType.ElementType()), nil
		}
		shuffled := make([]cty.Value, len(items))
		copy(shuffled, items)
		rand.Shuffle(len(shuffled), func(i, j int) {
			shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
		})
		return cty.ListVal(shuffled), nil
	},
})

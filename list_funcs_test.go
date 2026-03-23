package randcty

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zclconf/go-cty/cty"
)

// --- randchoice ---

func TestRandChoiceReturnsElement(t *testing.T) {
	list := cty.ListVal([]cty.Value{
		cty.StringVal("a"),
		cty.StringVal("b"),
		cty.StringVal("c"),
	})
	result, err := RandChoiceFunc.Call([]cty.Value{list})
	require.NoError(t, err)
	assert.Equal(t, cty.String, result.Type())
	v := result.AsString()
	assert.Contains(t, []string{"a", "b", "c"}, v)
}

func TestRandChoiceEmptyList(t *testing.T) {
	list := cty.ListValEmpty(cty.String)
	_, err := RandChoiceFunc.Call([]cty.Value{list})
	assert.Error(t, err)
}

func TestRandChoiceSingleElement(t *testing.T) {
	list := cty.ListVal([]cty.Value{cty.StringVal("only")})
	for range 10 {
		result, err := RandChoiceFunc.Call([]cty.Value{list})
		require.NoError(t, err)
		assert.Equal(t, "only", result.AsString())
	}
}

func TestRandChoiceNumberList(t *testing.T) {
	list := cty.ListVal([]cty.Value{
		cty.NumberIntVal(1),
		cty.NumberIntVal(2),
		cty.NumberIntVal(3),
	})
	result, err := RandChoiceFunc.Call([]cty.Value{list})
	require.NoError(t, err)
	assert.Equal(t, cty.Number, result.Type())
}

// --- randsample ---

func TestRandSampleReturnsK(t *testing.T) {
	list := cty.ListVal([]cty.Value{
		cty.StringVal("a"),
		cty.StringVal("b"),
		cty.StringVal("c"),
		cty.StringVal("d"),
		cty.StringVal("e"),
	})
	result, err := RandSampleFunc.Call([]cty.Value{list, cty.NumberIntVal(3)})
	require.NoError(t, err)
	assert.Equal(t, cty.List(cty.String), result.Type())
	var items []cty.Value
	for it := result.ElementIterator(); it.Next(); {
		_, v := it.Element()
		items = append(items, v)
	}
	assert.Len(t, items, 3)
}

func TestRandSampleKZero(t *testing.T) {
	list := cty.ListVal([]cty.Value{
		cty.StringVal("a"),
		cty.StringVal("b"),
	})
	result, err := RandSampleFunc.Call([]cty.Value{list, cty.NumberIntVal(0)})
	require.NoError(t, err)
	assert.True(t, result.LengthInt() == 0)
}

func TestRandSampleKExceedsLength(t *testing.T) {
	list := cty.ListVal([]cty.Value{
		cty.StringVal("a"),
		cty.StringVal("b"),
	})
	_, err := RandSampleFunc.Call([]cty.Value{list, cty.NumberIntVal(5)})
	assert.Error(t, err)
}

func TestRandSampleKNegative(t *testing.T) {
	list := cty.ListVal([]cty.Value{cty.StringVal("a")})
	_, err := RandSampleFunc.Call([]cty.Value{list, cty.NumberIntVal(-1)})
	assert.Error(t, err)
}

func TestRandSampleNoRepeats(t *testing.T) {
	list := cty.ListVal([]cty.Value{
		cty.StringVal("a"),
		cty.StringVal("b"),
		cty.StringVal("c"),
		cty.StringVal("d"),
		cty.StringVal("e"),
	})
	for range 20 {
		result, err := RandSampleFunc.Call([]cty.Value{list, cty.NumberIntVal(5)})
		require.NoError(t, err)
		seen := map[string]bool{}
		for it := result.ElementIterator(); it.Next(); {
			_, v := it.Element()
			s := v.AsString()
			assert.False(t, seen[s], "duplicate element: %s", s)
			seen[s] = true
		}
	}
}

// --- randshuffle ---

func TestRandShuffleSameLength(t *testing.T) {
	list := cty.ListVal([]cty.Value{
		cty.StringVal("a"),
		cty.StringVal("b"),
		cty.StringVal("c"),
	})
	result, err := RandShuffleFunc.Call([]cty.Value{list})
	require.NoError(t, err)
	assert.Equal(t, cty.List(cty.String), result.Type())
	assert.Equal(t, 3, result.LengthInt())
}

func TestRandShuffleSameElements(t *testing.T) {
	list := cty.ListVal([]cty.Value{
		cty.StringVal("a"),
		cty.StringVal("b"),
		cty.StringVal("c"),
	})
	result, err := RandShuffleFunc.Call([]cty.Value{list})
	require.NoError(t, err)
	seen := map[string]bool{}
	for it := result.ElementIterator(); it.Next(); {
		_, v := it.Element()
		seen[v.AsString()] = true
	}
	assert.True(t, seen["a"])
	assert.True(t, seen["b"])
	assert.True(t, seen["c"])
}

func TestRandShuffleEmptyList(t *testing.T) {
	list := cty.ListValEmpty(cty.String)
	result, err := RandShuffleFunc.Call([]cty.Value{list})
	require.NoError(t, err)
	assert.Equal(t, 0, result.LengthInt())
}

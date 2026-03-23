package randcty

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zclconf/go-cty/cty"
)

// --- random ---

func TestRandomRange(t *testing.T) {
	for range 100 {
		result, err := RandomFunc.Call([]cty.Value{})
		require.NoError(t, err)
		f, _ := result.AsBigFloat().Float64()
		assert.GreaterOrEqual(t, f, 0.0)
		assert.Less(t, f, 1.0)
	}
}

func TestRandomReturnsNumber(t *testing.T) {
	result, err := RandomFunc.Call([]cty.Value{})
	require.NoError(t, err)
	assert.Equal(t, cty.Number, result.Type())
}

// --- randint ---

func TestRandIntInRange(t *testing.T) {
	for range 100 {
		result, err := RandIntFunc.Call([]cty.Value{
			cty.NumberIntVal(5),
			cty.NumberIntVal(10),
		})
		require.NoError(t, err)
		n, _ := result.AsBigFloat().Int64()
		assert.GreaterOrEqual(t, n, int64(5))
		assert.LessOrEqual(t, n, int64(10))
	}
}

func TestRandIntSingleValue(t *testing.T) {
	// When a == b, always returns a
	for range 10 {
		result, err := RandIntFunc.Call([]cty.Value{
			cty.NumberIntVal(7),
			cty.NumberIntVal(7),
		})
		require.NoError(t, err)
		n, _ := result.AsBigFloat().Int64()
		assert.Equal(t, int64(7), n)
	}
}

func TestRandIntNegativeRange(t *testing.T) {
	result, err := RandIntFunc.Call([]cty.Value{
		cty.NumberIntVal(-10),
		cty.NumberIntVal(-1),
	})
	require.NoError(t, err)
	n, _ := result.AsBigFloat().Int64()
	assert.GreaterOrEqual(t, n, int64(-10))
	assert.LessOrEqual(t, n, int64(-1))
}

func TestRandIntBLessThanA(t *testing.T) {
	_, err := RandIntFunc.Call([]cty.Value{
		cty.NumberIntVal(10),
		cty.NumberIntVal(5),
	})
	assert.Error(t, err)
}

// --- randuniform ---

func TestRandUniformInRange(t *testing.T) {
	for range 100 {
		result, err := RandUniformFunc.Call([]cty.Value{
			cty.NumberFloatVal(1.5),
			cty.NumberFloatVal(3.5),
		})
		require.NoError(t, err)
		f, _ := result.AsBigFloat().Float64()
		assert.GreaterOrEqual(t, f, 1.5)
		assert.LessOrEqual(t, f, 3.5)
	}
}

func TestRandUniformBLessThanA(t *testing.T) {
	_, err := RandUniformFunc.Call([]cty.Value{
		cty.NumberFloatVal(5.0),
		cty.NumberFloatVal(1.0),
	})
	assert.Error(t, err)
}

// --- randgauss ---

func TestRandGaussReturnsNumber(t *testing.T) {
	result, err := RandGaussFunc.Call([]cty.Value{
		cty.NumberFloatVal(0.0),
		cty.NumberFloatVal(1.0),
	})
	require.NoError(t, err)
	assert.Equal(t, cty.Number, result.Type())
}

func TestRandGaussMeanShift(t *testing.T) {
	// With a very large mean, all results should be far positive
	for range 20 {
		result, err := RandGaussFunc.Call([]cty.Value{
			cty.NumberFloatVal(1e9),
			cty.NumberFloatVal(1.0),
		})
		require.NoError(t, err)
		f, _ := result.AsBigFloat().Float64()
		assert.Greater(t, f, 0.0)
	}
}

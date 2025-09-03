package mathtools

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxN(t *testing.T) {
	assert.Equal(t, math.MaxInt, MaxN[int]())
	assert.Equal(t, int8(math.MaxInt8), MaxN[int8]())
	assert.Equal(t, int16(math.MaxInt16), MaxN[int16]())
	assert.Equal(t, int32(math.MaxInt32), MaxN[int32]())
	assert.Equal(t, int64(math.MaxInt64), MaxN[int64]())

	assert.Equal(t, uint(math.MaxUint), MaxN[uint]())
	assert.Equal(t, uint8(math.MaxUint8), MaxN[uint8]())
	assert.Equal(t, uint16(math.MaxUint16), MaxN[uint16]())
	assert.Equal(t, uint32(math.MaxUint32), MaxN[uint32]())
	assert.Equal(t, uint64(math.MaxUint64), MaxN[uint64]())

	assert.Equal(t, float32(math.MaxFloat32), MaxN[float32]())
	assert.Equal(t, math.MaxFloat64, MaxN[float64]())
}

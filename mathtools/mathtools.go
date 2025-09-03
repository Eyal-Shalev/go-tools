package mathtools

import (
	"math"

	"golang.org/x/exp/constraints"
)

type Rational interface {
	constraints.Integer | constraints.Float
}

func MaxN[N Rational]() N {
	var zero N
	switch any(zero).(type) {

	case int:
		return any(math.MaxInt).(N)
	case int8:
		return any(int8(math.MaxInt8)).(N)
	case int16:
		return any(int16(math.MaxInt16)).(N)
	case int32:
		return any(int32(math.MaxInt32)).(N)
	case int64:
		return any(int64(math.MaxInt64)).(N)

	case uint:
		return any(uint(math.MaxUint)).(N)
	case uint8:
		return any(uint8(math.MaxUint8)).(N)
	case uint16:
		return any(uint16(math.MaxUint16)).(N)
	case uint32:
		return any(uint32(math.MaxUint32)).(N)
	case uint64:
		return any(uint64(math.MaxUint64)).(N)

	case float32:
		return any(float32(math.MaxFloat32)).(N)
	case float64:
		return any(math.MaxFloat64).(N)

	default:
		panic("unreachable")
	}
}

func ModFn[N Rational](modBy N) func(N) N {
	return func(val N) N {
		return N(math.Mod(float64(val), float64(modBy)))
	}
}

package functools

func Identity[T any](v T) T {
	return v
}

func ToAny[T any](v T) any {
	return v
}

func ToAny2[K, V any](k K, v V) (K, any) {
	return k, v
}

func MustAssertType[T any](v any) T {
	return v.(T)
}

func Not1[T any](f func(T) bool) func(T) bool {
	return func(v T) bool {
		return !f(v)
	}
}

func Not2[V1, V2 any](f func(V1, V2) bool) func(V1, V2) bool {
	return func(v1 V1, v2 V2) bool {
		return !f(v1, v2)
	}
}

func Not3[V1, V2, V3 any](f func(V1, V2, V3) bool) func(V1, V2, V3) bool {
	return func(v1 V1, v2 V2, v3 V3) bool {
		return !f(v1, v2, v3)
	}
}

func Combine[V1, V2, V3 any](f1to2 func(V1) V2, f2to3 func(V2) V3) func(V1) V3 {
	return func(v1 V1) V3 {
		return f2to3(f1to2(v1))
	}
}

func Combine3[V1, V2, V3, V4 any](f1to2 func(V1) V2, f2to3 func(V2) V3, f3to4 func(V3) V4) func(V1) V4 {
	return func(v1 V1) V4 {
		return f3to4(f2to3(f1to2(v1)))
	}
}

func Always[T any](v T) func() T {
	return func() T {
		return v
	}
}

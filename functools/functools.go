package functools

func Identity[V any](v V) V {
	return v
}

func Not1[V any](f func(V) bool) func(V) bool {
	return func(v V) bool {
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

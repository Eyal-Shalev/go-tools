package slicetools

func Select[V any, S ~[]V](s S, f func(V) bool) S {
	result := make(S, 0, len(s))
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

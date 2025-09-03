package maptools

func Select[K comparable, V any, M ~map[K]V](m M, f func(K, V) bool) M {
	result := make(M, len(m))
	for k, v := range m {
		if f(k, v) {
			result[k] = v
		}
	}
	return result
}

package maptools

func Partition[K comparable, V any, M ~map[K]V](s M, f func(K, V) bool) (M, M) {
	left := make(M, len(s))
	right := make(M, len(s))
	for k, v := range s {
		if f(k, v) {
			left[k] = v
		} else {
			right[k] = v
		}
	}
	return left, right
}

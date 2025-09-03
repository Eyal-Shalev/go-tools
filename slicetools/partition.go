package slicetools

func Partition[V any, S ~[]V](s S, f func(V) bool) (S, S) {
	left := make(S, 0, len(s))
	right := make(S, 0, len(s))
	for _, v := range s {
		if f(v) {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	return left, right
}

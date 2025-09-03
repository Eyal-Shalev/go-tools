package maptools

import "github.com/Eyal-Shalev/itertools-go/internal"

func Reject[K comparable, V any, M ~map[K]V](m M, f func(K, V) bool) M {
	return Select(m, internal.Not2(f))
}

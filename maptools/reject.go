package maptools

import "github.com/Eyal-Shalev/go-tools/functools"

func Reject[K comparable, V any, M ~map[K]V](m M, f func(K, V) bool) M {
	return Select(m, functools.Not2(f))
}

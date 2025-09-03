package maptools

import (
	"maps"

	"github.com/Eyal-Shalev/go-tools/functools"
	"github.com/Eyal-Shalev/go-tools/itertools"
)

func Map[M ~map[K1]V1, K1, K2 comparable, V1, V2 any](m M, f func(K1, V1) (K2, V2)) map[K2]V2 {
	return maps.Collect(itertools.Map2(maps.All(m), f))
}

func ToAny[M ~map[K]V, K comparable, V any](m M) map[K]any {
	return Map(m, functools.ToAny2)
}

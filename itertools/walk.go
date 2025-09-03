package itertools

import (
	"iter"
)

func Walk[V any](seq iter.Seq[V], fn func(V)) {
	for v := range seq {
		fn(v)
	}
}

func Walk2[K, V any](seq iter.Seq2[K, V], fn func(K, V)) {
	for k, v := range seq {
		fn(k, v)
	}
}

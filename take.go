package itertools

import (
	"iter"
	"math"
)

func Take[V any](seq iter.Seq[V], optFns ...takeOption[V]) iter.Seq[V] {
	options := takeOptions[V]{
		limit:       math.MaxUint,
		includeLast: false,
		predicate:   func(V) bool { return true },
	}
	for _, optFn := range optFns {
		optFn(&options)
	}
	return func(yield func(V) bool) {
		for idx, v := range Enumerate(seq) {
			if idx >= options.limit {
				return
			}
			if !options.includeLast && !options.predicate(v) {
				return
			}
			if !yield(v) {
				return
			}
			if !options.predicate(v) {
				return
			}
		}
	}
}

func Take2[K, V any](seq iter.Seq2[K, V], optFns ...takeOption[Entry[K, V]]) iter.Seq2[K, V] {
	options := takeOptions[Entry[K, V]]{
		limit:       math.MaxUint,
		includeLast: false,
		predicate:   func(Entry[K, V]) bool { return true },
	}
	for _, optFn := range optFns {
		optFn(&options)
	}
	return func(yield func(K, V) bool) {
		for idx, entry := range Enumerate2(seq) {
			if idx >= options.limit {
				return
			}
			if !options.includeLast && !options.predicate(entry) {
				return
			}
			if !yield(entry.Key, entry.Val) {
				return
			}
			if !options.predicate(entry) {
				return
			}
		}
	}
}

type takeOptions[V any] struct {
	limit       uint
	predicate   func(V) bool
	includeLast bool
}

type takeOption[V any] func(options *takeOptions[V])

//goland:noinspection GoExportedFuncWithUnexportedType
func WithLimit[V any](limit uint) takeOption[V] {
	return func(options *takeOptions[V]) {
		options.limit = limit
	}
}

//goland:noinspection GoExportedFuncWithUnexportedType
func WithPredicate[V any](predicate func(V) bool) takeOption[V] {
	return func(options *takeOptions[V]) {
		options.predicate = predicate
	}
}

//goland:noinspection GoExportedFuncWithUnexportedType
func WithIncludeLast[V any](includeLast bool) takeOption[V] {
	return func(options *takeOptions[V]) {
		options.includeLast = includeLast
	}
}

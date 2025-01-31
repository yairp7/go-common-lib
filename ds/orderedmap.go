package ds

import (
	"iter"
	"maps"
	"slices"
)

type OrderedMap[K string, V any] map[K]V

func (m OrderedMap[K, V]) Entries() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, k := range m.sortedKeys() {
			if !yield(k, m[k]) {
				return
			}
		}
	}
}

func (m OrderedMap[K, V]) Keys() iter.Seq[K] {
	return func(yield func(K) bool) {
		for _, k := range slices.Sorted(maps.Keys(m)) {
			if !yield(k) {
				return
			}
		}
	}
}

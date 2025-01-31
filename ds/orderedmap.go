package ds

import (
	"iter"
	"maps"
	"slices"
)

type OrderedMap[K string, V any] map[K]V

func (m OrderedMap[K, V]) sortedKeys() []K {
	slices.SortedStableFunc(maps.Keys(m), func(a, b K) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})
	return slices.Collect(maps.Keys(m))
}

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
		for _, k := range m.sortedKeys() {
			if !yield(k) {
				return
			}
		}
	}
}

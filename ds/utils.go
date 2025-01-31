package ds

type forEachFunc[K comparable, V any] func(k K, v V)

func forEachMap[K comparable, V any](s map[K]V, f forEachFunc[K, V]) {
	for k, v := range s {
		f(k, v)
	}
}

package slices

func Reduce[T any, R any](src []T, initValue R, f func(acc R, item T) R) R {
	acc := initValue
	for _, item := range src {
		acc = f(acc, item)
	}
	return acc
}

func Filter[T any](src []T, f func(item T) bool) []T {
	result := []T{}
	for _, item := range src {
		if f(item) {
			result = append(result, item)
		}
	}
	return result
}

func Map[T any, R any](src []T, f func(item T) R) []R {
	result := make([]R, len(src))
	for i := 0; i < len(src); i++ {
		result[i] = f(src[i])
	}
	return result
}

func ForEach[T any](src []T, f func(item T)) {
	for _, item := range src {
		f(item)
	}
}

func Keys[K comparable, V any](src map[K]V) []K {
	keys := make([]K, len(src))
	i := 0
	for k := range src {
		keys[i] = k
	}
	return keys
}

func Values[K comparable, V any](src map[K]V) []V {
	values := make([]V, len(src))
	i := 0
	for _, v := range src {
		values[i] = v
		i++
	}
	return values
}

func Unique[T any, V comparable](slice []T, f func(item T) V) []T {
	keys := make(map[V]bool)
	result := []T{}
	for _, entry := range slice {
		v := f(entry)
		if _, ok := keys[v]; !ok {
			keys[v] = true
			result = append(result, entry)
		}
	}
	return result
}

func MapToSet[T any, R comparable](src []T, f func(item T) R) map[R]bool {
	result := make(map[R]bool, len(src))
	for i := 0; i < len(src); i++ {
		result[f(src[i])] = true
	}
	return result
}

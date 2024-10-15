package slices

import "reflect"

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

func Contains[T comparable](src []T, query T) bool {
	for _, item := range src {
		if item == query {
			return true
		}
	}
	return false
}

func Map[T any, R any](src []T, f func(item T) R) []R {
	result := make([]R, len(src))
	for i := 0; i < len(src); i++ {
		result[i] = f(src[i])
	}
	return result
}

func MapI[T any, R any](src []T, f func(item T, index int) R) []R {
	result := make([]R, len(src))
	for i := 0; i < len(src); i++ {
		result[i] = f(src[i], i)
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

func Some[T any](src []T, f func(item T) bool) bool {
	for i := 0; i < len(src); i++ {
		if f(src[i]) {
			return true
		}
	}
	return false
}

func Index[T comparable](s []T, query T) int {
	for i, c := range s {
		if c == query {
			return i
		}
	}

	return -1
}

func Find[T comparable](s []T, query T) (result T, isFound bool) {
	for _, c := range s {
		if c == query {
			result = c
			isFound = true
			break
		}
	}
	return result, isFound
}

func FindGeneric[T any](s []T, query T, equalFunc func(a, b T) bool) (result T, isFound bool) {
	for _, c := range s {
		if equalFunc(c, query) {
			result = c
			isFound = true
			break
		}
	}
	return result, isFound
}

func FilterNils[T any](s []T) []T {
	return Filter[T](s, func(item T) bool {
		if v := reflect.ValueOf(item); v.Kind() == reflect.Ptr ||
			v.Kind() == reflect.Interface ||
			v.Kind() == reflect.Slice ||
			v.Kind() == reflect.Map ||
			v.Kind() == reflect.Chan ||
			v.Kind() == reflect.Func {
			return !v.IsNil()
		}
		return true
	})
}

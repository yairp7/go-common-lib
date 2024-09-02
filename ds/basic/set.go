package ds

type Set[K comparable] struct {
	data map[K]struct{}
	size int
}

type SetOption[K comparable] func(*Set[K])

func WithSetItems[K comparable](items ...K) SetOption[K] {
	return func(s *Set[K]) {
		for _, item := range items {
			s.Add(item)
		}
	}
}

func NewSet[K comparable](capacity int, opts ...SetOption[K]) *Set[K] {
	s := &Set[K]{
		data: make(map[K]struct{}, capacity),
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Set[K]) Add(keys ...K) {
	for _, key := range keys {
		s.data[key] = struct{}{}
		s.size++
	}
}

func (s *Set[K]) Remove(key K) {
	delete(s.data, key)
	s.size--
}

func (s Set[K]) Has(key K) bool {
	if _, ok := s.data[key]; ok {
		return true
	}
	return false
}

func (s Set[K]) Size() int {
	return s.size
}

func (s Set[K]) Copy(dest *Set[K]) {
	for k := range s.data {
		dest.Add(k)
	}
}

func (s Set[K]) Keys() []K {
	keys := make([]K, s.Size())
	i := 0
	for k := range s.data {
		keys[i] = k
		i++
	}
	return keys
}

func (s *Set[K]) ForEach(f func(K)) {
	forEachMap(s.data, func(k K, v struct{}) {
		f(k)
	})
}

func (s *Set[K]) Clear() {
	clear(s.data)
	s.size = 0
}

func (s *Set[K]) Merge(other *Set[K]) {
	if other == nil {
		return
	}

	for k := range other.data {
		s.Add(k)
	}
}

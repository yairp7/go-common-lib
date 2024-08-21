package ds

type Set[K comparable] struct {
	data map[K]struct{}
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
	set := Set[K]{
		data: make(map[K]struct{}, capacity),
	}

	for _, opt := range opts {
		opt(&set)
	}

	return &set
}

func (s *Set[K]) Add(keys ...K) {
	for _, key := range keys {
		s.data[key] = struct{}{}
	}
}

func (s *Set[K]) Remove(key K) {
	delete(s.data, key)
}

func (s Set[K]) Has(key K) bool {
	if _, ok := s.data[key]; ok {
		return true
	}
	return false
}

func (s Set[K]) Size() int {
	return len(s.data)
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

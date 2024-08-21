package ds

type Set[K comparable] struct {
	data map[K]struct{}
}

type SetOption[K comparable] func(*Set[K])

func WithSetCapacity[K comparable](capacity int) SetOption[K] {
	return func(s *Set[K]) {
		s.data = make(map[K]struct{}, capacity)
	}
}

func WithSetItems[K comparable](items ...K) SetOption[K] {
	return func(s *Set[K]) {
		for _, item := range items {
			s.Add(item)
		}
	}
}

func NewSet[K comparable](opts ...SetOption[K]) Set[K] {
	set := Set[K]{}

	for _, opt := range opts {
		opt(&set)
	}

	if set.data == nil {
		set.data = make(map[K]struct{}, 0)
	}

	return set
}

func (s *Set[K]) Add(key K) {
	s.data[key] = struct{}{}
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

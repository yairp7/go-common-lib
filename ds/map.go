package ds

type Map[K comparable, V any] struct {
	data map[K]V
	size int
}

type MapOption[K comparable, V any] func(*Map[K, V])

type MapEntry[K comparable, V any] struct {
	Key   K
	Value V
}

func WithMapItems[K comparable, V any](items ...MapEntry[K, V]) MapOption[K, V] {
	return func(m *Map[K, V]) {
		for _, item := range items {
			m.Set(item.Key, item.Value)
		}
	}
}

func NewMap[K comparable, V any](capacity int, opts ...MapOption[K, V]) *Map[K, V] {
	m := &Map[K, V]{
		data: make(map[K]V, capacity),
	}

	for _, opt := range opts {
		opt(m)
	}

	return m
}

func (m *Map[K, V]) Set(key K, value V) {
	m.data[key] = value
	m.size++
}

func (m Map[K, V]) Get(key K) V {
	return m.data[key]
}

func (m *Map[K, V]) Remove(key K) {
	delete(m.data, key)
	m.size--
}

func (m Map[K, V]) Has(key K) bool {
	if _, ok := m.data[key]; ok {
		return true
	}
	return false
}

func (m Map[K, V]) Size() int {
	return len(m.data)
}

func (m Map[K, V]) Keys() []K {
	keys := make([]K, m.Size())
	i := 0
	for k := range m.data {
		keys[i] = k
		i++
	}
	return keys
}

func (m Map[K, V]) Values() []V {
	values := make([]V, m.Size())
	i := 0
	for _, v := range m.data {
		values[i] = v
		i++
	}
	return values
}

func (m Map[K, V]) Entries() []MapEntry[K, V] {
	entries := make([]MapEntry[K, V], m.Size())
	i := 0
	for k, v := range m.data {
		entries[i] = MapEntry[K, V]{
			Key:   k,
			Value: v,
		}
		i++
	}
	return entries
}

func (m *Map[K, V]) ForEach(f forEachFunc[K, V]) {
	forEachMap(m.data, f)
}

func (m *Map[K, V]) Clear() {
	clear(m.data)
	m.size = 0
}

func (m *Map[K, V]) Merge(other *Map[K, V]) {
	if other == nil {
		return
	}

	for k, v := range other.data {
		m.Set(k, v)
	}
}

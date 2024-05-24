package ds

type Map[K comparable, V any] struct {
	data map[K]V
}

func NewMap[K comparable, V any](capacity int) Map[K, V] {
	return Map[K, V]{
		data: make(map[K]V, capacity),
	}
}

func (m *Map[K, V]) Set(key K, value V) {
	m.data[key] = value
}

func (m Map[K, V]) Get(key K) V {
	return m.data[key]
}

func (m *Map[K, V]) Remove(key K) {
	delete(m.data, key)
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

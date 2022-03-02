package aki

// M represents generic map with useful methods
type M[K comparable, V any] map[K]V

// Values returns values of map with no defined order
func (m M[K, V]) Values() []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// Keys returns keys of map with no defined order
func (m M[K, V]) Keys() []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

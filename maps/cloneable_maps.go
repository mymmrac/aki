package maps

func (m CloneableMap[K, T, V]) Clone() CloneableMap[K, T, V] {
	if m == nil {
		return nil
	}

	copyMap := make(CloneableMap[K, T, V], len(m))
	for key, value := range m {
		copyMap[key] = value.Clone().(V)
	}
	return copyMap
}

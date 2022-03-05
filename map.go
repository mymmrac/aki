package aki

// M represents generic map with useful methods
type M[K comparable, V any] map[K]V

// Values returns values of this map with no defined order
func (m M[K, V]) Values() []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// MValues returns values of specified map with no defined order
func MValues[K comparable, V any](m map[K]V) []V {
	return M[K, V](m).Values()
}

// Keys returns keys of this map with no defined order
func (m M[K, V]) Keys() []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// MKeys returns keys of specified map with no defined order
func MKeys[K comparable, V any](m map[K]V) []K {
	return M[K, V](m).Keys()
}

// MPredicate defines map predicate
type MPredicate[K comparable, V any] func(key K, values V) bool

// MPredicateByKey defines map predicate by key
type MPredicateByKey[K comparable] func(key K) bool

// MPredicateByValue defines map predicate by value
type MPredicateByValue[V any] func(values V) bool

// Filter returns new map from this, filtered by key and value using provided predicate
func (m M[K, V]) Filter(predicate MPredicate[K, V]) M[K, V] {
	filtered := make(M[K, V])
	for key, value := range m {
		if predicate(key, value) {
			filtered[key] = value
		}
	}
	return filtered
}

// MFilter returns new map from specified, filtered by key and value using provided predicate
func MFilter[K comparable, V any](m map[K]V, predicate MPredicate[K, V]) map[K]V {
	return M[K, V](m).Filter(predicate)
}

// FilterByKey returns new map from this, filtered by key using provided predicate
func (m M[K, V]) FilterByKey(predicate MPredicateByKey[K]) M[K, V] {
	filtered := make(M[K, V])
	for key, value := range m {
		if predicate(key) {
			filtered[key] = value
		}
	}
	return filtered
}

// MFilterByKey returns new map from specified, filtered by key using provided predicate
func MFilterByKey[K comparable, V any](m map[K]V, predicate MPredicateByKey[K]) map[K]V {
	return M[K, V](m).FilterByKey(predicate)
}

// FilterByValue returns new map from this, filtered by value using provided predicate
func (m M[K, V]) FilterByValue(predicate MPredicateByValue[V]) M[K, V] {
	filtered := make(M[K, V])
	for key, value := range m {
		if predicate(value) {
			filtered[key] = value
		}
	}
	return filtered
}

// MFilterByValue returns new map from specified, filtered by value using provided predicate
func MFilterByValue[K comparable, V any](m map[K]V, predicate MPredicateByValue[V]) map[K]V {
	return M[K, V](m).FilterByValue(predicate)
}

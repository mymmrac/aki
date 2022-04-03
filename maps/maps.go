/*
Package maps provides useful generic types, methods & functions for maps.
*/
package maps

// Values returns values of this map with no defined order
func (m Map[K, V]) Values() []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// Values returns values of specified map with no defined order
func Values[K comparable, V any](m Map[K, V]) []V {
	return m.Values()
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

// Keys returns keys of this map with no defined order
func (m Map[K, V]) Keys() []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Keys returns keys of specified map with no defined order
func Keys[K comparable, V any](m Map[K, V]) []K {
	return m.Keys()
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

// Filter returns new map from this, filtered by key and value using provided predicate
func (m Map[K, V]) Filter(predicate Predicate[K, V]) Map[K, V] {
	if m == nil {
		return nil
	}

	filtered := make(Map[K, V])
	for key, value := range m {
		if predicate(key, value) {
			filtered[key] = value
		}
	}
	return filtered
}

func (m Map[K, V]) FilterSelf(predicate Predicate[K, V]) Map[K, V] {
	for key, value := range m {
		if !predicate(key, value) {
			delete(m, key)
		}
	}
	return m
}

// Filter returns new map from specified, filtered by key and value using provided predicate
func Filter[K comparable, V any](m Map[K, V], predicate Predicate[K, V]) map[K]V {
	return m.Filter(predicate)
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

// FilterByKey returns new map from this, filtered by key using provided predicate
func (m Map[K, V]) FilterByKey(predicate PredicateByKey[K]) Map[K, V] {
	if m == nil {
		return nil
	}

	filtered := make(Map[K, V])
	for key, value := range m {
		if predicate(key) {
			filtered[key] = value
		}
	}
	return filtered
}

func (m Map[K, V]) FilterSelfByKey(predicate PredicateByKey[K]) Map[K, V] {
	for key := range m {
		if !predicate(key) {
			delete(m, key)
		}
	}
	return m
}

// FilterByKey returns new map from specified, filtered by key using provided predicate
func FilterByKey[K comparable, V any](m Map[K, V], predicate PredicateByKey[K]) map[K]V {
	return m.FilterByKey(predicate)
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

// FilterByValue returns new map from this, filtered by value using provided predicate
func (m Map[K, V]) FilterByValue(predicate PredicateByValue[V]) Map[K, V] {
	if m == nil {
		return nil
	}

	filtered := make(Map[K, V])
	for key, value := range m {
		if predicate(value) {
			filtered[key] = value
		}
	}
	return filtered
}

func (m Map[K, V]) FilterSelfByValue(predicate PredicateByValue[V]) Map[K, V] {
	for key, value := range m {
		if !predicate(value) {
			delete(m, key)
		}
	}
	return m
}

// FilterByValue returns new map from specified, filtered by value using provided predicate
func FilterByValue[K comparable, V any](m Map[K, V], predicate PredicateByValue[V]) map[K]V {
	return m.FilterByValue(predicate)
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

// Entries returns entries of this map
func (m Map[K, V]) Entries() []Entry[K, V] {
	entries := make([]Entry[K, V], 0, len(m))
	for key, value := range m {
		entries = append(entries, Entry[K, V]{
			Key:   key,
			Value: value,
		})
	}
	return entries
}

// Entries returns entries of specified map
func Entries[K comparable, V any](m Map[K, V]) []Entry[K, V] {
	return m.Entries()
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

// FillEntries fills entries into this map
func (m Map[K, V]) FillEntries(entries []Entry[K, V]) Map[K, V] {
	for _, entry := range entries {
		m[entry.Key] = entry.Value
	}
	return m
}

// FillEntries fills entries into specified map
func FillEntries[K comparable, V any](m Map[K, V], entries []Entry[K, V]) {
	m.FillEntries(entries)
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

// FillEntry fills entries into this map
func (m Map[K, V]) FillEntry(entries ...Entry[K, V]) Map[K, V] {
	return m.FillEntries(entries)
}

// FillEntry fills entries into specified map
func FillEntry[K comparable, V any](m Map[K, V], entries ...Entry[K, V]) {
	m.FillEntries(entries)
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

// FromEntries fills entries into specified map
func FromEntries[K comparable, V any](entries []Entry[K, V]) Map[K, V] {
	m := make(Map[K, V], len(entries))
	m.FillEntries(entries)
	return m
}

// FromEntry fills entries into specified map
func FromEntry[K comparable, V any](entries ...Entry[K, V]) Map[K, V] {
	return FromEntries(entries)
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

// Copy returns shallow copy of this map
func (m Map[K, V]) Copy() Map[K, V] {
	if m == nil {
		return nil
	}

	copyMap := make(Map[K, V], len(m))
	for key, value := range m {
		copyMap[key] = value
	}
	return copyMap
}

// Copy returns shallow copy of specified map
func Copy[K comparable, V any](m Map[K, V]) map[K]V {
	return m.Copy()
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

func (m Map[K, V]) Merge(other Map[K, V]) Map[K, V] {
	merged := m.Copy()
	for key, value := range other {
		merged[key] = value
	}
	return merged
}

// MergeSelf merges values from provided map into this
func (m Map[K, V]) MergeSelf(other Map[K, V]) Map[K, V] {
	for key, value := range other {
		m[key] = value
	}
	return m
}

// Merge merges values from provided map into specified
func Merge[K comparable, V any](this, other Map[K, V]) Map[K, V] {
	return this.Merge(other)
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

func (m Map[K, V]) MergeLeft(other Map[K, V]) Map[K, V] {
	merged := m.Copy()
	for key, value := range other {
		if _, found := m[key]; found {
			continue
		}

		merged[key] = value
	}
	return merged
}

func (m Map[K, V]) MergeSelfLeft(other Map[K, V]) Map[K, V] {
	for key, value := range other {
		if _, found := m[key]; found {
			continue
		}

		m[key] = value
	}
	return m
}

func MergeLeft[K comparable, V any](this, other Map[K, V]) Map[K, V] {
	return this.MergeLeft(other)
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

func (m Map[K, V]) ContainsKey(key K) bool {
	_, found := m[key]
	return found
}

func ContainsKey[K comparable, V any](m Map[K, V], key K) bool {
	return m.ContainsKey(key)
}

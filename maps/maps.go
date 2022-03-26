/*
Package maps provides useful generic types, methods & functions for maps.
*/
package maps

import "github.com/mymmrac/aki"

// M represents generic map with useful methods
type M[K comparable, V any] map[K]V

func ToM[K comparable, V any](m map[K]V) M[K, V] {
	return m
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

// Values returns values of this map with no defined order
func (m M[K, V]) Values() []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// Values returns values of specified map with no defined order
func Values[K comparable, V any](m M[K, V]) []V {
	return m.Values()
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

// Keys returns keys of this map with no defined order
func (m M[K, V]) Keys() []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Keys returns keys of specified map with no defined order
func Keys[K comparable, V any](m M[K, V]) []K {
	return m.Keys()
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

// Predicate defines map predicate
type Predicate[K comparable, V any] func(key K, values V) bool

// PredicateByKey defines map predicate by key
type PredicateByKey[K comparable] func(key K) bool

// PredicateByValue defines map predicate by value
type PredicateByValue[V any] func(values V) bool

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

// Filter returns new map from this, filtered by key and value using provided predicate
func (m M[K, V]) Filter(predicate Predicate[K, V]) M[K, V] {
	if m == nil {
		return nil
	}

	filtered := make(M[K, V])
	for key, value := range m {
		if predicate(key, value) {
			filtered[key] = value
		}
	}
	return filtered
}

func (m M[K, V]) FilterSelf(predicate Predicate[K, V]) M[K, V] {
	for key, value := range m {
		if !predicate(key, value) {
			delete(m, key)
		}
	}
	return m
}

// Filter returns new map from specified, filtered by key and value using provided predicate
func Filter[K comparable, V any](m M[K, V], predicate Predicate[K, V]) map[K]V {
	return m.Filter(predicate)
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

// FilterByKey returns new map from this, filtered by key using provided predicate
func (m M[K, V]) FilterByKey(predicate PredicateByKey[K]) M[K, V] {
	if m == nil {
		return nil
	}

	filtered := make(M[K, V])
	for key, value := range m {
		if predicate(key) {
			filtered[key] = value
		}
	}
	return filtered
}

func (m M[K, V]) FilterSelfByKey(predicate PredicateByKey[K]) M[K, V] {
	for key := range m {
		if !predicate(key) {
			delete(m, key)
		}
	}
	return m
}

// FilterByKey returns new map from specified, filtered by key using provided predicate
func FilterByKey[K comparable, V any](m M[K, V], predicate PredicateByKey[K]) map[K]V {
	return m.FilterByKey(predicate)
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

// FilterByValue returns new map from this, filtered by value using provided predicate
func (m M[K, V]) FilterByValue(predicate PredicateByValue[V]) M[K, V] {
	if m == nil {
		return nil
	}

	filtered := make(M[K, V])
	for key, value := range m {
		if predicate(value) {
			filtered[key] = value
		}
	}
	return filtered
}

func (m M[K, V]) FilterSelfByValue(predicate PredicateByValue[V]) M[K, V] {
	for key, value := range m {
		if !predicate(value) {
			delete(m, key)
		}
	}
	return m
}

// FilterByValue returns new map from specified, filtered by value using provided predicate
func FilterByValue[K comparable, V any](m M[K, V], predicate PredicateByValue[V]) map[K]V {
	return m.FilterByValue(predicate)
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

// Entry represents generic map entry
type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

func NewEntry[K comparable, V any](key K, value V) Entry[K, V] {
	return Entry[K, V]{
		Key:   key,
		Value: value,
	}
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

// Entries returns entries of this map
func (m M[K, V]) Entries() []Entry[K, V] {
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
func Entries[K comparable, V any](m M[K, V]) []Entry[K, V] {
	return m.Entries()
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

// FillEntries fills entries into this map
func (m M[K, V]) FillEntries(entries []Entry[K, V]) M[K, V] {
	for _, entry := range entries {
		m[entry.Key] = entry.Value
	}
	return m
}

// FillEntries fills entries into specified map
func FillEntries[K comparable, V any](m M[K, V], entries []Entry[K, V]) {
	m.FillEntries(entries)
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

// FillEntry fills entries into this map
func (m M[K, V]) FillEntry(entries ...Entry[K, V]) M[K, V] {
	return m.FillEntries(entries)
}

// FillEntry fills entries into specified map
func FillEntry[K comparable, V any](m M[K, V], entries ...Entry[K, V]) {
	m.FillEntries(entries)
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

// FromEntries fills entries into specified map
func FromEntries[K comparable, V any](entries []Entry[K, V]) M[K, V] {
	m := make(M[K, V], len(entries))
	m.FillEntries(entries)
	return m
}

// FromEntry fills entries into specified map
func FromEntry[K comparable, V any](entries ...Entry[K, V]) M[K, V] {
	return FromEntries(entries)
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

// Copy returns shallow copy of this map
func (m M[K, V]) Copy() M[K, V] {
	if m == nil {
		return nil
	}

	copyMap := make(M[K, V], len(m))
	for key, value := range m {
		copyMap[key] = value
	}
	return copyMap
}

// Copy returns shallow copy of specified map
func Copy[K comparable, V any](m M[K, V]) map[K]V {
	return m.Copy()
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

type MCloneable[K comparable, T any, V aki.Cloneable[T]] map[K]V

func ToMCloneable[K comparable, T any, V aki.Cloneable[T]](m map[K]V) MCloneable[K, T, V] {
	return m
}

func (m MCloneable[K, T, V]) Clone() MCloneable[K, T, V] {
	if m == nil {
		return nil
	}

	copyMap := make(MCloneable[K, T, V], len(m))
	for key, value := range m {
		copyMap[key] = value.Clone().(V)
	}
	return copyMap
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

func (m M[K, V]) Merge(other M[K, V]) M[K, V] {
	merged := m.Copy()
	for key, value := range other {
		merged[key] = value
	}
	return merged
}

// MergeSelf merges values from provided map into this
func (m M[K, V]) MergeSelf(other M[K, V]) M[K, V] {
	for key, value := range other {
		m[key] = value
	}
	return m
}

// Merge merges values from provided map into specified
func Merge[K comparable, V any](this, other M[K, V]) M[K, V] {
	return this.Merge(other)
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

func (m M[K, V]) MergeLeft(other M[K, V]) M[K, V] {
	merged := m.Copy()
	for key, value := range other {
		if _, found := m[key]; found {
			continue
		}

		merged[key] = value
	}
	return merged
}

func (m M[K, V]) MergeSelfLeft(other M[K, V]) M[K, V] {
	for key, value := range other {
		if _, found := m[key]; found {
			continue
		}

		m[key] = value
	}
	return m
}

func MergeLeft[K comparable, V any](this, other M[K, V]) M[K, V] {
	return this.MergeLeft(other)
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

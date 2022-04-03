package maps

import "github.com/mymmrac/aki/types"

// Map represents generic map with useful methods
type Map[K comparable, V any] map[K]V

func ToMap[K comparable, V any](m map[K]V) Map[K, V] {
	return m
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

type ComparableMap[K, V comparable] map[K]V

func ToComparableMap[K, V comparable](m map[K]V) ComparableMap[K, V] {
	return m
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

type CloneableMap[K comparable, T any, V types.Cloneable[T]] map[K]V

func ToCloneableMap[K comparable, T any, V types.Cloneable[T]](m map[K]V) CloneableMap[K, T, V] {
	return m
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

// Predicate defines map predicate
type Predicate[K comparable, V any] func(key K, values V) bool

// PredicateByKey defines map predicate by key
type PredicateByKey[K comparable] func(key K) bool

// PredicateByValue defines map predicate by value
type PredicateByValue[V any] func(values V) bool

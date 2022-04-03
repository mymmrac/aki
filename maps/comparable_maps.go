package maps

import "github.com/mymmrac/aki/types"

func (m ComparableMap[K, V]) Contains(value V) bool {
	for _, mapValue := range m {
		if mapValue == value {
			return true
		}
	}
	return false
}

func Contains[K, V comparable](m ComparableMap[K, V], value V) bool {
	return m.Contains(value)
}

// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====

func (m ComparableMap[K, V]) FindKeyOf(value V) (K, bool) {
	for key, mapValue := range m {
		if mapValue == value {
			return key, true
		}
	}
	return types.Empty[K](), false
}

func FindKeyOf[K, V comparable](m ComparableMap[K, V], value V) (K, bool) {
	return m.FindKeyOf(value)
}

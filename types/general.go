package types

func Empty[T any]() T {
	var empty T
	return empty
}

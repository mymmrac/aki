package is

import "github.com/mymmrac/aki/types"

func If[T any](is bool, ifTrue, ifFalse T) T {
	if is {
		return ifTrue
	}
	return ifFalse
}

func IfTrue[T any](is bool, ifTrue T) T {
	if is {
		return ifTrue
	}

	return types.Empty[T]()
}

func Or[T comparable](first, second T) T {
	if first != types.Empty[T]() {
		return first
	}
	return second
}

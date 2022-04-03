package aki

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

	var t T
	return t
}

func Or[T comparable](first, second T) T {
	var empty T
	if first != empty {
		return first
	}
	return second
}

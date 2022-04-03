package types

type Cloneable[T any] interface {
	Clone() Cloneable[T]
}

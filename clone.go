package aki

type Cloneable[T any] interface {
	Clone() Cloneable[T]
}

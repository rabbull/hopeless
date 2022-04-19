package hopeless

type Result[T any] interface {
	Val() T
	Err() error
	Result() (T, error)
}
type resultImpl[T any] Tuple[T, error]

func (r *resultImpl[T]) Val() T {
	return r.A
}

func (r *resultImpl[T]) Err() error {
	return r.B
}

func (r *resultImpl[T]) Result() (T, error) {
	return r.A, r.B
}

func Ok[T any](val T) Result[T] {
	return &resultImpl[T]{
		A: val,
	}
}

func Err[T any](err error) Result[T] {
	return &resultImpl[T]{
		B: err,
	}
}

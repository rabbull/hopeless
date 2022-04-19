package hopeless

type Result[T any] interface {
	Val() T
	Err() error
	Result() (T, error)
	Expect(err error) T
	ExpectErr(err error) error
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

func (r *resultImpl[T]) Expect(err error) T {
	if r.Err() != nil {
		safePanic(err)
	}
	return r.Val()
}

func (r *resultImpl[T]) ExpectErr(err error) error {
	if r.Err() == nil {
		safePanic(err)
	}
	return r.Err()
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

func JoinResult[T, S any](lhs Result[T], rhs Result[S]) Result[Tuple[T, S]] {
	if lhs.Err() != nil {
		return Err[Tuple[T, S]](lhs.Err())
	}
	if rhs.Err() != nil {
		return Err[Tuple[T, S]](rhs.Err())
	}
	return Ok(Tuple[T, S]{
		A: lhs.Val(),
		B: rhs.Val(),
	})
}

func safePanic(err error) {
	if err != nil {
		panic(err)
	}
	panic(ErrExpectWithNilError)
}

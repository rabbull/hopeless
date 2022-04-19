package hopeless

import "errors"

var ErrPanic = errors.New("unknown panic")

type Tuple[T, S any] struct {
	A T
	B S
}

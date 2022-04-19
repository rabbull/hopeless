package hopeless

import "errors"

var ErrPanic = errors.New("unknown panic")
var ErrExpectWithNilError = errors.New("expect with nil error")

type Tuple[T, S any] struct {
	A T
	B S
}

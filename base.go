package hopeless

import "errors"

var ErrUnknownPanic = errors.New("unknown panic")
var ErrTimeout = errors.New("timeout")
var ErrExpectWithNilError = errors.New("expect with nil error")

type Tuple[T, S any] struct {
	A T
	B S
}

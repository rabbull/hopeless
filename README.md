# Hopeless

Yet another Promise/Future implementation in Go.

## Get Started

```go
package main

import (
    "errors"

    future "github.com/rabbull/hopeless"
)

func Fibonacci(n uint64) (uint64, error) {
    if n < 0 {
        return errors.New("invalid arguments")
    }

    if n == 0 {
        return 0, nil
    } else if n == 1 {
        return 1, nil
    }

    return Fibonacci(n - 1) + Fibonacci(n - 2), nil
}

func main() {
    fut := future.New(func() (uint64, error) {
		return Fibonacci(20)
	}).Then(func(val uint64) {
		t.Logf("fibonacci(20)=6765=%v", val)
	}).Catch(func(err error) {
		t.Logf("fibonacci(20) failed: err=%v", err)
	})

    // do something else

    fut.Wait() // block until future resolved
}
```
package hopeless_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	future "github.com/rabbull/hopeless"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestHopeless(t *testing.T) {
	future.Then(
		future.Then(
			future.New(func() future.Result[uint64] {
				n, err := DangerousFibonacci(20)
				if err != nil {
					return future.Err[uint64](err)
				}
				return future.Ok(n)
			}),
			func(fibonacciResult future.Result[uint64]) future.Result[uint8] {
				if fibonacciResult.Err() != nil {
					t.Logf("fibonacci failed: %v", fibonacciResult.Err())
					return future.Err[uint8](fibonacciResult.Err())
				}
				t.Logf("fibonacci succeeded: %v", fibonacciResult.Val())
				return future.Ok(uint8(fibonacciResult.Val() % 7))
			},
		),
		func(modResult future.Result[uint8]) future.Result[any] {
			if modResult.Err() != nil {
				t.Logf("fibonacci mod failed: %v", modResult.Err())
				return future.Err[any](modResult.Err())
			}
			t.Logf("fibonacci mod 7: %v", modResult.Val())
			return future.Ok[any](nil)
		},
	).Wait()
}

func DangerousFibonacci(n uint64) (uint64, error) {
	if n == 1 || n == 2 {
		return 1, nil
	}

	// very dangerous
	dice := rand.Int31n(10000)
	if dice == 0 {
		return 0, fmt.Errorf("yay")
	}

	x, err := DangerousFibonacci(n - 1)
	if err != nil {
		return 0, err
	}

	y, err := DangerousFibonacci(n - 2)
	if err != nil {
		return 0, err
	}

	return x + y, err
}

package future_test

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"

	future "github.com/rabbull/hopeless"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestFuture(t *testing.T) {
	future.New(func() (uint64, error) {
		return DangerousFibonacci(20)
	}).Then(func(val uint64) {
		t.Logf("fibonacci(20)=6765=%v", val)
	}).Catch(func(err error) {
		t.Logf("fibonacci(20) failed: err=%v", err)
	})
}

func DangerousFibonacci(n uint64) (uint64, error) {
	if n == 1 || n == 2 {
		return 1, nil
	}

	// very dangerous
	dice := rand.Int31n(10000)
	if dice == 0 {
		return 0, errors.New(fmt.Sprintf("yay"))
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

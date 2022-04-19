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
	_, _ = future.Then(
		future.Then(
			future.New(func() (uint64, error) {
				return DangerousFibonacci(20)
			}), func(fibonacci uint64, err error) (uint64, error) {
				if err != nil {
					t.Logf("fibonacci failed: %v", err)
					return 0, err
				}
				t.Logf("fibonacci finished: %v", fibonacci)
				return fibonacci % 7, nil
			}), func(mod uint64, err error) (any, error) {
			if err != nil {
				t.Logf("fibonacci mod failed: %v", err)
				return nil, err
			}
			t.Logf("fibonacci mod 7: %v", mod)
			return nil, nil
		}).Wait()
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

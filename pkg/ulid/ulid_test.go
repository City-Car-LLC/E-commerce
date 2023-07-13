package ulid

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/oklog/ulid/v2"
)

// from ulid_test.go
func TestMonotonicSafe(t *testing.T) {
	t.Parallel()

	var (
		src       = rand.NewSource(time.Now().UnixNano())
		entropy   = rand.New(src)
		monotonic = ulid.Monotonic(entropy, 0)
		safe      = &safeMonotonicReader{monotonicEntropy: *monotonic}
		t0        = ulid.Timestamp(time.Now())
	)

	errs := make(chan error, 100)
	for i := 0; i < cap(errs); i++ {
		go func() {
			u0 := ulid.MustNew(t0, safe)
			u1 := ulid.MustNew(t0, safe)
			for j := 0; j < 1024; j++ {
				u0, u1 = u1, ulid.MustNew(t0, safe)
				if u0.String() >= u1.String() {
					errs <- fmt.Errorf(
						"%s (%d %x) >= %s (%d %x)",
						u0.String(), u0.Time(), u0.Entropy(),
						u1.String(), u1.Time(), u1.Entropy(),
					)
					return
				}
			}
			errs <- nil
		}()
	}
	for i := 0; i < cap(errs); i++ {
		if err := <-errs; err != nil {
			t.Fatal(err)
		}
	}
}

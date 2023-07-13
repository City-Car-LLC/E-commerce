package ulid

import (
	"crypto/rand"
	"strings"
	"sync"

	"github.com/oklog/ulid/v2"
)

type safeMonotonicReader struct {
	mu               sync.Mutex
	monotonicEntropy ulid.MonotonicEntropy
}

func (r *safeMonotonicReader) Read(p []byte) (n int, err error) {
	return r.monotonicEntropy.Read(p)
}

func (r *safeMonotonicReader) MonotonicRead(ms uint64, p []byte) (err error) {
	r.mu.Lock()
	err = r.monotonicEntropy.MonotonicRead(ms, p)
	r.mu.Unlock()
	return err
}

var globalSafeMonotonicReader = &safeMonotonicReader{monotonicEntropy: *ulid.Monotonic(rand.Reader, 1)}

func New() string {
	u := ulid.MustNew(ulid.Now(), globalSafeMonotonicReader).String()
	return strings.ToLower(u)
}

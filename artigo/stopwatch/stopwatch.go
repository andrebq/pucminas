package stopwatch

import (
	"sync/atomic"
	"time"
)

type (
	// S is a stopwatch
	S struct {
		ptr *int32
	}
)

// New returns a new stopwatch which will stop after
// the given timeout
func New(timeout time.Duration) *S {
	val := new(int32)
	atomic.StoreInt32(val, 0)
	time.AfterFunc(timeout, func() {
		atomic.StoreInt32(val, 1)
	})
	return &S{
		ptr: val,
	}
}

// Stop returns true if the expected timeout has passed
func (s *S) Stop() bool {
	return atomic.LoadInt32(s.ptr) == 1
}

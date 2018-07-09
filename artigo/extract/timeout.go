package extract

import (
	"context"
	"io"
	"time"

	"github.com/pkg/errors"
)

type (
	// ContextRead allows the system to read data and fail if a
	// given timeout is exceeded.
	ContextRead struct {
		in      io.Reader
		ctx     context.Context
		timeout time.Duration
	}
)

var (
	errTimeout = errors.New("i/o timeout")
)

// IsTimeout returns true if the error represents a timeout
func IsTimeout(err error) bool {
	return errors.Cause(err) == errTimeout
}

// IsCancel returns true if the error represents a context cancel
func IsCancel(err error) bool {
	return errors.Cause(err) == context.Canceled
}

// NewContextRead returns a reader which is controlled by the given
// context and allows for a timeout on each read operation.
//
// The timeout is reset before each read
func NewContextRead(ctx context.Context, timeout time.Duration, in io.Reader) *ContextRead {
	return &ContextRead{
		ctx:     ctx,
		timeout: timeout,
		in:      in,
	}
}

// Read implements io.Reader
func (r *ContextRead) Read(buf []byte) (n int, err error) {
	ctx, cancel := context.WithTimeout(r.ctx, r.timeout)
	defer cancel()
	var readok bool
	go func() {
		n, err = r.in.Read(buf)
		readok = true
		cancel()
	}()
	select {
	case <-ctx.Done():
		if readok {
			return
		}
		return n, errTimeout
	}
}

// Close closes the input stream
func (r *ContextRead) Close() error {
	if in, ok := r.in.(io.Closer); ok {
		return in.Close()
	}
	return nil
}

package internal

import (
	"time"
)

// NewReal creates a new *Real using sane defaults.
func NewReal() *Real {
	return &Real{}
}

// Real implements the `clock.Clock` interface using the `time` package
// from the golang standard library.
type Real struct{}

// Now returns the current Real time.
func (c *Real) Now() time.Time {
	return time.Now()
}

// Since returns the time elapsed since the given time.
func (c *Real) Since(t time.Time) time.Duration {
	return time.Since(t)
}

// Sleep pauses the current goroutine for at least the duration d.
// A negative or zero duration causes Sleep to return immediately.
func (c *Real) Sleep(d time.Duration) {
	time.Sleep(d)
}

// Until returns the duration until the given time.
func (c *Real) Until(t time.Time) time.Duration {
	return time.Until(t)
}

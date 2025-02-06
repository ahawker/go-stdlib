package clock

import (
	"time"
)

var (
	_ Clock = (*clock)(nil)
)

// clock implements the `clock.Clock` interface using the `time` package
// from the golang standard library.
type clock struct{}

// Now returns the current clock time.
func (c *clock) Now() time.Time {
	return time.Now()
}

// Since returns the time elapsed since the given time.
func (c *clock) Since(t time.Time) time.Duration {
	return time.Since(t)
}

// Sleep pauses the current goroutine for at least the duration d.
// A negative or zero duration causes Sleep to return immediately.
func (c *clock) Sleep(d time.Duration) {
	time.Sleep(d)
}

// Until returns the duration until the given time.
func (c *clock) Until(t time.Time) time.Duration {
	return time.Until(t)
}

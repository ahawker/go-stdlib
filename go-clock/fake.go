package clock

import (
	"sync"
	"time"
)

var (
	_ Clock = (*fake)(nil)
	_ Mock  = (*fake)(nil)
)

// fake implements the `clock.Clock` interface using a time source
// which can be manipulated using the `clock.Mock` interface.
type fake struct {
	t  time.Time
	mu sync.Mutex
}

// Now returns the current clock time.
func (f *fake) Now() time.Time {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.t
}

// Since returns the time elapsed since the given time.
func (f *fake) Since(t time.Time) time.Duration {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.t.Sub(t)
}

// Sleep pauses the current goroutine for at least the duration d.
// A negative or zero duration causes Sleep to return immediately.
func (f *fake) Sleep(d time.Duration) {
	panic("implement me")
}

// Until returns the duration until the given time.
func (f *fake) Until(t time.Time) time.Duration {
	return t.Sub(f.t)
}

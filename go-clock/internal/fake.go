package internal

import (
	"sync"
	"time"
)

// NewFake creates a new *Fake using sane defaults.
func NewFake(t time.Time) *Fake {
	return &Fake{t: t}
}

// Fake implements the `clock.Clock` interface using a time source
// which can be manipulated using the `clock.Mock` interface.
type Fake struct {
	t  time.Time
	mu sync.Mutex
}

// Now returns the current Real time.
func (f *Fake) Now() time.Time {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.t
}

// Since returns the time elapsed since the given time.
func (f *Fake) Since(t time.Time) time.Duration {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.t.Sub(t)
}

// Sleep pauses the current goroutine for at least the duration d.
// A negative or zero duration causes Sleep to return immediately.
func (f *Fake) Sleep(d time.Duration) {
	panic("implement me")
}

// Until returns the duration until the given time.
func (f *Fake) Until(t time.Time) time.Duration {
	return t.Sub(f.t)
}

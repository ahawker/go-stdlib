package clock

import (
	"github.com/ahawker/go-stdlib/go-clock/internal"
	"time"
)

// Interface Assertions
var (
	_ Clock     = (*internal.Real)(nil)
	_ Clock     = (*internal.Fake)(nil)
	_ FakeClock = (*internal.Fake)(nil)
)

// Type Aliases
type (
	Duration = time.Duration
	Time     = time.Time
	Timer    = time.Timer
)

// Real returns a `clock.Clock` implementation that uses the `time` package
// with the current time.
func Real() Clock {
	return internal.NewReal()
}

// Fake returns a `clock.Clock` implementation that uses a fake implementation.
//
// Cast this to `clock.FakeClock` to access helper functions for controlling
// the underlying time source.
func Fake() FakeClock {
	return internal.NewFake(time.Now())
}

// Clock represents an interface to the functions in the `time` package.
type Clock interface {
	// Now returns the current clock time.
	Now() time.Time
	// Since returns the time elapsed since the given time.
	Since(time.Time) time.Duration
	// Sleep pauses the current goroutine for at least the duration d.
	// A negative or zero duration causes Sleep to return immediately.
	Sleep(time.Duration)
	// Until returns the duration until the given time.
	Until(time.Time) time.Duration
}

// FakeClock represents a `clock.Clock` implementation with of helper functions
// for controlling the underlying time source.
type FakeClock interface {
	Clock
}

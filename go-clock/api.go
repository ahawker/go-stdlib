package clock

import (
	"time"
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
	return &clock{}
}

// Fake returns a `clock.Clock` implementation that uses a fake implementation.
//
// Cast this to `clock.Mock` to access helper functions for controlling
// the underlying time source.
func Fake() Mock {
	return &fake{t: time.Unix(0, 0)}
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

// Mock represents an interface of helper functions for controlling
// the time source of a mocked clock.
type Mock interface {
	Clock
}

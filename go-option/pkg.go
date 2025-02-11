// Package option contains functionality for managing
// functional options for fluent instance creation.
package option

// Opt defines a functional option for type `T`
// that mutates the instance.
type Opt[T any] func(t T) error

// Opts is a slice of `Opt` options.
type Opts[T any] []Opt[T]

// Apply applies all functional options for type `T`
// and returns the value or error if any fail to apply.
func Apply[T any](t *T, options ...Opt[*T]) (*T, error) {
	for _, o := range options {
		if err := o(t); err != nil {
			return t, err
		}
	}
	return t, nil
}

// Make creates a zero value for type `T`, applies all options,
// and returns the value or error if any fail to apply.
func Make[T any](options ...Opt[*T]) (*T, error) {
	return Apply(new(T), options...)
}

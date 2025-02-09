// Package option contains functionality for managing
// functional options for fluent instance creation.
package option

// Ref defines a functional option for reference type `T`
// that mutate the instance.
type Ref[T any] func(t T) error

// Refs is a slice of `Ref` options.
type Refs[T any] []Ref[T]

// Val defines a functional option for value type `T`
// that returns a copy the instance with the option applied.
type Val[T any] func(t T) (T, error)

// Vals is a slice of `Val` options.
type Vals[T any] []Val[T]

// Apply applies all functional options for reference type `T`
// and returns the value or error if any fail to apply.
func Apply[T any](t T, options ...Ref[T]) (T, error) {
	for _, o := range options {
		if err := o(t); err != nil {
			return t, err
		}
	}
	return t, nil
}

// Copy applies all functional builders for value type `T`
// and returns the value or error if any fail to apply.
func Copy[T any](t T, options ...Val[T]) (T, error) {
	var err error
	for _, o := range options {
		if t, err = o(t); err != nil {
			return *new(T), err
		}
	}
	return t, nil
}

// Make creates a zero value for reference type `T`, applies all options,
// and returns the value or error if any fail to apply.
func Make[T any](options ...Ref[T]) (T, error) {
	return Apply(*new(T), options...)
}

// New creates a zero value for value type `T`, applies all options,
// and returns the value or error if any fail to apply.
func New[T any](options ...Val[T]) (T, error) {
	return Copy(*new(T), options...)
}

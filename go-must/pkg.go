package must

import (
	"fmt"
	"reflect"
)

// Must panics if given value is nil.
func Must[T any](t T) T {
	if t == nil {
		panic(fmt.Sprintf("Must[%s] received nil", gt[T]()))
	}
	return t
}

// MustAlias panics if given value cannot be aliased to the type 'T'.
func MustAlias[T any](v any) T {
	t, ok := v.(T)
	if !ok {
		panic(fmt.Sprintf("MustAlias[%s] cannot alias %T", gt[T](), v))
	}
	return t
}

// MustE panics if given error is set otherwise 't' is returned.
func MustE[T any](t T, err error) T {
	if err != nil {
		panic(fmt.Sprintf("MustE[%s] (err=%v)", gt[T](), err))
	}
	return t
}

// MustF panics if the given func returns an error or returns a zero value.
func MustF[T any](fn func() (T, error)) T {
	t, err := fn()
	if err != nil {
		panic(fmt.Sprintf("MustF[%T] returned error %v", gt[T](), err))
	}
	return Must[T](t)
}

// MustZ panics if given value is equal to the zero value of the type.
func MustZ[T any](t T) T {
	if isZero[T](t) {
		panic(fmt.Sprintf("Must[%s] received zero value", gt[T]()))
	}
	return t
}

// isZero returns true if the given value is equal to the
// zero value of the type.
func isZero[T any](t T) bool {
	return reflect.DeepEqual(t, *new(T))
}

// GT returns the name of the generic type.
//
// This is helpful since '%T' in printf doesn't support generic types.
func gt[T any]() string {
	return reflect.TypeFor[T]().String()
}

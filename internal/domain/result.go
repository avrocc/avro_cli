package domain

import "fmt"

// Result represents a success or failure outcome.
type Result[T any] struct {
	value T
	err   error
	ok    bool
}

// Ok creates a successful result.
func Ok[T any](value T) Result[T] {
	return Result[T]{value: value, ok: true}
}

// Fail creates a failed result.
func Fail[T any](err error) Result[T] {
	return Result[T]{err: err, ok: false}
}

// Failf creates a failed result from a formatted string.
func Failf[T any](format string, args ...any) Result[T] {
	return Result[T]{err: fmt.Errorf(format, args...), ok: false}
}

// IsOk returns true if the result is successful.
func (r Result[T]) IsOk() bool { return r.ok }

// Value returns the success value. Panics if result is a failure.
func (r Result[T]) Value() T {
	if !r.ok {
		panic("called Value() on a failed Result")
	}
	return r.value
}

// Err returns the error. Returns nil if result is successful.
func (r Result[T]) Err() error { return r.err }

// ValueOr returns the value if ok, otherwise the fallback.
func (r Result[T]) ValueOr(fallback T) T {
	if r.ok {
		return r.value
	}
	return fallback
}

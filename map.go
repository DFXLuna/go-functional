package gofunctional

import (
	"errors"
	"fmt"
)

var (
	ErrMap = errors.New("error appling map")
)

// Transform provided array ts into a new array with elements of
// type U by applying function f to each element of ts and storing the result
func Map[T any, U any](ts []T, f func(T) U) []U {
	us := make([]U, 0, len(ts))
	for _, t := range ts {
		us = append(us, f(t))
	}
	return us
}

// Same as Map but with the possibility of returning an error
func MapE[T any, U any](ts []T, f func(T) (U, error)) ([]U, error) {
	us := make([]U, 0, len(ts))
	for _, t := range ts {
		u, err := f(t)
		if err != nil {
			return []U{}, fmt.Errorf("%w: %w", ErrApply, err)
		}
		us = append(us, u)
	}
	return us, nil
}

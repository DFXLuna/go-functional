package gofunctional

import (
	"errors"
	"fmt"
)

var (
	ErrFoldr = errors.New("error applying foldr")
)

// Applies a right fold over an array of type U, summing the results into the accumulator of type T by repeatedly applying
// function F to every member of us
func Foldr[T Summable, U any](acc T, f func(acc T, u U) T, us []U) T {
	for _, u := range us {
		acc += f(acc, u)
	}
	return acc
}

// Applies FoldR as above but allows function f to error and returns that error
func FoldrE[T Summable, U any](acc T, f func(acc T, u U) (T, error), us []U) (T, error) {
	for _, u := range us {
		v, err := f(acc, u)
		if err != nil {
			return acc, fmt.Errorf("%w: %w", ErrFoldr, err)
		}
		acc += v
	}
	return acc, nil
}

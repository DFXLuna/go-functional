package gofunctional

import (
	"errors"
	"fmt"
)

var (
	ErrFoldr = errors.New("error applying foldr")
	ErrFoldl = errors.New("error applying foldl")
)

// Applies a right fold over an array of type U, summing the results into the accumulator of type T by repeatedly applying
// function F to every member of us.
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

// Applies a left fold over an array of type U, summing the results into the accumulator of type T by repeatedly applying
// function F to every member of us. Equivalent to iteration from the end of an array to the beginning of an array
func Foldl[T Summable, U any](acc T, f func(acc T, u U) T, us []U) T {
	for i := len(us) - 1; i >= 0; i-- {
		acc += f(acc, us[i])
	}
	return acc
}

// Applies Foldl as above but allows function f to error and returns that error
func FoldlE[T Summable, U any](acc T, f func(acc T, u U) (T, error), us []U) (T, error) {
	for i := len(us) - 1; i >= 0; i-- {
		v, err := f(acc, us[i])
		if err != nil {
			return acc, fmt.Errorf("%w: %w", ErrFoldl, err)
		}
		acc += v

	}
	return acc, nil
}

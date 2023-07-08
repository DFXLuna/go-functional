package gofunctional

import (
	"errors"
	"fmt"
)

var (
	ErrApply = errors.New("error appling apply")
)

func Apply[T any](ts []T, f func(T)) {
	for _, t := range ts {
		f(t)
	}
}

func ApplyE[T any](ts []T, f func(T) error) error {
	for _, t := range ts {
		err := f(t)
		if err != nil {
			return fmt.Errorf("%w: %w", ErrApply, err)
		}
	}
	return nil
}

package test

import (
	"errors"
	"testing"

	gofunctional "github.com/DFXLuna/go-functional"
	"github.com/stretchr/testify/assert"
)

var ErrExpected = errors.New("error is expected")

func TestFoldr(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3}
	sumIfGreater := func(acc int, val int) int {
		if val > acc {
			return val
		}
		return 0
	}
	// Type inference allows skipping the type parameters
	sum := gofunctional.Foldr(0, sumIfGreater, arr)
	assert.Equal(sum, 3, "sum should be 1 + 2")
}

func TestFoldrE(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3}
	sumIfGreater := func(acc int, val int) (int, error) {
		if val > acc {
			return val, nil
		}
		return 0, nil
	}

	sum, err := gofunctional.FoldrE(0, sumIfGreater, arr)
	assert.Nil(err, "should not error with basic test")
	assert.Equal(sum, 3, "basic sum should work the same with foldre")

	errFunc := func(_ int, _ int) (int, error) {
		return 0, ErrExpected
	}

	sum, err = gofunctional.FoldrE(0, errFunc, arr)
	assert.NotNil(err, "should error with errFunc")
	assert.ErrorIs(err, gofunctional.ErrFoldr, "should have foldr error type")
	assert.Zero(sum, "should have errored before any summing")
}

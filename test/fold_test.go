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
	arr := []int{}
	sumIfGreater := func(acc int, val int) int {
		if val > acc {
			return val
		}
		return 0
	}
	testAcc := 456
	sum := gofunctional.Foldr(testAcc, sumIfGreater, arr)
	assert.Equal(testAcc, sum, "sum should not change iterating over empty array")

	arr = []int{1, 2, 3}

	// Type inference allows skipping the type parameters
	sum = gofunctional.Foldr(0, sumIfGreater, arr)
	assert.Equal(3, sum, "sum should be 1 + 2")

	strArr := []string{"a", "b", "c"}
	sumStrFunc := func(_ string, val string) string {
		return val
	}
	sumStr := gofunctional.Foldr("", sumStrFunc, strArr)
	assert.Equal("abc", sumStr, "summing string should make larger string")

	compArr := []complex128{1 + 1i, 2 + 2i, 3 + 3i}
	sumCompFunc := func(_ complex128, val complex128) complex128 {
		return val
	}

	sumComp := gofunctional.Foldr(0+0i, sumCompFunc, compArr)
	assert.Equal(6+6i, sumComp, "summing complex should work")
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
	assert.Equal(3, sum, "basic sum should work the same with foldre")

	errFunc := func(_ int, _ int) (int, error) {
		return 0, ErrExpected
	}

	sum, err = gofunctional.FoldrE(0, errFunc, arr)
	assert.NotNil(err, "should error with errFunc")
	assert.ErrorIs(err, gofunctional.ErrFoldr, "should have foldr error type")
	assert.Zero(sum, "should have errored before any summing")
}

func TestFoldl(t *testing.T) {
	assert := assert.New(t)
	arr := []int{}
	sumIfGreater := func(acc int, val int) int {
		if val > acc {
			return val
		}
		return 0
	}
	testAcc := 456
	sum := gofunctional.Foldl(testAcc, sumIfGreater, arr)
	assert.Equal(testAcc, sum, "sum should not change iterating over empty array")

	arr = []int{1, 2, 4}

	// Type inference allows skipping the type parameters
	sum = gofunctional.Foldl(0, sumIfGreater, arr)
	assert.Equal(4, sum, "sum should be 4")

	strArr := []string{"a", "b", "c"}
	sumStrFunc := func(_ string, val string) string {
		return val
	}
	sumStr := gofunctional.Foldl("", sumStrFunc, strArr)
	assert.Equal("cba", sumStr, "summing string should make larger string")

	compArr := []complex128{1 + 1i, 2 + 2i, 3 + 3i}
	sumIfGreaterCompFunc := func(acc complex128, val complex128) complex128 {
		if real(val) > real(acc) {
			return val
		}
		return 0 + 0i
	}

	sumComp := gofunctional.Foldl(3+3i, sumIfGreaterCompFunc, compArr)
	assert.Equal(3+3i, sumComp, "summing complex should work")
}

func TestFoldlE(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 4}
	sumIfGreater := func(acc int, val int) (int, error) {
		if val > acc {
			return val, nil
		}
		return 0, nil
	}

	sum, err := gofunctional.FoldlE(0, sumIfGreater, arr)
	assert.Nil(err, "should not error with basic test")
	assert.Equal(4, sum, "basic sum should work the same with foldre")

	errFunc := func(_ int, _ int) (int, error) {
		return 0, ErrExpected
	}

	sum, err = gofunctional.FoldlE(0, errFunc, arr)
	assert.NotNil(err, "should error with errFunc")
	assert.ErrorIs(err, gofunctional.ErrFoldl, "should have foldr error type")
	assert.Zero(sum, "should have errored before any summing")
}

package test

import (
	"testing"

	gofunctional "github.com/DFXLuna/go-functional"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	assert := assert.New(t)
	arr := []int{}
	ret := func(i int) int {
		return i
	}

	testArr := gofunctional.Map(arr, ret)
	assert.Equal(arr, testArr, "testArr should not change iterating over empty array")

	arr = []int{1, 2, 3}

	testArr = gofunctional.Map(arr, ret)
	assert.Equal(arr, testArr, "testArr should be the same as arr")

	strArr := []string{"a", "b", "c"}
	addY := func(s string) string {
		return s + "Y"
	}

	yStrs := gofunctional.Map(strArr, addY)
	assert.NotEqual(strArr, yStrs, "yStr should be different from starting array")
	assert.Equal([]string{"aY", "bY", "cY"}, yStrs)

	str := "def"
	setX := func(r rune) rune {
		return rune('X')
	}
	xRs := gofunctional.Map[rune]([]rune(str), setX)
	xStr := string(xRs)
	assert.NotEqual(str, xStr, "xStr should be different from starting string")
	assert.Equal("XXX", xStr)
}

func TestMapE(t *testing.T) {
	assert := assert.New(t)
	arr := []int{}
	fErr := func(_ int) (int, error) {
		return 0, ErrExpected
	}

	testArr, err := gofunctional.MapE(arr, fErr)
	assert.Nil(err, "should not error with empty array")
	assert.Empty(testArr, "test array should be empty")

	arr = []int{1, 2, 3}
	testArr, err = gofunctional.MapE(arr, fErr)
	assert.NotNil(err, "should not error with non-empty array")
	assert.Empty(testArr, "test array should be empty")

	f := func(i int) (int, error) {
		return i, nil
	}

	testArr, err = gofunctional.MapE(arr, f)
	assert.Nil(err, "should not error applying map")
	assert.Equal(arr, testArr, "arrays should be equal")
}

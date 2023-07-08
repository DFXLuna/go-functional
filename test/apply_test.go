package test

import (
	"testing"

	gofunctional "github.com/DFXLuna/go-functional"
	"github.com/stretchr/testify/assert"
)

func TestApply(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3}
	funcCalls := 0
	f := func(_ int) {
		funcCalls++
	}

	gofunctional.Apply(arr, f)
	assert.Equal(len(arr), funcCalls, "f should get called 3 times")

	arr = []int{}
	funcCalls = 0
	gofunctional.Apply(arr, f)
	assert.Equal(len(arr), funcCalls, "f should get called 0 times")
}

func TestApplyE(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3}
	funcCalls := 0
	f := func(_ int) error {
		funcCalls++
		return nil
	}

	err := gofunctional.ApplyE(arr, f)
	assert.Nil(err, "should not error")
	assert.Equal(len(arr), funcCalls, "f should get called 3 times")

	arr = []int{}
	funcCalls = 0
	gofunctional.ApplyE(arr, f)
	assert.Nil(err, "should not error")
	assert.Equal(len(arr), funcCalls, "f should get called 0 times")

	f = func(_ int) error {
		return ErrExpected
	}

	err = gofunctional.ApplyE(arr, f)
	assert.Nil(err, "should not error with empty array")

	arr = []int{1, 2, 3}
	err = gofunctional.ApplyE(arr, f)
	assert.NotNil(err, "should error")
}

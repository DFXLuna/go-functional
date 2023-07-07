package gofunctional

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Float interface {
	~float32 | ~float64
}

type Complex interface {
	~complex64 | ~complex128
}

// Represents types for which operator + (and implicitly, operator -) are defined.
// Useful for generics
type Summable interface {
	Integer | Float | Complex | ~string
}

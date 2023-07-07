package gofunctional

// Pair is a convenience type for holding two related elements
type Pair[T any, U any] struct {
	First  T
	Second U
}

package util

func Foldr[T Summable, U any](acc T, f func(acc T, u U) T, us []U) T {
	for _, u := range us {
		acc += f(acc, u)
	}
	return acc
}

## go-functional
Golang lacks a lot of functional stuff and I'm tired of reading Stackoverflow answers that say "These are trivial to implement".



## Constraints
- Contains some type constraints for use in generic type parameters
- Example(Summable):
    ```go
    func Foldr[T Summable, U any](acc T, f func(acc T, u U) T, us []U) T ...
    ```


## Fold
- Contains various generic fold functions.
    - These are typically equivalent to different styles of loops but are a little more convenient and easier to compose.
- Example
    ```go
    arr := []int{1, 2, 3}
    sumIfGreater := func(acc int, val int) int {
        if val > acc {
            return val
        }
        return 0
    }
    // Type inference allows skipping the type parameters
    sum := gofunctional.Foldr(0, sumIfGreater, arr)

    ```
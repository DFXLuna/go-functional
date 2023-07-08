## go-functional
Golang lacks a lot of functional stuff and I'm tired of reading Stackoverflow answers that say "These are trivial to implement".

## Errors
- Most functions have two variants, `Func` and `FuncE`. `Func` is the basic function that works as expected, `FuncE` is the equivalent function that can also handle and return errors.


## Functions:
- Foldl, Foldr
- Map
- Apply

## Types:
- Pair
    - Convenience type for storing two variables of arbitrary type
    - Slighty more structure than using an anonymous struct

## Constraints
- Some type constraints for use in generic type parameters
- Example(Summable):
    ```go
    func Foldr[T Summable, U any](acc T, f func(acc T, u U) T, us []U) T ...
    ```



package gmath 
import (
    
    "golang.org/x/exp/constraints"
)
type Number interface {
    constraints.Integer | constraints.Float
}

func Max[T Number](a, b T) T {
    if a > b {
        return a
    }
    return b
}
func Min[T Number](a, b T) T {
    return -Max(-a, -b)
}

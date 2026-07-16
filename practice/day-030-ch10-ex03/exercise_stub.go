package adder

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

func Add[T Number](a, b T) T {
	var zero T
	return zero
}

package main

type ValidTypes interface {
	~int | ~float64
}

func Doubler[T ValidTypes](value T) T {
	var zero T
	return zero
}

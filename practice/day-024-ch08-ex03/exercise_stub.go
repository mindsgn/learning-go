package main

type node[T comparable] struct {
	value T
	next  *node[T]
}

type List[T comparable] struct {
	head *node[T]
}

func (l *List[T]) Add(value T)             {}
func (l *List[T]) Insert(value T, pos int) {}
func (l *List[T]) Index(value T) int       { return -1 }

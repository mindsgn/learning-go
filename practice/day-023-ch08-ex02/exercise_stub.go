package main

import "fmt"

type Printable interface {
	fmt.Stringer
}

type PrintInt int
type PrintFloat float64

func (p PrintInt) String() string   { return "" }
func (p PrintFloat) String() string { return "" }

func PrintIt[T Printable](value T) {}

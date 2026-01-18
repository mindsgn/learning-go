package main

import (
	"fmt"
	"math"
)

func NumberAssignment() (byte, int32, uint64) {
	var b byte = math.MaxInt8
	var smallI int32 = math.MaxInt32
	var bigI uint64 = math.MaxUint64

	return b + 1, smallI + 1, bigI + 1
}

func AssignConstant() (float32, int) {
	const value = 10
	const i int = value
	const f float32 = value

	return f, i
}

func AssignFloat() (float32, int) {
	var i int = 20
	var f float32

	f = float32(i)

	return f, i
}

func main() {
	b, smallI, bigI := NumberAssignment()
	fmt.Println(b, smallI, bigI)
}

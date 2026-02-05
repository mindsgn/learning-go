package main

import "fmt"

func Arrays() {
	// arrays
	var x [3]int
	var y = [5]int{10, 1, 2, 3, 4}
	var z = [12]int{1, 5: 4, 6, 10: 100, 15}
	var a = [...]int{10: 12}
	var b [3][5]int
	var c [10]int

	//Arrays
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(a)
	fmt.Println(b)

	c[0] = 10
	fmt.Println(c)
	fmt.Println(len(c))
}

func Slices() {
	var a = []int{10, 1}
	var b = []int{10: 100}
	var c []int
	fmt.Println(c)
	c = append(c, 10)
	fmt.Println(c)
	c = append(c, 20)
	fmt.Println(c)
	c = append(c, 20, 100, 10)
	fmt.Println(c)
	c = append(c, c...)
	fmt.Println(c)
	fmt.Println(a)
	fmt.Println(b)

}

func Capacity() {
	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))
}

func Make() {
	x := make([]int, 5)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
}

func EmptySlice() {
	s := []string{"one", "two", "three"}
	fmt.Println(s, len(s))
	clear(s)
	fmt.Println(s, len(s))
}

func main() {
	// Arrays()
	//Slices()
	//Capacity()
	//Make()
	EmptySlice()
}

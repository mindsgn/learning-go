package main

import "fmt"

func main() {
	if 7%2 == 0 {
		println("7 is even")
	} else {
		println("7 is odd")
	}

	if 8%4 == 0 {
		println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		println("either 8 or 7 is even")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has 2 digits or more")
	}
}

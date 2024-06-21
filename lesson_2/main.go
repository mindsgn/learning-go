package main

import "fmt"

func AddString(word string) string {
	return word + "lang"
}

func AddStringAndNumber(number uint) string {
	sum := 1 + number
	return "1+1=" + fmt.Sprint(sum)
}

func main() {
	fmt.Println(AddString("coffee"))
	fmt.Println(AddStringAndNumber(1))
}

package main

import (
	"fmt"
	"testing"
)

func TestAddString(t *testing.T) {
	want := "golang"
	msg := AddString("go")

	if want != msg {
		fmt.Println(`expected output: "golang"`)
		fmt.Println(`got:`, msg)
		t.Fatal("Failed")
	}
}

func TestADDStringandNumber(t *testing.T) {
	want := "1+1=2"
	msg := AddStringAndNumber(1)
	if want != msg {
		fmt.Println(`expected output:"1+1=2"`)
		fmt.Println(`got:`, msg)
		t.Fatal("Failed")
	}
}

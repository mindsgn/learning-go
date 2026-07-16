package main

import (
	"os"

	practice "learninggo/practice"
)

func main() {
	os.Exit(practice.RunCLI(os.Stdout, os.Stderr, os.Args[1:]))
}

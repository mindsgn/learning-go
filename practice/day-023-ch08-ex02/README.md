
# Day 023 - Chapter 8 Exercise 2

**Chapter:** Generics

## Original Prompt

2. Define a generic interface called Printable that matches a type that implements fmt.Stringer and has an underlying type of int or float64. Define types that meet this interface. Write a function that takes in a Printable and prints its value to the screen using fmt.Println.

## Acceptance

Define the `Printable` constraint, concrete printable types, and `PrintIt` so the output matches the reference behavior.

Run `go test` inside this folder when you want feedback.

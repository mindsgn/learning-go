
# Day 035 - Chapter 12 Exercise 2

**Chapter:** Concurrency in Go

## Original Prompt

2. Create a function that launches two goroutines. Each goroutine writes 10 num‐ bers to its own channel. Use a for-select loop to read from both channels, printing out the number and the goroutine that wrote the value. Make sure that your function exits after all values are read and that none of your goroutines leak.

## Acceptance

The acceptance test checks the values printed by the `for-select` loop without caring about goroutine scheduling order.

Run `go test` inside this folder when you want feedback.

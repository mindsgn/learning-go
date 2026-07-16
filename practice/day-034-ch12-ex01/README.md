
# Day 034 - Chapter 12 Exercise 1

**Chapter:** Concurrency in Go

## Original Prompt

1. Create a function that launches three goroutines that communicate using a channel. The first two goroutines each write 10 numbers to the channel. The third goroutine reads all the numbers from the channel and prints them out. The function should exit when all values have been printed out. Make sure that none of the goroutines leak. You can create additional goroutines if needed.

## Acceptance

You can implement this in any goroutine order you like. The acceptance test sorts the printed numbers and checks the expected multiset.

Run `go test` inside this folder when you want feedback.


# Day 018 - Chapter 6 Exercise 3

**Chapter:** Pointers

## Original Prompt

3. Write a program that builds a []Person with 10,000,000 entries (they could all be the same names and ages). See how long it takes to run. Change the value of GOGC and see how that affects the time it takes for the program to complete. Set the environment variable GODEBUG=gctrace=1 to see when garbage collections happen and see how changing GOGC changes the number of garbage collections. What happens if you create the slice with a capacity of 10,000,000?

## Acceptance

The original exercise is performance-focused, so this offline pack adds a helper contract: create `buildPeople(count int) []Person` and use it from `main` for the large allocation experiment.

Run `go test` inside this folder when you want feedback.

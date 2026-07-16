
# Day 036 - Chapter 12 Exercise 3

**Chapter:** Concurrency in Go

## Original Prompt

3. Write a function that builds a map[int]float64 where the keys are the numbers from 0 (inclusive) to 100,000 (exclusive) and the values are the square roots of those numbers (use the math.Sqrt function to calculate square roots). Use sync.OnceValue to generate a function that caches the map returned by this function and use the cached value to look up square roots for every 1,000th number from 0 to 100,000.

## Acceptance

Expose the expensive work through `buildSquareRootMap()` so the acceptance test can validate the generated data directly.

Run `go test` inside this folder when you want feedback.

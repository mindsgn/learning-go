
# Day 040 - Chapter 14 Exercise 1

**Chapter:** The Context

## Original Prompt

1. Create a middleware-generating function that creates a context with a time‐ out. The function should have one parameter, which is the number of milli‐ seconds that a request is allowed to run. It should return a func(http.Handler) http.Handler.

## Acceptance

Implement `Timeout(ms int) func(http.Handler) http.Handler`. The acceptance test checks that the middleware adds a context deadline with the requested timeout.

Run `go test` inside this folder when you want feedback.

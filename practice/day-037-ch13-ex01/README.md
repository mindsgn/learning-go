
# Day 037 - Chapter 13 Exercise 1

**Chapter:** The Standard Library

## Original Prompt

1. Write a small web server that returns the current time in RFC 3339 format when you send it a GET command. You can use a third-party module if you’d like.

## Acceptance

The acceptance test expects a helper named `createServeMux()` that returns an `http.Handler` responding with the current time in RFC 3339 format. The optional chi version is still available in `reference/` for inspiration.

Run `go test` inside this folder when you want feedback.

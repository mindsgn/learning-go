
# Day 038 - Chapter 13 Exercise 2

**Chapter:** The Standard Library

## Original Prompt

2. Write a small middleware component that uses JSON structured logging to log the IP address of each incoming request to your web server.

## Acceptance

Expose `createChiRouter(logger *slog.Logger)` so the test can verify both the time response and the structured IP logging middleware.

Run `go test` inside this folder when you want feedback.

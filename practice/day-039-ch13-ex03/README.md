
# Day 039 - Chapter 13 Exercise 3

**Chapter:** The Standard Library

## Original Prompt

3. Add the ability to return the time as JSON. Use the Accept header to control whether JSON or text is returned (default to text). The JSON should be struc‐ tured as follows: { "day_of_week": "Monday", "day_of_month": 10, "month": "April", "year": 2023, "hour": 20, "minute": 15, "second": 20 }

## Acceptance

Implement `buildText`, `buildJSON`, and `createChiRouter(logger *slog.Logger)` so the handler can switch between text and JSON based on the `Accept` header.

Run `go test` inside this folder when you want feedback.

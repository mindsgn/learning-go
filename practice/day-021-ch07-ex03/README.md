
# Day 021 - Chapter 7 Exercise 3

**Chapter:** Types, Methods, and Interfaces

## Original Prompt

3. Define an interface called Ranker that has a single method called Ranking that returns a slice of strings. Write a function called RankPrinter with two parame‐ ters, the first of type Ranker and the second of type io.Writer. Use the io.Write String function to write the values returned by Ranker to the io.Writer, with a newline separating each result. Call this function from main.

## Acceptance

Define `Ranker` and `RankPrinter` so the printer can write one team name per line to an `io.Writer`.

Run `go test` inside this folder when you want feedback.

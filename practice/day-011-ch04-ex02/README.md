
# Day 011 - Chapter 4 Exercise 2

**Chapter:** Blocks, Shadows, and Control Structures

## Original Prompt

2. Loop over the slice you created in exercise 1. For each value in the slice, apply the following rules: a. If the value is divisible by 2, print “Two!” b. If the value is divisible by 3, print “Three!” c. IIf the value is divisible by 2 and 3, print “Six!”. Don’t print anything else. d. Otherwise, print “Never mind”.

## Acceptance

The original loop is easier to test offline if you extract the classification logic into `describeNumber(int) string`. The function should return `Two!`, `Three!`, `Six!`, or `Never mind`.

Run `go test` inside this folder when you want feedback.

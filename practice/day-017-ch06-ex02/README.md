
# Day 017 - Chapter 6 Exercise 2

**Chapter:** Pointers

## Original Prompt

2. Write two functions. The UpdateSlice function takes in a []string and a string. It sets the last position in the passed-in slice to the passed-in string. At the end of UpdateSlice, print the slice after making the change. The GrowSlice function also takes in a []string and a string. It appends the string onto the slice. At the end of GrowSlice, print the slice after making the change. Call these functions from main. Print out the slice before each function is called and after each function is called. Do you understand why some changes are visible in main and why some changes are not?

## Acceptance

Implement `UpdateSlice` and `GrowSlice` so the tests can observe the visible slice mutation behavior from the caller.

Run `go test` inside this folder when you want feedback.


# Day 016 - Chapter 6 Exercise 1

**Chapter:** Pointers

## Original Prompt

1. Create a struct named Person with three fields: FirstName and LastName of type string and Age of type int. Write a function called MakePerson that takes in firstName, lastName, and age and returns a Person. Write a second function MakePersonPointer that takes in firstName, lastName, and age and returns a *Person. Call both from main. Compile your program with go build -gcflags="-m". This both compiles your code and prints out which values escape to the heap. Are you surprised about what escapes?

## Acceptance

Define `Person`, `MakePerson`, and `MakePersonPointer` exactly as the book asks so the tests can inspect both the value and pointer versions.

Run `go test` inside this folder when you want feedback.

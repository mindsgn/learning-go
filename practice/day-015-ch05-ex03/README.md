
# Day 015 - Chapter 5 Exercise 3

**Chapter:** Functions

## Original Prompt

3. Write a function called prefixer that has an input parameter of type string and returns a function that has an input parameter of type string and returns a string. The returned function should prefix its input with the string passed into prefixer. Use the following main function to test prefixer: func main() { helloPrefix := prefixer("Hello") fmt.Println(helloPrefix("Bob")) // should print Hello Bob fmt.Println(helloPrefix("Maria")) // should print Hello Maria }

## Acceptance

Expose a function named `prefixer(prefix string) func(string) string` and make the returned closure prefix incoming strings.

Run `go test` inside this folder when you want feedback.

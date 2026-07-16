
# Day 046 - Chapter 16 Exercise 1

**Chapter:** Reflection, Unsafe, and cgo

## Original Prompt

1. Use reflection to create a simple minimal string-length validator for struct fields. Write a ValidateStringLength function that takes in a struct and returns an error if one or more of the fields is a string, has a struct tag called minStrlen, and the length of the value in the field is less than the value specified in the struct tag. Nonstring fields and string fields that don’t have the minStrlen struct tag are ignored. Use errors.Join to report all invalid fields. Be sure to validate that a struct was passed in. Return nil if all fields are of the proper length.

## Acceptance

Implement `ValidateStringLength(any) error` so it uses reflection and `minStrlen` tags to reject short string fields.

Run `go test` inside this folder when you want feedback.

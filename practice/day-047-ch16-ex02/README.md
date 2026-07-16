
# Day 047 - Chapter 16 Exercise 2

**Chapter:** Reflection, Unsafe, and cgo

## Original Prompt

2. Use unsafe.Sizeof and unsafe.Offsetof to print out the size and offsets for the OrderInfo struct defined in ch16/tree/main/sample_code/orders. Create a new type, SmallOrderInfo, that has the same fields, but reordered to use as little memory as possible.

## Acceptance

Define both `OrderInfo` and `SmallOrderInfo`. The acceptance test checks the original layout and verifies that your reordered version uses less memory.

Run `go test` inside this folder when you want feedback.

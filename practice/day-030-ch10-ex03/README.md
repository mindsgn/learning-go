
# Day 030 - Chapter 10 Exercise 3

**Chapter:** Modules, Packages, and Imports

## Original Prompt

3. Change Add to make it generic. Import the golang.org/x/exp/constraints package. Combine the Integer and Float types in that package to create an interface called Number. Rewrite Add to take in two parameters of type Number and return a value of type Number. Version your module again. Because this is a backward-breaking change, this should be v2.0.0 of your module.

## Acceptance

This offline version uses a local `VERSION` file instead of public git tags. Make `Add` generic, define `Number`, and set `VERSION` to `v2.0.0`.

Run `go test` inside this folder when you want feedback.

# Learning Go Practice Pack

This folder turns the chapter-end exercises from *Learning Go, 2nd Edition* into a terminal-first offline practice pack.

Each exercise lives in its own folder:

- `day-001-ch01-ex01`
- `day-002-ch01-ex02`
- ...
- `day-048-ch16-ex03`

Each folder includes:

- the original exercise text from the book
- either an official reference solution or a local offline testing contract
- a `go test` acceptance test so you can check your work in the terminal

Some book prompts were normalized for offline use. This mostly affects exercises that originally depended on public repositories, downloadable sample apps, playground sharing, or open-ended manual exploration.

## Commands

List the full catalog:

```sh
make list
```

Run one exercise:

```sh
make day DAY=day-004-ch02-ex01
```

Run from inside a single exercise folder:

```sh
go test
```

Run the whole pack:

```sh
go test ./...
```

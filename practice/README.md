# Learning Go Practice Pack

This folder turns the chapter-end exercises from *Learning Go, 2nd Edition* into a terminal-first offline practice pack.

It now also includes a small SQLite-backed terminal app so you can track progress locally.

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

List the exercise folders:

```sh
make list
```

Show progress from the terminal app:

```sh
make progress
```

Show the next exercise:

```sh
make next
```

Start a day and record it in SQLite:

```sh
make start DAY=4
```

Mark a day done:

```sh
make done DAY=4
```

Run tests for a day and record pass/fail:

```sh
make test-day DAY=day-004-ch02-ex01
```

Run the CLI directly:

```sh
go run ./cmd/practice help
```

Run one exercise without the tracker:

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

## SQLite Tracking

The CLI stores progress in `practice/.practice.db`.

You can run it from the `practice/` root or from inside an individual day folder. Useful commands:

```sh
go run ./cmd/practice list
go run ./cmd/practice status
go run ./cmd/practice next
go run ./cmd/practice show 4
go run ./cmd/practice start 4
go run ./cmd/practice done 4
go run ./cmd/practice test 4
```

If you want the database somewhere else, set `PRACTICE_DB=/your/path.db`.

# Learning Go Practice Workspace

This repository is a local practice workspace for *Learning Go, 2nd Edition* by Jon Bodner.

The main deliverable is the offline exercise pack in [practice/README.md](/Users/sibongiseni/projects/learning-go/practice/README.md). It extracts the chapter-end exercises into individual day-by-day folders so you can practice in the terminal and check your work with tests.

## What Is Here

- `practice/`
  A 48-day offline practice pack with one folder per exercise.
- `practice/SCHEDULE.md`
  A simple day-by-day roadmap.
- `practice/catalog.json`
  A machine-readable list of all extracted exercises.
- `Bodner J Learning Go An Idiomatic ApproachReal-world Go Programming 2ed 2024.pdf`
  The source book PDF used for extraction.

## Quick Start

Move into the practice pack:

```sh
cd practice
```

List all exercise folders:

```sh
make list
```

Run one day:

```sh
make day DAY=day-004-ch02-ex01
```

Or run tests directly inside a specific folder:

```sh
cd day-004-ch02-ex01
go test
```

## How The Pack Works

Each day folder includes:

- the original exercise prompt
- a local acceptance test
- either an official reference solution or an offline testing contract

Some exercises were adapted slightly so they still work offline. This mostly affects prompts that originally depended on public repositories, the Go Playground, downloadable sample apps, or manual exploration.

## Recommended Flow

1. Pick the next folder from `practice/SCHEDULE.md`.
2. Read that folder's `README.md`.
3. Implement your solution in the exercise folder.
4. Run `go test`.
5. If you get stuck, compare with the local `reference/` folder when one exists.

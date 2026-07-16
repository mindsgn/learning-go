.DEFAULT_GOAL := list

.PHONY: list test day cli progress next start done test-day

list:
	@find . -maxdepth 1 -mindepth 1 -type d -name 'day-*' | sort

test:
ifdef DAY
	go test ./$(DAY)
else
	go test ./...
endif

day: test

cli:
	go run ./cmd/practice $(CMD)

progress:
	go run ./cmd/practice status

next:
	go run ./cmd/practice next

start:
	go run ./cmd/practice start $(DAY)

done:
	go run ./cmd/practice done $(DAY)

test-day:
	go run ./cmd/practice test $(DAY)

.DEFAULT_GOAL := list

.PHONY: list test day

list:
	@find . -maxdepth 1 -mindepth 1 -type d -name 'day-*' | sort

test:
ifdef DAY
	go test ./$(DAY)
else
	go test ./...
endif

day: test

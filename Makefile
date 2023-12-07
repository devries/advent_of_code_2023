GOFILES := $(filter-out run.go, $(wildcard **/*.go))
TEMPLATES := $(wildcard templates/*.tmpl)
GO ?= go
.PHONY: run runall test build clean start help

.DEFAULT_GOAL := build

run.go: $(GOFILES) $(TEMPLATES)
	$(GO) generate

run: main.go run.go ## run the most recently edited day
	$(GO) run .

runall: main.go run.go ## Run all days
	$(GO) run . -a

test: ## Run all tests
	$(GO) test -cover ./utils
	$(GO) test -cover ./day*

aoc_run: main.go run.go
	$(GO) build -o aoc_run .

build: aoc_run ## Build binary executable aoc_run

clean: ## Clean run.go and aoc_run
	- rm run.go
	- rm aoc_run

help: ## Show this help
	@echo "These are the make commands for the solutions to this Advent of Code repository.\n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

day%p1:
	$(GO) run ./start -d $(shell echo $* | sed 's/^0*//')

day%p2: day%p1
	mkdir $@
	- sed -E 's/^package day(.*)p1$$/package day\1p2/' day$(*)p1/solution.go > day$(*)p2/solution.go
	- sed -E 's/^package day(.*)p1$$/package day\1p2/' day$(*)p1/solution_test.go > day$(*)p2/solution_test.go

start: day$(shell date +%d)p1  ## Start today

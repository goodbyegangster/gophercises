.DEFAULT_GOAL := help

.PHONY: help
help:
	@printf "%-30s %-60s\n" "[Sub command]" "[Description]"
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "%-30s %-60s\n", $$1, $$2}'

.PHONY: ex01
ex01: ## ex01
	@go build -o ./bin/ ./ex01/quiz-master/exercise
	./bin/ex01 -csv ./ex01/quiz-master/exercise/problems.csv -limit 10 -shuffle true

.PHONY: ex02
ex02: ## ex02
	@go build -o ./bin/ ./ex02/urlshort-master/exercise
	./bin/ex02

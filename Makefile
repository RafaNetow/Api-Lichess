PROJECT_NAME := rafaProject

-include .env
export 

.PHONY: run
run:
    go run ./cmd/myapp.go
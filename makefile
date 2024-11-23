.DEFAULT_GOAL := run

.PHONY:fmt vet build run
fmt:
	go fmt ./main.go
vet: fmt
	go vet ./main.go
build: vet
	go build -o example.com/loan ./main.go
run: build
	./example.com/loan
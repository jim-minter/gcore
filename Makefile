gcore:
	go build -ldflags '-extldflags -static' ./cmd/gcore

.PHONY: gcore

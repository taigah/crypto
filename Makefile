.PHONY: build
build:
	go build -o crypto cmd/crypto/main.go

.PHONY: clean
clean:
	rm crypto

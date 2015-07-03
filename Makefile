build:
	@go build main.go

install:
	@go get .

.PHONY: install build

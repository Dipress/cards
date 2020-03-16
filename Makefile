SHELL := /bin/sh

test:

	go test -v -race `go list ./...`

cover:

	go test --race `go list ./... | grep -v /vendor | grep -v /cmd/cards ` -coverprofile cover.out.tmp && \
	cat cover.out.tmp > cover.out && \
	go tool cover -func cover.out && \
	rm cover.out.tmp && \
	rm cover.out

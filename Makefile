all: test

test:
	# go vet ./...
	go test ./...

bench:
	go test -bench=.
	cd math32 && go test -bench=.

.PHONY: test bench

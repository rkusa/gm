all: test

test:
	# go vet ./...
	go test ./...

bench:
	cd scalar && go test -bench=.
	cd simd && go test -bench=.

.PHONY: test bench

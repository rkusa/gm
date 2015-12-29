all: test

test:
	# go vet ./...
	go test ./...

bench:
	@echo "# math32"
	@cd math32 && go test -bench=.
	@echo "# vec3"
	@cd vec3 && go test -bench=.
	@echo "# vec4"
	@cd vec4 && go test -bench=.
	@echo "# mat3"
	@cd mat3 && go test -bench=.
	@echo "# mat4"
	@cd mat4 && go test -bench=.

.PHONY: test bench

.PHONY: all
all: db/seed_grids.sql

db/seed_grids.sql:
	go run ./cmd/seed/main.go $@

.PHONY: test
test:
	go test -v ./pkg/...
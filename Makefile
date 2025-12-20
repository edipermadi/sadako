.PHONY: all
all: db/seed.sql

db/seed.sql:
	go run ./cmd/seed/main.go $@

.PHONY: test
test:
	go test -v ./pkg/...
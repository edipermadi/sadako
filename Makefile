.PHONY: all
all: db/seed.sql.gz

db/seed.sql.gz:
	go run ./cmd/seed/main.go $@

.PHONY: test
test:
	go test -v ./pkg/...
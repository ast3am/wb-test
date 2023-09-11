.PHONY: all run test

all: test run

run:
	docker-compose up -d
	go run ./cmd/main.go

test:
	docker-compose -f ./test/docker-compose.yml down -v
	docker-compose -f ./test/docker-compose.yml up -d
	go test -count=1 ./api/...
	go test -count=1 ./internal/...
	docker-compose -f ./test/docker-compose.yml down -v
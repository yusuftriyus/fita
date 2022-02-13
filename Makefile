test:
	go test ./...

build:
	go build -o ./bin/checkout_service ./cmd/graphql

run:
	./bin/checkout_service


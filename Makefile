build:
	go build -o app ./cmd

test:
	go test ./...

docker-build:
	docker build -t user-service .

run:
	go run ./cmd
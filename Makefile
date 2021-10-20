

init:
	go run github.com/99designs/gqlgen init

generate:
	go run github.com/99designs/gqlgen

tags:
	go run ./models/model_tags/model_tags.go

run-server-dev:
	go run main.go

test:
	go test -v ./...

clean:
	go clean -cache
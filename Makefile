.PHONY: protogen

build:
	go build -o bin/main

run:
	./bin/main

run-watch:
	./bin/air

protogen:
	protoc --go_out=:. --go-grpc_out=. src/proto/*.proto

build:
	go build -o bin/main cmd/main.go

start:
	./bin/main

run:
	go run cmd/main.go

test:
	go test

clean:
	rm -rf bin

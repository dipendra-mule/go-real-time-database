build: 
	go build -o bin/go-real-time-database cmd/main.go

run:
	./bin/go-real-time-database

test: 
	go test -v ./...
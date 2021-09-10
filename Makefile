hello:
	echo "Hello"

build:
	go build -o bin/main cmd/main.go

run:
	go run cmd/main.go

migrations up:
	./migrate.linux-amd64 -database cassandra://${DB_HOST}:${DB_PORT}/${DB_KEYSPACE}?x-multi-statement=true -path pkg/database/migrations up

migrations down:
	./migrate.linux-amd64 -database cassandra://${DB_HOST}:${DB_PORT}/${DB_KEYSPACE}?x-multi-statement=true -path pkg/database/migrations down

compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 cmd/main.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 cmd/main.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 cmd/main.go

all: hello build
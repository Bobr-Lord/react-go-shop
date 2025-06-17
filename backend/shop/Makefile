run:
	go run cmd/main.go

up:
	migrate -path ./migrations -database "postgres://user:secret@localhost:5432/postgres?sslmode=disable" up

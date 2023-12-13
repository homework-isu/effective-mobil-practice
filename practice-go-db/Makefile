run: build run-app

migrate:
	@go run ./cmd/migrator/main.go
	
build:
	@go build -o main.exe ./cmd/app/main.go

run-app:
	@main.exe

docker-up:
	@docker-compose up -d

db-connect:
	@docker exec -it test-storage psql -U test-user -d test
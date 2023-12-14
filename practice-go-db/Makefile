run: build run-app

migrate-up:
	@go run ./cmd/migrator/main.go --up=1

migrate-down:
	@go run ./cmd/migrator/main.go --down=1

migrate:
	@go run ./cmd/migrator/main.go -actual
	
build:
	@go build -o main.exe ./cmd/app/main.go

run-app:
	@main.exe

docker-up:
	@docker-compose up -d

db-connect:
	@docker exec -it test-storage psql -U test-user -d test
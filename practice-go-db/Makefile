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

redis-connect:
	@docker exec -it redis redis-cli

mock-gen:
	@go get github.com/golang/mock/gomock
	@mockgen -source=internal/core/ports/repository/department.go -destination=pkg/mocks/repository/department_mock.go
	@mockgen -source=internal/http/service/department.go -destination=pkg/mocks/service/department_mock.go

run-test-env:
	@docker-compose -f test_env/docker-compose.test.yaml up  -d 

db-test-connect:
	@docker exec -it test-env-storage psql -U test-user -d test_db
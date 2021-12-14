SERVICE_NAME = article-app

TEST_FLAGS = ./...

DB_URL = 'postgres://postgres:postgres@0.0.0.0:5432/postgres?sslmode=disable'

build:
	docker-compose build $(SERVICE_NAME)	
run: 
	docker-compose up $(SERVICE_NAME)	

test:
	go test -v $(TEST_FLAGS)

migrate:
	migrate -path ./schema -database $(DB_URL) up

clean:
	rm -rf article-app

.PHONY: build, run, migrate, test, clean
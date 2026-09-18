postgres:
	docker run --name school-postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=school_db -p 5434:5432 -d postgres:15-alpine

createdb:
	docker exec -it school-postgres createdb --username=postgres --owner=postgres school_db

dropdb:
	docker exec -it school-postgres dropdb school_db

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5434/school_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5434/school_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server:
	go run main.go

down:
	docker compose down 

up:
	docker compose up -d 

build:
	docker build -t school-api:latest .
.PHONY: postgres createdb dropdb migrateup migratedown sqlc server down up build
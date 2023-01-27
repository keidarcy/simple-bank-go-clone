DB_NAME = "simple_bank"

postgres:
	docker run -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=root --name=pg -d -p 5432:5432 postgres:14-alpine

createdb:
	# docker exec -it postgres psql -U postgres -c "CREATE DATABASE $(DB_NAME);"
	docker exec -it pg createdb --username=root --owner=root $(DB_NAME)

dropdb:
	# docker exec -it postgres psql -U postgres -c "DROP DATABASE $(DB_NAME);"
	docker exec -it pg dropdb ${DB_NAME}

migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc
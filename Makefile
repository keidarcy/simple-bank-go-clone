DB_NAME = "simple_bank"

postgres:
	docker run -e POSTGRES_PASSWORD=password -e POSTGRES_USER=postgres --name=pg -d -p 5432:5432 postgres:14-alpine

createdb:
	# docker exec -it pg psql -U postgres -c "CREATE DATABASE $(DB_NAME);"
	docker exec -it pg createdb --username=postgres --owner=postgres $(DB_NAME)

dropdb:
	# docker exec -it pg psql -U postgres -c "DROP DATABASE $(DB_NAME);"
	docker exec -it pg dropdb ${DB_NAME}

migrateup:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc

# create migration
# migrate create -ext sql -dir db/migration -seq init_schema
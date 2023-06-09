DB_NAME="simple_bank"
DB_URL="postgresql://postgres:password@localhost:5432/$(DB_NAME)?sslmode=disable"

postgres:
	docker run -e POSTGRES_PASSWORD=password -e POSTGRES_USER=postgres --name=pg -d -p 5432:5432 postgres:14-alpine

createdb:
	# docker exec -it pg psql -U postgres -c "CREATE DATABASE $(DB_NAME);"
	docker exec -it pg createdb --username=postgres --owner=postgres $(DB_NAME)

dropdb:
	# docker exec -it pg psql -U postgres -c "DROP DATABASE $(DB_NAME);"
	docker exec -it pg dropdb ${DB_NAME}

migrateup:
	migrate -path db/migrations -database $(DB_URL) -verbose up

migrateup1:
	migrate -path db/migrations -database $(DB_URL) -verbose up 1

migratedown:
	migrate -path db/migrations -database $(DB_URL) -verbose down

migratedown1:
	migrate -path db/migrations -database $(DB_URL) -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v ./...

server:
	go run main.go

mock:
	mockgen -destination db/mock/store.go -package mockdb github.com/keidarcy/simple-bank/db/sqlc Store

dbdocs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

proto:
	rm -rf pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=pb --grpc-gateway_opt paths=source_relative \
		--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
    proto/*.proto
		statik -src=./doc/swagger -dest=./doc

evans:
	evans --host localhost --port 9090 -r repl

tui:
	go run tui/main.go "user='postgres' password='password' sslmode=disable"

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc mock dbdocs dbschema proto evans tui

# create migration
# migrate create -ext sql -dir db/migration -seq init_schema

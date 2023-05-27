### Tools

```bash
brew install protobuf
go install github.com/golang/protobuf/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
volta install dbdocs
volta install @dbml/cli
```

### Quick commands

- create migration

```bash
migrate create -ext sql -dir db/migrations -seq <migration_name>
```
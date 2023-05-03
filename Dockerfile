# Build stage
FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Install go-migrate
RUN apk add --no-cache curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

# download wait-for from https://github.com/eficode/wait-for/releases
RUN curl -Lo wait-for https://raw.githubusercontent.com/eficode/wait-for/v2.2.3/wait-for
RUN chmod +x wait-for


# Run stage
FROM alpine:3.16
WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY --from=builder /app/wait-for .

COPY ./db/migrations ./db/migrations
COPY ./app.env ./start.sh ./

EXPOSE 8080

## start app
### with CMD and ENTRYPOINT, will pass CMD to ENTRYPOINT
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]
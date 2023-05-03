#!/bin/sh

set -e

echo "run db migrations"
/app/migrate -path /app/db/migrations -database "$DB_SOURCE" -verbose up
echo "start app"
exec "$@"

#!/bin/sh

set -e

echo "running db migration"
/app/migrate -path /app/migration -database "postgresql://vijay:zmxmcmvbn@postgres:5432/echo_rest?sslmode=disable" -verbose up

echo "Starting app"

exec "$@"
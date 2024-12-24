#!/bn/sh

set -e

echo "Run the db migration"
/app/migrate -path /app/migration -database "$DB_SOURCE" --verbose up

echo "start the app"
exec "$@"
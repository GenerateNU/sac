#!/bin/bash

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

PGHOST="127.0.0.1"
PGPORT="5432"
PGUSER="postgres"
PGPASSWORD="postgres"
PGDATABASE="sac"

DATABASES=$(psql -h "$PGHOST" -p "$PGPORT" -U "$PGUSER" -d "$PGDATABASE" -t -c "SELECT datname FROM pg_database WHERE datistemplate = false AND datname != 'postgres' AND datname != 'template0' AND datname != '$PGDATABASE' AND datname !='$(whoami)';")


for db in $DATABASES; do
  echo "Dropping database $db"
  psql -h "$PGHOST" -p "$PGPORT" -U "$PGUSER" -d "$PGDATABASE" -c "DROP DATABASE $db;"
done

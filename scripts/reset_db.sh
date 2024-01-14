#!/bin/bash

# PostgreSQL connection parameters
PGHOST="127.0.0.1"
PGPORT="5432"
PGUSER="postgres"
PGPASSWORD="postgres"
PGDATABASE="sac"

# Check if there are tables to drop
table_count=$(psql -h $PGHOST -p $PGPORT -U $PGUSER -d $PGDATABASE -t -c "SELECT COUNT(*) FROM pg_tables WHERE schemaname = 'public';")
if [ "$table_count" -eq 0 ]; then
  echo "No tables to drop. The database is empty."
  exit 0
fi

echo "Generating DROP TABLE statements..."
if ! psql -h $PGHOST -p $PGPORT -U $PGUSER -d $PGDATABASE -t -c \
  "SELECT 'DROP TABLE IF EXISTS \"' || tablename || '\" CASCADE;' FROM pg_tables WHERE schemaname = 'public';" > drop_tables.sql; then
  echo "Error generating DROP TABLE statements."
  exit 1
fi

echo "Dropping tables..."
if ! psql -q -h $PGHOST -p $PGPORT -U $PGUSER -d $PGDATABASE -a -f drop_tables.sql > /dev/null 2>&1; then
  echo "Error dropping tables."
  rm drop_tables.sql
  exit 1
fi

rm drop_tables.sql

echo "All tables dropped successfully."
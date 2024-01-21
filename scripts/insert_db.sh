#!/bin/bash

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
PGHOST="127.0.0.1"
PGPORT="5432"
PGUSER="postgres"
PGPASSWORD="postgres"
PGDATABASE="sac"
INSERTSQL="../backend/src/migrations/data.sql"
CHECK_TABLES_QUERY="SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'public' LIMIT 1);"

# Change the working directory to the specified location
cd "$SCRIPT_DIR" || { echo "Error: Could not change directory to $SCRIPT_DIR"; exit 1; }

# Check if tables exist in the database
if psql -h "$PGHOST" -p "$PGPORT" -U "$PGUSER" -d "$PGDATABASE" -t -c "$CHECK_TABLES_QUERY" | grep -q "t";  then
  echo "Database $PGDATABASE exists with tables."
else
  echo "Error: Database $PGDATABASE does not exist or has no tables. Running database migration."
  go run ../backend/src/main.go --only-migrate # TODO
  sleep 3

  # Find the process running on port 8080 and kill it
  PROCESS_ID=$(lsof -i :8080 | awk 'NR==2{print $2}')
  if [ -n "$PROCESS_ID" ]; then
    kill -INT $PROCESS_ID
    echo "Killed process $PROCESS_ID running on port 8080."
  else
    echo "No process running on port 8080."
    exit 0
  fi
fi

# Insert data from data.sql
if psql -h "$PGHOST" -p "$PGPORT" -U "$PGUSER" -d "$PGDATABASE" -a -f "$INSERTSQL" > /dev/null 2>&1; then
  echo "Data inserted successfully."
else
  echo "Error: Failed to insert data."
  exit 1
fi
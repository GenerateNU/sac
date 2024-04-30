#!/usr/bin/env bash

set -x
set -eo pipefail

if ! [ -x "$(command -v redis-cli)" ]; then
  echo >&2 "Error: redis-cli is not installed."
  exit 1
fi

if ! [ -x "$(command -v docker)" ]; then
  echo >&2 "Error: docker is not installed."
  exit 1
fi

# Redis Ports and Server Name
REDIS_STACK_PORT="${REDIS_PORT:=6379}"
REDIS_SERVER_PORT="${REDIS_SERVER_PORT:=8001}"

# REDIS_DATA_DIR="${REDIS_DATA_DIR:=/path/to/redis/data}"
# HOST_DATA_DIR="${HOST_DATA_DIR:=/data/redis}"


if [[ -z "${SKIP_DOCKER}" ]]  
then
  RUNNING_REDIS_CONTAINER=$(docker ps --filter 'name=redis' --format '{{.ID}}')
  if [[ -n $RUNNING_REDIS_CONTAINER ]]; then
    echo >&2 "there is a redis container already running, killing it..."
    docker kill ${RUNNING_REDIS_CONTAINER}
    docker rm ${RUNNING_REDIS_CONTAINER}
  fi
  # Start Redis
  docker run \
          -d \
          --name "redis_$(date '+%s')" \
          -p "${REDIS_STACK_PORT}":"${REDIS_STACK_PORT}" \
          -p "${REDIS_SERVER_PORT}":"${REDIS_SERVER_PORT}" \
          -e REDIS_PASSWORD="${REDIS_PASSWORD}" \
          redis/redis-stack:latest
          # Mount volume for data persistence (optional)
          # -v "${HOST_DATA_DIR}":"${REDIS_DATA_DIR}" \
fi
>&2 echo "Redis server is up and running on port ${REDIS_PORT}"

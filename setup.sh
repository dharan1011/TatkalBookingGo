#!/usr/bin/env bash
container_id=$(docker run --name postgres-local -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres:14)
echo "ContainerId: $container_id"
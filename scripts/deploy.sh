#!/bin/bash

COMPOSE_FILE=./deployments/docker-compose.yml

docker-compose --file $(COMPOSE_FILE) build
docker-compose --file $(COMPOSE_FILE) up -d zookeeper broker

until docker-compose --file $(COMPOSE_FILE) exec broker nc -z localhost 9092; do
  sleep 1
done

docker-compose --file $(COMPOSE_FILE) exec broker kafka-topics --create --if-not-exists --topic events --bootstrap-server broker:9092 --partitions 1 --replication-factor 1

docker-compose --file $(COMPOSE_FILE) up -d redis collector storage

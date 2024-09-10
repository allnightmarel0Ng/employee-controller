docker-compose --file ./deployments/docker-compose.yml build
docker-compose --file ./deployments/docker-compose.yml up -d zookeeper broker

until docker-compose --file ./deployments/docker-compose.yml exec broker nc -z localhost 9092; do
  sleep 1
done

docker-compose --file ./deployments/docker-compose.yml exec broker kafka-topics --create --if-not-exists --topic events --bootstrap-server broker:9092 --partitions 1 --replication-factor 1

docker-compose --file ./deployments/docker-compose.yml up -d redis collector storage

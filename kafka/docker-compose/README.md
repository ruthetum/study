# Kafka docker compose
## Installation
### Single broker

```shell
docker-compose -f docker-compose-single.yml up -d
```

### Multi broker
kafka broker 3개, zookeeper 1개

```shell
docker-compose -f docker-compose-multi.yml up -d
```

## Test
### Kafka CLI
```shell
brew install kafka
```

### Create topic
```shell
# single broker
docker exec -t kafka /usr/bin/kafka-topics --bootstrap-server kafka:9092 --create --topic wilump-topic
```

### Describe topic
```shell
# single broker
docker exec -t kafka /usr/bin/kafka-topics --bootstrap-server kafka:9092 --describe --topic wilump-topic
```

## Reference
- https://devocean.sk.com/blog/techBoardDetail.do?ID=164007
- https://devocean.sk.com/blog/techBoardDetail.do?ID=164016
- https://sup2is.github.io/2020/06/10/apache-kafka-02.html
- https://www.conduktor.io/kafka/how-to-install-apache-kafka-on-mac-with-homebrew/
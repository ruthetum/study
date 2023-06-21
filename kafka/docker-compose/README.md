# Kafka docker compose
## Installation
### Single broker

```shell
docker-compose -f docker-compose-single.yml up -d
```

### Multi broker
zookeeper 3개 인스턴스, kafka broker 3개

```shell
docker-compose -f docker-compose-multi.yml up -d
```

## Test
### Create topic
```shell
docker-compose exec kafka kafka-topics --create --topic wilump-topic --bootstrap-server kafka:9092 --replication-factor 1 --partitions 1
```

### Describe topic
```shell
docker-compose exec kafka kafka-topics --describe --topic wilump-topic --bootstrap-server kafka:9092 
```

## Reference
- https://devocean.sk.com/blog/techBoardDetail.do?ID=164007
- https://devocean.sk.com/blog/techBoardDetail.do?ID=164016
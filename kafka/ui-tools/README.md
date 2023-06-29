# UI Monitoring Tools
![image](https://miro.medium.com/v2/resize:fit:1400/format:webp/1*MZLPKPUqwyAGMc88k6IzSQ.png)

## Kafka-ui
### compose
```yaml
version: '2'

services:
  kafka-ui:
    image: provectuslabs/kafka-ui
    container_name: kafka-ui
    ports:
      - "18989:8080"
    restart: always
    environment:
      - KAFKA_CLUSTERS_0_NAME=local
      - KAFKA_CLUSTERS_0_BOOTSTRAPSERVERS=kafka-1:29092,kafka-2:29093,kafka-3:29094
      - KAFKA_CLUSTERS_0_ZOOKEEPER=zookeeper-1:22181
```

ref. https://github.com/provectus/kafka-ui/blob/master/documentation/compose/DOCKER_COMPOSE.md

### helm

ref. https://docs.kafka-ui.provectus.io/configuration/helm-charts/quick-start

## Redpanda console (Kowl)
### compose

ref. https://docs.redpanda.com/docs/reference/docker-compose/

### helm

ref. https://docs.redpanda.com/docs/deploy/deployment-option/self-hosted/kubernetes/eks-guide/

## References
- comparison: https://towardsdatascience.com/overview-of-ui-tools-for-monitoring-and-management-of-apache-kafka-clusters-8c383f897e80
- https://github.com/provectus/kafka-ui
  - https://docs.kafka-ui.provectus.io/overview/readme
- https://github.com/redpanda-data/console
  - https://docs.redpanda.com/docs/manage/console/
  - https://docs.redpanda.com/docs/deploy/deployment-option/self-hosted/kubernetes/eks-guide/
- https://devocean.sk.com/blog/techBoardDetail.do?ID=163980
- https://github.com/redpanda-data/console/blob/master/docs/local/docker-compose.yaml
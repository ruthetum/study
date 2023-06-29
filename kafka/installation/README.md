# Installation

## helm
```shell
helm repo add bitnami https://charts.bitnami.com/bitnami

helm install kafka bitnami/kafka

helm install kafka --set replicaCount=3 bitnami/kafka
```

## Reference
- https://github.com/bitnami/charts
  - https://github.com/bitnami/charts/tree/main/bitnami/kafka
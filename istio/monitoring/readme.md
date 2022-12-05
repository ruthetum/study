# Grafana
- 오픈소스 메트릭 데이터 시각화 도구
- 데이터 수집은 다른 도구(data source)에 맡기고, 시각화에 집중

## Dashboard
- https://grafana.com/grafana/dashboards/ 에서 원하는 템플릿을 다운받을 수 있음

# Prometheus
- Golang으로 개발된 TSDB(시계열 데이터베이스)
- 오픈소스 모니터링 시스템
- 시계열 데이터 수집
- 다차원 데이터 모델을 활용한 유연한 쿼리 PromQL
- 성능이 좋아 한 대로 초당 수백만개의 데이터 샘플 처리

## PromQL
- 사용자가 실시간으로 시계열 데이터를 선택하고 집계할 수 있는 쿼리

## Metric Type
- Counter
- Gauge
- Histogram

### Counter
- 가장 많이 수집되는 유형
- 증가하는 이벤트의 누적 개수 또는 크기 (리셋 시 제로)
- Ex. 초당 요청 개수

### Gauge
- 간단한 수로 표현되는 값
- 증가할 수도 있고, 감소할 수도 있는 값
- Ex. 메모리, 프로세스 할당량

### Histogram
- 특정 버킷에 속하는 값의 관찰빈도 측정
- 다소 복잡한 유형, istio metric에서 이슈가 종종 발생함
- 분위수 계산(Quantiles)할 때 유용 -> p95 latency

## Range Selector
- 쿼리에 범위 기간을 추가해서 각 메트릭의 값을 추출 (Range Vector라고 함)
# Kafka Lag
- partition이 1개인 토픽을 가정하고 프로듀서가 partition에 데이터를 넣게 되면 0부터 차례대로 offset이 붙여짐
- 프로듀서는 계속해서 데이터를 넣게 되고, 컨슈머는 계속해서 데이터를 가져감
- 만약 프로듀서가 데이터를 넣어주는 속도보다 컨슈머가 가져가는 속도보다 빠르게 된다면?

    ![image](https://user-images.githubusercontent.com/59307414/155877884-1496c912-67f9-47b4-8a8e-610281a37954.png)

- `프로듀서가 넣은 offset`과 `컨슈머가 읽은 offset` 간의 차이가 발생
    - 이 차이가 `Kafka consumer lag`
- lag은 적을수도 있고, 많을수도 있음
- 컨슈머 그룹이 1개이고, 파티션이 2개인 토픽에서 데이터를 가져간다면 lag은 2개가 측정
- 이처럼 한 개의 토픽과 컨슈머 그룹에 대한 lag이 여러 개 존재할 수 있을 때 그 중 높은 숫자의 lag을 `records-lag-max`라고 부름

# 실시간 모니터링
- kafka clients를 이용해서 구현한 Kafka 컨슈머를 통해 lag 정보를 실시간으로 가져올 수 있음
- lag을 모니터링하고 싶다면 데이터를 Elasticsearch나 InfluxDB 같은 저장소에 넣은 뒤 Grafana 대시보드를 통해 확인할 수 있음

    > ## Elasticsearch
    > - Elasticsearch는 Apache Lucene기반의 Java 오픈소스 분산 검색 엔진
    > - 방대한 양의 데이터를 신속하게, 거의 실시간(NRT, Near Real Time)으로 저장, 검색, 분석할 수 있음
    > - [Document](https://www.elastic.co/guide/index.html)
    > - [관련 POST](https://victorydntmd.tistory.com/308)
    > - cf. ELK (Elasticsearch / Logstatsh / Kibana)

    > ## InfluxDB
    > - 많은 쓰기 작업과 쿼리 부하를 처리하기 위해 개발된 시계열 데이터베이스
    > - 손쉽게 scale-out할 수 있으며, Restful API를 제공하고 있어 API 통신이 가능
    > - [Document](https://docs.influxdata.com/influxdb/v2.1/)
    > - [관련 POST](https://mangkyu.tistory.com/190)
    > - cf. Tick Stack (Telegraf + InfluxDB + Chronograf + Kapacitor)

    > ## Grafana
    > - 시계열 데이터에 대한 대시보드를 제공해주는 Data Visualization Tool
    > - [Document](https://grafana.com/docs/grafana/latest/)
    > - [관련 POST](https://medium.com/finda-tech/grafana%EB%9E%80-f3c7c1551c38)

- 컨슈머 단위에서 lag을 모니터링하고 관리하는 것은 위험하고 운영요소가 많이 들어감
- 컨슈머 로직단에서 lag을 수집하다보니 컨슈머 상태에 디펜던시가 걸림
    - 컨슈머에 장애가 발생하는 경우 lag 정보를 보낼 수 없음
    - 컨슈머가 추가될 때마다 lag 정보를 특정 저장소에 저장할 수 있도록 로직을 추가로 개발해야 함

# Burrow
- 따라서 컨슈머의 로직단에서 처리하지 말고 Burrow를 활용하자
- https://github.com/linkedin/Burrow

## 특징
1. Multi kafka cluster 지원
2. Sliding window를 통한 Consumer의 status 
    - `ERROR`, `WARNING`, `OK`로 표현
3. HTTP API 제공

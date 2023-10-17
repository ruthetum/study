# Discord Migrates Trillions of Messages from Cassandra to ScyllaDB
> https://www.infoq.com/news/2023/06/discord-cassandra-scylladb/

## Discord 시작: MongoDB
> [MongoDB → Cassandra](https://discord.com/blog/how-discord-stores-billions-of-messages): What we were doing

- 2015년 개발 당시 Discord는 MongoDB 를 사용
  - a single MongoDB replica set

- 2015년 11월, 사용이 증가함에 따라 저장된 메세지가 1억 건에 달했고 이 때 성능 문제가 발생
  - 데이터/인덱스가 메모리(RAM) 크기를 넘어섬
  - latency가 예측 불가능해짐

- 메세지의 수는 지속적으로 증가
  - 2016년 7월: 일 4천만 개
  - 2016년 12월: 일 1억 개

## MongoDB 를 사용했을 때 발생했던 문제점과 새로운 DB 선택 요구사항
> [MongoDB → Cassandra](https://discord.com/blog/how-discord-stores-billions-of-messages): Choosing the Right Database

### MongoDB 를 사용했을 때 발생했던 문제점
- 매우 랜덤하게 읽기 작업이 발생
  - 읽기/쓰기 비율이 5:5
- 음성 채팅이 많은 서버(디스코드 내 organization/workspace): 연간 1000개 메세지 발생
  -  소수 메세지 조회 → 디스크에 많은 랜덤 검색 발생 → 캐시 효율 감소
- 비공개 채팅이 많은 서버(디스코드 내 organization/workspace): 연간 10~100만개의 메세지 발생
  - 요청은 최근 메세지만 조회, 회원 수는 100명 미만 → 요청 비율이 낮고, 캐시 적중률(히트)이 낮음
- 대규모 공개 Discord 서버(디스코드 내 organization/workspace): 연간 수백만 개의 메세지 발생
  - 최근 한 시간 내에 발생하는 메세지(데이터)를 조회하는 경우가 많음 → 캐시 히트가 높음
- 또한 앞으로 랜덤 검색을 유발하는 기능들이 추가될 예정

### DB 선택 요구사항
- 선형 수평 확장: 솔루션 재검토하거나 데이터를 수동으로 재배치(re-shard)하고 싶지 않음
- 자동 장애 조치: 스스로 복구될 수 있는 시스템
- 낮은 유지보수 비용: 데이터가 증가하면 노드만 추가
- 검증된 기술: 새 기술을 좋아하지만 너무 새롭지 않은 기술
- 예측 가능한 성능: 95p의 응답 시간이 80ms 이하, Redis 또는 Memcached에 메세지를 캐시하고 싶지 않음
- blob 저장소 아닌 저장소: 초당 수천 개의 메세지가 작성되기 때문에 blob 직렬화/역직렬화는 비효율적임
- 오픈 소스: 타사에 의존하고 싶지 않음

## Cassandra 선택
> [Cassandra → ScyllaDB](https://discord.com/blog/how-discord-stores-trillions-of-messages): Our Cassandra Troubles

- 2017년: 12개 노드로 시작, 수 십억 개의 메세지를 저
- 적용 후 좋았으나 GC가 10초동안 발생하는 문제가 발생하기도 함
  - 툼스톤 때문에 발생한 문제
    - 툼스톤 기간: 10일에서 2일로 축소
    - 빈 버킷을 조회하지 않도록 함

- 2022년: 177개 노드, 메세지는 수 조개에 도달
    - 많은 동시 읽기 → 핫 파티션 → 성능 문제 발생: 대기 시간 예측 불가
      - 카산드라: 읽기 비용 > 쓰기 비용
        - 읽기는 Memtable(메모리)에 데이터가 없으면 SSTable(파일)을 조회

    - 유지보수 비용 증가
      - SSTable 압축에 따른 성능 문제 → gossip dance 운영 작업
      - gossip dance: 클러스터 내 노드 중 한 대를 가져와서 트래픽을 받지 않고, 파일을 압축하고, 다시 클러스터에 돌려보내는 작업을 반복

## 아키텍처 변경
> [Cassandra → ScyllaDB](https://discord.com/blog/how-discord-stores-trillions-of-messages): Changing Our Architecture

- Cassandra → ScyllaDB
  - GC로 인해 발생하는 지연 시간 문제 감소

- 데이터 서비스 API(레이어) 추가 ([Cassandra → ScyllaDB](https://discord.com/blog/how-discord-stores-trillions-of-messages):Data Services Serving Data)
  - 동일 데이터에 대한 여러 요청을 한 번에 DB로 보냄 → DB로 갈 쿼리 수를 줄여서 DB 부하를 감소
  - 일관성 해시([consistence hashing](../distributed-system/consistence-hashing.md)) 사용으로 동일 데이터에 대한 요청은 동일 데이터 서비스로 보냄

## 결과
- 2022년 5월
  - Cassandra 177개 노드 → ScyllaDB 72개 노드
    - 노드 별 평균 4TB 디스크 → 9TB 디스크
  - 메세지 히스토리 읽기 p99: 40-125ms → 15ms
  - 메시지 작성 p99: 5-70ms → 5ms
  - 온콜 대응이 줄어듦

## Reference
- https://www.infoq.com/news/2023/06/discord-cassandra-scylladb/
- https://discord.com/blog/how-discord-stores-billions-of-messages
- https://discord.com/blog/how-discord-stores-trillions-of-messages
- https://www.youtube.com/watch?v=mU3JiOI31Ao
- https://blog.voidmainvoid.net/469
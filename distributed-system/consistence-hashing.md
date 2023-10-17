# Consistent Hashing
> (데이터) 서버 수평 확장

## 데이터 수평 확장
- 데이터 규모가 커질 때 확장하는 방법
  - 스토리지 용량 늘리기 → 금방 한계에 도달
  - 데이터를 각 서버에 나눠서 저장하기 → 파티션 / 샤딩

## 데이터를 나누는 쉬운 방법
- 해시 함수, 나머지 연산 → 데이터를 어떤 서버에 저장할지 결정
  - 서버 인덱스 = hash(키) % 서버 개수
  - e.g. hash(x) = x, 서버 개수 = 3
    - 서버 0: 0, 3, 6, 9, ...
    - 서버 1: 1, 4, 7, 10, ...
    - 서버 2: 2, 5, 8, 11, ...

- 하지만 서버 개수 변동 시, 많은 데이터의 재배치가 필요(데이터의 위치가 변경되어야 함)

## Consistent Hashing
- 해시 테이블 크기 변경 시 평균 k/n 개만 재배치되는 해시 기법
  - k: 키의 개수, n: 슬롯 개수

- 기본 원리
  - 해시 공간: x0 ~ xn / 슬롯(서버): s0 ~ sn
  - 해시 공간에 슬롯을 위치시킴
  - hash(키) 결과로 저장할 슬롯 선택

- 데이터가 많고, 서버 개수가 작으면 서버 추가/삭제 시 많은 데이터 이동 발생
  - 가상 노드로 이동 개수를 줄일 수 있음

### 가상 노드
- 실제 서버를 늘리지 않고 가상으로 노드를 늘림

## 사용
- 데이터를 분산 저장하는 서비스에 사용 고려
  - e.g. 슬랙 채널 서버 관리, Cassandra, DynamoDB
    - Redis cluster는 hash slot을 사용 ([Redis Cluster does not use consistent hashing](https://redis.io/docs/management/scaling/))
  - 분산 캐시에도 적합

## Reference
- https://en.wikipedia.org/wiki/Consistent_hashing
- https://binux.tistory.com/119
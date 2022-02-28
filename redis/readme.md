# Redis 야무지게 사용하기
### content
1. Redis 캐시로 사용하기
2. Redis 데이터 타입 활용하기
3. Redis에서 데이터를 영구 저장하려면? (RDB vs AOF)
4. Redis 아키텍처 선택 노하우 (Replication vs Sentinel vs Cluster)
5. Redis 운영 Tip 및 장애 포인트
---

# 1. Redis 캐시로 사용하기
> ### 캐시(Cache) : Temporary Location For Speed
> 데이터의 원래 소스보다 더 빠르고 효율적으로 엑세스할 수 있는 임시 데이터 저장소
> - 원본(Origin)보다 빠른 접근 속도
> - 같은 데이터에 대한 반복적인 엑세스
> - 잘 변하지 않는 데이터

![image](https://user-images.githubusercontent.com/59307414/155966314-fc882829-36ea-4181-9562-68ec8f9b756a.png)

## Redis as a cache
- Redis는 캐시로 활용하기에 가장 효과적이고 효율적인 소프트웨어
- 단순한 key-value 구조
- In-memory 데이터 저장소(RAM)
- 빠른 성능 
    - 평균 작업속도 < 1ms
    - 초당 수백만건의 작업 가능

## 캐싱 전략
### 읽기 전략 : Look-Aside(Lazy Loading)
- 애플리케이션에 데이터를 읽는 작업이 많을 때 활용되는 전략
- Redis를 캐시로 사용할 때 가장 많이 쓰는 전략

    ![image](https://user-images.githubusercontent.com/59307414/155966426-f1a1bb93-65ac-487b-9e7d-5607025b9fcc.png)

- 애플리케이션은 데이터를 찾을 때 redis를 먼저 확인하고 있는 경우 해당 데이터를 redis에서 가져옴

    ![image](https://user-images.githubusercontent.com/59307414/155966664-b7b7a731-aa7e-4684-9ea7-afa83539f941.png)

- 만약 redis에 찾는 키가 없다면 DB에서 조회해서 값을 가져온 뒤, 다시 redis에 저장
- 따라서 찾는 데이터가 없을 때에만 redis에 값을 저장하기 때문에 <b>Lazy Loading</b>이라고도 부름
- 이 구조는 redis가 다운되더라도 바로 장애가 일어나지 않고 DB에서 값을 조회할 수 있음
- 대신 cache로 붙이있던 connection이 많았다면 이 connection이 모두 DB로 연결되기 때문에 DB에 갑자기 많은 부하가 발생할 수 있음
- 따라서 캐시를 새로 투입하거나 DB에 새로운 값을 저장한 경우 `Cache Miss`가 증가하기 때문에 성능 저하를 유발할 수 있음

    ![image](https://user-images.githubusercontent.com/59307414/155968126-6d38bcd6-bdca-47bd-8146-531880460ae3.png)

- 이럴 때에는 미리 DB에서 cache로 데이터를 밀어 넣어주는 `Cache Warming` 작업을 진행


### 쓰기 전략

![image](https://user-images.githubusercontent.com/59307414/155968521-11397b02-43f9-4cb0-ab71-6e3d27b7b3fd.png)

#### Write-Around
- 모든 데이터는 일단 DB에만 저장
- cache miss 발생할 때는 cache에 DB의 데이터를 끌어옴
- cache 내 데이터와 DB 데이터가 다를 수 있음

#### Write-Through
- DB에 데이터 저장할때, cache에도 같이 저장
- cache는 항상 최신 정보를 가지고 있지만 저장할때마다 2단계 과정을 거치기 때문에 상대적으로 느림
- cache에 넣은 데이터가 저장만 되고 재사용되지 않을 가능성이 있어서 일종의 리소스 낭비라고 볼 수 있음
- 따라서 데이터를 저장할 때에 "몇 분", "몇 시간" 등을 세팅하는 `expire time`를 정하는 게 좋음
- 값의 관리에 따라 장애 포인트가 발생할 수 있음

---

# 2. Redis 데이터 타입 활용하기

Redis는 자체적으로 다양한 자료구조를 제공

![image](https://user-images.githubusercontent.com/59307414/155970237-fa8deced-00d9-40ca-a62b-3e83943dd523.png)

|자료형|특징|
|---|---|
|Strings|- 가장 기본적인 형태<br>- Set command이용해 저장되는 데이터는 모두 string 형태|
|Bitmaps|- String의 변형<br>- bit 단위 연산 가능|
|Lists|- 데이터를 순서대로 저장<br>- queue로 사용하기 적절|
|Hashes|- 하나의 key안에 다시 key-value 형태 저장|
|Sets|- 중복되지 않는 문자열의 집합|
|Sorted Sets|- Sets과 기본적으로 동일<br>- 모든 값은 score라는 숫자 값으로 정렬<br>- 데이터가 저장될 때부터 score 순으로 정렬<br>- score가 같을 때에는 사전순으로 정렬|
|HyperLogLogs|- 굉장히 많은양의 데이터를 다룰 때 사용<br>- 중복되지 않는 데이터를 count할때 주로 많이 사용|
|Streams|- 로그를 저장하기 가장 좋은 자료 구조|

---

## Best Practice - Counting
### Strings
- 단순 증감 연산
- `INCR`, `INCRBY`

### Bits
- 데이터 저장공간 절약
- 예를 들어 오늘 접속한 사용자의 수를 저장할 때 날짜 키 하나를 만들고 userId의 비트를 증가
- 사용자가 1000만명일 때 1000만 bit = 1.2MB
- `SETBIT`, `BITCOUNT`

### HyperLogLogs
- 대용량 데이터를 카운팅할 때 적절
- set과 비슷하지만 저장되는 용량은 매우 작음(12KB 고정)
- 저장된 데이터는 다시 확인할 수 없음
- `PFADD`, `PFCOUNT`, `PFMERGE`

## Best Practice - Messaging
### Lists
- Message queue로 활용하기에 적절
- 자체적으로 Blocking 기능을 제공하기 때문에 불필요한 polling process를 방지할 수 있음
- `BRPOP`, `LPUSH`
- 키가 있을 때만 저장 가능 - `LPUSHX` / `RPUSHX`
    - 이미 캐싱되어 있는 피드에만 신규 트윗을 저장

### Streams
- 로그를 저장하기 가장 적절한 자료구조
- append-only
- 시간 범위로 검색 / 신규 추가 데이터 수신 / 소비자별 다른 데이터 수신(소비자 그룹)
- `XADD`

---

# 3. Redis에서 데이터를 영구 저장하려면? (AOF vs RDB)

## Redis Persistance
### Redis는 In-memory 데이터 스토어
- 서버 재시작 시 모든 데이터 유실
- 복제 기능을 사용해도 사람의 실수 발생 시 데이터 복원 불가
- Redis를 캐시 이외의 용도로 사용한다면 적절한 데이터 백업이 필요

## Redis Persistance Option : AOF vs RDB
- redis에서는 데이터를 영구 저장하기 위해 두 가지 옵션 제공

### AOF(Append Only File)
- 데이터를 변경하는 command가 들어오면 command를 그대로 파일로 관리
- AOF 파일은 계속해서 추가되기만 해서 대부분 RDB파일보다 커지게 됨
- 따라서 AOF 파일은 주기적으로 압축해서 재작성해야 됨
- 실제 저장은 Redis protocol 형태로 저장

### RDB(snapshat)
- command를 저장하지 않고 저장 당시에 메모리를 그대로 파일로 관리
- 실제 저장은 binary 형태로 저장

## 자동/수동 파일 저장 방법 : AOF rewrite & RDB save

### AOF
- 자동 : redis.conf 파일에서 `auto-aof-rewrite-percentege` 옵션(크기 기준)
- 수동 : `BGREWRITEAOF` 커맨드를 이용해 CLI 창에서 수동으로 AOF 파일 재작성

### RDB
- 자동 : redis.conf 파일에서 `SAVE` 옵션(시간 기준)
- 수동 : `BGSAVE` 커맨드를 이용해서 CLI 창에서 수동으로 RDB 파일 저장
    - `SAVE` 커맨드는 절대 사용 X

## RDB vs AOF 선택 기준
> Redis를 단순히 캐시로만 사용하는 경우 백업은 불필요

### 백업은 필요하지만 어느 정도의 데이터 손실이 발생해도 괜찮은 경우
- <b>RDB</b> 단독 사용
- redis.conf 파일에서 `SAVE` 옵션을 적절히 사용
- ex. `SAVE 900 1`

### 장애 상황 직전까지의 모든 데이터가 보장되어야 할 경우
- <b>AOF</b> 사용
- `APPENDSYNC` 옵션이 `everysec`인 경우 최대 1초 사이의 데이터 유실 가능(기본 설정)

### 제일 강력한 내구성이 필요한 경우
- <b>RDB</b>와 <b>AOF</b> 모두 사용

---

# 4. Redis 아키텍처 선택 노하우 (Replication vs Sentinel vs Cluster)

![image](https://user-images.githubusercontent.com/59307414/155981739-da87f35f-8d32-4715-a0a4-f5a9726cf836.png)

## Replication
### 단순한 복제 연결
- `replicaof` 커맨드를 이용해 간단하게 복제 연결
- 비동기식 복제
- HA 기능이 없으므로 장애 상황 시 수동 복구
    - `replicaof no one`
    - 애플리케이션에서 연결 정보 변경

> HA(High Availablity)?
> 고가용성, 백업 또는 장애극복 처리를 의미

## Sentinel
### 자동 Failover 가능한 HA 구성
- sentinel 노드가 다른 노드 감시
- 마스터가 비정상 상태일 때 자동으로 failover
- 연결 정보 변경 필요 없음
- 이렇게 구성하기 위해서는 sentinal 프로세스를 추가로 실행시켜야 함
- sentinel 노드는 하상 3개 이상의 홀수로 존재해야 함
    - 과반수 이상의 sentinel이 동의해야 failover 진행

## Cluster
### Scale Out과 HA 구성
- 키를 여러 노드에 자동으로 분할해서 저장(샤딩)
- 모든 노드가 서로를 감시하여 마스터가 비정상 상태일 때 자동으로 failover
- 최소 3대의 마스터 노드가 필요

## 아키텍처 선택 기준

![image](https://user-images.githubusercontent.com/59307414/155982769-f24268a5-77f4-4b81-8d33-443f6f035022.png)

### 자동 failover 필요 & Scale Out 필요(샤딩)
- <b>Cluster</b>

### 자동 failover 필요 & Scale Out 필요 X(샤딩)
- <b>Sentinel</b>
    
### 자동 failover 필요 X & 복제 필요
- <b>Replication</b>

### 자동 failover 필요 X & 복제 필요 X
- <b>Stand-Alone</b>

---

# 5. Redis 운영 Tip 및 장애 포인트
## 사용하면 안 되는 커맨드
### Redis는 Single Thread로 동작
- `keys *` → `scan`으로 대체
- Hash나 Sorted Set 등 자료구조
    - 키 내부에 아이템이 많아질수록 성능이 저하
    - 성능을 고려한다면 키를 나누어서 하나의 키에 최대 100만개까지만 아이템을 저장
    - `hgetall` → `hscan`
    - 키에 많은 아이템이 저장되어 있을 때 `del`로 지우면 해당 키를 지우는동안 아무런 동작을 할 수가 없음
        - 이 때 `unlink` 커맨드를 사용하면 백그라운드에서 키를 삭제

## 변경하면 장애를 막을 수 있는 기본 설정값
### `STOP-WRITES-ON-BGSAVE-ERROR = NO`
- Default 설정은 yes
- RDB 파일 저장 실패 시 redis로 들어오는 모든 write를 차단하는 옵션

### `MAXMEMORY-POLICY = ALLKEYS-LRU`
- redis를 캐시로 사용할 때 expire time 설정 권장
- 메모리가 가득 찼을 때 MAXMEMORY-POLICY 정책에 의해 키 관리
    - `NOEVICTION`(default) : 삭제 안함 → 더 이상 redis에 새로운 키를 저장하지 않는 것을 의미 
    - `VOLATILE-LRU` : LRU 정책, 가장 최근에 사용하지 않았던 key부터 삭제, expire 설정에 있는 key만 삭제
    - `ALLKEYS-LRU` : 모든 key에 대해 expire에 상관없이 LRU 정책 적용

## Cache Stampede
### TTL 값을 너무 작게 설정한 경우

![image](https://user-images.githubusercontent.com/59307414/155986102-b102835f-b9a5-48ca-80db-37c5fc555828.png)

- 대규모 트래픽 환경에서 TTL 값을 너무 작게 설정한 경우 `cache stampede` 현상 발생

## MaxMemory 값 설정
### Persistence / 복제 사용 시 MaxMemory 설정 주의
- RDB 저장 & AOF rewrite 시 `fork()`
- Copy-on-Write로 인해 메모리를 두 배로 사용하는 경우 발생 가능
- Persistence / 복제 사용 시 MaxMemory는 실제 메모리의 절반으로 설정

## Memory 관리
### 물리적으로 사용되고 있는 메모리를 모니터링
- <b>`used_memory`보다 `used_memory_rss`를 보는 것이 더 중요</b>
- `used_memory` : 논리적으로 redis가 사용하는 메모리
- `used_memory_rss` : OS가 redis에 할당하기 위해 사용한 물리적 메모리 양
- 삭제되는 키가 많으면 fragmentation 증가
    - 특정 시점에 피크를 찍고 다시 삭제되는 경우
    - TTL로 인해 eviction이 많이 발생하는 경우
- 이 때 `CONFIG SET activedefrag yes` 옵션을 켜서 해결

---

## Reference
https://www.youtube.com/watch?v=92NizoBL4uA
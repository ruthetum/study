# Kafka

## 배경
![before](https://user-images.githubusercontent.com/59307414/155870791-d9add106-e406-4da2-adfb-f42001b1f93e.png)

- 기존 링크드인의 데이터 처리 시스템은 각 파이프라인이 파편화되고 시스템 복잡도가 높아 새로운 시스템을 확장하기 어려운 상황이였음
- 기존 메시징 큐 시스템인 ActiveMQ를 사용했지만, 링크드인의 수많은 트래픽과 데이터를 처리하기에는 한계가 있었음
- 이로 인해 새로운 시스템의 개발 필요성이 높아졌고, 링크드인의 몇몇 개발자가 다음과 같은 목표를 가지고 새로운 시스템을 개발

![after](https://user-images.githubusercontent.com/59307414/155870795-2abbab39-459a-4c84-bafc-a28d08f2275e.png)

- 카프카를 적용함으로써 모든 이벤트/데이터의 흐름을 중앙에서 관리할 수 있게 됨.
- 서비스 아키텍처가 기존에 비해 관리하기 간단해짐

## 메시징 시스템

- Publisher와 Subscriber로 이루어진 비동기 메시징 전송 방식
- 메시지라고 불리는 데이터 단위를 보내는 측(Publisher/Producer)에서 메시지 시스템에 메시지를 저장하면 가져가는 측(Subscriber/Consumer)는 데이터를 수신
발신자의 메시지에는 수신자가 정해져 있지 않은 상태로 발행(Publish)
- 구독(Subscribe)을 신청한 수신자만이 정해진 메시지를 받음

![image](https://user-images.githubusercontent.com/59307414/155870932-ee3b34ab-004d-4945-8e86-37c825931c81.png)

- 기존 연결의 경우 서비스가 많아짐에 따라 서비스간에 복합적으로 연결되면서 N:N 연결이 발생
- 요청이 많아지거나 사용자가 많아질수록 느려질 수 있음
- 서비스나 클라이언트가 다운되는 경우 메세지가 유실될 수 있음

![image](https://user-images.githubusercontent.com/59307414/155870974-192aaf95-7062-4cde-a8b5-6285cbc2b523.png)

- 프로듀서가 메시지를 컨슈머에게 직접 전달하는 것이 아니라, 중간의 메시징 시스템에 전달 (수신처 ID포함)
- 메시징 시스템의 교환기가 메시지의 수신처 ID값을 통해 컨슈머들의 큐에 메시지를 전달 (push)
- 컨슈머는 큐를 모니터링하다가 큐에 메시지가 있을 경우 값을 회수

---

## 카프카 아키텍처
![image](https://user-images.githubusercontent.com/59307414/155871215-b7cb45c2-0af3-4f73-80de-e2461bae946a.png)

### 프로듀서(Producer)
- 메시지를 생산하여 브로커의 토픽으로 전달하는 역할
    - Topic에 해당하는 메세지를 생성
    - 특정 Topic으로 데이터를 publish
    - 처리 실패/재시도
        - 전송 성공 여부를 알 수 있음

### 브로커(Broker)
- 카프카 애플리케이션이 설치되어 있는 서버 또는 노드를 지칭
- 보통 3개 이상으로 구성하는 것을 권장

### 컨슈머(Consumer)
- 브로커의 토픽으로부터 저장된 메시지를 전달받는 역할
    - Topic의 partition으로부터 데이터 polling
    - Partition offset 위치 기록(commit)
    - Consumer group을 통해 병렬 처리
- Consumer의 개수는 partition의 개수보다 적거나 같아야 함

### 주키퍼(Zookeeper)
- 분산 애플리케이션 관리를 위한 코디네이션 시스템
- 분산된 노드의 정보를 중앙에 집중하고 구성관리, 그룹 네이밍, 동기화 등의 서비스 수행

## Dependency
- 카프카 프로듀서와 컨슈머를 사용하기 위해서 아래 아파치 카프카 라이브러리를 추가

```gradle
// gradle
compile group: 'org.apache.kafka', name: 'kafka-clinets', version: '2.3.0'
```

- 카프카는 브로커 버전과 클라이언트 버전의 하위 호환성을 완벽하게 지원하지 않음
- 따라서 브로커 버전과 클라이언트 버전의 호환성을 확인

## 카프카 작동방식
- 프로듀서는 새 메시지를 카프카에 전달
- 전달된 메시지는 브로커의 토픽이라는 메시지 구분자에 저장
- 컨슈머는 구독한 토픽에 접근하여 메시지를 가져옴 (pull 방식)

---

## Kafka Replication
![image](https://user-images.githubusercontent.com/59307414/155871797-4c7bb189-bb1b-4d4f-a2f5-38107ad954f6.png)

- 서버에 장애가 발생할 것을 대비해 카프카는 replication 기능을 제공
- 카프카의 replication은 토픽 자체를 복제하는 것이 아니라 토픽을 이루는 각각의 파티션을 복제함
> 위 그림에 대한 설명
> - partition : 4, replication : 3
> - 파란색 상자는 Leader partition, 주황색 상자는 Follower partition

- Leader partition과 Follower partition을 합쳐서 <strong>ISR</strong>(In Sync Replica)

- 최대로 설정할 수 있는 replication 값은 브로커 수와 동일
    - 브로커가 4개이면 replication 값은 4보다 클 수 없음

### Why Replicate?
- Replication은 partition의 고가용성을 위해 사용
- 만약 replication이 1이고, partition이 1인 topic이 존재할 때, 브로커의 장애가 발생한 경우 해당 서비스를 복구할 수 없음
- 만약 replication이 2라면, 브로커 1개가 죽더라도 Follower partition이 존재하므로 복제본을 통해 복구 가능
    - Follower partition이 Leader partition의 역할을 승계받음

### Replication & ack
- 프로듀서에는 ack라는 옵션이 존재하고, 이를 통해 고가용성을 유지할 수 있음
- ack는 0, 1, all 옵션 3개 중 하나를 골라 사용 가능

#### ack = 0
- 프로듀서는 Leader partition에 데이터를 전송하고 응답값을 받지 않음
- 따라서 Leader partition에 데이터가 정상적으로 전송되었는지, 나머지 partition에 정상적으로 복제되었는지 알 수 없음
- 속도는 빠르지만 데이터 유실 가능성이 있음

#### ack = 1
- 프로듀서는 Leader partition에 데이터를 전송하고, Leader partition이 정상적으로 받았는지 응답값을 받음
- 하지만 나머지 partition에 정상적으로 복제되었는지 확인하지 않음
- 만약 Leader partition이 데이터를 받은 즉시 브로커에 장애가 발생하면 나머지 partition에 데이터가 아직 전송되지 않은 상태이므로 `ack=0` 일 때와 마찬가지로 데이터 유실 가능성이 있음

#### ack = all
- 프로듀서가 Leader partition에 데이터를 전송하고, Follower partition에 복제가 잘 이루어졌는지 응답값을 받음
- `ack=all`을 사용하는 경우 데이터 유실은 없음
- 하지만 `ack=0, 1`에 비해 확인하는 부분이 많기 때문에 속도가 현저히 느림

### Replication Count
- replication의 수가 커진다고 무조건 좋지는 않음
- replication의 수가 커진만큼 브로커의 resource 사용량이 커짐
- 따라서 카프카에 들어오는 데이터의 양과 retention date(저장 시간)을 고려해서 수를 선택하는 것이 좋음
    - broker의 수가 3개 이상일 때 replication은 3으로 설정하는 것을 권장


## Reference
- https://kafka.apache.org/
- https://epicdevs.com/17
- https://velog.io/@jaehyeong/Apache-Kafka%EC%95%84%ED%8C%8C%EC%B9%98-%EC%B9%B4%ED%94%84%EC%B9%B4%EB%9E%80-%EB%AC%B4%EC%97%87%EC%9D%B8%EA%B0%80
- https://www.youtube.com/watch?v=7QfEpRTRdIQ
# 프로듀서 (Producer)

![image](https://github.com/ruthetum/study/assets/59307414/726ed680-6715-48e2-b9d2-b5cbd4700532)

카프카애서 데이터의 시작은 프로듀서

프로듀서 애플리케이션은 카프카에 필요한 데이터를 선언하고, 브로커의 특정 토픽의 파티션에 전송

프로듀서는 데이터를 전송할 때 리더 파티션을 가지고 있는 카프카 브로커와 직접 통신

프로듀서는 카프카 브로커로 데이터를 전송할 때 내부적으로 파티셔너, 배치 생성 단계를 거침

### 프로듀서 애플리케이션 개발
기본적으로 프로듀서 애플리케이션은 Java를 포함해서 Go, Python 등으로 개발할 수 있음

하지만 Kafka에서 공식적으로 제공하는 클라이언트는 Java만 지원
- 다른 언어의 경우 공식적으로 제공하는 것이 아님

따라서 성능적인 면이나 지원하는 범위에서 차이가 발생할 수 있음
- 그렇기 때문에 프로듀서 애플리케이션을 개발할 때는 Java 애플리케이션으로 개발하는 것이 좋음

## 프로듀서 내부 구조
![image](https://github.com/ruthetum/study/assets/59307414/b68d6a50-af7b-40e5-b9d3-64368cfbcb86)

ProducerRecord: 프로듀서에서 생성하는 레코드. 오프셋은 미포함(파티션에 데이터가 저장될 때 오프셋이 지정)

send(): 레코드 전송 요청 메서드
- send()를 호출한다고 바로 전송되는 것은 아님. partitioner, accumulator, sender를 통해 레코드 전송

Partitioner: 어느 파티션으로 전송할지 지정하는 파티셔너. 기본값으로 DefaultPartitioner 사용
- 메세지 키에 따라 파티션을 분배하는 게 파티셔너의 역할

Accumulator: 배치로 묶어 전송할 데이터를 모으는 버퍼
- 파티셔너에 의해 어떤 파티션으로 전송할지 결정됐다면 Accumulator에서는 배치로 데이터를 묶음
- 배치로 묶여진 데이터는 특정 시점 또는 크기에 따라 Sender에서 전송

## 파티셔너 (Partitioner)
프로듀서 API를 사용하면 UniformStickyPartitioner, RoundRobinPartitioner 2개의 파티셔너를 제공함

카프카 클라이언트 라이브러리 2.5.0 버전부터 파티셔너를 지정하는 경우 UniformStickyPartitioner가 기본 파티셔너로 기본 설정

### 메세지 키가 있을 경우
- UniformStickyPartitioner, RoundRobinPartitioner 둘 다 메세지 키가 있을 때는 메세지 키의 해시값과 파티션을 매칭하여 레코드를 전송
- 동일한 메세지 키가 존재하는 레코드는 동일한 파티션 번호에 전달됨
- 만약 파티션 개수가 변경될 경우 메세지 키와 파티션 번호 매칭은 깨지게 됨

### 메세지 키가 없는 경우
- 메세지 키가 없을 때는 파티션에 최대한 동일하게 분배하는 로직이 포함
- UniformStickyPartitioner, RoundRobinPartitioner의 단점을 개선

#### RoundRobinPartitioner
- ProducerRecord가 들어오는 대로 파티션을 순화하면서 전송
- Accumulator에 묶이는 정도가 작기 떄문에 전송 성능이 낮음

#### UniformStickyPartitioner
- Accumulator에서 레코드들이 배치로 묶일 때까지 기다렸다가 전송
- 배치로 묶일 뿐 결국 파티션을 순회하면서 보내기 때문에 모든 파티션에 분배되어 전송

### Custom Partitioner
제공되는 Partitioner 인터페이를 이용하여 활용

## 주요 옵션
### 필수 옵션
| 옵션 | 설명                                    |
| --- |---------------------------------------|
| bootstrap.servers | 카프카 클러스터의 브로커 정보를 설정. 1개 이상의 브로커 정보 작성 |
| key.serializer | 레코도의 메세지 키를 직렬화할 클래스 설정               |
| value.serializer | 레코드의 메세지 값을 직렬화할 클래스 설정             |

serializer의 경우 특별한 상황이 아니면 string(StringSerializer) 활용
- kafka-console-consumer.sh, kafka-console-producer.sh를 사용할 경우 기본값으로 string을 사용하기 때문에 string을 사용하는 것이 좋음
- 다양하게 consumer를 활용하는 경우, 어떤 방식으로 직렬화되어있는지 모를 수도 있음

### 선택 옵션
| 옵션 | 설명                                                                                                                     |
| --- |------------------------------------------------------------------------------------------------------------------------|
| acks | 프로듀서가 전송한 데이터가 브로커들에 정상적으로 저장되었는지 전송 성공 여부를 확인하는 데에 사용하는 옵션.<br/>0, 1, all(-1)로 설정 가능. 기본 값은 1                         |
| linger.ms | 배치를 전송하기 전까지 기다리는 최소 시간. 기본 값은 0.                                                                                      |
| retries | 브로커로부터 에러를 받고난 뒤 재전송을 시도하는 횟수. 기본 값은 2147483647.                                                                       |
| max.in.flight.requests.per.connection | 한 번에 요청하는 최대 커넥션 개수. 설정된 값만큼 동시에 전달 요청을 수행. 기본 값은 5                                                                    |
| partitioner.class | 레코드를 파티션에 전송할 때 적용하는 파티셔너 클래스를 설정.<br/>기본 값은 org.apache.kafka.clients.producer.internals.DefaultPartitioner.           |
| enable.idempotence | 멱등성 프로듀서로 동작할지 여부를 설정. 기본 값은 false. (3.x.x부터는 true)                                                                    |
| transactional.id | 프로듀서가 레코드를 전송할 때 레코드를 트랜잭션 단위로 묶을지 여부 설정.<br/>enable.idempotence 옵션을 true로 설정하면 자동으로 idempotence는 true로 설정. 기본 값은 null |

## ISR (In-Sync-Replicas)와 acks 옵션
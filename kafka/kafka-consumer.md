# 컨슈머 (Consumer)

<img width="771" alt="스크린샷" src="https://github.com/ruthetum/study/assets/59307414/836a8892-1dd8-439b-bbc9-dd38307d4e0c">

프로듀서가 전송한 데이터는 카프카 브로커에 적재

컨슈머는 적재된 데이터를 사용하기 위해 브로커로부터 데이터를 가져와서 필요한 처리를 진행
- 예를 들어 마케팅 문자를 고객에게 보내는 기능이 있다면 컨슈머는 토픽으로부터 고객 데이터를 가져와서 문자 발송 처리를 진행

## 컨슈머 내부 구조

<img width="802" alt="스크린" src="https://github.com/ruthetum/study/assets/59307414/a095c1fa-4846-48c6-b274-0abe5560f352">

- Fetcher: 리더 파티션으로부터 레코들을 미리 가져와서 대기
- poll(): Fetcher에 있는 레코드들을 리턴하는 레코드
  - completedFetches에서 데이터를 충분히(최대 max.poll.records 만큼) 가져오는 경우 poll 호출
- ConsumerRecords: 처리하고자 하는 레코드들의 모음. 오프셋 포함
  - 처리가 완료된 경우 commit을 통해 처리한 offset을 판단

## 컨슈머 그룹

<img width="1086" alt="스크린샷" src="https://github.com/ruthetum/study/assets/59307414/02ee3a3a-125f-41f7-a5a3-cf28a66dc7f2">

컨슈머 그룹으로 운영하는 방법은 컨슈머를 각 컨슈머 그룹으로부터 격리된 환경에서 안전하게 운영할 수 있도록 도와주는 카프카의 독특한 방식

컨슈머 그룹으로 묶인 컨슈머들은 토픽의 1개 이상 파티션들에 할당되어 데이터를 가져갈 수 있음

컨슈머 그룹으로 묶인 컨슈머가 토픽을 구독해서 데이터를 가져갈 때, 1개의 파티션은 최대 1개의 컨슈머에 할당 가능함

1개 컨슈머는 여러 개의 파티션에 할당될 수 있음. 따라서 컨슈머 그룹의 컨슈머 개수는 가져가고자 하는 토픽의 파티션 개수보다 같거나 작아야 함

#### 컨슈머 그룹의 컨슈머가 파티션 개수보다 많을 경우
1개의 파티션은 최대 1개의 컨슈머에 할당 가능하기 때문에, 파티션을 할당받지 못 하고 유휴 상태로 남는 컨슈머가 존재

파티션을 할당받지 못한 컨슈머는 스레드만 차지하고 실질적인 데이터 처리를 하지 못하므로 애플리케이션 실행에 있어 불필요한 스레드로 남게 됨

### 컨슈머 그룹을 활용하는 이유

<img width="631" alt="스크린샷" src="https://github.com/ruthetum/study/assets/59307414/47da2721-050e-4337-a8cb-714457de0a4b">

운영 서버의 주요 리소스인 CPU, Memory 정보를 수집하는 데이터 파이프라인을 구축

실시간 리소스를 시간 순으로 확인하기 위해 데이터를 엘라스틱서치에 저장하고, 이와 동시에 대용량 적재를 위해 하둡에 적재

만약 카프카를 활용한 파이프라인이 아니라면 서버에서 실행되는 리소스 수집 및 전송 에이터전트는 수집한 리소르를 엘라시틱서치와 하둡에 적재하기 위해 동기적으로 적재 요청

이렇게 동기로 실행되는 에이전트는 엘라스틱서치 또는 하둡 둘 중에 하나에 장애가 발생한다면 더는 적재가 불가능할 수 있음

<img width="805" alt="스크린샷" src="https://github.com/ruthetum/study/assets/59307414/51ed47cd-9cc0-4f18-8d2e-e9e17015aab0">

카프카는 파이프라인을 운영함에 있어 최종 적재되는 저장소의 장애에 유연하게 대응할 수 있도록 각각 다른 저장소에 저장하는 컨슈머를 다른 컨슈머 그룹으로 묶음으로써 각 저장소의 장애에 격리되어 운영 가능

엘라스틱서치의 장애로 인해 더는 적재하지 못 하더라도 하둡으로 데이터를 적재하는 데에는 문제가 없음

엘라스틱서치의 장애가 해소되면 엘라스틱서치로 적재하는 컨슈머의 컨슈머 그룹은 마지막으로 적재 완료한 데이터 이후부터 다시 적재를 수행하여 최종적으로 정상화 가능

## 리밸런싱

<img width="866" alt="스크린샷" src="https://github.com/ruthetum/study/assets/59307414/765777b8-b84e-4062-895b-e61bd0b496ff">

컨슈머 그룹으로 이루어진 컨슈머들 중 일부 컨슈머에 장애가 발생하면, 장애가 발생하면 자애가 발생한 컨슈머에 할당된 파티션은 장애가 발생하지 않은 컨슈머에 소유권이 넘어감
- 이러한 과정을 리밸런싱이라고 부름

리밸런싱은 크게 두 가지 상황에서 일어나는데, 첫 번째는 컨슈머가 추가되는 상황이고, 두 번째는 컨슈머가 제외되는 상황

이슈가 발생한 컨슈머를 컨슈머 그룹에서 제외하여 모든 파티션이 지속적으로 데이터를 처리할 수 있도록 가용성을 높여줌

리밸런싱은 컨슈머가 데이터를 처리하는 도중에 언제든지 발생할 수 있기 때문에 데이터 처리 중 발생한 리밸런싱에 대응하는 코드를 작성해야 함

## 커밋

<img width="826" alt="스크린샷" src="https://github.com/ruthetum/study/assets/59307414/dd14b0d0-03f8-400a-89ad-c965d7362463">

컨슈머는 카프카 브로커로부터 데이터를 어디까지 가져갔는지 커밋(commit)을 통해 기록

특정 토픽의 파티션을 어떤 컨슈머 그룹이 몇 번째로 가져갔느지 카프카 브로커 내부에서 사용되는 내부 토픽(`__consumer_offsets`)에 기록

컨슈머 동작 이슈가 발생하여 `__consumer_offsets` 토픽에 어느 레코드까지 읽어갔는지 오프셋 커밋이 기록되지 못 하면 데이터 처리의 중복이 발생할 수 있음

따라서 데이터 처리의 중복이 발생하지 않게 하기 위해서는 컨슈머 애플리케이션이 오프셋 커밋을 정상적으로 처리했는지 검증해야 함

## 어사이너

컨슈머와 파티션 할당 정책은 컨슈머의 어사이너(Assignor)에 의해 결정

카프카에서는 RangeAssignor, RoundRobinAssignor, StickyAssignor 세 가지 어사이너를 제공
- RangeAssignor: 각 토픽에서 파티션을 숫자로 정렬, 컨슈머를 사전 순서로 정렬하여 할당
- RoundRobinAssignor: 모든 파티션을 컨슈머에서 번갈아가면서 할당
- StickyAssignor: 최대한 파티션을 균등하게 배분하면서 할당

카프카 2.5.0은 RangeAssignor가 기본값으로 설정

## 컨슈머 주요 옵션
### 필수 옵션
- `bootstrap.servers`: 프로듀서가 데이터를 전송할 대상 카프카 클러스터에 속한 브로커의 호스트 이름:포트를 1개 이상 작성
- `key.deserializer`: 레코드의 메세지 키를 역직렬화하는 클래스를 지정
- `value.deserializer`: 레코드의 메세지 값을 역직렬화하는 클래스를 지정

### 선택 옵션
- `group.id`: 컨슈머 그룹 아이디를 지정. 기본 값은 null
- `auto.offset.reset`: 컨슈머 그룹이 특정 파티션을 읽을 때 저장된 컨슈머 오프셋이 없는 경우 어느 오프셋이 읽을지 선택. 컨슈머 오프셋이 있는 경우 옵션 값은 무시. 기본 값은 latest
- `enable.auto.commit`: 컨슈머 오프셋을 자동으로 커밋할지 여부를 지정. 기본 값은 true
- `auto.commit.interval.ms`: 자동 커밋일 경우 오프셋 커밋 주기를 지정. 기본 값은 5000(5초)
- `max.poll.records`: poll() 메서드를 통해 반환되는 레코드 개수를 지정. 기본 값은 500
- `session.timeout.ms`: 컨슈머가 브로커와 연결이 끊기는 최대 시간. 기본 값은 10000(10초)
- `heartbeat.interval.ms`: 컨슈머가 브로커에 하트비트를 보내는 주기. 기본 값은 3000(3초)
- `max.poll.interval.ms`: poll() 메서드를 호출하는 간격의 최대 시간. 기본 값은 300000(5분)
- `isololation.level`: 트랜잭션 프로듀서가 레코드를 트랜잭션 단위로 보낼 경우 사용

#### auto.offset.reset
컨슈머 그룹이 특정 파티션을 읽을 때 저장된 컨슈머 오프셋이 없는 경우 어느 오프셋이 읽을지 선택하는 옵션

컨슈머 오프셋이 있다면 옵션 값은 무시

옵션은 latest, earliest, none 중 1개를 설정할 수 있음
- latest: 설정하면 가장 높은(가장 최근에 넣은) 오프셋부터 읽기 시작
- earliest: 설정하면 가장 낮은(가장 오래된) 오프셋부터 읽기 시작
- none: 설정하면 컨슈머 그룹이 커밋한 기록이 있는지 찾아봄. 만약 커밋 기록이 없으면 오류를 반환하고, 커밋 기록이 있다면 기존 커밋 기록 이후 오프셋부터 읽기 시작

## 리밸런스 리스너를 가진 컨슈머
리밸런스 발생을 감지하기 위해 카프카 라이브러리는 ConsumerRebalanceListener 인터페이스를 지원

ConsumerRebalanceListener 인터페이스로 구현된 클래스는 onPartitionAssigned() 메서드와 onPartitionRevoked() 메서드로 구성

- onPartitionAssigned(): 리밸런스가 끝난 뒤에 파티션이 할당 완료되면 호출되는 메서드
- onPartitionRevoked(): 리밸런스가 시작되기 직전에 호출되는 메서드. 마지막을 처리한 레코드를 기준으로 커밋을 하기 위해서는 리밸런스가 시작하기 직전에 커밋을 하면 되므로 onPartitionRevoked() 메서드에서 커밋을 구현하여 처리

```java
public class RebalanceListener implements ConsumerRebalanceListener {

    ...
    
    @Override
    public void onPartitionsAssigned(Collection<TopicPartition> partitions) {
        // 리밸런스가 끝난 뒤에 파티션이 할당 완료되면 호출되는 메서드
    }

    @Override
    public void onPartitionsRevoked(Collection<TopicPartition> partitions) {
        // 리밸런스가 시작되기 직전에 호출되는 메서드
        // 마지막을 처리한 레코드를 기준으로 커밋을 하기 위해서는 리밸런스가 시작하기 직전에 커밋을 하면 됨
        // onPartitionsRevoked() 메서드에서 커밋을 구현하여 처리
        consumer.commitSync(currentOffsets);
    }
}
```

## 파티션 할당 컨슈머 애플리케이션

일반적으로 사용하기보다는 특별한 파티션에 할당해야 하는 경우 사용

```java
private final static int PARTITION_NUMBER  = 0;

public static void main(String[] args) {

        Properties configs = new Properties();
        configs.put(ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG, BOOTSTRAP_SERVERS);
        configs.put(ConsumerConfig.KEY_DESERIALIZER_CLASS_CONFIG, StringDeserializer.class.getName());
        configs.put(ConsumerConfig.VALUE_DESERIALIZER_CLASS_CONFIG, StringDeserializer.class.getName());

        KafkaConsumer<String, String> consumer = new KafkaConsumer<>(configs);
        consumer.assign(Collections.singleton(new TopicPartition(TOPIC_NAME, PARTITION_NUMBER)));
        while (true) {
            ConsumerRecords<String, String> records = consumer.poll(Duration.ofSeconds(1));
            for (ConsumerRecord<String, String> record : records) {
                logger.info("record:{}", record);
            }
        }
}
```

## 컨슈머의 종료
```java
public static void main(String[] args) {

  // shutdown hook 등록을 통해 종료 설정
  Runtime.getRuntime().addShutdownHook(new ShutdownThread());
  
  ...
  
  cosumer = new KafkaConsumer<>(configs);
  consumer.subscribe(Arrays.asList(TOPIC_NAME));
  
  try {
    while (true) {
      ConsumerRecords<String, String> records = consumer.poll(Duration.ofSeconds(1));
      for (ConsumerRecord<String, String> record: records) {
        ...
      }
    }
  } catch (WakeupException e) {
    logger.warn("WakeupException");
  } finally {
    consumer.close();
  }
}

class ShutdownThread extends Thread {
  public void run() {
    consumer.wakeup();
  }
}
```

컨슈머 애플리케이션이 정상적으로 종료되지 않는 경우 세션 타임아웃이 발생할 때가지 컨슈머 그룹에 남게 됨

컨슈머를 안전하게 종료하기 위해 KafkaConsumer 클래스는 wakeup() 메서드를 지원

wakeup() 메서드가 실행된 이후 poll() 메서드가 호출되면 WakeupException 예외 발생

WakeupException 예외를 받은 뒤에는 데이터 처리를 위해 사용한 자원들을 해제

## 멀티스레드 컨슈머 애플리케이션
<img width="680" alt="스크린샷" src="https://github.com/ruthetum/study/assets/59307414/5cfca86c-42a7-4654-9756-38ae80694d74">

카프키는 처리량을 늘리기 위해 파티션과 컨슈머 개수를 늘려서 운영할 수 있음

파티션을 여러 개로 운영하는 경우 데이터를 병렬처리하기 위해 파티션 개수와 컨슈머 개수를 동일하게 맞추는 것이 가장 좋은 방법
- 토픽의 파티션은 1개 이상으로 이루어져 있으며 1개의 파티션은 1개 컨슈머가 할당되어 데이터를 처리할 수 있음

파티션 개수가 n개 라면 동일 컨슈머 그룹으로 묶인 컨슈머 스레드를 최대 n개 운영할 수 있음
- n개의 스레드를 가진 1개의 프로세스를 운영하거나 1개의 스레드를 가진 프로세를 n개 운영하는 방법도 존재

## 컨슈머 랙 (Consumer Lag)
<img width="616" alt="스크린샷" src="https://github.com/ruthetum/study/assets/59307414/cbae530b-2786-464e-853f-3839f6f04292">

> 파티션의 최신 오프셋(LOG-END-OFFSET): 4
> 
> 컨슈머 오프셋(CURRENT-OFFSET): 2
> 
> 컨슈머 랙: 2
> - 레코드가 2개만큼 지연이 발생했다는 의미


컨슈머 랙은 파티션의 최신 오프셋(LOG-END-OFFSET)과 컨슈머 오프셋(CURRENT-OFFSET) 간의 차이

프로듀서는 계속해서 새로운 데이터를 파티션에 저장하고, 컨슈머는 자신이 처리할 수 있는 만큼 데이터를 가져감

컨슈머 랙은 컨슈머가 정상 동작하는지 여부를 확인할 수 있기 때문에 컨슈머 애플리케이션을 운영한다면 필수적으로 모니터링해야 하는 지표

### 모니터링할 수 있는 컨슈머 랙의 수
<img width="632" alt="스크린샷" src="https://github.com/ruthetum/study/assets/59307414/61b44a1d-f3ac-4c94-aa9d-31c34e857785">

> (모니터링할 수 있는 컨슈머 랙의 수) = (토픽의 파티션 개수) * (컨슈머 그룹 개수)
> 
> 1개의 토픽에 3개의 파티션 있고, 1개의 컨슈머 그룹이 토픽을 구독하여 데이터를 가져가면, 모니터링할 수 있는 컨슈머 랙은 총 3개

컨슈머 랙은 컨슈머 그룹과 토픽, 파티션 별로 생성
  
### 컨슈머 랙의 발생
컨슈머 랙은 프로듀서와 컨슈머의 데이터 처리량의 차이로 인해 발생

프로듀서가 보내는 데이터 양이 컨슈머의 데이터 처리량보다 크다면 컨슈머 랙은 늘어남

반대로 프로듀서가 보내는 데이터 양이 컨슈머의 데이터 처리량보다 적으면 컨슈머 랙은 줄어들고, 최솟값은 0으로 지연이 없음을 의미

### 컨슈머 랙을 모니터링하는 이유

컨슈머 랙을 모니터링하는 것은 카프카를 통한 데이터 파이프라인을 운영하는 데에 핵심적인 역할을 수행

컨슈머 랙을 모니터링함으로써 컨슈머의 장애를 확인할 수있고 파티션 개수를 정하는 데에 참고할 수 있음

### 컨슈머 랙 관련 이슈
#### 1. 처리량 이슈
프로듀서의 데이터 양이 늘어나는 경우 컨슈머 랙이 늘어날 수 있음

이 때는 파티션 개수와 컨슈머 개수를 늘려 병렬 처리량을 증가시켜 컨슈머 랙을 줄일 수 있음

#### 2. 파티션 이슈
프로듀서의 데이터 양이 일정함에도 불구하고 컨슈머의 장애로 인해 컨슈머 랙이 증가할 수 있음

컨슈머는 파티션 개수만큼 늘려서 병렬 처리하며 파티션마다 컨슈머가 할당되어 데이터를 처리

e.g. 2개의 파티션으로 구성된 토픽에 2개의 컨슈머가 각각 할당되어 데이터를 처리한다고 가정할 때

- 프로듀서가 보내는 데이터 양은 동일한데 파티션 1번의 컨슈머 랙이 늘어나는 상황이 발생한다면, 1번 파티션에 할당ㅇ된 컨슈머에 이슈가 발생했음을 유추할 수 있음

### 컨슈머 랙 모니터링 방법
#### 1. 카프카 명령어 사용
`kafka-consumer-groups.sh` 명령어를 사용하면 컨슈머 랙을 포함한 특정 컨슈머 그룹의 상태를 확인할 수 있음

가장 기초적인 방법으로, 명령어를 통해 컨슈머 랙을 확인한느 방법은 일회성에 그치고, 지표를 지속적으로 기록하고 모니터링하는데 어려움이 있음

따라서 일시적으로나 테스트용으로 활용

#### 2. metrics() 메서드 사용
컨슈머 애플리케이션에서 KafkaConsumer 인스턴스의 `metrics()` 메서드를 활용하면 컨슈머 랙 지표를 확인할 수 있음

컨슈머 인스턴스가 제공하는 컨슈머 랙 관련 모니터링 지표는 `records-lag-max`, `records-lag-avg`, `records-lag` 세 가지를 제공

하지만 metrics() 메서드를 사용할 경우 아래의 이슈 존재
- 컨슈머가 정상 동작할 때에만 지표를 확인할 수 있음. metrics() 메서드는 컨슈머가 정상적으로 실행될 경우에만 호출
  - 따라서 컨슈머 애플리케이션이 비정상적으로 종료되면 더는 컨슈머 랙을 모니터링할 수 없음
- 모든 컨슈머 애플리케이션에 컨슈머 랙 모니터링 코드를 중복해서 작성해야 함
  - 컨슈머 애플리케이션을 여러 종류로 운영할 경우 각기 다른 컨슈머 애플리케이션에 metrics() 메서드를 호출하여 컨슈머 랙을 수집하는 로직을 중복해서 작성해야 함
  - 특정 컨슈머 그룹에 해당하는 애플리케이션이 수집하는 컨슈머 랙은 자기 자신 컨슈머 그룹에 대해서만 한정되기 때문
- 컨슈머 랙을 모니터링하는 코드를 추가할 수 없는 서드 파트 애플리케이션는 컨슈머 랙 모니터링이 불가

#### 3. 외부 모니터링 툴 사용
컨슈머 랙을 모니터링하는 가장 최선의 방법은 외부 모니터링 툴을 사용하는 것

data dog, confluent control center와 같은 카프카 클러스터 종합 모니터링 툴을 사용하면 카프카 운영에 필요한 다양한 지표를 모니터링할 수 있음

모니터링 지표에는 컨슈머 랙도 포함되어 있기 때문에 클러스터 모니터링과 함께 컨슈머 랙을 함께 모니터링하기에 적합함

컨슈머 랙 모니터링만을 위한 툴로는 버로우(Burrow)가 존재

## 버로우
> https://github.com/linkedin/Burrow

버로우는 링크드인에서 개발하여 오픈소스로 공개한 컨슈머 랙 체크 툴로 REST API를 통해 컨슈머 그룹 별로 컨슈머 랙을 확인할 수 있음

외부 모니터링 툴을 사용하면 카프카 클러스터에 연결된 모든 컨슈머, 토픽들의 랙 정보를 한 번에 모니터링할 수 있음

또한, 모니터링 툴들은 클러스터와 연동되어 컨슈머의 데이터 처리와는 별개로 지표를 수집하기 때문에 데이터를 활용하는 프로듀서난 컨슈머의 동작에 영향을 미치지 않음

버로우는 다수의 카프카 클러스터를 동시에 연결하여 컨슈머 랙을 확인하고, 일반적인 기업 환경에서는 카프카 클러스터를 여러 개 운영하는 경우가 많기 때문에 한 번의 설정으로 다수의 카프카 클러스터 컨슈머 랙을 확인할 수 있음

### 버로우의 컨슈머 랙 평가
버로우는 단순 임계치가 아닌 슬라이딩 윈도우 계산을 통해 문제가 생긴 파티션과 컨슈머의 상태를 표현
- 이렇게 버로우에서 컨슈머 랙의 상태를 표현하는 것을 컨슈머 랙 평가(Evaluation)이라고 함

컨슈머 랙과 파티션의 오프셋을 슬라이딩 윈도우로 계산하면 상태가 정해짐

결과적으로 파티션의 상태 OK, STALLED, STOPPED 로 표현하고, 컨슈머의 상태를 OK, WARNING, ERROR 로 표현함
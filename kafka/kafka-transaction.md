# 트랜잭션 프로듀서

## 트랜잭션 프로듀서의 동작

<img width="654" alt="스크린샷" src="https://github.com/ruthetum/study/assets/59307414/8ff58290-2e9f-4143-8963-33670b5715a3">

카프카에서 트랜잭션은 다수의 파티션에 데이터를 저장할 경우 모든 데이터에 대해 동일한 원자성(atomic)을 만족시키기 위해 사용

원자성을 만족시킨다는 의미는 다수의 데이터를 동일 트랜잭션으로 묶음으로써 전체 데이터를 처리하거나 전체 데이터를 처리하지 않도록 하는 것을 의미

트래잭션 프로듀서는 사용자가 보낸 데이터를 레코드로 파티션에 저장할 뿐만 아니라 트랜잭션의 시작과 끝을 표현하기 위해 트랜잭션 레코드를 한 개 보냄

## 트랜잭션 컨슈머의 동작

<img width="520" alt="스크린샷" src="https://github.com/ruthetum/study/assets/59307414/6e3ecd41-f99e-42dc-a65a-c13f6f640892">

> 그림에서 위의 경우 트랜잭션 컨슈머는 데이터를 가져가지 않음
> 
> 아래의 경우처럼 commit이 있는 경우 데이터를 가져감

트랜잭션 컨슈머는 파티션에 저장된 트랜잭션 레코드를 보고 트랜잭션이 완료(commit)되었음을 확인하고 데이터를 가져감

트랜잭션 레코드는 실질적인 데이터는 가지고 있지 않으며 트랜잭션이 끝난 상태를 표시하는 정보만 가지고 있음

## 트랜잭션 프로듀서 설정
```Java
configs.put(ProducerConfig.TRANSACTIONAL_ID_CONFIG, UUID.randomUUID());

Producer<String, String> producer = new KafkaProducer<>(configs);

producer.initTransactions();

producer.beginTransaction();
producer.send(new ProducerRecord<>("topic", "test message 1"));
producer.send(new ProducerRecord<>("topic", "test message 2"));
producer.send(new ProducerRecord<>("topic", "test message 3"));
producer.commitTransaction();

producer.close();
```

트랜잭션 프로듀서로 동작하기 위해서는 transactional.id 설정이 필요

프로듀서 별로 고유한 ID 값을 사용 

init, begin, commit 순서대로 수행

## 트랜잭션 컨슈머 설정
```Java
configs.put(CommonClientConfigs.ISOLATION_LEVEL_CONFIG, "read_committed");

KafkaConsumer<String, String> consumer = new KafkaConsumer<>(configs);
```

트랜잭션 컨슈머는 커밋이 완료된 레코들만 읽기 위해 isolation.level 옵션을 read_committed로 설정

기본 값은 read_uncommitted로 트랜잭션 프로듀서가 레코드를 보낸 후 커밋 여부와 상관없이 모두 읽음

read_committed로 설정하면 커밋이 완료된 레코드만 읽음
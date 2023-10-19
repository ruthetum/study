# 멱등성 프로듀서 
> 멱등성: 연산을 여러 번 적용하더라도 동일한 결과를 나태내는 것
> 
> 멱등성 프로듀서: 동일한 데이터를 여러 번 전송하더라도 카프카 클러스터에 단 한 번만 저장됨을 의미

## 전달 신뢰성
> - At most once: 최대 한 번 전달
> - At least once: 적어도 한 번 이상 전달
> - Exactly once: 정확히 한 번 전달

기본 프로듀서의 동작 방식은 적어도 한 번 전달(at least once delivery)을 지원

적어도 한 번 전달한다는 것은 프로듀서가 클러스터에 데이터를 전송하여 저장할 때 적어도 한 번 이상 데이터를 적재할 수 있고, 데이터가 유실되지 않음을 의미

다만 두 번 이상 적재할 가능성이 있기 때문에 데이터 중복이 발생할 수 있음

## 멱등성 프로듀서
프로듀서가 보내는 데이터의 중복 적재를 막기 위해 enable.idempotence 옵션을 사용해서 exactly once delivery 지원

카프카 3.0.0부터는 enable.idempotence 옵션 값이 true(acks=all)로 변경되므로 필요에 따라 옵션값 수정 필요
- enable.idempotence 의 기본값(카프카 3.0.0 이전)은 false
- https://www.conduktor.io/kafka/idempotent-kafka-producer/#How-should-Kafka-producer-idempotence-be-enabled?-1


### 멱등성 프로듀서의 동작
멱등성 프로듀서는 기본 프로듀서와 달리 데이터를 브로커로 전달할 때 프로듀서 PID(Producer Unique ID)와 시퀀스 넘버(Sequence number)를 함께 전달

그러면 브로커는 프로듀서의 PID와 시퀀스 넘버를 확인하여 동일한 메세지의 적재 요청이 오더라도 단 한 번만 데이터를 적재함으로써 프로듀서의 데이터는 정확히 한번 브로커에 적재되도록 동작

- PID(Producer Unique ID): 프로듀서의 고유한 ID
- SID(Sequence ID): 레코드의 전달 번호 ID

### 멱등성 프로듀서의 한계
멱등성 프로듀서는 동일한 세션에서만 정확히 한번 전달을 보장
- 동일한 세션: PID의 생명 주기

만약 멱등성 프로듀서로 동작하는 프로듀서 애플리케이션에 이슈가 발생하여 종료되고 애플리케이션을 재시작하면 PID가 달라짐

동일한 데이터를 보내더라도 PID가 달라지면 브로커 입장에서 다른 프로듀서 애플리케이션이 다른 데이터를 보냈다고 판단하기 때문에 멱등성 프로듀서는 장애가 발생하지 않을 경우에만 정확히 한번 적재하는 것을 보장

### 멱등성 프로듀서로 설정할 경우 옵션
멱등성 프로듀서를 사용하기 위해 enable.idempotence를 true로 설정하면 exactly once delivery가 성립되기 위해 프로듀서의 일부 옵션들이 강제로 설정됨

프로듀서의 데이터 재전송 횟수를 정하는 retries는 기본값으로 Integer.MAX_VALUE로 설정되고 acks 옵션은 all로 설정됨
- 이렇게 설정되는 이유는 프로듀서가 적어도 한 번 이상 브로커에 데이터를 보냄으로써 브로커에 단 한 번만 데이터가 적재되는 것을 보장하기 위함

멱등성 프로듀서는 정확히 한 번 브로커에 데이터를 적재하기 위해 단 한번만 데이터를 전송하는 것은 아님
- 상황에 따라 프로듀서가 여러 번 데이터를 전송하지만 브로커가 전송된 데이터를 확인하고 중복된 데이터는 적재하지 않는 것

### 멱등성 프로듀서 사용 시 오류 확인
멱등성 프로듀서의 시퀀스 넘버는 0부터 시작하여 숫자를 1씩 더한 값이 전달

브로커에서 멱등성 프로듀서가 전송한 데이터의 PID와 시퀀스 넘버를 확인하는 과정에서 시퀀스 넘버가 일정하지 않은 경우 예외가 발생할 수 있음
- OutOfOrderSequenceException: 브로커가 예상한 시퀀스 넘버와 다른 번호의 데이터 적재 요청이 왔을 때 발생하는 예외

- OutOfOrderSequenceException이 발생했을 경우에는 시퀀스 넘버의 역전현상이 발생할 수 있기 때문에 순서가 중요한 데이터를 전송하는 프로듀서는 해당 Exception이 발생했을 경우 대응하는 방안을 고려해야 함
[apache kafka](https://kafka.apache.org/), [rabbitmq](https://www.rabbitmq.com/), [nats](https://nats.io/) 대상으로 진행합니다.

다양한 message queue가 존재하지만, 현재 기준으로는 요금이 발생하는 관리형 서비스(managed service)는 제외하고 무료로 운용 가능한 message queue를 고려합니다.

일반적으로 많이 사용하는 kafka, rabbitmq와 더불어 이전에 전체적으로 사용 경험이 있었던 nats를 포함합니다.

# 특징
## Apache Kafka
### 소개
2011년 링크드인에서 개발된 빅데이터를 처리하기 위한 오픈 소스 메시지 브로커(분산 메세징 시스템, 이벤트 스트리밍 플랫폼)

### 작성 언어
Java, Scala

### 주요 특징
High throughput
- 많은 양의 데이터를 묶음 단위로 처리하는 배치로 빠르게 처리하기 때문에 높은 처리량을 보장

Scalable
- 클러스터 내 브로커 수를 조절하여 가변적인 환경에서 확장성있게 운영 가능

Persistent Stroage
- 디스크에 메시지를 저장하기 때문에 메시지 손실이 없음(데이터를 생성한 프로그램이 종료되더라도 사라지지 않음)

High Availbility
- 클러스터 내 일부 서버에 장애가 발생하더라도 무중단으로 안전하고 지속적으로 데이터를 처리할 수 있음
- 데이터 복제(replication)를 통해 고가용성 보장

### 사용 회사
- Apple, Netflix, Uber, PayPal, Spotify, Slack ...
- More than 80% of all Fortune 100 companies trust, and use Kafka.

## RabbitMQ
### 소개
AMQP(Advanced Message Queuing Protocol)기반의 오픈 소스 메시지 브로커
-  AMQP: MQ의 오픈소스에 기반한 표준 프로토콜 (특정 플랫폼에 종속된 프로토콜이 아님)

메세지를 많은 사용자에게 전달하거나, 요청에 대한 처리 시간이 길 때, 요청을 다른 API에게 위임하고 빠른 응답을 할 때 많이 사용

### 작성 언어
Erlang

### 주요 특징
In-built Clustering
- 하나의 노드에 장애가 발생하더라도 소비자와 생산자는 이벤트에서 계속 작동하고 추가 노드를 추가하여 메시징 처리량을 선형적으로 확장 가능

Flexible Routing
- 일반적인 라우팅부터 복잡한 라우팅까지 다양하고 유연한 라우팅 방법을 작업

Reliability
- Persistence(지속성), delivery feedback, publisher confirms(브로커에게 도달 확인), and high availability(고가용성) 제공

Management & Monitoring
- 관리와 모니터링을 위한 편의 기능 제공

### 사용 회사
Raddit, CircleCI, Trivago, Runtastic ...

## NATS
### 소개
NATS는 클라우드 네이티브 애플리케이션, IoT 메시징 및 마이크로 서비스 아키텍처를 위한 단순하고 안전한 고성능 오픈 소스 메시징 브로커

Kafka와 같이 Streaming 기반의 전송으로 Pub/Sub 기반의 데이터 처리가 가능

### 작성 언어
Go

### 주요 특징
High throughput
- 타 오픈소스 메시지 브로커와 비교 시 우월한 성능 보유

Deploy anywhere
- 가볍고 간단하기 때문에 어떤 환경에서든 배포 가능하고 활용 가능

Scalable
- scale out 에 용이한 구조

### 사용 회사
Zenly, Teachable, LaunchDarkly ...


# 비교
![image](https://img1.daumcdn.net/thumb/R1280x0/?scode=mtistory2&fname=https%3A%2F%2Fblog.kakaocdn.net%2Fdn%2FdZZVgo%2FbtrtLVoCBsc%2FBKVzrag6LtGf21av6GkgkK%2Fimg.png)

메시지 큐를 선택하기 위해 아래 6가지 항목을 확인
- [x] 언어 지원
- [x] 처리량
- [x] 메세지 전달
- [x] 고가용성
- [x] 관리 및 모니터링
- [x] 참고자료

클러스터 내부에서 연결 및 통신을 진행하기 때문에 보안 관련 부분(AuthN, AuthZ)은 당장의 우선순위에서 높게 측정하지 않음

(개별적으로 보안 관련 부분은 지원하고 있기 때문에 필요 시 활용 가능할 것으로 판단)

## 언어 지원
고려하고 있는 모든 대상에 대해서 지원

| Project | Compatibility | Reference |
|:---:|:---:|:---:|
| Apache Kafka | O | https://docs.nestjs.com/microservices/kafka |
| RabbitMQ | O | https://docs.nestjs.com/microservices/rabbitmq |
| NATS | O | https://docs.nestjs.com/microservices/nats |

## 처리량
![image](https://s3.ap-northeast-2.amazonaws.com/zoyi-ghost/kr/2022/03/Untitled__1_-1648046959574.png)

| Project | Throughput |
|:---:|:---:|
| Apache Kafka | 100K msgs/s |
| RabbitMQ | 20-50K msgs/s | 
| NATS | 200K msgs/s | 

> NATS > Apache Kafka > RabbitMQ

성능 테스트 방법이나 메세지 크기에 따라 측정된 처리량 차이가 발생하지만 보편적으로 확인했을 때 위와 같음

NATS Streaming(JetStream)의 경우 Kafka와 비슷한 처리량을 보임
- ref. https://bravenewgeek.com/benchmarking-commit-logs/

당연하겠지만 메시지 크기가 크면 클수록 전체적으로 메세지 큐의 성능이 떨어짐
- 이에 대한 해법은 아니겠지만 zero payload 방식 적용 시 일정 부분 성능 개선 가능 (메시지 크기 감소)
- zero payload 방식의 경우 본래 목적은 메세지 크기를 줄이기 위한 방법이라기보다는 (큐의 비순차성 또는 중복 전송으로 인한) 멱등성을 보장하기 위한 방식

## 메세지 전달
![image](https://substackcdn.com/image/fetch/w_1456,c_limit,f_webp,q_auto:good,fl_progressive:steep/https%3A%2F%2Fbucketeer-e05bbc84-baa3-437e-9518-adb32be77984.s3.amazonaws.com%2Fpublic%2Fimages%2F933ba5a8-cf94-4da2-86f2-53e2dc57d5cc_1999x1215.png)

- At most once(최대 한번): 최대 한 번만 전송, 메시지를 한번만 전송하고 상대가 받았는지 받지 못했는지는 확인하지 않음
- At least once(최소 한번): 메시지를 전송하고 최소한 상대방이 하나의 메시지는 받았는지 확인
- Exactly once(정확히 한번): 메시지를 정확히 한번만 전송

| Project | QoS / Guarantees |
|:---:|:---:|
| Apache Kafka | At least once, Exactly once |
| RabbitMQ | At most once, At least once | 
| NATS (with JetStream) | At most once, At least once, Exactly once | 

## 고가용성
> Availability = Uptime / ( Uptime + Downtime )
- Availability(가용성): 애플리케이션 또는 시스템이 정상적으로 사용 가능한 정도
- High Availability(고가용성): 긴 시간동안 지속적으로 운영이 가능한 시스템이나 컴포넌트 (available 99.999%)

| Project | High Availability |
|:---:|:---:|
| Apache Kafka | 클러스터와 리플리케이션 구성으로 가용성 보장 |
| RabbitMQ | 클러스터를 활용하여 가용성 보장 | 
| NATS (with JetStream) | 클러스터를 활용하여 가용성 보장 | 


## 관리 및 모니터링

| Project | Management tools | Reference |
|:---:|:---:|:---:|
| Apache Kafka | 공식 관리 도구는 존재하지 않지만, 다양한 소스들이 존재 | https://towardsdatascience.com/overview-of-ui-monitoring-tools-for-apache-kafka-clusters-9ca516c165bd |
| RabbitMQ | 관리 및 모니터링을 하기 위한 대시보드 기능이 내장 | https://www.cloudamqp.com/blog/part3-rabbitmq-for-beginners_the-management-interface.html | 
| NATS (with JetStream) | API 형태로 metric을 제공하지만, 관리 및 모니터링에 용이하지 않음 | https://docs.nats.io/running-a-nats-service/nats_admin/monitoring |


## 참고자료

| Project | Github stars | Questions of Stackoverflow |
|:---:|:---:|:---:|
| Apache Kafka | 25.1K | 31,257 |
| RabbitMQ | 10.8K | 14,113 |
| NATS | 12.9k | 815 |

오랜 기간 지속적으로 많이 사용되었던 kafka와 rabbitmq가 래퍼런스는 많음

# 요약
세 가지 메세지큐 모두 현재 기준으로 도입하는데 문제는 없음
- 비동기 메세지 전송
- 처리량
- ...

일반적으로 kafka와 rabbitmq가 성숙(오랜기간 안정적으로 여러 곳에서 사용)하다고 평가
- 이에 따른 reference도 많이 존재

kafka와 rabbitmq를 비교할 때 트래픽의 크기가 많이 논의되었고, 안정성과 성능을 고려했을 때 대용량 처리를 위해서는 대부분 kafka를 권장
- 트래픽이 증가함에 따라 rabbit mq에서 kafka로 메세지큐를 변경했다는 경우도 존재

상대적으로 rabbitmq가 강점이 있다고 한 내용은 아래와 같음
- low latency
    - 물론 트래픽이 많아지면 처리량이 떨어지기 때문에 latency가 급증하지만, 트래픽이 수용 가능한 경우 상대적으로 latency가 낮음
    - 브로커/컨슈머간 접근 방법의 차이 때문에 발생
        - kafka: pull-based approach (smart consumer dumb broker)
        - rabbitmq: push-based approach (smart broker, dumb consumer)

- flexible routing
    - rabbitmq의 다양한 Exchange pattern(consumer들을 대상으로 메세지를 발행하는 방법)을 제공
        - consumer가 topic(subject)를 구독했을 때 다양한 패턴 및 조합으로 메세지를 전달할 수 있음
    - kafka의 경우 단순한 라우팅 방법만 지원하기 때문에 개별 마이크서비스에서 복잡한 방식으로 topic을 구독해야 하는 경우 rabbitmq가 kafka보다 유연하게 메세지를 전달할 수 있음

- user-friendly interface for management
    - rabbitmq의 경우 관리 및 모니터링을 하기 위한 대시보드 기능이 내장
    - kafka는 외부 OSS를 활용하여 지원

- 그 외에 kafka는 자체 프로토콜, rabbitmq는 AMQP 기반이기 때문에 상대적으로 다양한 플랫폼과 호환될 수 있음

# Reference
> recommendation
> - https://www.cloudamqp.com/blog/when-to-use-rabbitmq-or-apache-kafka.html
> - https://www.projectpro.io/article/kafka-vs-rabbitmq/451


- kafka: https://kafka.apache.org/
    - docs: https://kafka.apache.org/documentation/
    - who uses: https://stackshare.io/kafka
- rabbitmq: https://www.rabbitmq.com/
    - docs: https://www.rabbitmq.com/documentation.html
    - who uses: https://stackshare.io/rabbitmq
- nats: https://nats.io/
    - docs: https://docs.nats.io/
    - who uses: https://stackshare.io/nats
    - ko: https://www.joinc.co.kr/w/CNCF_NATS

- comparison
    - compare with nats: https://docs.nats.io/nats-concepts/overview/compare-nats
    - kafka vs rabbitmq
        - https://hevodata.com/learn/kafka-vs-rabbitmq/
        - https://tech.kakao.com/2021/12/23/kafka-rabbitmq/
        - https://www.upsolver.com/blog/kafka-versus-rabbitmq-architecture-performance-use-case
        - https://www.openlogic.com/blog/kafka-vs-rabbitmq
        - https://www.confluent.io/blog/kafka-fastest-messaging-system/
    - rabbitmq vs nats
        - https://streamnative.io/blog/comparison-of-messaging-platforms-apache-pulsar-vs-rabbitmq-vs-nats-jetstream

- performance
    - dissecting message queues: https://bravenewgeek.com/dissecting-message-queues/
    - kafka vs nats: https://bravenewgeek.com/benchmarking-commit-logs/
    - nats performance
        - https://www.slideshare.net/nats_io/microservices-meetup-san-francisco-august-2017-talk-on-nats
        - https://channel.io/ko/blog/real-time-chat-server-2-redis-pub-sub

- message delivery guarantees
    - https://blog.bytebytego.com/p/at-most-once-at-least-once-exactly
    - https://mentha2.tistory.com/259

- high availability
    - https://en.wikipedia.org/wiki/High_availability
    - https://www.techtarget.com/searchdatacenter/definition/high-availability\
    - kafka: https://firststep-de.tistory.com/30
    - rabbitmq: https://www.nakjunizm.com/2019/10/30/RabbitMQ_Cluster/
    - nats: https://www.joinc.co.kr/w/CNCF_NATS

- monitoring & dashboard
    - kafka monitoring tools: https://towardsdatascience.com/overview-of-ui-monitoring-tools-for-apache-kafka-clusters-9ca516c165bd
    - kafka-ui: https://github.com/provectus/kafka-ui
        - https://towardsdatascience.com/overview-of-ui-tools-for-monitoring-and-management-of-apache-kafka-clusters-8c383f897e80
        - https://devocean.sk.com/blog/techBoardDetail.do?ID=163980

    - nats monitoring: https://docs.nats.io/running-a-nats-service/nats_admin/monitoring
        - https://demo.nats.io:8222/

- event-driven
    - 배민.회원시스템 이벤트기반 아키텍처 구축하기: https://techblog.woowahan.com/7835/
    - zero payload
        - https://sowhat4.tistory.com/71
        - https://backtony.github.io/spring/2022-06-06-spring-event/
    - Event driven architecture with Kafka and RabbitMQ: https://www.redoxengine.com/blog/event-driven-architecture-with-kafka-and-rabbitmq/

---

https://www.codenary.co.kr/techblog/list?tag=kafka

https://www.codenary.co.kr/techblog/list?tag=rabbitmq

---

## Is there any reason to use RabbitMQ over Kafka

I wrote an answer on Stackoverflow a while ago to answer the question, “Is there any reason to use RabbitMQ over Kafka?”. The answer is just a few lines, but it has proven to be an answer many people have found helpful.

저는 얼마 전 스택 오버플로에 "래빗을 사용할 이유가 있습니까?"라는 질문에 답하기 위해 답을 썼습니다MQ over Kafka?" 답은 몇 줄에 불과하지만, 많은 사람들이 도움이 된다고 생각하는 답임이 증명되었습니다.

I will try to break down the answer into sub answers and try to explain each part.

답을 하위 답으로 나누고 각 부분을 설명하도록 하겠습니다.

First of all, I wrote - “RabbitMQ" is a solid, mature, general-purpose message broker that supports several protocols such as AMQP, MQTT, STOMP, and more.

먼저, 저는 "RabbitMQ"는 AMQP, MQTT, STOMP 등과 같은 여러 프로토콜을 지원하는 견고하고 성숙한 범용 메시지 브로커라고 썼습니다.

RabbitMQ can handle high throughput. A common use case for it is to handle background jobs or to act as a message broker between microservices.

RabbitMQ는 높은 처리량을 처리할 수 있습니다. 일반적으로 백그라운드 작업을 처리하거나 마이크로서비스 간의 메시지 브로커 역할을 수행하는 것이 이 솔루션의 일반적인 사용 사례입니다.

Kafka is a message bus optimized for high-ingress data streams and replay. Kafka can be seen as a durable message broker where applications can process and re-process streamed data on disk."

Kafka는 높은 입력 데이터 스트림 및 재생에 최적화된 메시지 버스입니다. Kafka는 애플리케이션이 디스크에서 스트리밍된 데이터를 처리하고 다시 처리할 수 있는 내구성 있는 메시지 브로커로 간주될 수 있습니다."

Regarding the term “mature”; RabbitMQ has simply been on the market for a longer time then Kafka (2007 vs 2011, respectively).

"성숙한"이라는 용어와 관련하여, 토끼MQ는 Kafka(각각 2007년과 2011년)보다 더 오랫동안 시장에 나와 있었습니다.

Both RabbitMQ and Kafka are “mature”, which means they both are considered to be reliable and scalable messaging systems.

RabbitMQ와 Kafka는 모두 "성숙한" 제품입니다. 즉, 둘 다 안정적이고 확장 가능한 메시징 시스템으로 간주됩니다.
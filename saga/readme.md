# Saga 패턴

## MSA에서 트랜잭션
- Monolithic 아키텍처에서는 모든 코드가 물리적으로 같은 공간에 있었기 때문에 Spring Framework의 도움을 받으면 `@Transactional` 등을 사용해 트랜잭션 기능을 손쉽게 활용할 수 있음
- 하지만 MSA에서는 각 마이크로 서비스가 독립적으로 존재하기 때문에 각 서비스에서 발생하는 작업들을 하나의 트랜잭션으로 관리하기 위해 추가적인 고민이 필요

> ### cf. 2PC (Two Phase Commit)
> - 분산 DB 환경에서 쓰는 방법으로 2단계에 거쳐서 데이터를 영속화 하는 작업
> - 트랜잭션을 조율하는 조정자(Coordinator)를 통해 트랜잭션 요청이 들어왔을 때 두 단계를 거쳐 트랜잭션을 진행을 담당
>
>   ![2pc1](https://img1.daumcdn.net/thumb/R1280x0/?scode=mtistory2&fname=https%3A%2F%2Fblog.kakaocdn.net%2Fdn%2FbSolTg%2FbtqBsGTE73r%2FzbQ4TkSYtegKkkhUVrUTv0%2Fimg.png)
>   ![2pc2](https://img1.daumcdn.net/thumb/R1280x0/?scode=mtistory2&fname=https%3A%2F%2Fblog.kakaocdn.net%2Fdn%2Fexd7h6%2FbtqBr748j48%2F3AQBbQfUhL7UojYbXfcpfk%2Fimg.png)
>
> #### MSA 환경에서 Two Phase Commit 문제점
> - Two Phase Commit은 DBMS 간 분산 트랜잭션을 지원해야 적용이 가능하고 또한 DBMS가 동일 제품군이어야 함
>   - NoSQL 제품군에는 이를 지원하지 않고, 함께 사용되는 DBMS가 동일 제품군(Oracle, MySQL, Postgres)이여야함
>   - 따라서 DBMS polyglot 구성은 어려움
> - 구현의 어려움
>   - Two Phase Commit은 보통 하나의 API 엔드포인트를 통해 서비스 요청이 들어오고 내부적으로 DB가 분산되어있을 때 사용돰
>   - 하지만 MSA 환경에서는 각 서비스에서 API간으로 통신을 통해 서비스 요청이 이루어지기 때문에 구현이 용이하지 않음

## Saga Pattern
- Saga 디자인 패턴은 분산 트랜잭션 시나리오에서 마이크로 서비스에서 데이터 일관성을 관리하는 방법
- Saga는 각 서비스를 업데이트하고 메시지 또는 이벤트를 게시하여 다음 트랜잭션 단계를 트리거
- 단계가 실패하면 Saga는 이전 트랜잭션을 중화시키는 보상 트랜잭션을 실행
- Choreography 방식과 Orchestration 방식이 존재

## Choreography-Based Saga
![choreography](https://velog.velcdn.com/images%2Fdvmflstm%2Fpost%2Fcb16712e-86f7-45df-a3e8-66d7ae2b0063%2Fimage.png)

### 특징
- 의사결정과 순서를 참가자에게 맡기고, 주로 이벤트 교환 방식으로 통신

### 장점
- 이해하기 쉽고 간단함
- 구축하기 쉬움

### 단점
- 어떤 서비스가 어떤 이벤트를 수신하는지 추측하기 어려움
- 트랜잭션이 많은 서비스를 거쳐야 할 때 상태를 인지하기 어려움
- 모든 서비스는 호출되는 각 서비스의 이벤트를 확인해야 함

## Orchestration-Based Saga
![orchestration](https://velog.velcdn.com/images%2Fdvmflstm%2Fpost%2Faff6ce96-fcc5-4f3e-a59c-08bfb2f51250%2Fimage.png)

### 특징
- orchestrator에 의해 중앙화해서 통제
- saga orchestrator가 참여자에게 커맨드 메시지를 보내 수행할 작업을 지시

### 장점
- 서비스 간의 종속성이 없고 orchestrator가 호출하기 때문에 분산 트랜잭션의 중앙 집중화가 이루어짐
- 서비스의 복잡성이 줄어들고, 구현 및 테스트 용이
- 롤백을 쉽게 관리할 수 있음

### 단점
- 모든 트랜잭션을 orchestrator가 관리하기 때문에 로직이 복잡해 질 수 있음
- orchestrator 서비스가 추가적으로 들어가고 이를 관리해야 하며 인프라의 복잡성이 증가

## Reference
- https://docs.microsoft.com/ko-kr/azure/architecture/reference-architectures/saga/saga
- https://velog.io/@dvmflstm/SAGA-pattern%EC%9D%84-%EC%9D%B4%EC%9A%A9%ED%95%9C-%EB%B6%84%EC%82%B0-%ED%8A%B8%EB%9E%9C%EC%9E%AD%EC%85%98-%EA%B5%AC%ED%98%84%ED%95%98%EA%B8%B0
- https://jjeongil.tistory.com/1100
- https://sarc.io/index.php/development/2128-saga-pattern
- https://cla9.tistory.com/22
- https://solace.com/blog/microservices-choreography-vs-orchestration/
- https://stackoverflow.com/questions/4127241/orchestration-vs-choreography
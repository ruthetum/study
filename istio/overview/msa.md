# 현대 서비스 아키텍처의 변화
## 모놀리틱 아키텍처 (하나의 머신, 하나의 서버)
### 장점
- 기술의 단일화
- 관리/운영에 유리

### 단점
- 여러 언어/기술 스택을 가져가기 어려움
- 빌드/배포 시간이 오래 걸림
- 모든 전체 로직을 알고 있는 사람이 필요
- 스케일링에 문제가 많음
- 의존성이 커서 잘못된 코드 배포 시 서버 전체에 장애가 발생할 확률이 높음

## 마이크로서비스 아키텍처
### 장점
- 기술스택의 다양화
- 팀별 전문성 및 개발속도 향상
- 적절한 기술과 적절한 배치
- 전체 장애가 아닌 부분 장애

## 단점
- 운영관점에서 여러 기술을 다뤄야함
- 요청 추적이 어려워짐
- 각 서비스별 방어로직이 적을 때 장애 여파
- 보안, 통제의 어려움

## 넷플릭스의 해결책
`MSA를 도입하면서 넷플릭스가 선택한 해결책`
- 서킷브레이커 패턴의 Hystrix
- 서비스 디스커버리 패턴의 Eureka
- 모니터링 서비스 Turbine
- API 게이트웨이 패턴의 Zuul -> Spring Cloud Gateway

## MSA 구조의 공통적인 문제들
<img width="419" alt="스크린샷 2022-12-05 오후 10 04 23" src="https://user-images.githubusercontent.com/59307414/205644156-02dc8c6b-a4ca-47e2-8a00-6c2feaf52ef2.png">

- 모니터링
- 중앙 로그 수집
- 분산 트랜잭션 추적
- 중복 인증 처리
- ...

> 이를 해결하기 위해 전통적인 방식에서는 **API Gateway** 도입

<img width="533" alt="스크린샷 2022-12-05 오후 10 04 36" src="https://user-images.githubusercontent.com/59307414/205644190-d25c0840-df94-4679-8cee-94ef03cd1e02.png">

- 앞단에서 공통적인 문제를 해결
- 클라이언트로부터 들어오는 요청을 GW가 받고, 중앙집중형으로 문제를 해결
    - 요청 로그 수집, 트랜잭션 관리, GW를 모니터링, 마이크로서비스별 latency 측정
- GW에서 요청을 핸들링하면서 문제 발생
    - GW에 집중 부하 발생
    - GW에 비즈니스 로직 추가 및 인프라 복잡도 증가
    - GW에 장애 발생 시 전체 서비스 장애로 이어짐

## 서비스 메쉬(Service Mesh)의 등장
- 컨테이너 오케스트레이션 툴 Kubernetes의 등장
- 여러 사이드카 컨테이너를 함께 띄울 수 있는 Pod
- API 기반 무중단 리로드가 가능한 proxy 기술의 발달
- SRE: 일관된 서비스 네트워크 정책 관리

<img width="567" alt="sm" src="https://user-images.githubusercontent.com/59307414/205605578-e24cc25e-e497-4b08-9f10-70d556e1a669.png">

- 클라우드/컨테이너 환경과 어우러지면서 API Gateway와는 또 다른 분산 책임 형태로 문제들을 해결
- 각 서비스마다 proxy 컨테이너를 소유하면서 네트워크 레벨에서 해결할 수 있는 common한 문제들을 손쉽게 해결
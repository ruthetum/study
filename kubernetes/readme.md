# Kubernetes

## Prologue
- 서버의 상태를 관리하기 위한 노력이 필요

### 도커를 사용하기까지 서버 관리 방법
1. 처음에는 문서화를 잘 해서 이를 해결하기 위한 시도 진행
    - 문서를 아무리 잘 만들어도 버전, 환경에 따른 제약이 발생
2. 문서로 관리하는 게 힘들기 때문에 chef, puppet, ansible 같은 서버 관리 도구를 사용
    - 설정 관리 도구가 좋기는 도구 자체를 배워야 하는 비용
    - 서버 구조 및 설정이 복잡해짐에 따라 도구 설정 또한 복잡해짐 (ex. 한 서버에서 여러 개의 버전, 애플리케이션을 실행)
3. 가상 머신 사용
    - 서버 하나에 가상 머신을 여러 개 사용해서 문제를 해결
    - 조금 느리고, 관리도 완전히 직관적이지는 않지만 그동안의 문제를 해결
    - 하지만 클라우드 환경, 벤더에 따라서 제약이 발생, 그리고 일단 속도가 느림
4. 도커 컨테이너 사용
    - 모든 실행 환경을 컨테이너 기반으로, 어디서든 동작하고 쉽고 효율적으로 동작

### 컨테이너의 특징
- 가상머신과 비교했을 때 컨테이너 생성이 쉽고 효율적(CPU, Memory)
- 컨테이너 이미지를 이용한 배포와 롤백이 간단
- 언어나 프레임워크에 상관없이 애플리케이션을 동일한 방식으로 관리
- 개발, 테스트, 운영 환경은 물론 로컬 피시와 클라우드까지 동일한 환경 구축
- 특정 클라우드 벤더에 종속적이지 않음

### 도커의 등장 
- 도커가 등장하면서 모든 서비스를 컨테이너로 사용(containerization)
    - mysql, redis, rabbitmq, jenkins, wordpress 등 모든 서비스르 컨테이너로 활용
    - 그만큼 편하게 사용할 수 있음
- 그래서 개발자는 아래의 흐름으로 작업
    
    `Develop(코드 작성)` - `Bulid(이미지화)` - `Ship(push)` - `Run(실행)`

- 개발자는 이미지만 잘 생성하면 실행지의 과정이 매우 편해짐, 그러나 서비스가 커지고 다양해지면서 점점 컨테이너의 수가 증가

- 도커를 이용해서 쉽게 실행을 할 수 있으나 관리 측면에서 어려움이 발생

### 도커를 이용한 배포
- ex. 3개의 서버에 띄우기 위한 상황
> 1.  1번 서버 접속(ssh) -> docker stop app && docekr run
> 2.  2번 서버 접속(ssh) -> docker stop app && docekr run
> 3.  3번 서버 접속(ssh) -> docker stop app && docekr run

#### 문제점
- 애플리케이션을 띄우기 위해 하나하나 서버에 들어가서 실행
- 어떤 서버가 비어있는지 확인 필요 -> 추가적인 모니터링 시스템 필요
- 버전을 업그레이드하는 경우 롤아웃, 롤백을 진행할 때 하나하나 모두 실행
- 서비스 검색
    - 프록시와 웹 애플리케이션 서비스가 있는 경우, 서비스가 증가한다면 그 사이에 load balancer 추가
    - 마이크로서비스가 많아지는 상황에서 추가 작업이 증가
- 서비스 노출
    - 앞단에 nginx가 있을 때 서비스별로 도메인이 다르면 nginx 설정 파일에 도메인이 추가될 때마다 설정 파일을 수정해줘야 함
- 모니터링
    - 자동화, 통합적인 모니터링에 어려움이 있음
    - 특정 서비스에 부하가 많이 발생하는 경우 자동화에 어려움이 있음

### Container Orchestration
> 컨테이너라는 기술은 좋지만, 복잡한 컨테이너 환경에서 관리의 어려움이 있음
> 
> 이를 해결하기 위해 복잡한 컨테이너 환경을 효과적으로 관리하기 위한 도구가 필요 -> Container Orchestration

- 서버 관리자가 수작업으로 진행했던 일들을 Container Orchestration가 대신 해줌

#### Container Orchestration 특징
- cluster: 서버를 하나하나 관리하는 것이 아닌 클러스터 단위로 관리
    - 마스터 서버를 하나 두고, 마스터 서버에 명령어를 실행시키면 마스터 노드에서 클러스터로 전달
    - 클러스터 내부의 네트워킹 설정 지원
    - 노드 스케일 지원

- state: 상태 관리
    - replica count를 이용해서 pod 수를 관리
    - 문제가 생기면 해당 pod을 삭제하고 다시 생성

- scheduling: 배포 관리
    - 리소스에 따라 적절한 서버에 배치, 필요하다면 추가적인 서버를 더 띄워서 실행

- rollout, rollback: 버전 관리
    - rollout, rollback을 하나하나 직접해주는 것이 아니라 전체적으로 관리

- service discovery: 서비스 검색
    - 새로운 서비스가 떴을 때 서비스를 등록, 설정을 자동 변경

- volume: 스토리지 연결
    - 서버에 따라서 스토리지를 연결 및 마운트해야할 수 있는데 수작업이 아닌 설정으로 관리

## Kubernetes
- Container Orchestration은 기술이기 때문에 해당 기능을 지원하는 다양한 도구 존재
- 쿠버네티스는 컨테이너를 쉽고 빠르게 배포/확장하고 관리를 자동화해주는 오픈소스 플랫폼

### kubectl 명령어
|명령어|설명|
|---|---|
|apply|원하는 상태를 적용합니다. 보통 -f 옵션으로 파일과 함께 사용합니다.|
|get|리소스 목록을 보여줍니다.|
|describe|리소스의 상태를 자세하게 보여줍니다.|
|delete|리소스를 제거합니다.|
|logs|컨테이너의 로그를 봅니다.|
|exec|컨테이너에 명령어를 전달합니다. 컨테이너에 접근할 때 주로 사용합니다.|
|config|kubectl 설정을 관리합니다.|


### HPA

> https://kubernetes.io/ko/docs/tasks/run-application/horizontal-pod-autoscale/#%EC%82%AC%EC%9A%A9%EC%9E%90-%EC%A0%95%EC%9D%98-%EB%A9%94%ED%8A%B8%EB%A6%AD%EC%9D%84-%EC%9D%B4%EC%9A%A9%ED%95%98%EB%8A%94-%EC%8A%A4%EC%BC%80%EC%9D%BC%EB%A7%81

**사용자 정의 메트릭을 이용하는 스케일링**

기능 상태: Kubernetes v1.23 [stable]

(이전에는 autoscaling/v2beta2 API 버전이 이 기능을 베타 기능으로 제공했었다.)

**autoscaling/v2beta2** API 버전을 사용하는 경우, (쿠버네티스 또는 어느 쿠버네티스 구성 요소에도 포함되어 있지 않은) 커스텀 메트릭을 기반으로 스케일링을 수행하도록 HorizontalPodAutoscaler를 구성할 수 있다. 이 경우 HorizontalPodAutoscaler 컨트롤러가 이러한 커스텀 메트릭을 쿠버네티스 API로부터 조회한다.

요구 사항에 대한 정보는 메트릭 API를 위한 지원을 참조한다.

**복수의 메트릭을 이용하는 스케일링**

기능 상태: Kubernetes v1.23 [stable]

(이전에는 autoscaling/v2beta2 API 버전이 이 기능을 베타 기능으로 제공했었다.)

**autoscaling/v2** API 버전을 사용하는 경우, HorizontalPodAutoscaler는 스케일링에 사용할 복수의 메트릭을 설정할 수 있다. 이 경우 HorizontalPodAutoscaler 컨트롤러가 각 메트릭을 확인하고 해당 단일 메트릭에 대한 새로운 스케일링 크기를 제안한다. HorizontalPodAutoscaler는 새롭게 제안된 스케일링 크기 중 가장 큰 값을 선택하여 워크로드 사이즈를 조정한다(이 값이 이전에 설정한 '총 최대값(overall maximum)'보다는 크지 않을 때에만).

ref. 
- https://kubernetes.io/ko/docs/tasks/run-application/horizontal-pod-autoscale/
- https://saramin.github.io/2022-05-17-kubernetes-autoscaling/
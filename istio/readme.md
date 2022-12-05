# Istio
## 서비스 메쉬(Service Mesh)의 등장
- 컨테이너 오케스트레이션 툴 Kubernetes의 등장
- 여러 사이드카 컨테이너를 함께 띄울 수 있는 Pod
- API 기반 무중단 리로드가 가능한 proxy 기술의 발달
- SRE: 일관된 서비스 네트워크 정책 관리

<img width="567" alt="sm" src="https://user-images.githubusercontent.com/59307414/205605578-e24cc25e-e497-4b08-9f10-70d556e1a669.png">

- 클라우드/컨테이너 환경과 어우러지면서 API Gateway와는 또 다른 분산 책임 형태로 문제들을 해결
- 각 서비스마다 proxy 컨테이너를 소유하면서 네트워크 레벨에서 해결할 수 있는 common한 문제들을 손쉽게 해결
- 그 툴 중 하나가 **istio**

## 서비스 메쉬 아키텍처의 기본 전략
### 프록시 사이드카 패턴(Proxy sidecar pattern)
<img width="382" alt="스크린샷 2022-12-05 오후 10 11 08" src="https://user-images.githubusercontent.com/59307414/205645390-1e7c9934-c266-4f52-8420-936c7bdddf37.png">

- 애플리케이션 레이어와 인프라 레이어 간의 간극을 줄임
- 애플리케이션 전용 프록시를 각 서비스별로 줌
  - 모든 inbound, outbound 트래픽이 proxy를 통해 처리
  - proxy만 잘 관리하면 일관되게 처리 가능
- 일관된 메트릭 처리와 일관된 네트워크 정책관리

### 컨트롤 플레인(Control plane)과 데이터 플레인(Data plane)
<img width="567" alt="sm" src="https://user-images.githubusercontent.com/59307414/205605578-e24cc25e-e497-4b08-9f10-70d556e1a669.png">

- 컨트롤 플레인을 이용해 마이크로서비스들의 프록시를 정책 관리
- 각 마이크로서비스들은 소유한 프록시를 활용
  - 세밀한 네트워크 정책 수정

### Istio 구조
<img width="507" alt="스크린샷 2022-12-05 오후 10 19 01" src="https://user-images.githubusercontent.com/59307414/205646968-4273aebf-b6df-4de5-bbda-d214613391ba.png">

- 컨트롤 플레인: **istiod** (istio daemon)
- 데이터 플레인: istio-proxy **(envoy)**

## Istio & Envoy
### Proxy
- 중계
- 보안 강화
- 확장성, 유연성
- 로드밸런싱
- Ex. Nginx, Envoy, HAProxy

### Envoy
- C++로 구현된 고성능 프록시
- **네트워크의 투명성**을 목표
- 다양한 필터체인 지원
- L3/L4 필터, HTTP L7 필터
- 동적 configuration API 제공
- CNCF OSS로 벤더사가 없음

### Istio-proxy
- envoy를 래핑한 프록시
  - 래핑 아이템으로 envoy가 선택된 이유
    - API 기반 Hot reload 기능
    - 다양한 기능(L3, L4, L7)
    - 친절한 네트워크 메트릭 제공
- envoy를 활용해 서비스앱의 사이드카 컨테이너로 활용
- 중앙 컨트롤 플레인인 istiod와 통신하고 서비스 앱의 트래픽을 정해진 정책/필터에 따라 처리
- 관찰가시성을 위한 메트릭 제공

**Pod / sidecar container**

<img width="392" alt="스크린샷 2022-12-05 오후 10 45 51" src="https://user-images.githubusercontent.com/59307414/205652241-7fe1ba98-5101-4a1c-9dfa-48bb9f8e8700.png">

- 파드 내 네트워크 인터페이스를 컨테이너끼리 공유 (=IP가 동일)
- `iptables`를 조작하여 inbound traffic과 outbound traffice 모두를 istio-proxy가 선처리 및 후처리
  ```
  - 예를 들어 80으로 요청이 들어왔다.
  - iptables을 조작해서 80으로 들어오는 패킷을 istio-proxy로 향하게 매핑
  - istio-proxy가 먼저 처리
  - istio-proxy 처리 후 서비스 앱에 전달
  - 서비스 앱에서 요청 처리
  - 서비스 앱에서 반환한 응답을 다시 istio-proxy에게 전달
  - istio-proxy 후처리 후 클라이언트에 반환
  ```
- 서비스 앱의 세밀한 **네트워킹 설정에 대한 책임**을 가져옴

## Related
- [모놀리틱, 마이크로서비스의 장단점과 서비스 메쉬의 등장](./overview/msa.md)
- [DevOps와 SRE 개념](./overview/devops.md)
- [쿠버네티스](./overview/kubernetes.md)

## Reference
- [istio를 써야할까?](https://maily.so/anarcher/posts/de7cd0)
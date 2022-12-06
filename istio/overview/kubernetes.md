# 컨테이너 오케스트레이션

<img width="578" alt="스크린샷 2022-12-05 오후 9 43 12" src="https://user-images.githubusercontent.com/59307414/205640029-3a2a91aa-ddd8-41b9-9eba-d9d4554b4345.png">

### Traditional Deployment
- 전통적인 방식은 하나의 OS 위에 app 프로세스들을 여러 개 띄우는 방식
- 하나의 물리서버 자원을 공유하다보니 app간 영향도가 큼

### Virtualized Deployment
- 가상화를 도입하면 VM을 이용해서 격리 방식
- 하이퍼 바이저 위에 OS 하나하나를 다 만듦
- 그만큼 비용이 발생

### Container Deployment
- 컨테이너는 VM과 유사하지만 실제 머신의 **OS 격리 속성**을 활용
- VM 이미지를 생성하는 것보다 비용이 저렴
- 운영자는 각 VM들을 관리하기보다 **애플리케이션** 중심으로 운영/관리 가능

# Kubernetes
- 컨테이너 오케스트레이션 툴
- **선언적 구성**과 자동화에 대해 용이하도록 설계

## 구성 요소
- 클러스터를 관리해주는 **마스터 노드**와 실제 서비스 컨테이너들이 뜨는 **워커 노드**들로 분리
- 내결함성과 고가용성
- API 서버를 지원

## 핵심 패러다임
- 선언적 배포방식
- 레이블을 이용한 간편한 워크로드 구성
- API 서버를 지원

### 선언적 배포방식
- **Desired state**를 정의하고 운영자가 원하는 방식을 정의
- 쿠버네티스는 해당 스테이트를 만들기 위해 지속적 노력(Sync loop)
- 과거 전통적인 명령형 배포방식(server start, down)과 다른 패러다임
  - 운영 코스트 감소

### 레이블을 이용한 간편한 워크로드 구성
- pod 집합에서 실행 중인 애플리케이션을 네트워크 서비스로 노출시키는 방법 (pod -> service)
- 새 배포에도 라우팅을 건드릴 필요가 없음 (레이블을 활용한 느슨한 커플링)

### API 서버를 지원
- 쿠버네티스의 상태를 핸들링할 수 있는 Control plane의 API server 제공
- kubectl 명령어의 도구
- 쿠버네티스의 오브젝트들의 질의하고 조작할 수 있음 (ex. `kubectl get pods -A`)

# 쿠버네티스 리소스 관리
## kubernetes 리소스를 편리하게 관리
- kube-ps1: https://github.com/jonmosco/kube-ps1
- kubectx: https://github.com/ahmetb/kubectx
  - 대화식: `brew install fzf`



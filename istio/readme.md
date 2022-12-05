# istio
## 서비스 메쉬(Service Mesh)의 등장
- 컨테이너 오케스트레이션 툴 Kubernetes의 등장
- 여러 사이드카 컨테이너를 함께 띄울 수 있는 Pod
- API 기반 무중단 리로드가 가능한 proxy 기술의 발달
- SRE: 일관된 서비스 네트워크 정책 관리

<img width="567" alt="sm" src="https://user-images.githubusercontent.com/59307414/205605578-e24cc25e-e497-4b08-9f10-70d556e1a669.png">

- 클라우드/컨테이너 환경과 어우러지면서 API Gateway와는 또 다른 분산 책임 형태로 문제들을 해결
- 각 서비스마다 proxy 컨테이너를 소유하면서 네트워크 레벨에서 해결할 수 있는 common한 문제들을 손쉽게 해결
- 그 툴 중 하나가 **istio**


## Reference
- [모놀리틱, 마이크로서비스의 장단점과 서비스 메쉬의 등장](./overview/msa.md)
- [DevOps와 SRE 개념](./overview/devops.md)
- [쿠버네티스](./overview/kubernetes.md)
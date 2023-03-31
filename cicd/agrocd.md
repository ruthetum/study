# ArgoCD
- Gitops CD tools for k8s
- automate the deployment of the desired state in target env
- k8s menifests can be specified in serveral ways
  - kubstomize
  - helm
  - ksonnet
  - jsonnet
  - yaml, manifest

## ArgoCD components
![image](https://user-images.githubusercontent.com/59307414/227162496-c1dafd4f-8d6b-4a64-92f7-f5694611795c.png)

### API server
- 사용자(or 관리자) UI 또는 CLI를 통해 접근
- 다른 서비스에서 gRPC 또는 REST API를 통해 접근
- Git webhook event를 통해 접근
- 철학적으로 ArgoCD는 외부의 identity provider에게 권한을 위임해서 인증/인가 구현
- ArgoCD는 RBAC을 통해 권한을 관리

### Repository Server
- git repository에 대한 캐싱
  - Git에 있는 일련의 현상을 k8s 워크로드에 싱크해주는 게 ArgoCD의 목적
- helm, yaml 파일들을 k8s manifest로 변경

### Application Controller
- k8s controller
- reconciliation(지속적으로 비교해서 current state을 desired state로)을 담당

## Project & Application

![image](https://user-images.githubusercontent.com/59307414/227167075-2e8002cd-8037-4fd4-b2fb-34706a0673db.png)

- Project: 여러 application의 묶음 (k8s의 namespace)
- Application: k8s의 workload에 맵핑

## How to request
![image](https://user-images.githubusercontent.com/59307414/227168408-f84b8330-4297-4447-a71f-08eab541c6b7.png)

- ArgoCD는 Project 내부에 Role을 생성하고 권한을 부여
- Account level의 token을 발급해서 사용할 수도 있음

## Example
- https://github.com/wilump-labs/argo-in-actions
# Istio

> Quick Start: https://github.com/wilump-labs/istio-in-action

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

## istio 환경 구성
### istioctl을 이용한 환경 구성
- download: https://istio.io/latest/docs/setup/getting-started/#download
  ```shell
  curl -L https://istio.io/downloadIstio | sh -
  cd istio-1.16.1
  export PATH=$PWD/bin:$PATH
  ```

- install: https://istio.io/latest/docs/setup/install/istioctl/
  ```shell
  istioctl install
  // istioctl install --set meshConfig.accessLogFile=/dev/stdout
  ```
  - install configuration: https://istio.io/latest/docs/setup/additional-setup/config-profiles/
  - 보통 default로 설치하고 커스텀하는 방식

- profile 확인: `istioctl profile dump [profile name]`
  - profile list 확인: `istioctl profile list`

- profile diff: `istioctl profile diff [profile1] [profile2]`
  - 프로필간 차이 확인

- manifest 파일 생성: `istioctl manifest generate > $HOME/generated-manifest.yaml`

## istioOperator를 이용한 환경 구성
- install: https://istio.io/latest/docs/setup/install/operator/
- IstioOperator options: https://istio.io/latest/docs/reference/config/istio.operator.v1alpha1/

### istio operator 적용
```shell
istioctl operator init
```

### 확인
`istio-operator`, `istio-system` ns 생성
```shell
$ kubectl get ns
NAME              STATUS   AGE
default           Active   259d
istio-operator    Active   42s 
istio-system      Active   19m
kube-node-lease   Active   259d
kube-public       Active   259d
kube-system       Active   259d
```

`istio-operator`, `istio-system` pod 생성 생성
```shell
$ kubectl get pods -A
istio-operator   istio-operator-5964f5c8d8-sl6mp         1/1     Running   0             2m10s
istio-system     istio-ingressgateway-55df5d8468-xzhwm   1/1     Running   0             21m
istio-system     istiod-7fb4bc46ff-cfmfs                 1/1     Running   0             21m
kube-system      coredns-787d4945fb-vnjtq                1/1     Running   0             33m
kube-system      etcd-minikube                           1/1     Running   0             33m
kube-system      kube-apiserver-minikube                 1/1     Running   0             33m
kube-system      kube-controller-manager-minikube        1/1     Running   0             33m
kube-system      kube-proxy-z7vnh                        1/1     Running   0             33m
kube-system      kube-scheduler-minikube                 1/1     Running   0             33m
kube-system      storage-provisioner                     1/1     Running   3 (32m ago)   259d
```

`istio-operator`, `istio-system` pod 생성 생성
```shell
$ kubectl get pods -A
istio-operator   istio-operator-5964f5c8d8-sl6mp         1/1     Running   0             2m10s
istio-system     istio-ingressgateway-55df5d8468-xzhwm   1/1     Running   0             21m
istio-system     istiod-7fb4bc46ff-cfmfs                 1/1     Running   0             21m
kube-system      coredns-787d4945fb-vnjtq                1/1     Running   0             33m
kube-system      etcd-minikube                           1/1     Running   0             33m
kube-system      kube-apiserver-minikube                 1/1     Running   0             33m
kube-system      kube-controller-manager-minikube        1/1     Running   0             33m
kube-system      kube-proxy-z7vnh                        1/1     Running   0             33m
kube-system      kube-scheduler-minikube                 1/1     Running   0             33m
kube-system      storage-provisioner                     1/1     Running   3 (32m ago)   259d
```

### istio operator를 설정 파일로 적용
```yaml
# https://istio.io/latest/docs/setup/install/operator/#install-istio-with-the-operator
apiVersion: install.istio.io/v1alpha1
kind: IstioOperator
metadata:
  namespace: istio-system
  name: example-istiocontrolplane
spec:
  profile: default
```

### 커스텀하게 리소스 수정하기
예를 들어 ingressgateway의 설정값을 수정하고 싶은 경우

```shell
# pod 검색
$ k get po -n istio-system
NAME                                    READY   STATUS    RESTARTS   AGE
istio-ingressgateway-55df5d8468-xzhwm   1/1     Running   0          33m
istiod-7fb4bc46ff-cfmfs                 1/1     Running   0          34m

# pod 설정을 yaml로 조회
$ k get po istio-ingressgateway-55df5d8468-xzhwm -n istio-system -o yaml

# 파일로 떨구기
$ k get po istio-ingressgateway-55df5d8468-xzhwm -n istio-system -o yaml > default-istio-ingressgateway.yaml
```

기본적인 설정에 대한 프로필은 istio-x.x.x/manifests/profiles 디렉토리에 있음. 해당 설정 파일을 복붙해서 수정해도 됨

istio operator manifest 파일에 원하는 부분 추가/수정 후 적용

```yaml
# sample-istio-operator.yaml
apiVersion: install.istio.io/v1alpha1
kind: IstioOperator
metadata:
  namespace: istio-system
  name: example-istiocontrolplane
spec:
  profile: demo
  components:
    ingressGateways:
      - name: istio-ingressgateway
        enabled: true
        k8s:
          resources:
            requests:
              cpu: 100m
              memory: 400Mi
```

## 사이드카 인젝션 적용하기
### Manual하게 deployment에 사이드카 배포(injection)
`istioctl kube-inject` 활용

```shell
$ istioctl kube-inject -f deployment-nginx.yaml | kubectl apply -f -
deployment.apps/nginx-deployment configured
```

#### 확인
```shell
# pod 조회: pod 내 container 1개(nginx)
$ k get po -n default
NAME                              READY   STATUS    RESTARTS   AGE
nginx-deployment-54bbcfcd-9lqn5   1/1     Running   0          53s

# 사이드카 배포
$ istioctl kube-inject -f deployment-nginx.yaml | kubectl apply -f -
deployment.apps/nginx-deployment configured

# pod 조회: pod 내 container 2개(nginx)
$ k get po -n default
NAME                              READY   STATUS    RESTARTS   AGE
nginx-deployment-54bbcfcd-9lqn5   2/2     Running   0          103s

# pod 상세 조회
$ k describe po nginx-deployment-54bbcfcd-9lqn5 -n default
...
Containers:
  nginx:
    Container ID:   docker://5dc7d98593715e770475eda1610283cd5f06207bc94f3ca442734e3a1fa7d095
    Image:          nginx:1.19.6
    Image ID:       docker-pullable://nginx@sha256:8e10956422503824ebb599f37c26a90fe70541942687f70bbdb744530fc9eba4
    Port:           80/TCP
    Host Port:      0/TCP
    State:          Running
      Started:      Sun, 07 May 2023 17:49:31 +0900
    Ready:          True
    Restart Count:  0
    Environment:    <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from kube-api-access-97m2q (ro)
  istio-proxy:
    Container ID:  docker://b64c16436135577d7438cf1b0955a51036cc406f57753ad8b6da1260f4c77a88
    Image:         docker.io/istio/proxyv2:1.16.1
    Image ID:      docker-pullable://istio/proxyv2@sha256:a861ee2ce3693ef85bbf0f96e715dde6f3fbd1546333d348993cc123a00a0290
    Port:          15090/TCP
    Host Port:     0/TCP
    Args:
      proxy
      sidecar
      --domain
      $(POD_NAMESPACE).svc.cluster.local
      --proxyLogLevel=warning
      --proxyComponentLogLevel=misc:error
      --log_output_level=default:info
      --concurrency
      2
    State:          Running
      Started:      Sun, 07 May 2023 17:49:32 +0900
    Ready:          True
    Restart Count:  0
    Limits:
      cpu:     2
      memory:  1Gi
    Requests:
      cpu:      10m
      memory:   40Mi
    Readiness:  http-get http://:15021/healthz/ready delay=1s timeout=3s period=2s #success=1 #failure=30
    Environment:
      JWT_POLICY:                    third-party-jwt
      PILOT_CERT_PROVIDER:           istiod
      CA_ADDR:                       istiod.istio-system.svc:15012
      POD_NAME:                      nginx-deployment-54bbcfcd-9lqn5 (v1:metadata.name)
      POD_NAMESPACE:                 default (v1:metadata.namespace)
      INSTANCE_IP:                    (v1:status.podIP)
      SERVICE_ACCOUNT:                (v1:spec.serviceAccountName)
      HOST_IP:                        (v1:status.hostIP)
      PROXY_CONFIG:                  {}
                                     
      ISTIO_META_POD_PORTS:          [
                                         {"containerPort":80}
                                     ]
      ISTIO_META_APP_CONTAINERS:     nginx
      ISTIO_META_CLUSTER_ID:         Kubernetes
      ISTIO_META_INTERCEPTION_MODE:  REDIRECT
      ISTIO_META_MESH_ID:            cluster.local
      TRUST_DOMAIN:                  cluster.local
    Mounts:
      /etc/istio/pod from istio-podinfo (rw)
      /etc/istio/proxy from istio-envoy (rw)
      /var/lib/istio/data from istio-data (rw)
      /var/run/secrets/credential-uds from credential-socket (rw)
      /var/run/secrets/istio from istiod-ca-cert (rw)
      /var/run/secrets/kubernetes.io/serviceaccount from kube-api-access-97m2q (ro)
      /var/run/secrets/tokens from istio-token (rw)
      /var/run/secrets/workload-spiffe-credentials from workload-certs (rw)
      /var/run/secrets/workload-spiffe-uds from workload-socket (rw)
...
```

### namespace를 통해 사이드카 배포(injection)
해당 ns에 뜨는 모든 pod들은 istio 사이드카가 자동으로 injection됨
```shell
# 적용
$ kubectl label namespace default istio-injection=enabled

# 제거
$ kubectl label namespace default istio-injection-
```

### label을 통해 사이드카 배포(injection)
ns 내의 모든 pod에 injection하는 것은 비효율적

label을 통해 injection 대상을 설정할 수 있음

```shell
# Enabled
sidecar.istio.io/inject: "true"

# Disabled
sidecar.istio.io/inject: "false"
```

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 1
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
        sidecar.istio.io/inject: "true" // pod의 label에 설정
    spec:
      containers:
        - name: nginx
          image: nginx:1.19.6
          ports:
            - containerPort: 80
```

## Related
- [모놀리틱, 마이크로서비스의 장단점과 서비스 메쉬의 등장](./overview/msa.md)
- [DevOps와 SRE 개념](./overview/devops.md)
- [쿠버네티스](./overview/kubernetes.md)

## Reference
- [istio를 써야할까?](https://maily.so/anarcher/posts/de7cd0)
- [토스ㅣSLASH 22 - 은행 앱에도 Service Mesh 도입이 가능한가요?](https://www.youtube.com/watch?v=ftFHZwyUN38)
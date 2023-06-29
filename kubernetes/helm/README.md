# Helm
## Concept
> Helm helps you manage Kubernetes applications — Helm Charts help you define, install, and upgrade even the most complex Kubernetes application.
- The package manager for kubernetes
- The tool for managing Charts
- Charts are packages of pre-configured k8s resources

## Install
```shell
$ curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3
$ chmod 700 get_helm.sh
$ ./get_helm.sh

$ helm version
version.BuildInfo{Version:"v3.4.0", GitCommit:"7090a89efc8a18f3d8178bf47d2462450349a004", GitTreeState:"clean", GoVersion:"go1.14.10"}
```

## Usage
### 생성
```shell
$ helm create sample-helm
```

### 구조
```shell
.
├── Chart.yaml
├── Makefile
├── templates
│   ├── _helpers.tpl
│   └── resources.yaml
└── values.yaml
```

ref. [sample-1](./sample-1)

### 파일 출력
```shell
$ helm template . > app.yaml
```

### 명령어
![command](https://img1.daumcdn.net/thumb/R1280x0/?scode=mtistory2&fname=https%3A%2F%2Fblog.kakaocdn.net%2Fdn%2Fdv59BR%2Fbtq1ov6BZP1%2F00zUE828PtxOm0tiBlefO1%2Fimg.png)

## Reference
- https://helm.sh/
- https://helm.sh/docs/topics/charts/
- https://helm.sh/ko/docs/intro/using_helm/
- https://helm.sh/ko/docs/howto/charts_tips_and_tricks/
---
- https://devocean.sk.com/blog/techBoardDetail.do?ID=163262
- https://tech.osci.kr/2019/11/13/helm-chart%EB%A5%BC-%EC%9D%B4%EC%9A%A9%ED%95%9C-kubernetes%EB%B0%B0%ED%8F%AC-%EA%B4%80%EB%A6%AC/
- https://kycfeel.github.io/2019/04/15/helm%EC%9C%BC%EB%A1%9C-%EC%86%90%EC%89%BD%EA%B2%8C-Kubernetes%EC%97%90-%EC%95%A0%ED%94%8C%EB%A6%AC%EC%BC%80%EC%9D%B4%EC%85%98-%EB%B0%B0%ED%8F%AC%ED%95%98%EA%B8%B0/
- https://daeunnniii.tistory.com/170
- https://gruuuuu.github.io/cloud/l-helm-basic/
- https://velog.io/@lee20h/Helm-Chart%EC%99%80-Template%EC%97%90-%EB%8C%80%ED%95%B4-%EC%95%8C%EC%95%84%EB%B3%B4%EC%9E%90
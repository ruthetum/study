# Kustomize
## Kustomize
> kustomize lets you customize raw, template-free YAML files for multiple purposes, leaving the original YAML untouched and usable as is.
>
> kustomize targets kubernetes; it understands and can patch kubernetes style API objects. It's like make, in that what it does is declared in a file, and it's like sed, in that it emits edited text.

kubernetes manifest 파일을 관리하기 위해 만들어진 오픈소스 도구

원본 YAML 파일을 보존한 채로 목적에 따라 변경본을 만들어 사용할 수 있도록 하는 것을 목표로 함

#### docs
- https://kustomize.io/
- https://github.com/kubernetes-sigs/kustomize
- https://kubectl.docs.kubernetes.io/references/kustomize/

### Helm과의 관계성
- Helm은 kustomize와 대조돼서 언급되는 패키지 매니저로 일반적으로 Chart와 함께 활용
  - Chart 패키지 설치할 수 있도록 도와주는 애드온(add-on)이 Helm
- Helm은 template 방식으로 template 위에 변수를 대입해서 완성본 manifest를 생성
- kustomize는 변경본(patch, overlay)을 만들어서 사용
  - kustomize는 원본 YAML 파일을 유지하기 때문에 바로 apply가 가능

## kustomization.yaml
Kustomize가 사용하는 YAML manifest 파일

#### Base Manifests
Kustomization에 의해 참조되는 Kustomization

보통 기본 설정으로 구성된 kubernetes manifest 파일의 묶음

#### Overlay Manifests
Base Manifests에 변형을 가하기 위해 사용되는 Kustomization

Overlay도 다른 Overlay의 Base가 될 수 있음

## 주요 명령어
### 생성
```shell
kustomize create
```

### manifest 출력
```shell
kustomize build .
```

#### 적용 및 삭제
```shell
kusomize build . | kubectl apply -f -
kustomize build . | kubectl delete -f -
```

### kubectl 통합
```shell
# kustomize build .
kubectl kustomize .

# kustomize build . | kubectl apply -f -
kubectl apply -k .

# kustomize build . | kubectl delete -f -
kubectl delete -k .
```

## Reference
- https://kustomize.io/
- https://github.com/kubernetes-sigs/kustomize
- https://kubectl.docs.kubernetes.io/references/kustomize/
- https://kubernetes.io/docs/tasks/manage-kubernetes-objects/kustomization/

- https://chhanz.github.io/kubernetes/2021/12/01/kustomize/
- https://blog.container-solutions.com/deep-dive-deployment-automation-for-applications-on-kubernetes-part-2
- https://bcho.tistory.com/1392
- https://malwareanalysis.tistory.com/402
- https://devocean.sk.com/blog/techBoardDetail.do?ID=164526&boardType=techBlog
- https://devocean.sk.com/blog/techBoardDetail.do?ID=164522
- https://velog.io/@pullee/Kustomize%EB%A1%9C-K8S-%EB%A6%AC%EC%86%8C%EC%8A%A4-%EA%B4%80%EB%A6%AC%ED%95%98%EA%B8%B0
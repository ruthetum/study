# Volume

## PV / PVC

- PV(PersistentVolume) : 데이터를 저장할 볼륨. 볼륨을 생성하고 이를 클러스터에 등록한 것
  - 볼륨 그 자체
- PVC(PersistentVolumeClaim) : 필요한 저장 공간·RW모드 등 요청사항을 기술한 명세로서 PV에 전달하는 요청. PV와 바인딩을 하는 목적으로 사용
  - 사용자가 PV에 하는 요청 (e.g. 사용하고 싶은 용량은 얼마인지, 읽기/쓰기는 어떤 모드로 설정하고 싶은지 )

### 개념
영속성을 보장할 수 없는 파드에 데이터를 저장할 경우 언제든 데이터가 사라질 가능성 존재

따라서 파드의 생명주기와 무관하게 저장이 유지되는 데이터 저장소가 필요한데 이런 요구사항을 만족하기 위한 개념이 PV(PersistentVolume)와 PVC(PersistentVolumeClaim)

![image](https://velog.velcdn.com/images/_zero_/post/45f7201a-f832-45f0-8a35-aeab908ecf4e/image.png)


![image](https://miro.medium.com/v2/resize:fit:1400/format:webp/1*eUpYUJz3bTBAcyMCI6ThOw.png)


쿠버네티스 볼륨을 파드에 직접 할당하는 방식이 아니라 중간에 PVC를 두어 파드와 파드가 사용할 스토리지를 분리
- 이런 구조는 파드 각각의 상황에 맞게 다양한 스토리지를 사용할 수 있게 함

클라우드 서비스를 사용할 때는 본인이 사용하는 클라우드 서비스에서 제공해주는 볼륨 서비스를 사용할 수도 있고, 직접 구축한 스토리지를 사용할 수도 있음
- 이렇게 다양한 스토리지를 PV로 사용할 수 있지만 파드에 직접 연결하는 것이 아니라 PVC를 거쳐서 사용하므로 파드는 어떤 스토리지를 사용하는지 신경 쓰지 않아도 됨

### 라이프사이클

프로비저닝 -> 바인딩 -> 사용 -> 반환

![image](https://img1.daumcdn.net/thumb/R1280x0/?scode=mtistory2&fname=https%3A%2F%2Fblog.kakaocdn.net%2Fdn%2FcGG4N9%2FbtrrKIfEPPV%2FchFy6chDn2PC23MbvC13t1%2Fimg.png)

#### 프로비저닝 (Provisioning)
PV를 만드는 단계

프로비저닝 방법에는 두 가지 존재
- 정적(static): PV를 미리 만들어 두고 사용하는 방법
- 동적(dynamic): 요청이 있을 때 마다 PV를 만드는 방법

정적(static) 프로비저닝
- 미리 적정 용량의 PV를 만들어 두고 사용자의 요청이 있으면 미리 만들어둔 PV를 할당
- 사용할 수 있는 스토리지 용량에 제한이 있을 때 유용
- 사용하도록 미리 만들어 둔 PV의 용량이 100GB라면 150GB를 사용하려는 요청들은 실패합니다. 1TB 스토리지를 사용하더라도 미리 만들어 둔 PV 용량이 150GB 이상인 것이 없으면 요청이 실패

동적(dynamic) 프로비저닝
- 사용자가 PVC를 거쳐서 PV를 요청했을 때 생성해서 제공
- 사용자가 원하는 용량만큼을 생성해서 사용 가능

#### 바인딩 (Binding)
프로비저닝으로 만든 PV를 PVC와 연결하는 단계

PVC에서 원하는 스토리지의 용량과 접근방법을 명시해서 요청하면 거기에 맞는 PV가 할당
- PVC에서 원하는 PV가 없다면 요청은 실패
- PVC는 원하는 PV가 생길 때까지 대기하다가 바인딩

PV와 PVC의 매핑은 1대1 관계

#### 사용 (Using)
PVC는 파드에 설정되고 파드는 PVC를 볼륨으로 인식해서 사용

할당된 PVC는 파드를 유지하는 동안 계속 사용하며 시스템에서 임의로 삭제할 수 없음
- Storage Object in Use Protection (사용 중인 스토리지 오브젝트 보호)

#### 반환 (Reclaiming)
사용이 끝난 PVC는 삭제되고 PVC를 사용하던 PV를 초기화(reclaim)하는 과정

초기화 정책은 3가지 존재
- Retain: PV를 그대로 보존
- Delete: PV를 삭제하고 연결된 외부 스토리지 쪽의 볼륨도 삭제
- Recycle: PV의 데이터들을 삭제하고 다시 새로운 PVC에서 PV를 사용

### 생성
#### PV
```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: sample-pv
spec:
  capacity:
    storage: 100Mi
  accessModes:
    - ReadWriteMany
  hostPath:
    path: "/pv/log"
  persistentVolumeReclaimPolicy: Retain
```

capacity : 볼륨 크기

accessModes : 볼륨 RW 모드
- ReadWriteOnce : 하나의 노드에서만 RW 가능
- ReadOnlyMany : 여러 노드에서 R 가능
- ReadWriteMany : 여러 노드에서 RW 가능

persistentVolumeReclaimPolicy : PV 릴리즈(사용 종료) 시 볼륨에 저장된 데이터 삭제 옵션
- Retain : PVC가 삭제되어도 PV(볼륨)의 데이터를 보존. 하지만 해당 PV를 다른 PVC가 사용하지 못하고, 재사용하기 위해서는 수동으로 PV를 반환해야 함
- Delete : PVC가 삭제되면 PV(볼륨)의 데이터를 비롯해 PV(볼륨) 자체를 삭제
- Recycle : PVC가 삭제되면 PV(볼륨)의 데이터만 삭제하고 볼륨 자체는 보존하여 곧바로 다른 PVC에 사용 가능. 하지만 현재는 deprecated되어 사용을 권하지 않음

#### PVC
```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: sample-pvc
spec:
  accessModes:
    - ReadWriteMany
  resources:
    requests:
      storage: 50Mi
  storageClassName: ''
```

accessModes : 사용하고자 하는 PV의 accessModes와 동일한 옵션을 사용해야 bound 가능

requests : 사용을 원하는 볼륨의 요구조건을 명시
- storage : 사용하고자 하는 최소한의 크기로서 명시한 용량보다 큰 PV도 상관 없음

storageClassName: ' ' : 미리 생성한 PV들 (정적 프로비저닝 한) 안에서 가능한 PV를 바인딩하겠다는 뜻

#### 확인
목록 및 상태 확인:
- pv: `kubectl get persistentvolume`
- pvc: `kubectl get persistentvolumeclaims`

세부 정보:
- pv, pvc: `kubectl describe persistentvolume [PV-이름] | [PVC-이름]`

### 마운트

볼륨 마운트
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: webapp
spec:
  containers:
  - name: webapp
    image: nginx
    volumeMounts:
    - name: log-vol
      mountPath: "/log"
  volumes:
  - name: log-vol
    hostPath:
      path: "/var/log/webapp"
```

PVC 마운트
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: webapp
spec:
  containers:
  - name: webapp
    image: nginx
    volumeMounts:
    - name: log-vol
      mountPath: "/log"
  volumes:
  - name: log-vol
    persistentVolumeClaim:
      claimName: sample-pvc
```

## Reference
- pv/pvc
  - https://medium.com/devops-mojo/kubernetes-storage-options-overview-persistent-volumes-pv-claims-pvc-and-storageclass-sc-k8s-storage-df71ca0fccc3
  - https://kimjingo.tistory.com/153
  - https://velog.io/@_zero_/%EC%BF%A0%EB%B2%84%EB%84%A4%ED%8B%B0%EC%8A%A4-PVPVC-%EA%B0%9C%EB%85%90-%EB%B0%8F-%EC%84%A4%EC%A0%95
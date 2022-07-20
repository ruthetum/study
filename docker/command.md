# 도커 명령어 활용
## 도커 이미지 명령어

도커 컨테이너는 일반적으로 [도커 허브](https://hub.docker.com)에서 제공(public image)하는 이미지를 기반으로 실행

도커 이미지는 `docker search`를 통해 조히할 수 있다.

로컬 서버 및 데스크톱에 도커 이미지를 저장하기 위해서는 **Dockerfile**을 통해 새로운 이미지를 생성(`docker build`)하거나 도커 허브로부터 이미지를 내려받는(`docker pull`) 방법이 있다.

내려받은 이미지는 `docker image ls` 또는 `docker images` 명령어를 통해 조회할 수 있다.

Dockerfile로 생성된 이미지는 도커 허브에 로그인(`docker login`)을 한 후 업로드(`docker push`) 할 수 있다.

<br>

### 도커 이미지 내려받기
```console
docker [image] pull [OPTIONS] name[:TAG | @IMAGE_DIGEST]
```

- `docker pull` : 도커 허브 레지스트리에서 로컬로 도커 이미지 내려받기
- `docker push` : 로컬에 있는 도커 이미지를 도커 허브 레지스트리에 업로드하기
- `docker login` : 업로드를 하기 전에 도커 허브 계정으로 로그인하기
- `docker logout` : 도커 허브에서 로그아웃하기

#### docker pull 명령 옵션
- `--all-tags`, `-a` : 저장소에 태그로 지정된 여러 이미지를 모두 다운로드
- `--disable-content-trust` : 이미지 검증 작업 건너뛰기
- `--platform` : 플랫폼 지정(윈도우 도커에서 리눅스 이미지를 받아야 하는 경우 : `--platform=linux`)
- `--quiet`, `-q` : 이미지 다운로드 과정에서 화면에 나타나는 상세 출력 숨김

#### digest
다이제스트 값은 원격 도커 레지스트리(도커 허브)에서 관리하는 이미지의 고유 식별값을 의미한다.

이 값을 포함한 조회는 `docker images --digest`옵션을 사용한다.

<br>

### 도커 이미지 세부 정보 조회
```console
docker image inspect [OPTIONS] IMAGE [IMAGE...]
```

### 도커 이미지 태그 설정
```console
docker tag 원본 이미지[:태그] 참조 이미지[:태그]
```

```console
# 태그 지정
docker image tag b2c2ab6dcf2e debin-httpd:1.0

# 도커 허브와 같은 레지스트리에 업로드하는 경우 저장소명과 함께 태그 지정
docker image tag httpd:latest ruthetum/httpd:2.0
```

### 도커 로그인 및 로그아웃
```console
docker login
docker logout
```

```console
# 로그인
$ docker login
Username : ruthetum
Passworkd :
...
Login Succeeded

# 도커 허브에 업로드
$ docker push ruthetum/httpd:2.0
The push refers to repository [docker.io/ruthetum/httpd]
...

# 로그아웃
$ docker logout
```

### 도커 이미지를 파일로 관리
```console
# 도커 이미지를 tar 파일로 저장
docker image save [옵션] <파일명> [image명]

# docker save로 저장한 tar 파일을 이미지로 불러오기
docker image load [옵션]
```

`docker image save` 명령을 통해 도커 원본 이미지의 레이어 구조까지 포함한 복제를 수행하여 tar파일로 이미지를 저장한다.

- 도커 허브로부터 이미지를 내려받아 내부망으로 이전하는 경우
- 신규 애플리케이션 서비스를 위해 Dockerfile로 새롭게 생성한 이미지를 저장 및 배포하는 경우
- 컨테이너를 완료(commit)하여 생성한 이미지를 저장 및 배포하는 경우

네트워크 제한 등으로 도커 허브를 이용하지 못하는 경우 이미지 이전과 배포를 위해 로컬에 저장된 이미지 파일로 저장하거나 불러올 수 있다.

```console
# 이미지 다운로드
$ docker pull mysql:5.7

# tar 파일로 저장
$ docker image save mysql5.7 > my-mysql57.tar

# tar 파일 용량을 줄이고 싶은 경우 gzip 옵션 사용
$ docker image save mysql:5.7 | gzip > my-mysql57.tar.gz

# tar 파일 내용 확인
$ tar tvf my-mysql57.tar

# 이미지 삭제
$ docker image rm mysql:5.7

# 파일로 만들어진 이미지 불러오기
$ docker image load < my-mysql57.tar>
```

### 도커 이미지 삭제
```console
# 정식 명령
dokcer image rm [옵션] {이미지 이름[:태그] | 이미지 ID}

# 단축
docker rmi [옵션] {이미지 이름[:태그] | 이미지 ID}
```

`docker image prune` 명령을 통해 컨테이너가 연결되지 않은 모든 이미지를 제거할 수도 있다.

```console
# -a 옵션을 이용하면 사용 중이지 않은 모든 이미지를 제거
$ docker image prune -a
```

## 도커 컨테이너 명령어
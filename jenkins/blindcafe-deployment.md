# Blindcafe 배포 과정

![image](https://raw.githubusercontent.com/Blind-Cafe/BlindCafe/main/assets/server/blindcafe-cicd.png)

## 로직
1. Github에 코드 푸시
2. Main branch에 커밋되었을 때 Webhook으로 Jenkins에 알림
3. Github의 코드를 받아 빌드, 테스트를 진행
4. Dockerfile을 이용해서 이미지를 빌드하고 DockerHub에 푸시
5. deploy.sh 실행
    - deploy.sh에는 DockeHub에서 이미지를 받아오고 실행시키는 코드 포함
6. deploy.sh 실행을 통해서 docker로 spring project 실행

## Jenkins 인스턴스 세팅


## 운영용 인스턴스
### Docker 세팅
```
패키지 업데이트
sudo yum -y upgrade

도커 설치
sudo yum -y install docker

도커 설치 작업이 잘 되었는지 버전 확인
docker -v

도커 시작
sudo service docker start

도커 그룹에 사용자 추가 -> docker가 그룹명, ec2-user가 사용자명
sudo usermod -aG docker ec2-user
```

### deploy.sh 작성
```
# 가동중인 awsstudy 도커 중단 및 삭제
sudo docker ps -a -q --filter "name=blindcafe" | grep -q . && docker stop blindcafe && docker rm blindcafe | true

# 기존 이미지 삭제
sudo docker rmi ruthetum/blindcafe:1.1.0

# 도커허브 이미지 pull
sudo docker pull ruthetum/blindcafe:1.1.0

# 도커 run
docker run -d -p 8000:8000 /home/ec2-user:/config --name blindcafe ruthetum/blindcafe:1.1.0

# 사용하지 않는 불필요한 이미지 삭제
docker rmi -f $(docker images -f "dangling=true" -q) || true
```

## 배포 인스턴스
> **프리티어의 경우 swapfile 설정을 통해 메모리 마련할 것**

### Docker 세팅
```
패키지 업데이트
sudo yum -y upgrade

도커 설치
sudo yum -y install docker

도커 설치 작업이 잘 되었는지 버전 확인
docker -v

도커 시작
sudo service docker start

도커 그룹에 사용자 추가 -> docker가 그룹명, ec2-user가 사용자명
sudo usermod -aG docker ec2-user
```

### Dockerfile 생성
- `sudo vi Dockerfile`
```
FROM jenkins/jenkins:jdk11
# 도커를 실행하기 위한 root 계정으로 전환
USER root
#도커 설치
COPY docker_install.sh /docker_install.sh
RUN chmod +x /docker_install.sh
RUN /docker_install.sh
# 도커 그룹에 사용자 추가
RUN usermod -aG docker jenkins
USER jenkins
```

### docker_install.sh 생성
- jenkins 컨테이너 내부에서 Docker hub에 푸시를 해야되기 때문에 도커 컨테이너 내부에 도커 컨테이너가 필요함
- `sudo vi docker_install.sh`
```
#!/bin/sh
apt-get update && \
apt-get -y install apt-transport-https \
  ca-certificates \
  curl \
  gnupg2 \
  zip \
  unzip \
  software-properties-common && \
curl -fsSL https://download.docker.com/linux/$(. /etc/os-release; echo "$ID")/gpg > /tmp/dkey; apt-key add /tmp/dkey && \
add-apt-repository \
"deb [arch=amd64] https://download.docker.com/linux/$(. /etc/os-release; echo "$ID") \
$(lsb_release -cs) \
stable" && \
apt-get update && \
apt-get -y install docker-ce
```

### docker.sock 설정
- 도커 안에 있는 도커는 host의 docker.sock을 빌려서 사용하기 때문에 권한 수정
- `sudo chmod 666 /var/run/docker.sock`

### Dockerfile 빌드해서 이미지 생성
- `docker build -t jenkins .`

### 메모리 마운트
- 컨테이너가 죽으면 메모리가 날라가기 때문에 마운트해서 사용하자
```
jenkins 폴더 만들기
mkdir jenkins

해당 폴더에 대해 권한 부여하기
sudo chown -R 1000 ./jenkins
```

### 컨테이너 실행
- jenkins 컨테이너를 9090 포트로 접근
```
sudo docker run -d --name jenkins \
-v /home/ec2-user/jenkins:/var/jenkins_home \
-v /var/run/docker.sock:/var/run/docker.sock \
-p 9090:8080 \
-e TZ=Asia/Seoul \
jenkins
```

### 초기 패스워드 및 계정 설정
- 컨테이너 내부로 들어가서 `initialAdminPassword` 확인하고 계정 생성
```
패스워드를 확인하기 위해 해당 컨테이너로 진입
sudo docker exec -it jenkins bash

패스워드 출력
cat /var/jenkins_home/secrets/initialAdminPassword
```

### 플러그인 세팅
- `Jenkins 관리` - `System Configuration` - `플러그인 관리`에서 아래 플러그인 설치
```
gradle
github integration
post build task
publish over ssh
```

### Credentials 세팅
- `Jenkins 관리` - `Security` - `Manage Credentials`
- Github 계정, Docker Hub 계정 입력
- Github secret token 추가 - `secret text`
    - Gtihub Webhook

### 시스템 설정 - Github Server 설정
- `Jenkins 관리` - `System Configuration` - `시스템 설정`에서 GitHub Servers 설정 추가

![image](https://user-images.githubusercontent.com/59307414/161944496-5f4a76da-c9cf-489d-ab80-33b93acedf5e.png)


### SSH Servers 세팅
- 운영용 인스턴스로 알려주기 위함
- `Jenkins 관리` - `System Configuration` - `시스템 설정`에서 ssh servers 설정 추가

![image](https://user-images.githubusercontent.com/59307414/161944913-31ee3b5c-33cf-47f2-bcfd-7f334545c057.png)

- `고급` - pem 추가

![image](https://user-images.githubusercontent.com/59307414/161945145-c25da704-5c2c-4cd8-8b5a-6f0f0813aab7.png)

### Gradle 세팅
- `Jenkins 관리` - `Global Tool Configuration` - `Gradle`

![image](https://user-images.githubusercontent.com/59307414/161945437-e9b8544e-2cd2-4519-b6b0-b6b20cfd0eac.png)

### 프로젝트 세팅
- `New item` - `Freestyle project`로 프로젝트 생성
- `소스코드 관리` - Github Repository 설정, crendential 설정

![image](https://user-images.githubusercontent.com/59307414/161945832-90b9dad5-fc9e-4126-820b-d576f669b53a.png)

- 빌드 유발에 GitHub hook trigger for GITscm polling을 클릭

![image](https://user-images.githubusercontent.com/59307414/161946094-144970c0-beae-4946-909d-2c6728d4bdc1.png)

- 빌드 환경에서 Use secret text or file을 선택하고 Bindings의 Add 드롭박스를 클릭하고 Username and password (separated)를 선택
    - Docker Hub 계정 설정해주는 과정

![image](https://user-images.githubusercontent.com/59307414/161946293-908cec3f-165e-49cc-8bd8-a440ebbc4364.png)

- Build에서 Invoke Gradle script 클릭, gradle script 작성
    - -Pprofile=prod 설정을 통해 profile 설정할 수 있음

![image](https://user-images.githubusercontent.com/59307414/161946539-0bebb1a8-1520-4f2e-93b3-6b0671bc35cb.png)

- Add build step에서 Execute shell 선택, 빌드를 진행한 후 동작할 명령어들을 입력

![image](https://user-images.githubusercontent.com/59307414/161947106-943a489d-8d94-4771-99f3-9910897e7831.png)

```shell
# 도커파일에 -t로 태그명을 주고 이미지를 빌드
docker build -t ruthetum/blindcafe:1.1.0 .

# 도커 로그인
echo $PASSWORD | docker login -u $USERNAME --password-stdin

# 도커허브에 빌드된 이미지를 푸시
docker push ruthetum/blindcafe:1.1.0

# 푸시한 이미지를 삭제
docker rmi ruthetum/blindcafe:1.1.0
```
```
참고사항 
docker build -t <dockerUserName>/<repository>:<tag> .
docker push <dockerUserName>/<repository>
```

- 빌드 후 조치를 선택하고 send build artifacts over SSH 선택
    - 운영용 인스턴스에 deploy.sh 실행을 알리기 위해

![image](https://user-images.githubusercontent.com/59307414/161947629-a9286e4f-e901-4730-acd0-fa5de805e2d5.png)

```
echo $PASSWORD | docker login -u $USERNAME --password-stdin
sh deploy.sh
```

- 최종적으로 작업이 끝난 후 workspace를 비우도록 설정
    - `빌드 후 조치` - `delete workspace`

![image](https://user-images.githubusercontent.com/59307414/161947955-1b8a520b-bc14-4d98-afb3-15cd12138346.png)
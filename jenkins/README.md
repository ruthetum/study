# [Jenkins](https://www.jenkins.io)
- Continuous Integration Tools의 일종
- 무료 오픈소스 자동화 서버
- 빌드, 배포 자동화를 지원하는 수백개의 플러그인 제공
    - 빌드: 코드를 실행할 수 있는 상태로 만드는 일
    - 배포: 사용자 접근이 가능하도록 배치시키는 일

## 기능
- 빌드 자동화
- 테스트 자동화
- 코드 표준 준수 여부 검사 (정적 코드 분석)
- Build Pipeline 구성 :프로젝트에 2개 이상 모듈이 있을 경우 참조 관계에 따라 순차적 빌드 가능
- 다양한 인증 기반과 결합한 인증, 권한 관리 기능
- 자동화된 배포 관리

## 설치
<details>
<summary>Windows</summary>
<div markdown="1">

1. www.jenkins.io/download/
2. 원하는 버전 설치
3. Account, Password 입력 후 Test Credentials
    - Logon Type 차이: stackoverflow.com/questions/63537185/what-are-the-differences-on-the-logon-types-choices-on-jenkins
        - Jenkins를 실행하기 위해 Local에서만 쓸지, 특정 서비스 계정을 만들지
    - `this account either does not have the privilege to logon as a service` 에러
        - Windows 관리 도구 - 로컬 보안 정책 - 로컬 정책- 사용자 권한 할당 - 서비스로 로그온 - 사용자 또는 그룹 추가 에서 사용자 추가
4. port 선택
5. 자신의 화면에 나온 경로에 들어가 initialAdminPassword를 확인하고 입력
6. 계정 생성 후 사용

</div>
</details>

<details>
<summary>Linux</summary>
<div markdown="1">

```shell
$ java -version //버전 확인
$ sudo apt-get install openjdk-11-jdk //java11 설치

$ wget -q -O - https://pkg.jenkins.io/debian/jenkins-ci.org.key | sudo apt-key add -
$ echo deb https://pkg.jenkins.io/debian-stable binary/ | sudo tee /etc/apt/sources.list.d/jenkins.list

$ sudo apt-get update

$ sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 대문자+숫자조합

$ sudo apt-get install jenkins

$ sudo systemctl start jenkins //Jenkins 실행
```

</div>
</details>



## Reference
- Jenkins 설치 : https://yeonyeon.tistory.com/56
- Jenkins 사용 : https://krksap.tistory.com/1806
- EC2 프리티어에서 메모리 부족으로 Jenkins가 돌아가지 않을 때 : https://tape22.tistory.com/22
- github & Jenkins & Docker 활용 CI/CD 구축 : https://backtony.github.io/spring/aws/2021-08-08-spring-cicd-1/
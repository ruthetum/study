# nGrinder
- 애플리케이션의 부하 테스트 용도로 많이 사용되는 성능 테스트 툴

## 구성
- 컨트롤러와 에이전트로 구성

### 컨트롤러
- 관리를 위한 Admin, 부하 스크립트 작성, 부하 테스트 작성 기능을 지원
- 관리할 에이전트를 승인하고 각 에이전트들이 부하를 발생시킬 수 있도록 제어
- cf. https://github.com/naver/ngrinder/wiki/Controller-Configuration-Guide

### 에이전트
- 컨트롤러의 제어에 따라 실제 부하를 발생
- cf. https://github.com/naver/ngrinder/wiki/Agent-Configuration-Guide

## 다운로드
- https://github.com/naver/ngrinder/releases/
- 실행시키기 위해서는 java 설치 필요
- war 파일 다운로드

## 실행
- `java -jar ngrinder/ngrinder-controller-3.5.5-p1.war --port=8300`
- `localhost:8300` 접속

### 로그인
- 초기 ID : `admin` / PW : `admin`

![home](https://user-images.githubusercontent.com/59307414/147868388-98fec377-925b-4872-a4e7-28b1fa0185d0.png)

### 부하 스크립트 생성
![create-script-1](https://user-images.githubusercontent.com/59307414/147868407-5672748d-4d78-4299-8cb6-51d1f74c03e1.png)
![create-script-1](https://user-images.githubusercontent.com/59307414/147868475-c63b85cf-a201-46bb-b562-b7fa26ab7772.png)

### 에이전트 다운로드 및 실행
![download-agent](https://user-images.githubusercontent.com/59307414/147868548-ac69fe30-1135-4ccc-bfaf-5ab6c4820827.png)
- __agent.conf (설정 파일)
```
common.start_mode=agent
agent.controller_host=localhost
agent.controller_port=16001
```
- 운영체제에 따라 `run_agent.sh` or `run_agent.bat` 실행

![run-agent](https://user-images.githubusercontent.com/59307414/147868650-6ee3395d-b570-4f64-981c-562253b89276.png)

- `Agent Management` 확인

![agent-management-1](https://user-images.githubusercontent.com/59307414/147868665-da4ab0d4-3566-419f-9216-fab6a14094ef.png)
![agent-management-2](https://user-images.githubusercontent.com/59307414/147868672-2dfe5128-c531-4c43-b2db-5bf247510222.png)

### 부하 테스트 작성 & 실행 & 분석
![load-testing-1](https://user-images.githubusercontent.com/59307414/147868708-fccb5446-3b2a-4651-9737-7f2d04f7ad97.png)
- `Create Test` 클릭

![script-save](https://user-images.githubusercontent.com/59307414/147868780-cb4b5f26-38ce-488b-bd6c-308173dd5ab2.png)
```
- Agent 개수: 1
- Vuser per agent: 2
- Script 종류: svn
- Script 이름: TestScript.groovy
- Duration: 0:00:10 (시:분:초로 10초 의미)
```
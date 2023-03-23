# Github Actions 
- event를 통해 workflow를 자동화할 수 있도록 도와주는 도구
- ex)
  - push event를 받을 때마다 Docker Image 빌드
  - Pull request가 생성되었을 때 Test 구동
  - master branch에 병합이 됐을 때 CI/CD 구동
  - 주기적으로 특정 task 실행
  - ...

## Github Actions Components
![image](https://user-images.githubusercontent.com/59307414/227134102-a2b45fa8-45af-44f6-8234-21a96516db41.png)

### Workflows
- 자동화 과정(Automated procedure: build, test, package, release, deploy, ...)
- 하나 이상의 job으로 구성
- event에 의해 트리거되거나 스케줄링

### Events
- workflow를 트리거하는 행위 (e.g. push, PR, Issue created)
- https://docs.github.com/ko/actions/using-workflows/events-that-trigger-workflows

### Jobs
- 동일한 Runner 내의 step들의 묶음
- 기본적인 설정으로 job들은 parallel하게 실행

### Steps
- 개별적인 task
- action 또는 shell command를 실행할 수 있음
- 같은 job 내에 포함된 step은 동일한 runner 안에서 작동하기 때문에 step 간 데이터를 공유할 수 있음

### Actions
- 전체 workflow에서 가장 작은 단위
- standalone commands

### Runners
- 태스크들을 수행하는 서버
- Github hosted(Azure) or Self hosted
# CI / CD

## CI (Continuous Integration)
- **빌드/테스트 자동화 과정**
- CI를 성공적으로 구현할 경우 애플리케이션에 대한 새로운 코드 변경 사항이 정기적으로 빌드 및 테스트되어 공유 리포지토리에 통합되므로 여러 명의 개발자가 동시에 애플리케이션 개발과 관련된 코드 작업을 할 경우 서로 충돌할 수 있는 문제를 해결
- **커밋할 때마다 빌드와 일련의 자동 테스트가 이루어져 동작을 확인하고 변경으로 인해 문제가 생기는 부분이 없도록 보장**
- Compile, Build, Unit test, Integration Test, Static analysis

### CI common practices
- Maintain a code repository
  - Application should be buildable from a fresh checkout without requiring additional dependencies
- Automate the build & Keep the build fast
- Everyone commits to the baseline every day & Every commit (to baseline) should be built
- Every bug-fix commit should come with a test case
- Test in a clone of the production environment(Staging)

## CD (Continuous Deployment)
- **배포 자동화 과정**
- 지속적인 서비스 제공(Continuous Delivery) 또는 지속적인 배포(Continuous Deployment)를 의미
- 코드 변경이 파이프라인의 이전 단계를 모두 성공적으로 통과하면 수동 개입 없이 해당 변경 사항이 프로덕션에 자동으로 배포
- 간단한 코드 변경이 정기적으로 마스터에 커밋되고, 자동화된 빌드 및 테스트 프로세스를 거치며 다양한 사전 프로덕션 환경으로 승격되며, 문제가 발견되지 않으면 최종적으로 배포

## CI/CD 종류
- Jenkins
- CircleCI
- TravisCI
- Github Actions
- Argo

## CI/CD  적용 전과후 비교
### CI/CD를 적용하기 전
1. 개발자들이 개발하여 코드를 수정합니다.
2. 각자의 feature 브랜치에 코드를 push합니다. (but, 어느 한 부분에서 에러가 났지만 개발자들은 눈치채지 못합니다.)
3. 각자의 코드를 git에 올리고 통합(Intergration)합니다.
4. 에러가 발생했지만 어느 부분에서 에러가 났는지 모르므로 다시 어디부분에 에러가 있는지 디버깅하고 코드를 수정합니다.
5. (1) ~ (4)의 과정을 반복합니다.
6. 많은 시간을 할애하여 에러가 해결되었으면 배포를 시작합니다. 하지만 배포과정 또한, 개발자가 직접 배포과정을 거치므로 많은 시간을 소요합니다.

### CI/CD를 적용 후의 과정
1. 개발자들이 개발하여 feature브랜치에 코드를 push합니다.
2. git push를 통해 Trigger되어 CI서버에서 알아서 Build, Test, Lint를 실행하고 결과를 전송합니다.
3. 개발자들은 결과를 전송받고 에러가 난 부분이 있다면 에러부분을 수정하고 코드를 master 브랜치에 merge합니다.
4. master 브랜치에 코드를 merge하고 Build, Test가 정상적으로 수행이 되었다면 CI서버에서 알아서 Deploy 과정을 수행합니다.

## Reference
- https://www.redhat.com/ko/topics/devops/what-is-ci-cd
- https://seosh817.tistory.com/104
- https://en.wikipedia.org/wiki/Continuous_integration
- https://en.wikipedia.org/wiki/Continuous_delivery
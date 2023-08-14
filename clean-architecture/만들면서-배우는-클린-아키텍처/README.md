# 만들면서 배우는 클린 아키텍처
'만들면서 배우는 클릭 아키텍처'를 읽고 실습 및 정리한 내용입니다.

> https://www.yes24.com/Product/Goods/105138479?pid=123487&cosemkid=go16373101893711165&gclid=Cj0KCQjw2eilBhCCARIsAG0Pf8sMhcrVAgCxjs5PbGMjUaaRa2Zu8PJLABGze_tqwJf1pxOKkv7lk9kaAt9PEALw_wcB

![image](https://image.yes24.com/goods/105138479/XL)

## [Contents](./contents)

1. [계층형 아키텍처의 문제는 무엇일까?](./contents/ch01-layer-architecture.md)
2. [의존성 역전하기](./contents/ch02-ioc.md)
3. [코드 구성하기](./contents/ch03-code.md)
4. [유스케이스 구현하기](./contents/ch04-usecase.md)
5. [웹 어뎁터 구현하기](./contents/ch05-web-adapter.md)
6. [영속성 어뎁터 구현하기](./contents/ch06-persistence-adapter.md)
7. [아키텍처 요소 테스트하기](./contents/ch07-architecture-test.md)
8. [경계 간 매핑하기](./contents/ch08-mapper.md)
9. 애플리케이션 조립하기
10. 아키텍처 경계 강제하기
11. 의식적으로 지름길 사용하기
12. 아키텍처 스타일 결정하기

## Structure
```
.
├── main
│   ├── java
│   │   └── com
│   │       └── architecture
│   │           └── clean
│   │               ├── CleanApplication.java
│   │               ├── account
│   │               │   ├── adapter
│   │               │   │   ├── in
│   │               │   │   │   └── web
│   │               │   │   │       └── SendMoneyController.java
│   │               │   │   └── out
│   │               │   │       └── persistence
│   │               │   │           ├── AccountJpaEntity.java
│   │               │   │           ├── AccountMapper.java
│   │               │   │           ├── AccountPersistenceAdapter.java
│   │               │   │           ├── AccountRepository.java
│   │               │   │           ├── ActivityJpaEntity.java
│   │               │   │           └── ActivityRepository.java
│   │               │   ├── application
│   │               │   │   ├── port
│   │               │   │   │   ├── in
│   │               │   │   │   │   ├── GetAccountBalanceQuery.java
│   │               │   │   │   │   ├── SendMoneyCommand.java
│   │               │   │   │   │   └── SendMoneyUseCase.java
│   │               │   │   │   └── out
│   │               │   │   │       ├── AccountLock.java
│   │               │   │   │       ├── LoadAccountPort.java
│   │               │   │   │       └── UpdateAccountStatePort.java
│   │               │   │   └── service
│   │               │   │       ├── GetAccountBalanceService.java
│   │               │   │       ├── MoneyTransferProperties.java
│   │               │   │       ├── NoOpAccountLock.java
│   │               │   │       ├── SendMoneyService.java
│   │               │   │       └── ThresholdExceededException.java
│   │               │   └── domain
│   │               │       ├── Account.java
│   │               │       ├── Activity.java
│   │               │       ├── ActivityWindow.java
│   │               │       └── Money.java
│   │               ├── common
│   │               │   └── SelfValidating.java
│   │               └── config
│   │                   ├── CleanArchitectureConfiguration.java
│   │                   └── CleanArchitectureConfigurationProperties.java
│   └── resources
│       ├── application.yml
│       ├── static
│       └── templates
└── test
    └── java
        └── com
            └── architecture
                └── clean
                    ├── CleanApplicationTests.java
                    ├── account
                    │   ├── adapter
                    │   │   ├── in
                    │   │   │   └── web
                    │   │   │       └── SendMoneyControllerTest.java
                    │   │   └── out
                    │   │       └── persistence
                    │   │           └── AccountPersistenceAdapterTest.java
                    │   ├── application
                    │   │   └── service
                    │   │       └── SendMoneyServiceTest.java
                    │   └── domain
                    │       ├── AccountTest.java
                    │       └── ActivityWindowTest.java
                    ├── common
                    │   ├── AccountTestData.java
                    │   └── ActivityTestData.java
                    └── system
                        └── SendMoneySystemTest.java
```

## Reference
source code: https://github.com/wikibook/clean-architecture
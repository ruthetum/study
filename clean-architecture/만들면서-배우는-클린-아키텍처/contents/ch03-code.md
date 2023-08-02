# ch03. 코드 구성하기

## 아키텍처적으로 표현력 있는 패키지 구조
```
main
└── account
    ├── adapter
    │   ├── in
    │   │   └── web
    │   │       └── AccountController
    │   └── out
    │       └── persistence
    │           ├── AccountPersistenceAdapter
    │           └── SpringDataAccountRepository
    ├── domain
    │   ├── Account
    │   └── Activity
    └── application
        ├── SendMoneyService
        └── port
            ├── in
            │   └── SendMoneyUseCase
            └── out
                ├── LoadAccountPort
                └── UpdateAccountStatePort
```
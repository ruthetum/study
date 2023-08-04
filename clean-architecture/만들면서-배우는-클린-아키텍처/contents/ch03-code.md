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

<img width="591" alt="스크린샷" src="https://github.com/ruthetum/study/assets/59307414/f679e692-3e33-4165-abe0-a50caa97cb5c">

- 컨트롤러가 서비스에 의해 구현된 인커밍 포트를 호출
- 서비스는 어뎁터에 의해 구현된 아웃고잉 포트를 호출
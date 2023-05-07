# CompletableFuture 예제 코드
```shell
.
└── src
    ├── main.java.com.exmaple.demo
    │                  ├── blocking
    │                  │   ├── repository
    │                  │   │   ├── ArticleRepository.java
    │                  │   │   ├── FollowRepository.java
    │                  │   │   ├── ImageRepository.java
    │                  │   │   └── UserRepository.java
    │                  │   └── service
    │                  │       └── UserBlockingService.java
    │                  ├── common
    │                  │   ├── domain
    │                  │   │   ├── Article.java
    │                  │   │   ├── Image.java
    │                  │   │   └── User.java
    │                  │   └── repository
    │                  │       ├── ArticleEntity.java
    │                  │       ├── ImageEntity.java
    │                  │       └── UserEntity.java
    │                  └── future
    │                      ├── repository
    │                      │   ├── ArticleFutureRepository.java
    │                      │   ├── FollowFutureRepository.java
    │                      │   ├── ImageFutureRepository.java
    │                      │   └── UserFutureRepository.java
    │                      └── service
    │                          └── UserFutureService.java
    │                          
    └── test.java.com.example.demo
                        ├── blocking
                        │   └── service
                        │       └── UserBlockingServiceTest.java
                        └── future
                            └── service
                                └── UserFutureServiceTest.java
```

- `blocking` package: 기존 블로킹 코드
- `future` package: CompletableFuture를 활용 비동기 코드
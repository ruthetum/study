# Spring Cloud 맛보기

## MSA
- 기능 별로 구성된 작은 애플리케이션들의 묶음
- *MSA 개념 정리는 추후에 자세히 하자, 여기서는 Spring Cloud 기반의 MSA를 목표로*
    - *cf.*
        - cloud native architecture, application
        - 12 factors
        - monolithic vs microservice
        - SOA vs MSA

## [Spring Cloud](https://spring.io/projects/spring-cloud)
- 분산 시스템의 빠른 개발을 위해 지원
    - 환경 설정, 회복성, etc.
- `Spring Cloud Config`, `Spring Cloud Netflix`, `Spring Cloud Security`, `Spring Cloud Gateway`, etc.
- 어떻게?
    - `Spring Cloud Config Server ` : 환경 설정 관리
        - Gateway의 IP, token에 대한 정보를 한 곳을 통해 참조
        - 다른 서비스들을 다시 빌드할 필요없이 적용 가능
    - `Spring Cloud Netflix (Eureka)` : 서비스 등록 및 관리
        - Naming server
    - `Spring Cloud Gateway` : 분산 (Load balancing)

## Service Discovery
- MSA 에서는 여러 서비스 간의 호출로 구성
- 클라우드 환경에서는 IP가 동적으로 변경되는 일이 많기 때문에 정확한(유효한) 위치를 알아내는 기능이 필요 (K-V 형태)
- 이 일을 Service discovery가 담당

### Spring Cloud Netflix Eureka Server & Client
#### Server
- build.grdle
    ```gradle
    implementation 'org.springframework.cloud:spring-cloud-starter-netflix-eureka-server'
    ```
- application
    ```
    @EnableEurekaServer
    ```
#### Client
- build.grdle
    ```gradle
    implementation 'org.springframework.cloud:spring-cloud-starter-netflix-eureka-client'
    ```
- application
    ```
    @EnableDiscoveryClient
    ```

## API Gateway
- 인증 및 권한 부여
- 서비스 검색
- 응답 캐싱
- 정책 설정
- 부하 분산
- 로깅, 추적, 상관 관계
    - *로깅은 자체적으로 mongo에 남길 수도, ELK 활용하거나 외부 SW 활용할 수도 있는데 이번에는 ELK 활용하는 걸 목표로*

### Spring Cloud에서 MSA간 통신
1) RestTemplate
    ```java
    RestTemplcate restTemplate = new RestTemplate();
    restTemplcate.getForObject("http://localhost:8080/users/", User.class, 200)
    ```
2) Feign Client
    ```java
    @FeignClient("orders")
    public interface OrderClient {
        @RequestMapping(method = GET, value = "/orders)
        List<Order> getOrders();
    }
    ```

### Ribbon
- <b>client side</b> load balancer
- 클라이언트 사이드에서 관리
- 각각의 MSA 이름으로 호출
- 과거에는 사용했는데 비동기 처리가 안됨
- 사용은 가능(2.4. 밑 버전)하나 더 이상 지원하지 않음,`Spring Cloud Loadbalancer` 활용
    - [maintenance mode](https://cloud.spring.io/spring-cloud-netflix/multi/multi__modules_in_maintenance_mode.html)

### Zuul
- 사용은 가능(2.4. 밑 버전)하나 더 이상 지원하지 않음, `Spring Cloud Gateway` 활용

### Spring Cloud Gateway
#### yml 파일로 라우팅
```yml
spring:
  application:
    name: gateway-service
  cloud:
    gateway:
      routes:
        - id: first-service
          uri: http://localhost:8081/
          predicates:
            - Path=/first-service/**
        - id: second-service
          uri: http://localhost:8082/
          predicates:
            - Path=/second-service/**
```
- routes 설정을 통해 라우팅 가능
    - uri: 라우팅 대상
    - predicates: 조건
    - http://localhost:8080/first-service -> http://localhost:8081/first-service/**
- actuator 설정을 통해 endpoint mapping 확인
    ```yml
    implementation 'org.springframework.boot:spring-boot-starter-actuator'
    ```
    ```yml
    management:
        endpoints:
            web:
            exposure:
                include:
                - "gateway"
        endpoint:
            gateway:
            enabled: true
    ```
- `{hostname}:{port}/actuator/gateway/routes`로 확인 가능
    ```json
    [
        {
            "predicate": "Paths: [/first-service/**], match trailing slash: true",
            "route_id": "first-service",
            "filters": [],
            "uri": "http://localhost:8081/",
            "order": 0
        },
        {
            "predicate": "Paths: [/second-service/**], match trailing slash: true",
            "route_id": "second-service",
            "filters": [],
            "uri": "http://localhost:8082/",
            "order": 0
        }
    ]
    ```

#### Filter 동작
![image](https://user-images.githubusercontent.com/59307414/145999040-2a5fa9b9-82ce-4f06-aaa6-114b2fbf1845.png)

```java
@Configuration
public class FilterConfig {
    @Bean
    public RouteLocator gatewayRoutes(RouteLocatorBuilder routeLocatorBuilder) {
        return routeLocatorBuilder.routes()
                .route(r -> r.path("/first-service/**")
                        .filters(f -> f.addRequestHeader("first-request", "first-request-header")
                                .addResponseHeader("first-response", "first-response-header"))
                        .uri("http://localhost:8081/"))
                .route(r -> r.path("/second-service/**")
                        .filters(f -> f.addRequestHeader("second-request", "second-request-header")
                                .addResponseHeader("second-response", "second-response-header"))
                        .uri("http://localhost:8082/"))
                .build();
    }
}
```
- filter를 통해 헤더 조작 가능


#### Custom Filter
- CustomFilter.java
    ```java
    public class CustomFilter extends AbstractGatewayFilterFactory<CustomFilter.Config> {

        public CustomFilter() {
            super(Config.class);
        }

        @Override
        public GatewayFilter apply(Config config) {
            return ((exchange, chain) -> {
                ServerHttpRequest request = exchange.getRequest();
                ServerHttpResponse response = exchange.getResponse();

                log.info("Custom PRE filter: request uri -> {}", request.getId());
                return chain.filter(exchange).then(Mono.fromRunnable(() -> {
                    log.info("Custom POST filter: response code -> {}", response.getStatusCode());
                }));
            });
        }

        public static class Config { }
    }
    ```
- application.yml
    ```yml
    spring:
      cloud:
        gateway:
          routes:
            - id: first-service
              uri: http://localhost:8081/
              predicates:
                - Path=/first-service/**
              filters:
                - name: CustomFilter
    ```

#### Global Filer
- CustomFilter.java
    ```java
    public class GlobalFilter extends AbstractGatewayFilterFactory<GlobalFilter.Config> {
        public GlobalFilter() {
            super(Config.class);
        }

        @Override
        public GatewayFilter apply(Config config) {
            return ((exchange, chain) -> {
                ServerHttpRequest request = exchange.getRequest();
                ServerHttpResponse response = exchange.getResponse();

                log.info("Global Filter baseMessage: {}", config.getBaseMessage());

                if (config.isPreLogger()) {
                    log.info("Global Filter start: request id -> {}", request.getId());
                }
                return chain.filter(exchange).then(Mono.fromRunnable(() -> {
                    if (config.isPostLogger()) {
                        log.info("Global Filter end: request id -> {}", response.getStatusCode());
                    }
                }));
            });
        }

        @Data
        public static class Config {
            private String baseMessage;
            private boolean preLogger;
            private boolean postLogger;
        }
    }
    ```

- application.yml
    ```yml
    spring:
      cloud:
        gateway:
          default-filters:
            - name: GlobalFilter
              args:
                baseMessage: Spring Cloud Gateway Global Filter
                preLogger: true
                postLogger: true
    ```
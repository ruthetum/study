# @Cacheable
## Overview
- 캐시는 서버의 부담을 줄이고, 성능을 높이기 위해 사용되는 기술이다.
    - 예를 들어 어떤 요청을 처리하는데 계산이 복잡하거나 혹은 DB에서 조회하는게 오래 걸리는 등에 적용하여 결과를 저장해두고 가져옴으로써 빠르게 처리할 수 있다.
    - 캐시는 값을 저장해두고 불러오기 때문에 반복적으로 동일한 결과를 반환하는 경우에 용이하다. 
    - 만약 매번 다른 결과를 돌려줘야 하는 상황에 캐시를 적용한다면 오히려 성능이 떨어지게 된다. 
    - 오히려 캐시에 저장하거나 캐시를 확인하는 작업 때문에 부하가 생기기 때문이다.
    - 그러므로 캐시는 동일한 결과를 반환하는 반복적인 작업과 시간이 오래 걸려서 서버(애플리케이션)에 부담이 되는 경우에 적용하면 좋다.

- 스프링은 AOP 방식으로 편리하게 메소드에 캐시를 적용하는 기능을 제공하고 있다.

- 스프링에서 제공하는 기능을 활용해서 캐시를 적용하고 적용 전/후 성능을 비교해보자

<br>

<details>
<summary>Spring에서 캐시 적용하기</summary>
<div markdown="1">

- 스프링에서는 `@Cacheable`, `@CachePut`, `@CacheEvict`를 활용해서 캐시를 적용할 수 있다.

### 설정
#### @EnableCaching
- @Cacheable과 같은 어노테이션 기반의 캐시 기능을 사용하기 위해서 `@EnableCaching`를 추가한다.

```java
@EnableCaching
@Configuration
public class RedisConfig {
    ...
}
```

#### CacheManager Bean 추가
- 캐시를 관리해줄 CacheManager를 Bean으로 등록해주어야 한다.
- [스프링 공식 문서 - 캐시 추상화 (Cache Abstarction)](https://docs.spring.io/spring-framework/docs/5.0.0.M5/spring-framework-reference/html/cache.html) 에서 현재 지원되는 다양한 매니저의 종류를 확인할 수 있다.

### @Cacheable
- 캐시를 저장/조회를 설정한다.
  
- 보통 메소드에 적용한다.
    - 클래스나 인터페이스에 지정할 수도 있지만 그런 경우는 극히 적다.

- 캐시에 데이터가 없을 경우에는 기존의 로직을 실행한 후에 캐시에 데이터를 추가한다.

- 캐시에 데이터가 있으면 캐시의 데이터를 반환한다.

#### 메소드의 파라미터가 1개인 경우
```java
@Cacheable("todayWebtoon")
public Webtoon getTodayWebtoon(String webtoonNo) {
    // logic
}
```
- 파라미터로 넘어온 `webtoonId`를 기준으로 `todayWebtoon` 캐시에서 값을 조회한다.
    - 값이 없으면 logic을 실행하고, 반환 값을 저장한다.
    - 값이 있으면 저장된 값을 반환한다.
    
- 만약 파라미터가 없는 경우 디폴트 값을 key로 활용하면 된다.

#### 메소드의 파라미터가 여러 개인 경우
```java
@Cacheable(value = "todayWebtoon", key = "webtoonNo")
public Webtoon getTodayWebtoon(String webtoonNo, Date date) {
    // logic
}
```
- 파라미터가 여러 개인 경우 key를 지정해준다.

#### 메소드의 파라미터가 객체인 경우
```java
@Cacheable(value = "todayWebtoon", key = "#webtoon.webtoonNo")
public Webtoon getTodayWebtoon(Webtoon webtoon, Date date) {
    // logic
}
```
- Key값의 지정에는 SpEL이 사용된다. 그렇기 때문에 만약 파라미터가 객체라면 다음과 같이 하위 속성에 접근하면 된다.

### @CachePut
- 캐시에 값을 저장하는 용도로만 사용된다.
- 실행 결과를 캐시에 저장하지만, 조회 시에 저장된 캐시의 내용을 사용하지는 않고 항상 메소드의 로직을 실행한다.

```java
@CachePut(value = "todayWebtoon", key = "webtoonNo")
public Webtoon updateWebtoon(String webtoonNo) {
    // logic
}
```

### @CacheEvict
- 캐시를 제거하는 용도로 사용된다.
- 만약 값이 달라지거나 현재 저장된 값이 무의미한 값이라면 제거되어야 한다.
- 캐시를 제거하기 위해서는 크게 두 가지 방법이 존재한다.
    - 일정한 주기로 캐시를 제거
    - 값이 변할 때 캐시를 제거
- 예를 들어 하루에 한 번씩 바뀌는 정보라면 batch 또는 scheduled 기능을 이용하여 일정 시간 주기로 캐시를 삭제한다.

#### 특정 캐시 내의 값을 모두 제거하는 경우
```java
@CacheEvict(value = "todayWebtoon", allEntries = true)
public void clearTodayWebtoon() {
    // logic    
}
```

#### 캐시 내의 특정 값만 제거하는 경우
```java
@CacheEvict(value = "todayWebtoon", key = "webtoon.webtoonNo")
public void clearWebtoon(Webtoon webtoon) {
    // logic    
}
```

</div>
</details>

<br>

## Scenario
- 오늘의 웹툰을 조회하는 경우를 가정하고, 캐시를 적용했을 때와 적용하지 않았을 때의 성능을 비교한다.
- 두 가지 상황에 대해서 각각의 TPS를 측정하고, 추가로 평균, 최소, 최대 응답시간을 확인한다.

## Experiment
### Environment
- 애플리케이션, DB, Cache 모두 로컬에서 작동한다.
- 웹툰 정보를 저장하는 관계형 데이터베이스는 MySQL, Cache는 Redis를 사용한다.

### Method
- 두 가지 상황(캐시 적용/미적용)에 대한 API를 만들고 각각의 상황에 알맞게 API를 요청한다.
- 이후 nGrinder를 통해 TPS를 측정하고, JMeter를 이용해서 평균, 최소, 최대 응답시간을 측정한다.

### Result


## Conclusion

## Reference
- [스프링 공식 문서 - Cacheable](https://docs.spring.io/spring-framework/docs/current/javadoc-api/org/springframework/cache/annotation/Cacheable.html)
- [스프링 공식 문서 - 캐시 추상화 (Cache Abstarction)](https://docs.spring.io/spring-framework/docs/5.0.0.M5/spring-framework-reference/html/cache.html) 
- [Redis vs Memcached](https://github.com/ruthetum/study/blob/main/redis/compairson-redis-memcached.md)
- [스프링부트 Redis 적용하기](https://kukekyakya.tistory.com/11)
- [스프링 캐시(Cache) 추상화와 사용법(@Cacheable, @CachePut, @CacheEvict)](https://mangkyu.tistory.com/179)
- [baeldung - 스프링 evict cache](https://www.baeldung.com/spring-boot-evict-cache)
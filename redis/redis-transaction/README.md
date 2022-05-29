# Redis - Transaction
- Redis 트랜잭션 기본 명령어
- Redis CLI를 이용한 트랙잭션 테스트
- Spring 환경에서의 트랜잭션 테스트
---

## Redis 트랜잭션
트랜잭션을 유지하기 위해서는 순차성을 가져야 하고 도중에 명령어가 치고 들어오지 못하게 Lock이 필요하다.

Redis에서는 **MULTI**, **EXEC**, **DISCARD**, **WATCH** 명령어를 이용한다.

|Command|Description|
|---|---|
|MULTI|Redis의 트랜잭션을 시작하는 커맨드.<br>트랜잭션을 시작하면 이후 커맨드는 바로 실행되지 않고 queue에 쌓인다.|
|EXEC|정상적으로 처리되어 queue에 쌓여있는 명령어를 일괄적으로 실행한다.<br>RDBMS의 Commit과 동일하다.|
|DISCARD|queue에 쌓여있는 명령어를 일괄적으로 폐기한다.<br>RDMS의 Rollback과 동일하다.|
|WATCH|Redis에서 Lock을 담당하는 명령어입니다. 이 명령어는 낙관적 락(Optimistic Lock) 기반이다.<br>**Watch 명령어를 사용하면 이 후 UNWATCH 되기전에는 1번의 EXEC 또는 Transaction 아닌 다른 커맨드만 허용한다.**|

<details>
<summary>낙관적 락(Optimistic Lock)과 비관적 락(Pessimistic Lock)</summary>
<div markdown="1">

<br>

DB에서 충돌 상황을 방지하기 위해서는 두 가지 방법이 존재한다.

1. 테이블의 row에 접근 시 Lock을 걸고, Lock이 걸려 있지 않을 경우에만 수정이 가능하게 한다.

2. 수정할 때 해당 값을 수정했다고 명시하여, 다른 요청이 동일한 조건으로 값을 수정할 수 없게 한다.

### 비관적 락(Pessimistic Lock)
비관적 락은 Reeatable Read 또는 Serializable 정도의 격리성 수준에서 가능하다.

비관적 락이란 **트랜잭션이 시작될 때 Shared Lock 또는 Exclusive Lock을 걸고 시작하는 방법**이다.

즉, Shared Lock을 걸게 되면 write를 하기위해서는 Exclucive Lock을 얻어야하는데 Shared Lock이 다른 트랜잭션에 의해서 걸려 있으면 해당 Lock을 얻지 못해서 업데이트를 할 수 없게 된다.

수정을 하기 위해서는 해당 트랜잭션을 제외한 모든 트랜잭션이 종료(commit) 되어야한다.

### 낙관적 락(Optimistic Lock)
낙관적 락은 DB 충돌 상황을 개선할 수 있는 방법 중 2번째인 수정할 때 내가 먼저 이 값을 수정했다고 명시하여 다른 사람이 동일한 조건으로 값을 수정할 수 없게 하는 것이다.

특징은 DB에서 제공해주는 특징을 이용하는 것이 아닌 Application Level에서 잡아주는 Lock이다. version 등의 구분 컬럼을 이용해서 충돌을 예방한다.

### 비교
낙관적 락은 트랜잭션을 필요로 하지 않는다. 따라서 성능적으로 비관적 락보다 난곽전 락이 더 좋다.

또한 낙관적 락은 트랜잭션을 필요로 하지 않기 때문에 아래와 같은 로직의 흐름을 가질때도 충돌 감지를 할 수 있다. 만약 비관적 락이라면 1번에서 3번사이의 트랜잭션을 유지할 수가 없다.

> 1. 클라이언트가 서버에 정보를 요청
> 2. 서버에서는 정보를 반환
> 3. 클라이언트에서 이 정보를 이용하여 수정 요청
> 4. 서버에서는 수정 적용 ( 충돌 감지 가능 )

하지만 낙관적 락의 최대 단점은 롤백이다. 만약 충돌이 나는 경우 이를 해결하려면 개발자가 수동으로 롤백처리를 하나하나 해줘야한다.

비관적 락이라면 트랜잭션을 롤백하면 끝나는 작업이지만 낙관적 락은 그렇지 않다. 수동으로 롤백처리는 구현하기도 까다롭지만 성능적으로 보더라도 update를 한번씩 더 해줘야 한다.

따라서 결과적으로 비관적 락 보다 좋지 않을 수 있다. 이러한 단점 때문에 낙관적 락은 충돌이 많이 예상되거나 충돌이 발생했을 때 비용이 많이 들 것이라고 판단되는 곳에서는 사용하지 않는 것이 좋다.

<br>

</div>
</details>

> #### Redis Lock
> WATCH 명령어를 사용해서 특정 Key에 Lock을 걸면 Lock이 걸리게 된다.
> 
> 이 경우에 RDBMS 처럼 Lock을 계속 잡는게 아니라 트랜잭션이 시작된 상황에서 값 변경을 1번으로 제한하는 기능이다.

---
<br>

## Redis CLI를 이용한 트랙잭션 테스트
### MULTI
MULTI 커맨드를 이용하여 트랜잭션 시작한다.

이후에 들어오는 명령어는 바로 실행되는 것이 아니라 큐에 쌓이게("QUEUED") 된다.

그리고 마지막에 EXEC 커맨드를 통해 일괄적으로 실행된다. (GET 커맨드 또한 QUEUED로 쌓이게 된다.)


```shell
> MULTI                 # 트랜잭션 시작
OK
> SET apple iphone      # 1) 값 설정 
QUEUED
> SET samsung galaxy    # 2) 값 설정
QUEUED
> GET apple             # 3) 값 조회
QUEUED
> EXEC                  # 실행
1) OK
2) OK
3) "iphone"
```

![multi](https://user-images.githubusercontent.com/59307414/170666210-19bdcfb0-4d21-46ca-bc69-7da024dcaa74.png)

### Rollback
Discard 명령어를 이용해서 QUEUE에 쌓여있던 명령어가 일괄적으로 없어지게 할 수 있다.

![discard](https://user-images.githubusercontent.com/59307414/170668068-4f73237d-5e12-40e5-9ed2-8518d1e84fa8.png)

#### 잘못된 명령어를 입력하는 경우
잘못된 명령어를 입력하는 경우 QUEUE에 쌓인 모든 명령어가 DISCARD 된다.

![image](https://user-images.githubusercontent.com/59307414/170669304-2a898664-dbe8-4c4a-bd44-90c15940d940.png)

#### 잘못된 자료구조를 사용하는 경우
잘못된 자료구조를 사용하는 경우 QUEUE에 쌓인 모든 명령어가 DISCARD되지 않고 정상적으로 사용한 명령어에 대해서는 잘 적용된다.

이런 트랜잭션 방법이 사용된 이유는 공식 문서에 따르면 이러한 경우는 대부분 개발 과정에서 일어날 수 있는 에러이며 production 환경에서는 거의 발생하지 않는 에러이고, 또한 rollback을 채택하지 않음으로써 빠른 성능을 유지할 수 있다고 한다.

![image](https://user-images.githubusercontent.com/59307414/170671045-7aff6bcd-9a05-46c9-ae57-90c233de52d1.png)

### Lock
WATCH 명령어를 이용해서 특정 Key를 트랜잭션에서 값 변경을 1번으로 제한할 수 있다.

![image](https://user-images.githubusercontent.com/59307414/170672297-738cc9cc-dd83-4060-b797-a988e53c9db9.png)

UNWATCH 명령어를 이용하면 key에 걸린 Lock을 풀어줄 수 있다. (**각각의 key별로 UNWATCH를 직접 선언할 수는 없다.**)

![image](https://user-images.githubusercontent.com/59307414/170673064-beef1362-f44a-476f-a459-1d6a0152ad39.png)

트랜잭션이 여러 번 발생하는 경우 첫 번째를 제외한 나머지 트랜잭션에서는 명령이 성공적으로 수행된다.

이유는 EXEC가 호출 될 때 UNWATCH가 묵시적으로 호출되기 때문에 이후 트랜잭션에 대해서는 해당 키에 대해서는 WATCH가 걸려있지 않기 때문이다.

![image](https://user-images.githubusercontent.com/59307414/170673718-e83aa6ee-98a7-480b-911a-3e7deecc0d45.png)

WATCH 명령어 이후 하나의 트랜잭션에서 여러 번이 값이 수정될 때 값은 제대로 반영된다.

![image](https://user-images.githubusercontent.com/59307414/170674195-47ef958a-0b46-4d0d-81cb-66e5351e552b.png)

---
<br>

## Spring 환경에서의 트랜잭션 테스트
### dependency
```groovy
dependencies {
	implementation 'org.springframework.boot:spring-boot-starter-data-redis'
    // implementation 'org.springframework:spring-tx:{version}'
    implementation 'org.springframework:spring-tx:5.3.9'    // 트랜잭션
}
```

### 설정
```java
public class RedisConfig {

    ...

    @Bean
    public RedisTemplate<String, Object> redisTemplate() {
        RedisTemplate<String, Object> redisTemplate = new RedisTemplate<>();
        redisTemplate.setConnectionFactory(redisConnectionFactory());
        redisTemplate.setEnableTransactionSupport(true); // Redis 트랜잭션 설정
        redisTemplate.setKeySerializer(new StringRedisSerializer());
        redisTemplate.setValueSerializer(new Jackson2JsonRedisSerializer<>(String.class));
        return redisTemplate;
    }
}
```

### Operations 활용
CLI 환경에서 사용했던 것처럼 multi, watch, discard 등의 명령어를 이용하여 트랜잭션을 처리할 수 있다.

다만 transaction을 유지하기 위해서는 동일한 connection을 유지할 필요가 있다.

그러나 일반적인 RedisTemplate 명령어는 connection을 유지하지 않기 때문에 connection을 유지하기 위한 명령어로 SessionCallback를 사용할 필요가 있다.

```java
@Service
public class IndexService {

    private final RedisTemplate redisTemplate;

    public void useOperations() {
        try {
            redisTemplate.execute(new SessionCallback<List<Object>>() {
                public List<Object> execute(RedisOperations operations) {
                    // redis transaction 시작
                    operations.multi();

                    // operations.watch("apple"); // watch 사용 예

                    operations.opsForValue().set("apple", "iphone1");
                    operations.opsForValue().set("samsung", "galaxy1");

                    // operations.discard(); // discard 사용 예

                    // redis transaction 종료
                    return operations.exec();
                }
            });
        } catch (DataAccessException e) {
            e.printStackTrace();
        }
    }
}
```

### @Transactional
@Transactional 애노테이션을 이용하여 connection을 유지하기 위해서는 RedisTemplate 설정에서 setEnableTransactionSupport(true)를 추가해야 한다.

```java
@Configuration
public class RedisConfig {

    ...
    
    @Bean
    public RedisTemplate<String, Object> redisTemplate() {
        RedisTemplate<String, Object> redisTemplate = new RedisTemplate<>();
        redisTemplate.setConnectionFactory(redisConnectionFactory());
        redisTemplate.setEnableTransactionSupport(true); // Redis 트랜잭션 설정
        redisTemplate.setKeySerializer(new StringRedisSerializer());
        redisTemplate.setValueSerializer(new Jackson2JsonRedisSerializer<>(String.class));
        return redisTemplate;
    }
}
```

이후에는 원하는 메서드에 @Transactional 애노테이션을 붙여서 사용하면 된다.

@Transactional 애노테이션이 붙게 되면 메서드 시작 시 multi가 실행되고, 메서드 종료 시 exec이 작동된다. (ThreadLocal 기반으로 작동된다.)

예외가 발생하는 경우에는 discard가 실행된다.

```java
@Service
public class IndexService {

    private final RedisTemplate redisTemplate;

    @Transactional
    public void useTransactionalAnnotation() {
        redisTemplate.opsForValue().set("apple", "iphone2");
        redisTemplate.opsForValue().set("samsung", "galaxy2");
    }
}
```

> @Transactional은 ThreadLocal 기반이기 때문에 reactive 환경에서는 정상적으로 동작하지 않는다.
> 
> reactive 환경에서 Transactional을 유지하기 위해서는 Netty 기반의 Redisson을 이용해야 한다.

---
<br>

## Reference
- https://redis.io/docs/manual/transactions/
- https://redis.io/topics/transactions
- https://insanelysimple.tistory.com/344
- https://sabarada.tistory.com/177
- https://sabarada.tistory.com/178?category=856943
- https://sabarada.tistory.com/175
- https://minholee93.tistory.com/entry/Redis-Transaction
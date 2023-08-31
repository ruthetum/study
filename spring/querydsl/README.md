# QueryDSL
## QueryDSL 설정
```gradle
buildscript {
	ext {
		queryDslVersion = "5.0.0"
	}
}

plugins {
	id 'org.springframework.boot' version '2.6.2'
	id 'io.spring.dependency-management' version '1.0.11.RELEASE'
	//querydsl 추가
	id "com.ewerk.gradle.plugins.querydsl" version "1.0.10"
	id 'java'
}

group = 'com.example'
version = '0.0.1-SNAPSHOT'
sourceCompatibility = '11'

configurations {
	compileOnly {
		extendsFrom annotationProcessor
	}
}

repositories {
	mavenCentral()
}

dependencies {
	implementation 'org.springframework.boot:spring-boot-starter-data-jpa'
	implementation 'org.springframework.boot:spring-boot-starter-web'
	//querydsl 추가
	implementation "com.querydsl:querydsl-jpa:${queryDslVersion}"
	implementation "com.querydsl:querydsl-apt:${queryDslVersion}"
	compileOnly 'org.projectlombok:lombok'
	runtimeOnly 'com.h2database:h2'
	annotationProcessor 'org.projectlombok:lombok'
	testImplementation 'org.springframework.boot:spring-boot-starter-test'
}

//querydsl 추가
def querydslDir = "$buildDir/generated/querydsl"

querydsl {
	jpa = true
	querydslSourcesDir = querydslDir
}
sourceSets {
	main.java.srcDir querydslDir
}
configurations {
	compileOnly {
		extendsFrom annotationProcessor
	}
	querydsl.extendsFrom compileClasspath
}
compileQuerydsl {
	options.annotationProcessorPath = configurations.querydsl
}

test {
	useJUnitPlatform()
}
```
- compileQuerydsl 오류 관련 해결 : https://www.inflearn.com/questions/355723

## Q type 뽑기
- 우측 Gradle - Tasks - other - compileQuerydsl 클릭

## Querydsl vs JPQL
```java
public class QuerydslBasicTest {

    @PersistenceContext
    EntityManager em;

    JPAQueryFactory queryFactory;

    @BeforeEach
    public void before() {

        queryFactory = new JPAQueryFactory(em);
        ...
    }

    @Test
    public void startJPQL() {
        String qlString =
                "select m from Member m " +
                        "where m.username = :username";

        Member findMember = em.createQuery(qlString, Member.class)
                .setParameter("username", "member1")
                .getSingleResult();

        Assertions.assertThat(findMember.getUsername()).isEqualTo("member1");
    }

    @Test
    public void startQuerydsl() {
        QMember m = new QMember("m");

        Member findMember = queryFactory
                .select(m)
                .from(m)
                .where(m.username.eq("member1")) //파라미터 바인딩 처리
                .fetchOne();

        Assertions.assertThat(findMember.getUsername()).isEqualTo("member1");
    }
}
```

## JPQL 검색 조건
```
member.username.eq("member1") // username = 'member1'
member.username.ne("member1") //username != 'member1'
member.username.eq("member1").not() // username != 'member1'

member.username.isNotNull() //이름이 is not null

member.age.in(10, 20) // age in (10,20)
member.age.notIn(10, 20) // age not in (10, 20)
member.age.between(10,30) //between 10, 30

member.age.goe(30) // age >= 30
member.age.gt(30) // age > 30
member.age.loe(30) // age <= 30
member.age.lt(30) // age < 30

member.username.like("member%") //like 검색
member.username.contains("member") // like ‘%member%’ 검색
member.username.startsWith("member") //like ‘member%’ 검색
```

## 결과 조회
- `fetch()` : 리스트 조회, 데이터 없으면 빈 리스트 반환
- `fetchOne()` : 단 건 조회
    - 결과가 없으면 : null
    - 결과가 둘 이상이면 : `com.querydsl.core.NonUniqueResultException`
- `fetchFirst()` : `limit(1).fetchOne()`
- `fetchResults()` : 페이징 정보 포함, total count 쿼리 추가 실행 / `deprecated`
- `fetchCount()` : count 쿼리로 변경해서 count 수 조회 / `deprecated`

```java
// List
List<Member> fetch = queryFactory
        .selectFrom(member)
        .fetch();

//단 건
Member findMember1 = queryFactory
        .selectFrom(member)
        .fetchOne();

//처음 한 건 조회
Member findMember2 = queryFactory
        .selectFrom(member)
        .fetchFirst();

//페이징에서 사용
QueryResults<Member> results = queryFactory
        .selectFrom(member)
        .fetchResults();

//count 쿼리로 변경
long count = queryFactory
        .selectFrom(member)
        .fetchCount();
```

- `fetchResults()`, `fetchCount()` 는 5.0.0 버전에서 deprecated 됨
  > fetchResults() : Get the projection in QueryResults form. Make sure to use fetch() instead if you do not rely on the QueryResults.getOffset() or QueryResults.getLimit(), because it will be more performant. Also, count queries cannot be properly generated for all dialects. For example: in JPA count queries can’t be generated for queries that have multiple group by expressions or a having clause. Get the projection in QueryResults form. Use fetch() instead if you do not need the total count of rows in the query result.

  > fetchCount() : An implementation is allowed to fall back to fetch().size().

- 페이징 처리 방법
    ```java
    public Page<User> findUserWithPaging(Pageable pageable) {
    
        List<User> content = queryFactory
                .selectFrom(user)
                .where(user.username.like("user_"))
                .offset(pageable.getOffset()) // offset
                .limit(pageable.getPageSize()) // limit
                .fetch();
    
        return new PageImpl<>(content, pageable, content.size()); // 쿼리 결과로 페이징 객체 리턴
    }
    ```
- cf. https://devwithpug.github.io/java/querydsl-with-datajpa/

## 정렬
- `desc()` , `asc()` : 일반 정렬
- `nullsLast()` , `nullsFirst()` : null 데이터 순서 부여

## 집합
```java
/**
 * JPQL
 * select
 * COUNT(m), //회원수
 * SUM(m.age), //나이 합
 * AVG(m.age), //평균 나이
 * MAX(m.age), //최대 나이
 * MIN(m.age) //최소 나이
 * from Member m
 */
@Test
public void aggregation() throws Exception {
    List<Tuple> result = queryFactory
            .select(member.count(),
                    member.age.sum(),
                    member.age.avg(),
                    member.age.max(),
                    member.age.min())
            .from(member)
            .fetch();

    Tuple tuple = result.get(0);
    Assertions.assertThat(tuple.get(member.count())).isEqualTo(4);
    Assertions.assertThat(tuple.get(member.age.sum())).isEqualTo(100);
    Assertions.assertThat(tuple.get(member.age.avg())).isEqualTo(25);
    Assertions.assertThat(tuple.get(member.age.max())).isEqualTo(40);
    Assertions.assertThat(tuple.get(member.age.min())).isEqualTo(10);
}
```

### Group by
```java
List<Tuple> result = queryFactory
        .select(team.name, member.age.avg())
        .from(member)
        .join(member.team, team)
        .groupBy(team.name)
        .fetch();
```

## Join
### join vs fetch join
-  일반 Join
   > Fetch Join과 달리 연관 Entity에 Join을 걸어도 실제 쿼리에서 SELECT 하는 Entity는
   오직 JPQL에서 조회하는 주체가 되는 Entity만 조회하여 영속화
   >
   > 조회의 주체가 되는 Entity만 SELECT 해서 영속화하기 때문에 데이터는 필요하지 않지만 연관 Entity가 검색조건에는 필요한 경우에 주로 사용됨

- Fetch Join
  > 조회의 주체가 되는 Entity 이외에 Fetch Join이 걸린 연관 Entity도 함께 SELECT 하여 모두 영속화
  >
  > Fetch Join이 걸린 Entity 모두 영속화하기 때문에 FetchType이 Lazy인 Entity를 참조하더라도
  이미 영속성 컨텍스트에 들어있기 때문에 따로 쿼리가 실행되지 않은 채로 N+1문제가 해결됨

- cf. https://cobbybb.tistory.com/18

```java
List<Tuple> result = queryFactory
        .select(member)
        .from(member)
        .join(member.team, team).fetchJoin() // fetch join도 가능
        .fetch();
```

## 서브 쿼리
- `com.querydsl.jpa.JPAExpressions` 사용
- Q type만 다르게 잘 설정해주자
    ```java
    public void subQuery() {
        QMember memberSub = new QMember("memberSub");

        List<Member> result = queryFactory
                .selectFrom(member)
                .where(member.age.eq(
                        JPAExpressions
                            .select(memberSub.age.max())
                            .from(memberSub)
                ))
                .fetch();

        Assertions.assertThat(result).extracting("age").containsExactly(40);
    }
    ```

## Case
```java
List<String> result = queryFactory
    .select(member.age
            .when(10).then("열살")
            .when(20).then("스무살")
            .otherwise("기타"))
    .from(member)
    .fetch();
```
```java
NumberExpression<Integer> rankPath = new CaseBuilder()
        .when(member.age.between(0, 20)).then(2)
        .when(member.age.between(21, 30)).then(1)
        .otherwise(3);

List<Tuple> result = queryFactory
        .select(member.username, member.age, rankPath)
        .from(member)
        .orderBy(rankPath.desc())
        .fetch();
```

## 상수
- 상수가 필요하면 `Expressions.constant(xxx)` 사용
    ```
    Tuple result = queryFactory
            .select(member.username, Expressions.constant("A"))
            .from(member)
            .fetchFirst()
    ```

- concat
    ```java
    String result = queryFactory
            .select(member.username.concat("_").concat(member.age.stringValue()))
            .from(member)
            .where(member.username.eq("member1"))
            .fetchOne();
    ```
    - `stringValue()` 활용


## DTO 반환
1. 프로퍼티(Setter) 접근
```java
List<MemberDto> result = queryFactory
                .select(Projections.bean(MemberDto.class,
                        member.username,
                        member.age))
                .from(member)
                .fetch();
```
2. 필드 직접 접근
```java
List<MemberDto> result = queryFactory
        .select(Projections.fields(MemberDto.class,
        member.username,
        member.age))
        .from(member)
        .fetch();
```
3. 생성자 접근
```java
List<MemberDto> result = queryFactory
                .select(Projections.constructor(MemberDto.class,
                        member.username,
                        member.age))
                .from(member)
                .fetch();
```

- 별칭이 다른 경우
    - `as` 활용
    ```java
    List<UserDto> fetch = queryFactory
             .select(Projections.fields(UserDto.class,
                 member.username.as("name"),
                 ExpressionUtils.as(
                     JPAExpressions
                     .select(memberSub.age.max())
                     .from(memberSub), "age")
                 )
            ).from(member)
            .fetch();
    ```
  
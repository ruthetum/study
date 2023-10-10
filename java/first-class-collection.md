# 일급 컬렉션 (First Class Collection)

> **왜 작성하게 되었는가**
> 
> - [Nextstep java playground](https://edu.nextstep.camp/s/RFY359FE/ls/cPVALOmI) 진행 도중 일급 컬렉션이라는 잘 모르는 용어가 등장함
> 
> **그래서 무엇을 작성하는가**
> 
> - 일급 컬렉션에 대한 정의
> - 일급 컬렉션 사용 이점
> 
> **작성을 통해 무엇을 얻고자 하는가**
>
> - 현재 기준으로 책을 읽을 계획은 없기 때문에 관련 내용이 담긴 포스트를 바탕으로 일급 컬렉션에 대한 정의를 인지하고자함
> - 소트윅스 앤솔로지에서 언급한 해당 규칙(일급 컬렉션 사용)을 이해하고자 함

## 일급 컬렉션이란?
일급 컬렉션이란 단어는 [소트웍스 앤솔로지](https://www.yes24.com/Product/Goods/3290339)의 객체지향 생활체조 파트에서 언급됨

> #### 규칙 8. 일급 컬렉션 사용
> 이 규칙의 적용은 간단하다. 
> 
> 콜렉션을 포함한 클래스는 반드시 다른 멤버 변수가 없어야 한다.
>
> 각 콜렉션은 그 자체로 포장돼 있으므로 이제 콜렉션과 관련된 동작은 근거지가 마련된 셈이다.
> 
> 필터가 이 새 클래스의 일부가 됨을 알 수 있다.
> 
> 필터는 또한 스스로 함수 객체가 될 수 있다.
> 
> 또한 새 클래스는 두 그룹을 같이 묶는다든가 그룹의 각 원소에 규칙을 적용하는 등의 동작을 처리할 수 있다.
> 
> 이는 인스턴스 변수에 대한 규칙의 확실한 확장이지만 그 자체를 위해서도 중요하다. 
> 
> 콜렉션은 실로 매우 유용한 원시 타입이다.
>
> 많은 동작이 있지만 후임 프로그래머나 유지보수 담당자에 의미적 의도나 단초는 거의 없다.

### 일급 컬렉션

간단하게 설명하면 Collection을 Wrapping하면서 그 외의 다른 멤버 변수가 없는 상태를 일급 컬렉션이라고 함

```java
Map<String, String> map = new HashMap<>();

map.put("J", "Java");
map.put("K", "Kotlin");
map.put("G", "Go");
```

Collection을 Wrapping하면서 그 외 다른 멤버 변수가 없는 상태로 변환

```java
public Class ProgrammingLanguage {
    
    private Map<String, String> languages;
    
    public ProgrammingLanguage(Map<String, String> languages) {
        this.languages = languages;
    }
}
```

## 일급 컬렉션 사용 이점
1. 비즈니스에 종속적인 자료구조
2. Collection의 불변성을 보장
3. 상태와 행위를 한 곳에서 관리 
4. 이름이 있는 컬렉션

### 1. 비즈니스에 종속적인 자료구조

예를 들어 Linux의 Alias 커맨드를 객체로 구현하는 경우 커맨드 규칙으로 `띄어쓰기를 포함할 수 없다`와 `중복된 key를 가질 수 없다`를 설정할 수 있음

보통 이러한 제한 규칙은 서비스 메서드 내에서 진행

하지만 이 경우 커맨드가 필요한 모든 장소에서 검증 로직을 구현해야 하는 이슈 발생
- 모든 코드와 도메인을 알고 있지 않는 경우 언제든 문제가 발생할 여지가 있음

따라서 이 문제를 깔끔하게 해결하기 위해서는 앞선 제한 규칙을 갖는 자료구조를 직접 만들면 해결됨

```java
public class Alias {

    private Map<String, String> command;

    public Alias(Map<String, String> command) {
        validateKeyName(command);
        validateKeyDuplication(command);
        this.command = command;
    }

    private void validateKeyName(Map<String, String> command) {
        for (String key : command.keySet()) {
            if (StringUtils.containsWhitespace(key)) {
                throw new IllegalArgumentException("command에는 공백을 포함할 수 없습니다.");
            }
        }
    }

    private void validateKeyDuplication(Map<String, String> command) {
        if (command.size() != command.keySet().stream().distinct().count()) {
            throw new IllegalArgumentException("command에 중복된 key가 있습니다.");
        }
    }
}
```

이렇게 일급 컬렉션을 생성하는 경우 비즈니스에 종속적인 자료구조를 생성할 수 있음

### 2. Collection의 불변성을 보장

Java 에서 `final` 키워드는 불변 객체를 만들어주는 것이 아닌 재할당만을 금지

따라서 아래와 같이 값 변경은 가능함

```java
Public class CollectionTest {

    @Test
    void fianl은_값_변경이_가능() {
        //given
        final Map<String, Boolean> collection = new HashMap<>();

        //when
        collection.put("1", true);
        collection.put("2", true);
        collection.put("3", true);
        collection.put("4", true);

        //then
        assertThat(collection.size()).isEqualTo(4);
    }
}
```

Java 에서는 final로 이 문제를 해결할 수 없기 때문에 일급 컬렉션과 래퍼 클래스 등의 방법으로 해결해야 함

```java
public class Orders {

    private final List<Order> orders;
    
    public Orders(List<Order> orders) {
        this.orders = orders;
    }
    
    public void getAmountSum(Order order) {
        return this.orders.stream()
                .map(Order::getAmount)
                .sum();
    }
}
```

### 3. 상태와 행위를 한 곳에서 관리
일급 컬렉션은 값과 로직이 함께 존재한다는 장점이 있음

예를 들어 특정 과일의 수량을 구해야하는 경우 아래와 같이 작성

```java
public class FruitTest {
    @Test
    void 로직이_밖에_있는_경우() {
        //given
        List<Fruit> fruits = Arrays.asList(
                new Fruit(APPLE, 2),
                new Fruit(BANANA, 3),
                new Fruit(APPLE, 4),
                new Fruit(ORANGE, 1)
        );

        //when
        int appleAmount = fruits.stream()
                .filter(fruit -> fruit.getType().equals(APPLE))
                .mapToInt(Fruit::getAmount)
                .sum();


        //then
        assertThat(appleAmount).isEqualTo(6);
    }
}
```

이 경우 List에 데이터를 담고, Service 에서 필요한 로직을 수행
- `fruits` 라는 컬렉션과 계산 로직은 서로 관계가 있는데, 이를 코드로 표현할 수 있음

따라서 일급 컬렉션 생성을 통해 상태와 로직을 한 곳에서 처리

```java
public class FruitGroup {
    
    private List<Fruit> fruits;
    
    public FruitGroup(List<Fruit> fruits) {
        this.fruits = fruits;
    }

    public int getAppleAmount() {
        return getFilterAmount(APPLE);
    }

    public int getBananaAmount() {
        return getFilterAmount(BANANA);
    }

    public int getOrangeAmount() {
        return getFilterAmount(ORANGE);
    }
    
    public int getFilterAmount(FruitType type) {
        return fruits.stream()
                .filter(fruit -> fruit.getType().equals(type))
                .mapToInt(Fruit::getAmount)
                .sum();
    }
}

public class FruitTest {
    @Test
    void 로직과_값이_한_곳에_있는_경우() {
        //given
        List<Fruit> fruits = Arrays.asList(
                new Fruit(APPLE, 2),
                new Fruit(BANANA, 3),
                new Fruit(APPLE, 4),
                new Fruit(ORANGE, 1)
        );
        FruitGroup fruitGroup = new FruitGroup(fruits);

        //when
        int appleAmount = fruitGroup.getAppleAmount();

        //then
        assertThat(appleAmount).isEqualTo(6);
    }
}
```

### 4. 이름이 있는 컬렉션
마지막 장점은 컬렉션에 이름을 붙일 수 있다는 점

같은 과일(Fruit)들의 리스트이지만, 사과(APPLE) 리스트와 바나나(BANANA) 리스트는 다른 리스트

이 때 가장 흔하게 구분하는 방법은 아래와 같이 변수명을 다르게 하는 것

```java
public class FruitTest {
    @Test
    public void 컬렉션을_변수명으로_구분() {
        // given
        List<Fruit> apples = createApples();
        List<Fruit> bananas = createBananas();
        
        // when
        
        // then
    }
}
```

하지만 이 경우 아래의 단점 존재
- 검색의 어려움 
  - 사과 그룹이 어떻게 사용되는지 검색 시 변수명으로만 검색할 수 있음
  - 변수명은 상황 및 사람에 따라 다르게 지을 수 있기 때문에 검색의 어려움이 있음
- 명확한 표현이 불가능 
  - 변수명에 불과하기 때문에 의미를 부여하기가 어려움
  - 이는 개발팀/운영팀간에 의사소통 시 보편적인 언어로 사용하기가 어려움을 이야기함
  - 중요한 값임에도 이를 표현할 명확한 단어가 없기 때문

```java
public class FruitTest {
    @Test
    public void 일급컬렉션으로_구분하기() {
        // given
        Apples apples = new Apples(createApples());
        Bananas bananas = new Bananas(createBananas());

        // when

        // then
    }
}
```

## Reference
- https://jojoldu.tistory.com/412
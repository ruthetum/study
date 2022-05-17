# Item 33. 타입 안전 이종 컨테이너를 고려하라
## 타입 안전 이종 컨테이너 패턴 (Type Safe Heterogeneous Container Pattern)
- 제니릭은 **Set<E>**, **Map<K,V>** 등의 컬렉션과 **ThreadLocal<T>**, **AtomicReference<T>** 등의 단일원소 컨테이너에서 사용된다.
    - 이런 쓰임에서 매개변수화되는 대상은 원소가 아닌 컨테이너 자신이다.
    - 따라서 하나의 컨테이너에서 매개변수화할 수 있는 타입의 수가 제한된다.
    - Ex. Set : 원소의 타입, Map : Key, Value

- 하지만 예를 들어 데이터베이스의 값처럼 행과 임의 개수의 열을 가질 때, 모든 열의 타입을 안전하게 이용할려면 어떻게 해야할까?

- 컨테이너 대신 키를 매개변수화한 다음, 컨테이너 값을 넣거나 뺄 때 매개변수화한 키를 함께 제공하면 된다.

- 이렇게 하면 제네릭 타입 시스템이 값의 타입이 키와 같음을 보장해주고, 이러한 설계 방식을 **타입 안전 이종 컨테이너 패턴**이라 한다.

```java
public class Favorites {
    // 타입 안전 이종 컨테이너 패턴 - 구현
    private Map<Class<?>, Object> favorites = new HashMap<>();

    public <T> void putFavorites(Class<T> type, T instance) {
        favorites.put(Objects.requireNonNull(type), instance);
    }

    public <T> void getFavorites(Class<T> type) {
        return type.cast(favorites.get(type));
    }

    // 타입 안전 이종 컨테이너 패턴 - 클라이언트
    public static void main(String[] args) {
        Favorites f = new Favorites();

        f.putFavorites(String.class, "Java");
        f.putFavorites(Integer.class, 0xcafebabe);
        f.putFavorites(Class.class, Favorites.class);

        String favoritesString = f.getFavorites(String.class);
        int favoritesInteger = f.getFavorites(Integer.class);
        Class<?> favoritesClass = f.getFavorites(Class.class);

        System.out.printf("%s %x %s%n", favoriteString, favoriteInteger, favoriteClass.getName()); // Java cafebabe Favorites
    }
}
```
### Map<Class<?>, Object>
- Favorites가 사용하는 private Map 변수인 favorites의 타입은 Map<Class<?>, Object>이다.
    - 비한정적 와일드카드 타입이라 맵 안에 아무것도 넣을 수 없다고 생각할 수 있지만, 와일드카드 타입이 중첩(nested)된다.
        - 와일드카드 타입 : 제네릭 코드에서 물음표(?)로 표기되어 있는 타입
        - cf. https://snoop-study.tistory.com/113
    - 모든 키가 서로 다른 매개변수화 타입일 수 있다. (= 다양한 타입을 지원한다.)
        - ex. `Class<String>`, `Class<Integer>`

### 동적 형변환
- favorites 맵의 값 타입은 단순히 Object이다.
    - 맵이 키와 값 사이의 타입 관계를 보증하지 않는다. (= 모든 값이 키로 명시한 타입임을 보증하지 않는다.)

- putFavorites 구현의 경우 주어진 Class 객체와 즐겨찾기 인스턴스를 favorites에 추가해 관계를 지으면 끝이다.
        - 키와 값 사이의 타입 링크(Type Linkage) 정보는 버려진다.

- getFavorites의 경우 주어진 Class 객체에 해당하는 값을 favorites Map에서 꺼내고, Object 타입의 값을 cast 메소드를 통해 T로 캐스팅하면 된다. 
    - cast 메소드는 형변환 연산자의 동적 버전으로, 주어진 인수가 Calss 객체가 알려주는 타입의 인스턴인지 검ㅅ한 후 맞다면 인수를 그대로 반환하고, 아니면 ClassCastException을 던진다.

### 제약 조건
- 위 구현 예제의 Favorites 클래스에는 두 가지 문제점이 있다.
    - 1. 타입 안정성의 취약점이 존재한다.
    - 2. 실체화 불가 타입에는 사용할 수 없다.

#### 1. 타입 안정성의 취약점이 존재한다.
- 악의적인 클라이언트가 Class 객체를 제네릭이 아닌 로 타입(Raw Type)으로 넘기면 Favorites 인스턴스의 타입 안전성이 쉽게 깨진다.
    - 하지만 이 경우 클라이언트 코드를 컴파일할 때 비검사 경고가 뜬다.

```java
f.putFavorites((Class)Integer.class, "Non Integer");    // 여기까지는 문제없이 작동하지만
int favoritesInteger = f.getFavorites(Integer.class);   // 여기서 ClassCastException이 던져짐
```

- Favorites가 타입 불변식을 어기는 일이 없도록 보장하려면 putFavorites 메서드에서 인수로 주어진 instance 타입이 type으로 명시한 타입과 같은지 확인하면 된다.
    - 동적 형변환을 통해 런타임 타입의 안전성을 확보한다.

```java
public <T> void putFavorites(Class<T> type, T instance) {
    favorites.put(Objects.requireNonNull(type), type.cast(instance));
}
```

- java.util.Collections에 checkedSet, checkedList, checkedMap같은 메서드가 위 방식을 적용한 컬렉션 래퍼이다.

#### 2. 실체화 불가 타입에는 사용할 수 없다.
- String이나 String[]은 저장할 수 있어도 `List<String>`은 저장할 수 없다.
    - `List<String>`용 Class 객체를 얻을 수 없기 때문에 `List<String>`을 저장하는 코드는 컴파일되지 않는다.
    - `List<String>`과 `List<Integer>`는 List.class라는 같은 Class 객체를 공유하기 때문이다.

- 이 문제를 [슈퍼 타입 토큰](http://gafter.blogspot.com/2006/12/super-type-tokens.html)을 통해 해결할 수는 있지만 완벽하지 않기 때문에 주의해서 사용해야 한다.
    - cf. https://yangbongsoo.gitbook.io/study/super_type_token

### 한정적 타입 토큰 사용
- 현재 Favorites가 사용하는 타입 토큰은 비한정적이기 때문에 getFavorites, putFavorites은 어떤 Class 객체든 받아들일 수 있다.

- 때로는 이 메서드들이 허용하는 타입을 제한하고 싶을 수 있는데, 이 때 한정적 타입 토큰을 활용하면 가능하다. 
    - 한정적 타입 토큰이란 단순히 한정적 타입 매개변수(Item 29)나 한정적 와일드카드(Item 31)를 사용하여 표현 가능한 타입을 제한하는 타입 토큰이다.

- 애너테이션 API(Item 39)는 이런 한정적 타입 토큰을 적극적으로 사용한다.

```java
// asSubclass를 사용하여 한정적 타입 토큰을 안전하게 형변환한다.
static Annotation getAnnotation(AnnotatedElement element, String annotationTypeName) {
    
    Class<?> annotationType = null;    // 비한정적 타입 토큰

    try {
        annotationType = Class.forName(annotationTypeName);
    }  catch (Exception e) {
        throw new IllegalArgumentEception(e);
    }

    return element.getAnnotation(annotationType.asSubclass(Annotation.class));
}

```

- 컴파일 시점에서 타입을 알 수 없는 애너테이션을 asSubclass 메소드를 사용해 런타임에서 읽어낼 수 있다.


## 정리
- 컬렉션 API로 대표되는 일반적인 제네릭 형태에서는 한 컨테이너가 다룰 수 있는 타입 매개변수의 수가 고정되어 있다.

- 컨테이너 자체가 아닌 키를 타입 매개변수로 바꾸면 이런 제약이 없는 타입 안전 이종 컨테이너를 만들 수 있다.

- 타입 안전 이종 컨테이너는 Class를 키로 쓰고, 이런 식으로 쓰이는 Class 객체를 타입 토큰이라고 한다.
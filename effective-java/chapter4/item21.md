# Item 21. 인터페이스는 구현하는 쪽을 생각해 설계하라

## 자바 8부터 기존 인터페이스에 메서드를 추가하는 방법 : 디폴트 메서드
- 자바 8부터 기존 인터페이스에 메서드를 추가할 수 있게 되었다.
    - 자바 8 이전에는 기존 구현체를 깨뜨리지 않고는 인터페이스에 메서드를 추가할 수 없었다.

- 기존 인터페이스에 메서드를 추가하기 위해 **디폴트 메서드**를 선언하면 인터페이스를 구현한 후 디폴트 메서드를 재정의하지 않은 모든 클래스에서 디폴트 구현이 쓰이게 된다.

- 하지만 모든 기존 구현체들과 매끄럽게 연동된다는 보장이 없기 때문에 모든 상황에서 영향을 끼치지 않는 디폴트 메서드를 작성하는 것은 어렵다.

### removeIf
- predicate가 true를 반환하는 모든 원소를 제거한다.

```java
// 자바 8의 Collection 인터페이스에 추가된 디폴트 메서드
default boolean removeIf(Predicate<? super E> filter) {
    Objects.requireNonNull(filter);
    boolean removed = false;
    final Iterator<E> each = iterator();
    while (each.hasNext()) {
        if (filter.test(each.next())) {
            each.remove();
            removed = true;
        }
    }
    return removed;
}
```

```java
// removeIf 적용
public class Item21 {
    public static void main(String[] args) {
        List<Integer> list = new ArrayList<>(Arrays.asList(1, 2, 3, 4, 5));
        System.out.println(list.toString());        // [1, 2, 3, 4, 5]

        // 3보다 작은 수 제거
        list.removeIf(e -> e < 3);
        System.out.println(list.toString());        // [3, 4, 5]
    }
}
```

- 제공되고 있는 removeIf 메서드의 코드보다 더 범용적으로 구현하기도 어렵겠지만, 현존하는 모든 Collection 구현체와 잘 연동되는 것은 아니다.

### SynchronizedCollection
- 대표적인 예로 `org.apache.commons.collection4.collection.SynchronizedCollection`가 있다.
    - https://commons.apache.org/proper/commons-collections/apidocs/org/apache/commons/collections4/collection/SynchronizedCollection.html

- 이 클래스는 **java.util**의 **Collections.synchronizedCollection** 정적 팩토리 메서드가 반환하는 클래스와 유사하다.

- 아파치 버전은 클라이너트가 제공한 객체로 락을 거는 능력을 추가로 제공한다.

- 즉, 모든 메서드에서 주어진 락 객체로 동기화한 후 내부 컬랙션 객체에 기능을 위임하는 래퍼 클래스(item 18)다.

- 4.4 버전 이전에는 **removeIf** 메서드를 재정의하지 않고 있었는데, 이 클래스를 자바 8과 함께 사용하여 **removeIf**의 디폴트 구현을 물려받게 되면 모든 메서드 호출을 동기화해준다는 규약을 어기게 된다.

    - **SynchronizedCollection** 인스턴스를 여러 스레드가 공유한 상황에서 한 스레드가 **removeIf**를 호출하면 `ConcurrentModificationException`이 발생할 수 있다.

    - 현재 기준(22.04.17)으로 버전 4.4가 최신 버전으로 제공되고 있고, **removeIf**는 재정의되어있다. (하단의 부록 참조)

- 자바 플랫폼 라이브러리는 구현한 인터페이스의 디폴트 메서드를 재정의하고, 다른 메서드에서 디폴트 메서드를 호출할 전에 필요한 작업을 수행하도록 했다.
    - ex. **Collections.synchronizedCollection**이 반환하는 package-private 클래스들은 **removeIf**를 재정의하고, 이를 호출하는 다른 메서드들은 디폴트 구현을 호출하기 전에 동기화

- 하지만 자바 플랫폼에 속하지 않은 제 3의 기존 컬렉션 구현체들 중 일부는 여전히 수정되지 않고 있다.

## 정리
- 자바 8에 와서 디폴트 메서드를 통해 기존 인터페이스에 메서드를 추가할 수 있게 되었다.

- 하지만 앞서 언급한 위험 요소들 때문에 **디폴트 메서드는 컴파일에 성공하더라도 기존 구현체에 런타임 오류를 일으킬 수 있다.**

- 기존 인터페이스에 디폴트 메서드로 새 메서드를 추가하는 일은 꼭 필요한 경우가 아니면 피해야 한다.

- 새로운 인터페이스를 만드는 경우라면 사용해도 좋다.

- 디폴트 메서드라는 도구가 생겼더라도 인터페이스를 설계할 때는 여전히 세심한 주의를 기울여야 한다. 


#### 부록
- `org.apache.commons.collection4.collection.SynchronizedCollection`는 버전 4.4부터 removeIf가 override되어었다.

- 4.3 버전
    
    ![4.3](https://user-images.githubusercontent.com/59307414/163707195-b28f6b3e-4939-4076-9330-bcfaa10d37ac.png)

- 4.4 버전

    ![4.4](https://user-images.githubusercontent.com/59307414/163707218-c2bfc358-d329-4955-ae25-c8ab50e95627.png)
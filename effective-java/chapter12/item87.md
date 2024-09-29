# Item 87. 커스텀 직렬화 형태를 고려해보라

개발 일정에 쫓기는 상황이라면 현재 릴리스에서는 그냥 동작만 하도록 만들고, 다음 릴리스에서 제대로 구현한다. 그리고 보통은 크게 문제되지는 않는다.

하지만 **Serializable**을 구현하고 기본 직렬화 형태를 사용하게 되면 다음 릴리스 때 수정하려고 한 현재의 구현 형태에 발이 묶이게 된다. (= 기본 직렬화 형태를 수정하기 어려워 진다.)

따라서 **먼저 고민해보고 괜찮다고 판단될 때만 기본 직렬화 형태를 사용해야 한다.**

- 유연성, 성능, 정확성 측면에서 신중히 고민한 후 합당한 경우에 사용해야 한다.
- 본인이 커스텀해서 설계했을 때와 기본 직렬화 형태가 비슷한 결과 나오는 경우에만 기본 형태를 적용해야 한다.

## 기본 직렬화 형태에 적합한 경우

객체의 물리적 표현과 논리적 내용이 같다면 기본 직렬화 형태를 적용해도 괜찮다.

```java
public class Name implements Serializable {
    /**
     * 성, null이 아니어야 함
     * @serial
     */
    private final String lastName;

    /**
     * 이름, null이 아니어야 함
     * @serial
     */
    private final String firstName;

    /**
     * 중간 이름, 중간 이름이 없는 경우 null
     * @serial
     */
    private final String middleName;

    ...
}
```

이름은 이름, 성, 중간 이름이라는 3개의 문자열로 구성되고, 각 필드는 논리적 내용을 그대로 반영한다.

기본 직렬화 형태가 적합하다고 결정했더라도 불변식 보장과 보안을 위해 **readObject** 메서드를 제공해야 할 때가 많다.

- 위의 Name 클래스에서 readObject 메서드는 lastName과 firstName 필드는 null이 아님을 보장해야 한다.


## 기본 직렬화 형태에 적합하지 않은 경우

```java
public final class StringList implements Serializable {
    private int size = 0;
    private Entry head = null;

    private static class Entry implements Serializable {
        String data;
        Entry next;
        Entry previous;
    }

    ...
}
```

위의 코드는 논리적으로 일련의 문자열을 표현하고, 물리적으로는 문자연들을 이중 연결 리스트로 연결한다.

**객체의 물리적 표현과 논리적 표현의 차이가 클 때 기본 직렬화 형태를 사용하면 크게 네 가지 면에서 문제가 생긴다.**

1. 공개 API가 현재의 내부 표현 방식에 영구히 묶인다. 
- 앞의 코드에서 StringList.Entry가 공개 API가 되는데 다음 릴리즈에서 내부 표현 방식을 바꾸더라도 StringList는 여전히 연결 리스트로 표현된 입력도 처리할 수 있어야 한다.

2. 너무 많은 공간을 차지할 수 있다.
- 내부 구현에 해당되어 포함할 가치가 없는 정보까지 담는 경우, 직렬화 형태가 너무 커져서 디스크에 저장하거나 네트워크 전송 시 속도가 느려진다.

3. 시간이 너무 많이 걸릴 수 있다.
- 직렬화 로직은 객체 그래프의 위상(topology)에 관한 정보가 없기 때문에 그래프를 직접 순회할 수 밖에 없다.

4. 스택오버플로우를 일으킬 수 있다.
- 기본 직렬화 과정은 객체 그래프를 재귀 순회하는데, 자칫 스택오버플로우를 일으킬 수 있다.

## 합리적인 커스텀 직렬화 형태를 갖춘 StringList
```java
public final class StringList implements Serializable {
    private transient int size = 0;
    private transient Entry head = null;

    // 이제는 직렬화되지 않는다.
    private static class Entry {
        String data;
        Entry next;
        Entry previous;
    }

    // 지정한 문자열을 이 리스트에 추가한다.
    public final void add(String s) { ... }

    /**
     * {@code StringList} 인스턴스를 직렬화한다.
     *
     * @serialData 이 리스트의 크기(포함된 문자열의 개수)를 기록한 후
     * ({@code int}), 이어서 모든 원소를(각각은 {@code string}) 순서대로 기록한다.
     */
    private void writeObject(ObjectOutputStream s) throws IOException {
        s.defaultWriteObject();
        s.writeInt(size);

        // 모든 원소를 올바른 순서로 기록한다.
        for (Entry e = head; e != null; e = e.next)
            s.writeObject(e.data);
    }

    private void readObject(ObjectInputStream s) throws IOException, ClassNotFoundException {
        s.defaultReadObject();
        int numElements = s.readInt();

        // 모든 원소를 읽어 이 리스트에 삽입한다.
        for (int i = 0; i < numElements; i++)
            add((String) s.readObject());
    }

    ...
}
```

Stringlist의 필드 모두가 transient더라도 writeObject와 readObject는 각각 가장 먼저 defaultWriteObject와 defaultReadObject를 호출한다.
- 직렬화 명세에서 이 작업을 무조건 하라고 요구하고, 이렇게 해야 향후 릴리즈에서 transient가 아닌 인스턴스 필드가 추가되더라도 상호(상위,하위 모두) 호환되기 때문이다.
- 신버전에서 인스턴스를 직렬화한 후 구버전으로 역직렬화하면 새로 추가된 필드들은 무시될 것이다.
- 구버전 readObject 메서드에서 defaultReadObject를 호출하지 않는다면 역직렬화힐 떼  StreamCorruptedException이 발생할 것이다.

기본 직렬화를 수용하든 하지 않든 defaultWriteObject 메서드를 호출하면 transient로 선언하지 않은 모든 인스턴스 필드가 직렬화된다. 따라서 transient로 선언해도 되는 인스턴스 필드에는 모두 transient 한정자를 붙여야 한다.
- 해당 객체의 논리적 상태와 무관한 필드라고­ 확신할때만 transient 한정자를 생략해야한다.

기본 직렬화를 사용한다면 transient 필드들은 역직렬화될 때 기본값으로­ 초기화됨을 잊지 말아야 한다.
- 객체 참조 필드는 null로, 숫자 기본 타입 필드는 0으로, boolean 필드는 false로 초기화된다.
- 기본값을 그대로 사용해서는­ 안 된다면 readObject 메서드에서 defaultReadObject를 호출한 다음, 해당 필드를 원하는 값으로 복원해야 한다.

### 동기화
기본 직렬화 사용 여부와 상관없이 **객체의 전체 상태를 읽는 매서드에 적용­ 해야 하는 동기화 메커니즘을 직렬화에도 적용해야 한다.**

```java
// 기본 직렬화를 사용하는 동기화된 클래스를 위한 writeObject 매서드
private synchronized void writeObject(ObjectOutputStream s) throws IOException {
    s.defaultWriteObject() ;
}
```

모든 메서드를 synchronized로 선언해서 스레드 안전한 객체에서 기본 직렬화를 사용하기 위해서는 메서드도 writeObject도 위의 코드처럼 synchronized로 선언해줘야 한다.
- writeObject 메서드 안에서 동기화하고 싶다면 클래스의 다른 부분에서 사용하는 락 순서를 똑같이 따라야 한다.
- 그렇지 않으면 교착상태에 빠질 수 있다.

### 직렬 버전 UID

```java
private static final long serialVersionUID = <무작위로 고른 long 값>;
```

어떤 직렬화 형태를 택하든 직렬화 가능 클래스 모두에 직렬 버전 UID를 명시적으로 부여해야 한다.
- 이렇게 하면 직렬 버전 UID가 일으키는 잠재적인 호환성 문제가 사라지고, 성능도 좋아진다.
- 직렬 버전 UID를 명시하지 않으면 런타임에서 이 값을 생성하느라 추가 연산을 수행한다.

기본 버전 클래스와의 호환성을 끊고 싶다면 직렬 버전 UID를 바꿔주면 된다.
- 이렇게 하면 기존 버전의 직렬화된 인스턴스를 역직렬화할 때 InvalidClassException이 던져진다.
- 구버전으로 직렬화된 인스턴스들과의 호환성을 끊으려는 경우를 제외하고는 직렬 버전 UID를 절대 수정하지 말아야 한다.

## 정리
- 클래스를 직렬화하기로 했다면 어떤 직렬화 형태를 사용할지 고민하자.
- 자바의 기본 직렬화 형태는 객체를 직렬화한 결과가 객체의 논리적 표현에 부합할 때만 사용하고, 그렇지 않은 경우 객체를 적절히 설명할 수 있는 커스텀 직렬화 형태를 설계해야 한다.
- 한번 공개된 메서드는 향후 릴리스에서 제거할 수 없듯이, 직렬화 형태에 포함된 필드 또한 마음대로 제거할 수 없다. 즉, 직렬화 호환성을 유지하기 위해 영원히 지원해야 하고, 잘못된 직렬화 형태를 선택하는 경우 영구히 부정적인 영향을 남길 수 있다.
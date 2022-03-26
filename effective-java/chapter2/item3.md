# Item 3. private 생성자나 열거 타입으로 싱글턴임을 보증하라

### 싱글턴(Singleton)?
- 인스턴스를 오직 하나만 생성할 수 있는 클래스
- 그런데 <strong>클래스를 싱글턴으로 만들면 이를 사용하는 클라이언트를 테스트하기 어려워질 수 있다.</strong>
    - 타입을 인터페이스를 정의한 다음 그 인터페이스를 구현해서 만든 싱글턴이 아닌 경우 싱글턴 인스턴스를 mock 구현으로 대체할 수 없다.

## 싱글턴을 만드는 방식
- 싱글턴을 만드는 방식은 세 가지이다.
    - **1. public static final 필드 방식의 싱글턴**
    - **2. 정적 팩터리 방식의 싱글턴**
    - **3. 열거 타입 방식의 싱글턴**
- 앞의 두 방식 모두 생성자는 `private`으로 감춰두고, 유일한 인스턴스에 접근할 수 있는 수단 `public static` 멤버를 마련한다.

### 1. public static final 필드 방식의 싱글턴
```java
public class Elvis {
    public static final Elvis INSTANCE = new Elvis();
    private Elvis() { ... }
    
    public void leaveTheBuilding() { ... }
}
```
- private 생성자는 public static final 필드인 `Elvis.INSTANCE`를  **초기화할 때 딱한 번만 호출**된다.
- public이나 protected 생성자가 없으므로, Elvis 클래스가 초기화될 때 만들어진 인스턴스가 전체 시스템에서 하나뿐임이 보장된다.
- 클라이언트는 따로 접근할 수 있는 방법이 없고, 예외는 단 한 가지, 권한이 있는 클라이언트가 호출할 때이다.
    - 클라이언트가 권한이 있는 경우 리플렉션 API(Item 65)인 AccessibleObject.setAccessible을 사용하여 private 생성자를 호출할 수 있다.
    - 이러한 공격을 방어하기 위해서는 생성자를 수정하여 두 번째 객체가 생성되려할 때 예외를 던지면 된다.

#### 장점
1. 해당 클래스가 싱글턴임이 API에 명백히 드러난다.
2. 간결하다.

### 2. 정적 팩터리 방식의 싱글턴
```java
public class Elvis {
    private static final Elvis INSTANCE = new Elvis();
    private Elvis() { ... }
    public static Elvis getInstance() { return INSTANCE; }
    
    public void leaveTheBuilding() { ... }
}
```
- `Elvis.getInstance`는 항상 같은 객체의 참조를 반환하므로 제 2의 Elvis 인스턴스는 생성되지 않는다.
    - public 필드 방식과 마찬가지로 리플렉션 API를 통한 예외는 똑같이 적용된다.

#### 장점
1. 상황에 따라 API를 바꾸지 않고도 싱글턴이 아닌게 변경할 수 있다.
2. 원한다면 정적 팩터리를 제네릭 싱글턴 팩터리(Item 30)로 만들 수 있다.
    - 제네릭 싱글턴 팩토리 : 제네릭으로 타입 설정 가능한 인스턴스를 만들어두고, 반환 시에 제네릭으로 받은 타입을 이용해 타입을 결정하는 것
    - https://jake-seo-dev.tistory.com/13?category=909023
        ```java
        private static UnaryOperator<Object> IDENTITY_FN = (t) -> t;
        
        @SuppressWarnings("unchecked")
        public static <T> UnaryOperator<T> identityFunction() {
        return (UnaryOperator<T>) IDENTITY_FN;    
        }
        ```
3. 정적 팩터리의 메서드 참조를 공급자(supplier)로 사용할 수 있다.
    - ex. `Elvis::getInstance` → `Supplier<Elvis>`로 사용(Item 43, 44)

<br>

> ### 두 방법을 비교할 때 정적 팩터리 방식의 장점이 굳이 필요하지 않다면 public 필드 방식이 좋다.

<br>

#### 직렬화 이슈
- 둘 중 하나의 방식으로 만든 싱글턴 클래스를 직렬화(12장 참조)하려면 단순히 `Serializable`을 구현한다고 선언하는 것만으로 부족하다.
    - 클래스의 선언에 `implements Serializable`을 추가하는 순간 더 이상 싱글턴이 아니다.
- 모든 인스턴스 필드를 일시적(transient)이라고 선언하고 `readResolve` 메서드를 제공해야 한다.(Item 89)
    - `transient` : Serialize하는 과정에 제외하고 싶은 경우 선언하는 키워드
    - https://nesoy.github.io/articles/2018-06/Java-transient
- 이렇게 하지 않으면 직렬화된 인스턴스를 역직렬화할 때마다 새로운 인스턴스가 만들어진다.
    - 두 번째 코드 예에서라면 가짜 Elvis가 탄생한다는 뜻이다.
    - 가짜 Elvis의 탄생을 예방하고 싶다면 Elvis 클래스에 다음의 readResolve 메서드를 추가해야 한다.
        ```java
        // 싱글턴임을 보장해주는 readResolve 메서드
        private Object readResolve() {
            // 진짜 Elvis를 반환하고, 가짜 Elvis는 가비지 컬렉터에 맡긴다.
            return INSTANCE;
        }
        ```

### 3. 열거 타입 방식의 싱글턴
```java
public enum Elvis {
    INSTANCE;
    
    public void leaveTheBuilding() { ... }
}
```
- public 필드 방식과 비슷하지만 더 간결하고 추가 노력없이 직렬화할 수 있다.
- **대부분의 상황에서는 원소가 하나뿐인 열거 타입이 싱글턴을 만드는 가장 좋은 방법이다.**
- 다만 만들려는 싱글턴이 Enum 외의 클래스를 상속해야 한다면 이 방법은 사용할 수 없다.
    - 열거 타입이 다른 인터페이스를 구현하도록 선언할 수는 있다.

## 정리
### 1. public static 필드 방식
- 해당 클래스가 싱글턴임이 API에 명백히 드러난다.
- 간결하다.
- 정적 팩터리 방식의 장점이 필요없는 경우 public 필드 방식이 좋다.
- 리플랙션 API를 이용한 공격에 대해 예외 처리가 필요하다.
- 직렬화 과정에서 가짜 인스턴스 생성 방지를 위한 추가 작업이 필요하다.

### 2. 정적 팩터리 방식
- 상황에 따라 API를 바꾸지 않고도 싱글턴이 아닌게 변경할 수 있다.
- 원한다면 정적 팩터리를 제네릭 싱글턴 팩터리로 만들 수 있다.
- 정적 팩터리의 메서드 참조를 공급자로 사용할 수 있다.
- 리플랙션 API를 이용한 공격에 대해 예외 처리가 필요하다.
- 직렬화 과정에서 가짜 인스턴스 생성 방지를 위한 추가 작업이 필요하다.

### 3. 열거 타입 방식
- public 필드 방식과 비슷하지만 더 간결하고 추가 노력없이 직렬화할 수 있다.
-  **대부분의 상황에서는 원소가 하나뿐인 열거 타입이 싱글턴을 만드는 가장 좋은 방법이다.**
- 만들려는 싱글턴이 Enum 외의 클래스를 상속해야 한다면 사용할 수 없다.
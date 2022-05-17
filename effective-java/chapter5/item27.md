# Item 27. 비검사 경고를 제거하라

- 비검사 경고는 중요하니 무시하지 말자

- 모든 비검사 경고는 런타임에 예외(ClassCastException)를 일으킬 수 있는 잠재적 가능성을 의미하므로 최선을 다해 제거해라

## 비검사 경고를 제거할 수 있는 경우

- 예를 들어 아래 코드처럼 잘못 작성된 경우가 있다면

    ```java
    Set<Lark> exaltation = new HashSet();
    ```

- 컴파일러는 무엇이 잘못됐는지 설명해준다. (javac 명령줄 인수에 **-Xlint:unchecked** 옵션을 추가)

    ![image](https://user-images.githubusercontent.com/59307414/164968957-4405a668-274f-4095-a225-b0de794ac719.png)

- 이 경우에는 자바 7부터 지원하는 다이아몬드 연산자(<>)만으로 해결이 가능하다. 컴파일러는 올바른 실제 타입 매개변수(Lack)을 추론하게 된다.

    ```java
    Set<Lark> exaltation = new HashSet<>();
    ```

- 이와 같이 할 수 있는 한 모든 비검사 경고를 제거해야 한다.

- 비검사 경고를 제거하면 제거할수록 코드는 타입 안전성이 보장하고, 런타임에 ClassCastException이 발생하는 상황을 방지할 수 있다.
    
## `@SuppressWarnings` : 경고를 제거할 수는 없지만 안전한 경우

- **만약 경고를 제거할 수 없지만 타입이 안전하다고 확신할 수 있다면 `@SuppressWarnings("unchecked")` 애너테이션을 달아 경고를 숨기자.**

- 안전하다고 검증된 비검사 경고를 숨기지 않고 그대로 두면, 진짜 문제를 알리는 새로운 경고가 나와도 경고들에 파묻혀 눈치채지 못할 수 있다.

- `@SuppressWarnings("unchecked")` 애너테이션은 개별 지역변수 선언부터 클래스 전체까지 어떤 선언에도 달 수 있다.

- 하지만 **`@SuppressWarnings("unchecked")` 애너테이션은 항상 가능한 좁은 범위에 적용하자.**

    - 보통은 변수 선언, 아주 짧은 메서드, 생성자에 적용될 것이다.

    - 하지만 자칫 심각한 경고를 놓칠 수 있기 때문에 클래스 전체에 적용해서는 안 된다.

### 지역변수를 추가하여 `@SuppressWarnings`의 범위를 좁히자

- 아래의 ArrayList 내의 toArray 메서드를 예로 들면

    ```java
    public <T> T[] toArray(T[] a) {
        if (a.length < size)
            return (T[]) Arrays.copyOf(elements, size, a.getClass());
        System.arraycopy(elements, 0, a, 0, size);
        if (a.length > size)
            a[size] = null;
        return a;
    }
    ```

- ArrayList를 컴파일하면 toArray 메서드에서 다음의 경고가 발생한다.

    ```
    ArryList.java:305: warning: [unchecked] unchecked cast retrun (T[]) Arrays.copyOf(elements, size, a.getClass());

    required:   T[]
    found:      Object[]
    ```

- 애너테이션은 선언에만 달 수 있기 때문에 return 문에는 `@SuppressWarnings`를 다는 게 불가능하다.

- 메서드 전체에 다는 경우 범위가 필요 이상으로 넓이질 수 있기 때문에 좋은 선택은 아니다.

- 따라서 반환값을 담을 지역변수를 하나 선언하고, 그 변수에 애너테이션을 달아준다.

    ```java
    public <T> T[] toArray(T[] a) {
        if (a.length < size) {
            @SuppressWarnings("unchecked") T[] result = (T[]) Arrays.copyOf(elements, size, a.getClass()); 
            return result;
        }
        System.arraycopy(elements, 0, a, 0, size);
        if (a.length > size)
            a[size] = null;
        return a;
    }
    ```

- 이 경우 코드는 깔끔하게 컴파일되고, 비검사 경고를 숨기는 범위도 최소로 좁혀지게 된다. 

- 추가로 **`@SuppressWarnings("unchecked")` 애너테이션을 사용할 때면 그 경고를 무시해도 안전한 이유를 항상 주석으로 남겨야 한다.**

    - 다른 사람이 그 코드를 이해하는데 도움을 준다.

    - 다른 사람이 그 코드를 잘못 수정하여 타입 안전성을 잃는 상황을 줄여준다.


## 정리
- 비검사 경고는 중요하니 무시하지 말자.

- 모든 비검사 경고는 런타임에 ClassCastException을 일으킬 수 있는 잠재성을 가지고 있기 때문에 최선을 다해 제거하라.

- 경고를 없앨 방법을 찾지 못하겠다면, 해당 코드가 타입 안전함을 증명하고, 가능한 범위를 좁혀`@SuppressWarnings("unchecked")` 애너테이션으로 경고를 숨겨라.

- 이후에 경고를 숨긴 이유를 주석으로 남겨라.
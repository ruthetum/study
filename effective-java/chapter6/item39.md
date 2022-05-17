# Item 39. 명명 패턴보다 애너테이션을 사용하라

## 명명 패턴
- 명명 패턴은 변수나 메서드의 이름을 일관된 방식으로 작성하는 패턴이다.
- JUnit은 버전 3까지 테스트 메서드 이름을 **test**로 시작하게끔 했다.

### 명명 패턴의 단점
1. 오타에 취약하다.
- Junit 3은 실수로 메서드의 이름을 **testSafetyOverride**로 지으면 메서드를 무시하게 된다.
- 테스트를 무시해서 테스트가 실패하지 않게 되기 때문에 통과했다고 오해할 수 있다.

2. 올바른 프로그램 요소에서만 사용되리라 보증할 방법이 없다.
- 예를 들어 메서드가 아닌 클래스 이름을 **TestSafetyMechanisms**로 지으면 개발자 입장에서는 이 클래스에 정의된 테스트 메서드들을 수행하주길 기대하겠지만 JUnit은 클래스 이름에 관심이 없다.
- JUnit은 경고 메시지도 출력하지 않고 개발자가 의도한 테스트도 수행하지 않게 된다.

3. 프로그램 요소를 매개변수로 전달할 마당한 방법이 없다.
- 특정 예외를 던져야 성공하는 테스트가 있을 때 메서드 이름에 포함된 문자열로 예외를 알려주는 방법이 있지만 보기 나쁘고 컴파일러가 문자열이 예외 이름인지 알 도리가 없다.

## 애너테이션
- 명명 패턴에서 발생하는 문제들을 해결할 수 있는 개념이다.
- JUnit 4부터 도입 되었으며, 테스트 메서드 작성 시 명명 패턴 대신 애너테이션 방식 도입됐다.

### 마커 애너테이션
```java
/**
 * 테스트 메서드임을 선언하는 애너테이션이다.
 * 매개변수 없는 정적 메서드 전용이다.
 */
@Retention(RetentionPolicy.RUNTIME)
@Target(ElementType.METHOD)
public @interface Test {
}
```

- 마커 애너테이션은 멤버를 포함하지 않으며 데이터로 구성되지 않는 애너테이션으로 애너테이션 선언을 표시하는 목적으로 활용된다.
    - cf. https://velog.io/@kwj1270/%EC%96%B4%EB%85%B8%ED%85%8C%EC%9D%B4%EC%85%98

- `@Retention`과 `@Target`처럼 애너테이션 선언에 다는 애너테이션을 **메타 애너테이션**이라 한다.

- `@Retention(RetentionPolicy.RUNTIME)`은 `@Test`가 런타임에도 유지되어야 한다는 표시다.
    - `RetentionPolicy.CLASS` : 애노테이션 정보가 바이트 코드까지 남아 있다. (default)
    - `RetentionPolicy.SOURCE` : 컴파일 후에 사라진다.

- `@Target(ElementType.METHOD)`은 `@Test`가 반드시 메소드 선언에서만 사용해야 한다는 표시다.

- 위 코드의 메서드 주석에는 '매개변수 없는 정적 메서드 전용이다'라고 쓰여 있는데, 이 제약을 컴파일러가 강제할 수 있으면 좋겠지만 그렇게 하려면 **적절한 애너테이션 처리기**를 직접 구현해야 한다.
    - cf. [javax.annotation.processing API docs](https://docs.oracle.com/javase/7/docs/api/javax/annotation/processing/package-summary.html)

- 마커 애너테이션은 단순히 Test라는 이름에 오타가 있거나 메서드 선언 외의 프로그램 요소에 달면 컴파일 오류를 내어주는 역할을 할 뿐이다.

### 적용
```java
public class Sample {

    @Test
    public static void m1() { } // 성공해야 한다.

    public static void m2() { } // 실행되지 않는다.

    @Test
    public static void m3() {   // 실패해야 한다.
        throw new RuntimeException("실패");
    }

    public static void m4() { } // 실행되지 않는다.

    @Test
    public void m5() { }        // 잘못 사용한 예 : 정적 메서드가 아니다.

    public static void m6() { } // 실행되지 않는다.

    @Test
    public static void m7() {   // 실패해야 한다.
        throw new RuntimeException("실패");
    } 

    public static void m8() { } // 실행되지 않는다.
}
```

#### 애너테이션 처리기
```java
public class RunTests {
    public static void main(String[] args) throws Exception {
        int tests = 0;
        int passed = 0;
        Class<?> testClass = Class.forName(args[0]);

        for (Method m : testClass.getDeclaredMethods()) {
            if (m.isAnnotationPresent(Test.class)) {                // ⓐ
                tests++;
                try {
                    m.invoke(null);                                 // ⓑ
                    passed++;
                } catch (InvocationTargetException wrappedExc) {    // ⓒ
                    Throwable exc = wrappedExc.getCause();
                    System.out.println(m + " 실패: " + exc);
                } catch (Exception exc) {                           // ⓓ
                System.out.println("잘못 사용한 @Test: " + m);
                }
            }
        }
        System.out.printf("성공: %d, 실패: %d%n", passed, tests - passed);
    }
}
```

- ⓐ : `isAnnotationPresent`를 통해 `@Test` 애너테이션이 적용된 메서드인지 찾는다.

- ⓑ : `invoke`를 통해 `@Test` 메서드를 실행한다.

- ⓒ : 테스트 메서드가 예외를 던지면 리플렉션 메커니즘이 `InvocationTargetException`으로 감싸서 다시 던진다.
    - 이 프로그램은 `InvocationTargetException`을 잡아 원래 예외에 실패 정보를 추출해(getCause) 출력한다.

- ⓓ : `InvocationTargetException`외에 인스턴스 메서드, 매개변수가 있는 메서드 등에 `@Test` 애너테이션을 잘못 사용한 경우를 잡는다.

#### 매개변수가 존재하는 경우

<details>
<summary>더보기</summary>
<div markdown="1">

<br>

#### 매개변수 하나를 받는 애너테이션
- 특정 예외를 던져야만 성공하는 테스트를 지원한다면 새로운 애너테이션 타입이 필요하다.

```java
/**
 * 명시한 예외를 던져야만 성공하는 테스트 메서드용 애너테이션
 */
@Retention(RetentionPolicy.RUNTIME)
@Target(ElementType.METHOD)
public @interface ExceptionTest {
    Class<? extends Throwable> values(); 
}
```

- `Class<? extends Throwable>` : 와일드카드를 통해 Throwable을 상속한 모든 타입을 지정한다.

```java
/**
 * 매개변수 하나짜리 애너테이션을 사용한 프로그램
 */
public class Sample2 {
    // 성공해야 한다.
    @ExceptionTest(ArithmeticException.class)
    public static void m1() {   // 성공해야 한다.
        int i = 0;
        i = i / i;
    }

    // 실패해야 한다. (다른 예외 발생)
    @ExceptionTest(ArithmeticException.class)
    public static void m2() {   // 실패해야 한다. (다른 예외 발생)
        int[] a = new int[0];
        int i = a[1];
    }

    @ExceptionTest(ArithmeticException.class)
    public static void m3() { } // 실패해야 한다. (예외가 발생하지 않음)
}
```

```java
/**
 * 매개변수 하나짜리 애너테이션을 위한 처리기 수정
 */
public class RunTests {
    public static void main(String[] args) throws Exception {
        int tests = 0;
        int passed = 0;
        Class<?> testClass = Class.forName(args[0]);

        for (Method m : testClass.getDeclaredMethods()) {
            if (m.isAnnotationPresent(Test.class)) {                
                tests++;
                try {
                    m.invoke(null);                                 
                } catch (InvocationTargetException wrappedExc) {   
                    Throwable exc = wrappedExc.getCause();

                    Class<? extends Throwable> excType = m.getAnnotation(ExceptionTest.class).value();
                    if (excType.isInstance(exc)) {
                        passed++;
                    } else {
                        System.out.printf("테스트 %s 실패: 기대한 예외 %s, 발생한 예외 %s%n", m, excType.getName, exc);
                    }
                } catch (Exception exc) {
                System.out.println("잘못 사용한 @Test: " + m);
                }
            }
        }
        System.out.printf("성공: %d, 실패: %d%n", passed, tests - passed);
    }
}
```

#### 배열 매개변수를 받는 애너테이션
```java
/**
 * 배열 매개변수를 받는 애너테이션 타입
 */
@Retention(RetentionPolicy.RUNTIME)
@Target(ElementType.METHOD)
public @interface ExceptionTest{
    Class<? extends Throwable>[] values(); 
}
```
```java
/**
 * 배열 매개변수를 받는 애너테이션을 사용하는 코드
 */
@ExceptionTest({ IndexOutOfBoundException.class,
                 NullPointerException.class})
public static void doublyBad() {    // 성공해야 한다.
    List<String> list = new ArrayList<>();
    list.add(5, null);
}
```

#### 반복 가능한 애너테이션 타입
- 자바8 부터는 단일 요소에 애너테이션을 반복적으로 달 수 있는 `@Repeatable` 애너테이션을 제공한다.
- 배열 매개변수 대신 `@Repeatable` 애너테이션을 달면 단일 요소에 반복적으로 적용할 수 있다.
- 주의할 점
    - `@Repeatable`을 단 애너테이션을 반환하는 '컨테이너 애너테이션'을 하나 더 정의하고, `@Repeatable`에 이 컨테이너 애너테이션의 class 객체를 매개변수로 전달해야 한다.
    - 컨테이너 애너테이션은 내부 애너테이션 타입의 배열을 반환하는 value 메서드를 정의해야 한다.
    - 적절한 @Retention과 @Target을 명시해야 한다. 그렇지 않으면 컴파일이 되지 않는다.

```java
/**
 * 반복 가능한 애너테이션 타입
 */

// 반복 가능한 애너테이션
@Retention(RetentionPolicy.RUNTIME)
@Target(ElementType.METHOD)
@Repeatable(ExceptionTestContainer.class)
public @interface ExceptionTest {
    Class<? extends Throwable> value();
}

// 컨테이너 애너테이션
@Retention(RetentionPolicy.RUNTIME)
@Target(ElementType.METHOD)
public @interface ExceptionTestContainer {
    ExceptionTest[] value();
}
```
```java
/**
 * 반복 가능 애너테이션을 두 번 단 코드
 */
@ExceptionTest(IndexOutOfBoundsException.class)
@ExceptionTest(NullPointerException.class)
public static void doublyBad() { ... }
```

<br>

</div>
</details>


## 정리
- 명명 패턴보다는 애너테이션을 사용하자. 애너테이션으로 할 수 있는 일을 명명 패턴으로 처리할 이유는 없다.
- 자바 프로그래머라면 예외 없이 자바가 제공하는 애너테이션 타입들을 사용해야 한다.
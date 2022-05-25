# TDD

<details>
<summary>prologue</summary>
<div markdown="1">

### 용어 정리
- Production Code
```java
public class Calculator {
    int add(int i, int j) {
        return i + j;
    }

    int subtract(int i, int j) {
        return i - j;
    }

    int multiply(int i, int j) {
        return i * j;
    }

    int divide(int i, int j) {
        return i / j;
    }
}
```

- Test Code
```java
public static void main(String[] args) {
    Calculator cal = new Calculator();
    System.out.println(cal.add(3,2));
    System.out.println(cal.subtract(4,1));
    System.out.println(cal.multiply(5,3));
    System.out.println(cal.divide(8,4));
}
```

### TDD란?
- TDD = TDF(Test First Development) + 리팩토링
- TDD랑 단위테스트는 다르다.
- 프로덕션 코드보다 테스트 코드를 먼저 작성. 이후 리팩토링.
- TDD는 테스트 기반으로 개발하지만 테스트 기술이 아님. TDD는 분석 기술이며, 설계 기술이다.

### TDD를 하는 이유
- 디버깅 시간을 감소
- 동작하는 문서 역할

### TDD 사이클
- Test fails > Test passes > Refactor > Test fails > ...
- 실패하는 테스트를 구현
- 테스트가 성공하도록 프로덕션 코드를 구현
- 프로덕션 코드와 테스트 코드를 리팩토링
    - 리팩토링 시에는 프로덕션 코드만 리팩토링하는 게 아니라 테스트 코드도 리팩토링

### TDD 원칙
- 원칙 1 : 실패하는 단위 테스트를 작성할 때까지 프로덕션 코드를 작성하지 않는다.
- 원칙 2 : 컴파일은 실패하지 않으면서 실행이 실패하는 정도로만 단위 테스트르 작성
- 원칙 3 : 현재 실패하는 테스트를 통과할 정도로만 실제 코드를 작성한다.

</div>
</details>

<br>

- [테스트 주도 개발 시작하기(최범균 저)](./%ED%85%8C%EC%8A%A4%ED%8A%B8-%EC%A3%BC%EB%8F%84-%EA%B0%9C%EB%B0%9C-%EC%8B%9C%EC%9E%91%ED%95%98%EA%B8%B0/)
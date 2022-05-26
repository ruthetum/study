# 테스트 주도 개발 시작하기
[테스트 주도 개발 시작하기(최범균 저)](http://www.yes24.com/Product/Goods/89145195) 를 읽고 실습 및 정리한 내용입니다.

## TDD 흐름
> 테스트 -> 코딩 -> 리팩토링 -> 테스트

TDD는 테스트를 먼저 작성하고 테스트를 통과시킬만큼 코드를 작성하고 리팩토링으로 마무리하는 과정을 반복한다.

<details>
<summary>실습</summary>
<div markdown="1">

[chapter 2. 계산기](./src/test/java/chap02/calculator)

[chapter 2. 암호 강도 검사기](./src/test/java/chap02/password)

규칙
- 길이가 8글자 이상
- 0부터 9 사이의 숫자를 포함
- 대문자 포함

강도
- 강함 : 모든 규칙 충족
- 보통 : 2개의 규칙 충족
- 약함 : 그 외의 경우

</div>
</details>

## 테스트 코드 작성 순서
### 작성 순서
> - 쉬운 경우에서 어려운 경우로 진행
> - 예외적인 경우에서 정상인 경우로 진행

### 예외 상황을 먼저 테스트해야 하는 이유
다양한 예외 상황은 복잡한 **if-else** 블록을 동반할 때가 많다. 그렇기 때문에 예외 상황을 전혀 고려하지 않은 코드에 예외 상황을 반영하려면 코드의 구조를 뒤집거나 코드 중간에 예외 상황을 처리하기 위해 조건문을 추가해야 한다.

### 완급 조절
한 번에 얼마만큼의 코드를 작성할 것인가?

> 1. 정해진 값을 리턴
> 2. 값 비교를 이용해서 정해진 값을 리턴
> 3. 다양한 테스트를 추가하면서 구현을 일반화

### 지속적인 리팩토링
테스트를 통과한 뒤에는 리팩토링을 진행한다. 매번 리팩토링을 진행해야 하는 것은 아니지만 규칙성이나 흐름이 보일 경우 리팩토링을 진행한다.

<details>
<summary>실습</summary>
<div markdown="1">

[chapter 3. 서비스 만료 검사](./src/test/java/chap03/expiry)

규칙
- 서비스를 사용하려면 매달 1만원을 선불로 납부, 납부일 기준으로 한 달 뒤가 서비스의 만료일
- 2개월 이상 요금을 납부할 수 있음
- 10만원을 납부하면 서비스를 1년 제공

</div>
</details>

## TDDㆍ기능 명세ㆍ설계
### 기능 명세
설계는 기능 명세로부터 시작한다. 스토리 보드를 포함한 다양한 형태의 요구사항 문서를 이용해서 기능 명세를 구체화한다.

기능 명세를 구체화하는 동안 **입력**과 **결과**를 도출하고 이렇게 도출한 기능 명세를 코드에 반영한다.

기능 명세의 입력과 결과를 코드에 반영하는 과정에서 기능의 이름, 파라미터, 리턴 타입 등이 결정된다.

### 설계
테스트 코드를 작성하기 위해 **클래스 이름**, **메서드 이름**, **메서드 파라미터**, **실행 결과**를 결정하는 과정에서 이름을 고민하고 파라미터 타입과 리턴 타입을 고민하게 된다. 이는 곧 설계 과정이다.

#### 필요한 만큼만 설계하기
TDD는 테스트를 통과할 만큼만 코드를 작성한다. 필요할 것으로 예측해서 미리 코드를 작성하지 않는다.

<details>
<summary>실습</summary>
<div markdown="1">

[chapter 4. 로그인](./src/test/java/chap04/login)

단순히 값을 비교하는 것뿐만 아니라 exception을 결과로 사용할 수도 있다.

</div>
</details>

## Junit 5
### Junit 5 모듈 구성
- Junit 플랫폼 : 테스팅 프레임워크를 구동하기 위한 런처와 테스트 엔진을 위한 API 제공
- Junit 주피터 : Junit 5를 위한 테스트 API와 실행 엔진을 제공
- Junit 빈티지 : Junit 3과 4로 작성된 테스트를 Junit 5 플랫폼에서 실행하기 위한 모듈을 제공

### 주요 단언 메서드
Assertions 클래스는 assertEquals 메서드를 포함해서 아래의 단언 메서드를 제공한다.

|Method|Description|
|---|---|
|assertEquals(expected, actual)|실제 값(actual)이 기대하는 값(expected)과 같은지 검사|
|assertNotEquals(expected, actual)|실제 값(actual)이 기대하는 값(expected)과 같지 않은지 검사|
|assertSame(Object expected, Object actual)|두 객체가 동일한 객체인지 검사|
|assertNotSame(expected, actual)|두 객체가 동일하지 않은 객체인지 검사|
|assertTrue(boolean condition)|값이 true인지 검사|
|assertFalse(boolean condition)|값이 false인지 검사|
|assertNull(Object actual)|값이 null인지 검사|
|assertNotNull(Object actual)|값이 null이 아닌지 검사|
|fail()|테스트를 실패 처리|

#### Assertions가 제공하는 익셉션 발생 유무 검사 메서드
|Method|Description|
|---|---|
|assertThrows(Class<T> expectedType, Executable executable)|executabl을 실행한 결과로 지정한 타입의 익셉션이 발생하는지 검사|
|assertDoesNotThrow(Executable executable)|executabl을 실행한 결과로 익셉션이 발생하지 않는지 검사|

#### assertAll()
```java
public class SampleTest {
    
    @Test
    void 모든_검증_실행후_실패한_것이_있는지_확인() {
        assertAll(
                () -> assertEquals(3, 5/2),
                () -> assertEquals(4, 2*2),
                () -> assertEquals(6, 1+5)
        );
    }
}
```
`assertAll()` 메서드를 통해 모든 검증을 실행하고 그 중에 실패한 것이 있는지 확인할 수 있다.

### 테스트 라이프사이클
#### @BeforeEach, @AfterEach
Junit은 각 테스트 메서드마다 다음 순서대로 코드를 실행한다.

> 1. 테스트 메서드를 포함한 객체 생성
> 2. (@BeforeEach 애노테이션 존재 시) @BeforeEach 애노테이션이 붙은 메서드 실행
> 3. @Test 애노테이션이 붙은 메서드 실행
> 4. (@AfterEach 애노테이션 존재 시) @AfterEach 애노테이션이 붙은 메서드 실행


#### @BeforeAll
- 한 클래스의 모든 테스트 메서드가 실행되기 전에 특정 작업을 수행해야 하는 경우 활용
- 정적 메서드에 붙여서 사용하고, 클래스의 모든 테스트 메서드를 실행하기 전에 한 번 실행

#### @AfterAll
- 클래스의 모든 테스트 메서드를 실행한 뒤에 실행
- 마찬 가지로 정적 메서드에 적용

<details>
<summary>실습</summary>
<div markdown="1">

[chapter 5. 라이프사이클](./src/test/java/chap05/lifecycle)

</div>
</details>

### 테스트 메서드 간 실행 순서 의존과 필드 공유 방지
```java
public class BadTest {
    private FileOperator o = new FileOperator();
    private static File file; // 두 테스트가 데이터를 공유할 목적으로 필드 사용

    @Test
    void fileCreationTest() {
        File createdFile = op.createFile();
        assertTrue(createdFile.length() > 0);
        this.file = createdFile;
    }
    
    @Test
    void readFileTest() {
        long data = op.readData(file);
        assertTrue(data > 0);
    }
}
```
작성한 순서대로 테스트 메서드가 실행될 때는 문제없지만 `readFileTest()`가 먼저 실행되는 경우 file 필드가 null이기 때문에 테스트에 실패하게 된다.

각 테스트 메서드는 서로 독립적으로 동작해야 한다. 한 테스트 메서드의 결과에 따라 다른 테스트 메서드의 실행 결과가 달라지면 안 된다.

그런 의미에서 테스트 메서드가 서로 필드를 공유한다거나 실행 순서를 가정하고 테스트를 작성하지 말아야 한다.

### 추가 애노테이션
#### @DisplayName
테스트에 표시 이름을 붙일 수 있다.

#### @Disabled
특정 테스트를 실행하지 않고 싶을 때 활용한다.

## 테스트 코드 구성
### 구성 요소 : 상황, 실행, 결과 확인
> given, when, then

기능은 상황에 따라 결과가 달라진다. 테스트 코드는 기능을 실행하고 그 결과를 확인하므로 **상황**, **실행**, **결과 확인**의 세 가지 요소로 테스트를 구성할 수 있다.

어떤 상황이 주어지고, 그 상황에서 기능을 실행하고, 실행한 결과를 확인하는 세 가지가 테스트 코드의 기본 골격을 이루게 된다.

## 대역
### 대역의 필요성
테스트를 작성하다 보면 외부 요인이 필요한 시점이 있다

- 테스트 대상에서 파일 시스템을 사용
- 테스트 대상에서 DB로부터 데이터를 조회하거나 데이터를 추가
- 테스트 대상에서 외부의 HTTP 서버와 통신

테스트 대상이 이런 외부 요인에 의존하면 테스트를 작성하고 실행하기 어려워진다.

<details>
<summary>실습</summary>
<div markdown="1">

[chapter 7. 자동 이체 검사기](./src/test/java/chap07/autodebit)

</div>
</details>

### 대역의 종류
|대역 종류|설명|
|---|---|
|스텁(Stub)|구현을 단순한 것으로 대체한다. 테스트에 맞게 단순히 원하는 동작을 수행한다.|
|가짜(Fake)|제품에는 적합하지 않지만, 실제 동작하는 구현을 제공한다.|
|스파이(Spy)|호출된 내역을 기록한다. 기록한 내용은 테스트 결과를 검증할 때 사용한다.|
|모의(Mock)|기대한 대로 상호작용하는지 행위를 검증한다. 기대한 대로 동작하지 않으면 익셉션을 발생할 수 있다.|

- 비밀번호의 강도가 약한지 검사한다 -> Stub 활용
- 동일한 ID를 가진 회원이 존재한 경우 익셉션을 발생하는지 검사한다 -> Fake 활용 (Repository를 메모리, fake로 구현)
- 이메일 발송 여부를 확인한다 -> Spy 활용 (특정 이메일 주소를 사용했는지 확인)


## 테스트 가능한 설계
### 테스트가 어려운 코드
- 하드 코딩된 경로
  
    ```java
    public class PaySync {
        public void sync() throws IOException {
            Path path = Paths.get("D:\\data\\pay\\cp0001.csv");
            ...
        }
    }
    ```

- 의존 객체를 직접 생성
    ```java
    public class PaySync {
        // 의존 대상을 직접 생성
        private PayInfoDao payInfoDao = new PayInfoDao();
        
        public void sync() throws IOException {
            ...
            payInfo.forEach(pi -> payInfoDao.insert(pi));
        }
    }
    ```

- 정적 메소드 사용

- 실행 시점에 따라 달라지는 결과

- 역할이 섞여 있는 코드

- 메서드 중간에 소켓 통신 코드가 포함되어 있는 경우

- 콘솔에서 입력을 받거나 결과를 콘솔에 출력하는 경우

- 테스트 대상이 사용하는 의존 대상 클래스나 메서드가 final인 경우

### 테스트 가능한 설계
- 생성자나 메서드 파라미터로 받기
  
    ```java
    public class PaySync {
  
        private String filePath = "D:\\data\\pay\\cp0001.csv";
  
        public void setFilePath(String filePath) {
            this.filePath = filePath;
        }
        
        public void sync() throws IOException {
            Path path = Paths.get(filePath);
            ...
        }
    }
    ```
  
- 의존 대상을 주입받기

- 테스트하고 싶은 코드를 분리하기

- 시간이나 임의 값 생성 기능 분리하기

- 외부 라이브러리는 직접 사용하지 말고 감싸서 사용하기

## 테스트 범위와 종류
![image](https://user-images.githubusercontent.com/59307414/170519165-f7386af1-17d4-4d17-8275-17406252a98b.png)

# Item 9. try-finally보다는 try-with-resources를 사용하라

- 자바 라이브러리에는 `close` 메서드를 호출해서 직접 닫아줘야 하는 자원이 많다.
    - ex. InputStream, OutputStream, java.sql.Connection ...
- 자원 닫기는 클라이언트가 놓치기 쉬워서 성능 문제로 이어질 수 있다.
- 이에 대한 방지책으로 `finalizer`를 활용하지만 그리 믿을만하지 못하다.
    - cf. [Item 8](https://github.com/Effective-Java-Camp/effective-java-3rd/issues/7)
- 전통적으로 자원의 닫힘을 보장하기 위해 `try-finally` 가 활용되었지만 이제부터는 `try-with-resources` 를 활용하자.

## try-finally는 자원을 회수하는 최고의 방법이 아니다.
### Example 1. 자원이 하나인 경우
```java
static String firstLineOfFile(String path) throws IOException {
    BufferedReader br = new BufferedReader(new FileReader(path));
    try {
        return br.readLine();
    } finally {
        br.close();
    }
}
```
- 만약 자원을 하나 더 사용하는 경우에는 어떻게 될까?

<br>

### Example 2. 자원이 많아지는 경우
```java
static void copy(String src, String dst) throws IOException {
    InputStream in = new FileInputStream(src);
    try {
        OutputStream out = new FileOutputStream(dst);
        try {
            byte[] buf = new byte[BUFFER_SIZE];
            int n;
            while ((n = in.read(buf)) >= 0)
                out.write(buf, 0, n);
        } finally {
            out.close();
        }
    } finally {
        in.close();
    }
}
```

- 자원이 많아지면 try-finally 구문은 매우 지저분해진다.
- 추가로 try-finally 문을 제대로 사용한 위의 두 코드에서도 결점이 존재한다.
- 예외는 try 블록과 finally 블록 모두에서 발생할 수 있다.

    > 예를 들어 기기에 물리적인 문제가 생긴다면
    > 1. `firstLineOfFile` 메서드 안의 `readLine` 메서드가 예외를 던지고
    > 2. 같은 이유로 `close` 메서드도 실패한다.
    >
    > 이런 상황에서는 두 번째 예외가 첫 번째 예외를 삼키기 때문에 맨 처음 발생한 예외를 확인할 수 없다.

- 물론 두 번째 예외 대신 첫 번째 예외를 기록하도록 코드를 작성할 수 있지만 코드가 너무 지저분해질 수 있다.

<br>

## try-with-resoucres : 자원을 회수하는 최선책
- 이러한 문제는 자바 7에서 등장한 try-with-resources 를 통해 해결되었다.
- 이 구조를 사용하려면 해당 자원이 `AutoCloseable` 인터페이스를 구현해야 한다.
    - AutoCloseable : https://velog.io/@sa1341/AutoCloseable-%ED%81%B4%EB%9E%98%EC%8A%A4

### Example 1.
```java
static String firstLineOfFile(String path) throws IOException {
    try (BufferedReader br = new BufferedReader(new FileReader(path))) {
        return br.readLine();
    }
}
```
- `firstLineOfFile` 메서드에서 `readLine`과 `close` 호출 양쪽에서 예외가 발생하면, `close`에서 발생한 예외는 숨겨지고 `readLine`에서 발생한 예외가 기록된다.
- 이처럼 프로그래머에게 보여줄 예외 하나만 보존되고, 여러 개의 다른 예외가 숨겨질 수 있다.
- 숨겨질 예외들도 그냥 버려지지 않고, 스택 추적 내역에 숨겨졌다(supressed)는 꼬리표를 달고 출력된다.
    - `Throwable`의 `getSuppressed` 메서드를 활용하면 프로그램 코드로 가져올 수 있다.

### Example 2.
```java
static void copy(String src, String dst) throws IOException {
    try (InputStream in = new FileInputStream(src);
         OutputStream out = new FileOutputStream(dst)) {
        byte[] buf = new byte[BUFFER_SIZE];
        int n;
        while ((n = in.read(buf)) >= 0)
            out.write(buf, 0, n);
    }
}
```

- 자원이 많아졌을 때에도 try-finally 구문과 비교할 때 try-with-resoucres 버전이 짧고 읽기 수월할 뿐 아니라 문제를 진단하기도 더 좋다.
- try-with-resources 구문 또한 catch 절과 함께 사용할 수 있다.

#### try-with-resources를 catch절과 함께 쓰는 방법
```java
static String firstLineOfFile(String path, String defaultValue) {
    try (BufferedReader br = new BufferedReader(new FileReader(path))) {
        return br.readLine();
    } catch (IOException e) {
        return defaultValue;
    }
}
```

<br>

## 정리
- 꼭 회수해야 하는 자원을 다룰 때는 `try-finally` 말고, 예외 없이 무조건 `try-with-resources`를 사용하자.
- 코드는 더 짧고 분명해지고, 만들어지는 예외 정보도 훨씬 유용하다.
- `try-finally` 로 작성하면 실용적이지 못할 만큼 코드가 지저분해지는 경우라도, `try-with-resources`를 사용하면 정확하고 쉽게 자원을 회수할 수 있다.
# Reactive

## 비동기 programming
### Caller와 Callee
- Caller: 호출하는 함수
- Callee: 호출되는 함수

### 함수형 인터페이스
- 함수형 프로그래밍을 지원하기 위해 java 8부터 도입
- **1개의 추상 메서드**를 갖고 있는 인터페이스
- 함수를 1급 객체로 사용할 수 있음
  - 함수를 변수에 할당하거나 인자로 전달하고 반환값으로 사용 가능
- Function, Consumer, Supplier, Runnable 등이 있음
  - `@FunctionalInterface` 어노테이션을 사용해 함수형 인터페이스임을 명시

    ```java
    @FunctionalInterface
    public interface Function<T, R> {
        R apply(T t);
    }
    ```

- 함수형 인터페이스를 구현한 익명 클래스를 람다식으로 변경 가능
- 함수형 인터페이스는 호출한 스레드에서 실행

## 동기/비동기, Blocking/Non-Blocking
> 동기 vs 비동기 : 호출한 함수의 작업 완료 여부를 확인하느냐
- 동기: 함수 A가 함수 B를 호출한 뒤, 함수 B의 리턴값을 계속 확인
  - caller는 callee의 결과에 관심이 있다
  - caller는 callee의 결과를 이용해서 action을 수행
- 비동기: A가 함수 B를 호출할 때 콜백 함수를 함께 전달, B 작업이 완료될 때 콜백함수가 실행 (A는 B 함수의 작업 완료 여부를 신경쓰지 않음)
  - caller는 callee의 결과에 관심이 없다
  - callee는 결과를 이용해서 callback을 수행

> Blocking vs Non-Blocking : 제어권을 어떻게 처리하느냐
- Blocking: A 함수가 B 함수를 호출하면, 제어권을 A가 호출한 B 함수에게 전달
  - callee를 호출한 후, callee가 완료되기 전까지 caller가 아무것도 할 수 없음
  - 제어권을 callee에게 전달
  - caller와 다른 별도의 thread가 필요하지 않음(혹은 thread를 추가로 사용할 수도 있음)
- Non-Blocking: A 함수가 B 함수를 호출해도 제어권은 그대로 A가 가지고 있음
  - callee를 호출한 후, callee가 완료되지 않더라도 caller가 본인의 일을 할 수 있음
  - 제어권을 caller가 가지고 있음
  - caller와 다른 별도의 thread가 필요함

## Blocking의 종류
- blocking은 thread가 오랜 시간 일을 하거나 대기하는 경우 발생
- CPU-bound blocking: thread가 대부분의 시간을 CPU 점유(= 오랜 시간 일을 한다)
  - 연산이 많은 경우
  - 성능을 높이기 위해서는 추가적인 코어를 투입
- IO-bound blocking: thread가 대부분의 시간을 대기(= 오랜 시간 대기한다)
  - 파일 읽기/쓰기, network 요청 처리, 요청 전달 등
  - IO-bound non blocking을 통해 성능 개선 가능

## CompletableFuture
- java 8에서 처음 도입
- 비동기 프로그래밍 지원
- Lambda, Method reference 등 java 8의 새로운 기능 지원
- 자세한 설명은 [여기](./completable-future)
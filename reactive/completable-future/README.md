# CompletableFuture
java 8에서 처음 도입 

비동기 프로그래밍 지원 

Lambda, Method reference 등 java 8의 새로운 기능 지원
- method reference: `::` 연산자를 이용해서 함수에 대한 참조를 간결하게 표현 

  ```Java
  public static void main(String[] args) {
      Person target = new Person("f");
      Consumer<String> staticPrint = MethodReference::print;
    
      Stream.of("a", "b", "c", "d")
              .map(Person::new)           // constructor reference
              .filter(target::compareTo)  // method reference
              .map(Person::getName)       // instance method reference
              .forEach(staticPrint);      // static method reference
  }
  ```
  - sample code: [MethodReference.java](./completablefuture/src/MethodReference.java)

## ExecutorService
스레드 풀을 이용하여 비동기적으로 작업을 실행하고 관리 

별도의 스레드를 생성하고 관리하지 않아도 되기 때문에 코드를 간결하게 유지 가능

스레드 풀을 이용하여 자원을 효율적으로 관리

```Java
public interface ExecutorService extends Executor {
    void execute(Runnable command);
    <T> Future<T> submit(Callable<T> task);
    void shutdown();
}
```

- **execute**: Runnable 인터페이스를 구현한 작업을 스레드 풀에서 비동기적으로 실행
- **submit**: Callable 인터페이스를 구현한 작업을 스레드 풀에서 비동기적으로 실행, 해당 작업의 결과를 Future<T> 객체로 반환
- **shutdown**: ExecutorService를 종료, 더 이상 task를 받지 않음

### ExecutorService 생성
- **newSingleThreadExecutor**: 단일 스레드로 구성된 스레프 풀을 생성, 한 번에 하나의 작업만 실행
- **newFixedThreadPool**: 고정된 크기의 스레드 풀을 생성, 크기는 인자로 주어진 n과 동일
- **newCachedThreadPool**: 사용 가능한 스레드가 없다면 새로 생성해서 작업을 처리, 사용 가능한 스레드가 있다면 재사용, 스레드가 일정 시간 사용되지 않으면 회수
- **newScheduledThreadPool**: 스케줄링 기능을 갖춘 고정 크기의 스레드 풀을 생성, 주기적이거나 지연이 발생하는 작업을 실행
- **newWorkStealingPool**: work steal 알고리즘을 사용하는 ForkJoinPool을 생성

## CompletableFuture 클래스
```Java
public class CompletableFuture<T> implements Future<T>, CompletionStage<T> {
    // ...
}
```

#### Future
- 비동기적인 작업을 수행
- 해당 작업이 완료되면 결과를 반환하는 인터페이스

#### CompletionStage
- 비동기적인 작업을 수행
- 해당 작업이 완려되면 결과를 처리하거나 다른 CompletionStage를 연결하는 인터페이스

### Future 인터페이스
```Java
public interface Future<V> {
    boolean cancel(boolean mayInterruptIfRunning);
    boolean isCancelled();
    boolean isDone();
    V get() throws InterruptedException, ExecutionException;
    V get(long timeout, TimeUnit unit) throws InterruptedException, ExecutionException, TimeoutException;
}
```

#### Future: isDone(), isCancelled()
future의 상태를 반환
- **isDone**: task가 완료되었다면 원인과 상관없이 true를 반환
- **isCancelled**: task가 명시적으로 취소되었다면 true를 반환

#### Future: get()
결과를 구할 때까지 thread가 계속 block

future에서 무한 루프나 오랜 시간이 걸린다면 thread가 blocking 상태 유지

#### Future: get(long timeout, TimeUnit unit)
결과를 구할 때까지 timeout동안 thread가 계속 block

timeout이 넘어가도 응답이 반환되지 않으면 TimeoutException 발생

#### Future: cancel(boolean mayInterruptIfRunning)
future의 작업 실행을 취소

취소할 수 없는 상황이라면 false를 반환

mayInterruptIfRunning이 false라면 시작하지 않은 작업에 대해서만 취소

#### Future 인터페이스의 한계
`cancel()`을 제외하고 외부에서 future를 컨트롤할 수 없음

반환된 결과를 `get()`해서 접근하기 때문에 비동기 처리가 어려움

완료되거나 에러가 발생했는지 구분하기 어려움 (= 상태 구분이 어려움)

> 이를 해결하기 위해 CompletionStage 활용

### CompletionStage 인터페이스

내부적으로 `newWorkStealingPool`을 사용

```Java
public interface CompletionStage<T> {
    public <U> CompletionStage<U> thenApply(Function<? super T,? extends U> fn);
    public <U> CompletionStage<U> thenApplyAsync(Function<? super T,? extends U> fn);
    
    public CompletionStage<Void> thenAccept(Consumer<? super T> action);
    public CompletionStage<Void> thenAcceptAsync(Consumer<? super T> action);
    
    public CompletionStage<Void> thenRun(Runnable action);
    public CompletionStage<Void> thenRunAsync(Runnable action);
    
    public <U> CompletionStage<U> thenCompose(Function<? super T, ? extends CompletionStage<U>> fn);
    public <U> CompletionStage<U> thenComposeAsync(Function<? super T, ? extends CompletionStage<U>> fn);
    
    public CompletionStage<T> exceptionally(Function<Throwable, ? extends T> fn);
}
```

#### thenAccept[Async]
- Consunmer를 파라미터로 받음
- 이전 task로부터 값을 받지만 값을 넘기지는 않음
- 다음 task에게 null이 전달
- 값을 받아서 action만 수행하는 경우 유용
- sample code: [thenAccept, thenAcceptAsync](./completablefuture/src/CompletionStageThenAccept.java)

> thenAccept vs thenAcceptAsync
> - thenAccept: 현재 스레드에서 실행
> - thenAcceptAsync: 별도의 스레드(forkJoinPool)에서 실행

- done 상태에서 thenAccept는 caller(main)의 스레드에서 실행
  - **done 상태**의 completionStage에 **thenAccept**를 사용하는 경우, **caller의 스레드를 block** 할 수 있음

- done 상태가 아닌 thenAccept는 callee(forkJoinPool)의 스레드에서 실행
  - **done 상태가 아닌**의 completionStage에 **thenAccept**를 사용하는 경우, **callee의 스레드를 block** 할 수 있음

#### then*[Async]의 실행 스레드
<img width="1136" alt="스크린샷" src="https://user-images.githubusercontent.com/59307414/236623240-33bec398-0452-4806-8198-57ba87e32f22.png">

#### then*Async의 스레드풀 변경 방법
- 모든 then*Async 연산자는 executor를 추가 인자로 받음
- 이를 통해서 다른 스레드풀로 task를 실행할 수 있음
- sample code: [CompletionStageThenAcceptAsyncExecutor](./completablefuture/src/CompletionStageThenAcceptAsyncExecutor.java)

#### thenApply[Async]
- Function을 파라미터로 받음
- 이전 task로부터 T 타입의 값을 받아서 가공하고 U 타입의 값을 반환
- 다음 task에게 반환했던 값이 전달
- 값을 변형해서 전달해야 하는 경우 유용

#### thenCompose[Async]
- Function을 파라미터로 받음
- 이전 task로부터 T 타입의 값을 받아서 가공하고 U 타입의 CompletionStage를 반환
- 반환한 CompletionStage가 done 상태가 되면 값을 다음 task에 전달
- 다른 future를 반환해야하는 경우 유용

#### thenRun[Async]
- Runnable을 파라미터로 받음
- 이전 task로부터 값을 받지 않고 값을 반환하지도 않음
- 다음 task에게 null 전달
- future가 완료되었다는 이벤트를 기록할 때 유용

#### exceptionally
- Function을 파라미터로 받음
- 이전 task에서 발생한 exception을 받아서 처리하고 값을 반환
- 다음 task에게 반환된 값을 전달
- future 파이프에서 발생한 에러를 처리할 때 유용

## CompletableFuture 클래스
```Java
public class CompletableFuture<T> implements Future<T>, CompletionStage<T> {
    public static <U> CompletableFuture<U> supplyAsync(Supplier<U> supplier) { ... };
    public static CompletableFuture<Void> runAsync(Runnable runnable) { ... };
    
    public boolean complete(T value) { ... };
    public boolean isCompletedExceptionally() { ... };
    
    public static CompletalbeFuture<Void> allOf(CompletableFuture<?>... cfs) { ... };
    public static CompletableFuture<Object> anyOf(CompletableFuture<?>... cfs) { ... };
}
```

#### supplyAsync
- Supplier를 제공하여 CompletableFuture를 생성 가능 (아무런 값이 없는 상태에서 completableFuture를 생성하는 기능)
- Supplier의 반환값이 CompletableFuture의 결과로 반환

#### runAsync
- Runnable을 제공하여 CompletableFuture를 생성 가능
- 값을 반환하지 않음
- 다음 task에게 null이 전달

#### complete
- CompletableFuture가 완료되지 않았다면 주어진 값으로 채움
- complete에 의해서 상태가 바뀌었다면 true, 아니라면 false를 반환

#### isCompletedExceptionally
- Exception에 의해서 complete 되었는지 확인할 수 있음

#### allOf
- 여러 개의 CompletableFuture를 모아서 하나의 CompletableFuture로 변환
- 모든 CompletableFuture가 완료되면 상태가 done으로 변경
- Void를 반환하므로 각각의 값에 get으로 접근해야 함

#### anyOf
- 여러 개의 CompletableFuture를 모아서 하나의 CompletableFuture로 변환
- 주어진 future 중 하나라도 완료되면 상태가 done으로 변경
- 제일 먼저 done 상태가 되는 future의 값을 반환

### CompletableFuture의 한계
- 지연 로딩 기능을 제공하지 않음
  - CompletableFuture를 반환하는 함수르 호출 시 즉시 작업이 실행
- 지속적으로 생성되는 데이터를 처리하기 어려움
  - CompletableFuture에서 데이터를 반환하고 나면 다시 다른 값을 전달하기 어려움
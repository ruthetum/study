# Goroutine

## Goroutine vs Thread
- 스레드는 프로그램 또는 프로세서 내에서 작업을 처리하기 위한 흐름

- 자바 스레드를 생각해보면 OS의 커널 스레드와 1:1 매핑
  - 결국 자바 스레드는 OS 단의 스케줄링 정책을 따름, 멀티 스레드인 경우 스레드 개수가 CPU 코어보다 많아지는 경우 컨텍스트 스위칭 발생

- 고루틴은 go 프로그램 런타임 시에 생성 및 사용하고, 소거하는 형식, 따라서 상대적으로 가겹게 관리 가능

## 동시성 처리 (with Java)
### Go
- channel
- mutex(lock, waitGroup)
- atomic

### Java
- syncronized(암시적 락): 접근을 제한 (하나의 thread에서만 접근 가능)
- ReentrantLock(명시적 락): 직접 lock, unlock을 걸음 (golang mutex lock)
- atomic
- volatile: 변수를 Main Memory 에만 저장
    - 자원에 대한 read 나 write 작업이, CPU Cache Memory 가 아닌 Main Memory
    - 가시성 문제는 해결되나, 동시접근 해결 불가
- concurrency util 사용
- 불변객체 사용
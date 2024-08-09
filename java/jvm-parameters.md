# JVM 옵션

## Heap 관련 옵션
> InitialRAMPercentage, MinRAMPercentage, MaxRAMPercentage

### InitialRAMPercentage (-XX:InitialRAMPercentage)
> 초기 힙 크기 설정 시 사용

`InitialRAMPercentage` 파라미터는 Java 애플리케이션의 초기 힙 크기를 설정하는 데 사용

인스턴스나 컨테이너의 전체 메모리 백분율을 기준으로 설정
- 컨테이너 메모리가 1000Mi 일 때 -XX:InitialRAMPercentage=50.0 설정
    - 초기 힙 크기: 500Mi
- 값 미설정 시 기본 값: 1.5625

### MinRAMPercentage (-XX:MinRAMPercentage)
> 256MB 미만의 메모리를 갖는 JVM의 최대 힙 크기 설정

`MinRAMPercentage` 파라미터는 JVM 힙 크기의 상한을 결정하는 데 사용

이름만 보면 최소 힙 크기를 설정하는 것처럼 보이지만, Limit 이 맞음 

- 256MB 미만의 메모리에서는 -XX:MinRAMPercentage 사용
- 256MB 이상의 메모리에서는 -XX:MaxRAMPercentage 사용

### MaxRAMPercentage (-XX:MaxRAMPercentage)
> 256MB 이상의 메모리를 갖는 JVM의 최대 힙 크기 설정

`MaxRAMPercentage` 파라미터는 JVM 힙 크기의 상한을 결정하는 데 사용 

- 256MB 미만의 메모리에서는 -XX:MinRAMPercentage 사용
- 256MB 이상의 메모리에서는 -XX:MaxRAMPercentage 사용

### 주의 사항
- Xms JVM 파라미터가 전달되면 InitialRAMPercentage, MinRAMPercentage, MaxRAMPercentage 파라미터는 무시
- JVM은 작은 메모리 서버/컨테이너에서는 MaxRAMPercentage 파라미터를, 큰 메모리 서버/컨테이너에서는 MinRAMPercentage 파라미터를 무시

### 메모리 별 디폴트 값 적용 결과
> JDK 8~17 동일

| 메모리  | Heap memory Limit | Percentage |
|:----:|:-----------------:|:----------:|
| 100MB | 48.38M | 50% |
| 256MB | 121.81M | 50% |
| 500MB | 121.81M | 25% |
| 512MB | 123.75M | 25% |
| 1024MB | 247.50M | 25% |
| 4GB | 1G (JDK 8: 910M) | 25% |

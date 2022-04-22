# Spring AOP(Aspect Oriented Programming)
- AOP : OOP를 보완하는 수단으로, 흩어진 Aspect를 모듈화할 수 있는 프로그래밍 기법

![image](https://user-images.githubusercontent.com/59307414/164704721-65da1a71-b5d7-4ddc-af89-ca285010756a.png)

- 동일한 기능이 흩어져 있으면 유지보수하는데 어려움이 존재

- 각 클래스 내에 흩어진 관심사를 묶어서 모듈화

- 애플리케이션 전체에 걸쳐 사용되는 기능을 재사용

## AOP 주요 용어
- Aspect : 관심사를 모듈화한 것
- Target : 적용이 되는 대상
- Advice : 해야할 일(실제 수행되는 코드)
- Join point : 메서드 실행 시점(Advice를 실제로 실행하고자 하는 위치)
- Pointcut : 대상 내에 어디에 적용이 되어야 하는지에 대한 정보(Join point를 선정하는 방법)
- Weaving : Aspect가 target에 적용되는 전체적인 과정, PointCut으로 지정된 JoinPoint에 Advice가 적용되어 Target을 호출 시 AOP Proxy가 만들어지는 과정

![image](https://img1.daumcdn.net/thumb/R1280x0/?scode=mtistory2&fname=https%3A%2F%2Fblog.kakaocdn.net%2Fdn%2FpfN5M%2FbtqE0lZkKfa%2FF6PvfwluAhiRRAs94EF0v0%2Fimg.png)

## AOP 구현체
- https://ko.wikipedia.org/wiki/%EA%B4%80%EC%A0%90_%EC%A7%80%ED%96%A5_%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D#%EA%B5%AC%ED%98%84
    - Java
        - AspectJ
        - Spring AOP

## AOP 적용 방법
- 컴파일
- 로드 타임
- **런타임**

### Spring AOP가 사용하는 방법 → 런타임 (Dynamic Proxy 기법으로 구현)
> - A라는 Class 타입의 Bean을 생성할 때, A 타입의 Proxy Bean을 생성
> - AOP가 적용된 Target 메서드를 호출 할 때 실제 메서드가 호출되는 것이 아니라 Advice가 요청을 대신 랩핑(Wrraping) 클래스로써 받고 그 랩핑 클래스가 Target을 호출

## 참고
- https://jojoldu.tistory.com/71
- https://sabarada.tistory.com/94?category=803157
- https://www.baeldung.com/spring-aop
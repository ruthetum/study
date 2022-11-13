# Design Patterns
```
- Head First Design Patterns을 읽고 정리/실습한 내용입니다.
- 추가적인 패턴별 설명은 아래 reference를 읽어주세요.
```

## 목차

| Chapter |                       Pattern                       | Description                                                             |            Directory             | Classification |
|:-------:|:---------------------------------------------------:|:------------------------------------------------------------------------|:--------------------------------:|:--------------:|
|    1    |                    전략(Stratergy)                    | 동일 계열의 알고리즘군을 정의하고 캡슐화하여 상호 교환이 가능하게 함                                  |  [ch1-strategy](./ch1-strategy)  |   행동(행위) 패턴    |
|    2    |                    옵저버(Observer)                    | 객체 상태가 변할 때 관련 객체들이 그 변화를 전달받아 자동으로 캐싱                                  |  [ch2-observer](./ch2-observer)  |   행동(행위) 패턴    |
|    3    |                  데코레이터(Decorator)                   | 주어진 상황에 따라 객체에 다른 객체를 덧붙임                                               | [ch3-decorator](./ch3-decorator) |     구조 패턴      |
|    4    | 팩토리 메서드(Factory method)<br>추상 팩토리(Abstract factory) | 객체 생성을 서브클래스로 분리하여 위임(캡슐화)<br>구체적인 클래스를 지정하지 않고 인터페이스를 통해 연관되는 객체들을 보여줌 |   [ch4-factory](./ch4-factory)   |     생성 패턴      |
|    5    |                   싱글턴(Singleton)                    | 한 클래스마다 인스턴스를 하나만 생성하여 어디서든 참조                                          | [ch5-singleton](./ch5-singleton) |     생성 패턴      |
|    6    |                    커맨드(Command)                     | 요청을 객체의 형태로 캡슐화하여 재사용하거나 취소                                                                        |   [ch6-command](./ch6-command)   |   행동(행위) 패턴    | 



## GoF 디자인 패턴
#### 목적에 따른 분류
- 생성 패턴, 구조 패턴, 행동 패턴 총 3가지로 구분됨
- 각 패턴이 어떤 작업을 위해 생성되는 것인지에 따른 구분

|생성 패턴|구조 패턴|행동 패턴|
|---|---|---|
|Abstract Factory<br>Builder<br>Factory Method<br>Prototype<br>Singleton|Adapter<br>Bridge<br>Composite<br>Decorator<br>Facade<br>Flyweight<br>Proxy|Chain of Responsibility<br>Command<br>Interpreter<br>Iterator<br>Mediator<br>Memento<br>Observer<br>State<br>Stratergy<br>Template Method<br>Visitor|

#### 생성 패턴
- 생성 패턴은 객체의 생성과 관련된 패턴
- 특정 객체가 생성되거나 변경되어도 프로그램 구조에 영향을 최소화할 수 있도록 유연성 제공

#### 구조 패턴
- 구조 패턴은 프로그램 내 자료 구조나 인터페이스 구조 등 프로그램 구조를 설계하는데 사용되는 패턴
- 클래스나 객체를 조합하여 더 큰 구조를 만들 수 있게 해줌

#### 행동(행위 패턴)
- 행동 패턴은 반복적으로 사용되는 객체들의 커뮤니케이션을 패턴화
- 객체 사이에 알고리즘 또는 책임을 분배하는 방법에 대해 정의됨
- 결합도를 최소화하는 것이 주 목적

## Reference
- 디자인 패턴별 설명 : https://github.com/ruthetum/study/wiki/Design-Pattern
- https://refactoring.guru/design-patterns
- https://msyu1207.tistory.com/entry/1%EC%9E%A5-%ED%97%A4%EB%93%9C%ED%8D%BC%EC%8A%A4%ED%8A%B8-%EB%94%94%EC%9E%90%EC%9D%B8-%ED%8C%A8%ED%84%B4-%EC%A0%84%EB%9E%B5-%ED%8C%A8%ED%84%B4
    - 소스 코드 내에 담겨진 readme의 설명은 위의 블로그를 참고해서 정리됐습니다.
    - 해당 블로그에 설명 및 참고 이미지가 잘 정리되어 있습니다.
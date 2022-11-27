# Design Patterns
```
- Head First Design Patterns을 읽고 정리/실습한 내용입니다.
- 추가적인 패턴별 설명은 아래 reference를 읽어주세요.
```

## 목차
| Chapter |                       Pattern                       | Description                                                                                    |                     Directory                      |   Classification   |
|:-------:|:---------------------------------------------------:|:-----------------------------------------------------------------------------------------------|:--------------------------------------------------:|:------------------:|
|    1    |                    전략(Stratergy)                    | 동일 계열의 알고리즘군을 정의하고 캡슐화하여 상호 교환이 가능하게 함                                                         |           [ch1-strategy](./ch1-strategy)           |     행동(행위) 패턴      |
|    2    |                    옵저버(Observer)                    | 객체 상태가 변할 때 관련 객체들이 그 변화를 전달받아 자동으로 캐싱                                                         |           [ch2-observer](./ch2-observer)           |     행동(행위) 패턴      |
|    3    |                  데코레이터(Decorator)                   | 주어진 상황에 따라 객체에 다른 객체를 덧붙임                                                                      |          [ch3-decorator](./ch3-decorator)          |       구조 패턴        |
|    4    | 팩토리 메서드(Factory method)<br>추상 팩토리(Abstract factory) | 객체 생성을 서브클래스로 분리하여 위임(캡슐화)<br>구체적인 클래스를 지정하지 않고 인터페이스를 통해 연관되는 객체들을 보여줌                        |            [ch4-factory](./ch4-factory)            |       생성 패턴        |
|    5    |                   싱글턴(Singleton)                    | 한 클래스마다 인스턴스를 하나만 생성하여 어디서든 참조                                                                 |          [ch5-singleton](./ch5-singleton)          |       생성 패턴        |
|    6    |                    커맨드(Command)                     | 요청을 객체의 형태로 캡슐화하여 재사용하거나 취소                                                                    |            [ch6-command](./ch6-command)            |     행동(행위) 패턴      | 
|    7    |             어댑터(Adpater)<br>파사드(Facade)             | 클래스의 인터페이스를 어떤 클래스에서든 이용할 수 있도록 변환<br>서브 시스템에 있는 인터페이스 집합에 대해 통합된 인터페이스 제공                     |     [ch7-adapter-facade](./ch7-adapter-facade)     |       구조 패턴        |
|    8    |              템플릿 메서드(Template method)               | 상위 클래스는 알고리즘의 골격만을 작성하고 구체적인 처리를 서브 클래스로 위임                                                    |    [ch8-template-method](./ch8-template-method)    |     행동(행위) 패턴      |
|    9    | 반복자(Iterator)<br>복합(Composite) | 컬렉션 구현 방법을 노출하지 않으면서 모든 항목에 접근할 수 있는 방법을 제공<br>객체들의 관계를 트리 구조로 표현하는 방식으로 복합 객체와 단일 객체를 구분없이 다룸 | [ch9-iterator-composite](./ch9-iterator-composite) | 행동(행위) 패턴<br>구조 패턴 |
|   10    | 상태(State) | 객체의 상태에 따라 동일한 동작을 다르게 처리                                                                      |            [ch10-state](./ch10-state)              | 행동(행위) 패턴) |
|   11    | 프록시(Proxy) | 실제 기능을 수행하는 객체 대신 가상의 객체를 사용해 로직의 흐름을 제어                                                       | [ch11-proxy](./ch11-proxy) | 구조 패턴 |
|   12    | 복합(Compound) | 두 개 이상의 패턴을 결합하여 일반적으로 자주 등장하는 문제들에 대한 해법을 제공하는 패턴  | [ch12-compound](./ch12-compound) | - |


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
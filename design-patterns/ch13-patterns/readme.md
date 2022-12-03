# 디자인 패턴
- 패턴이란 특정 **컨텍스트** 내에서 주어진 **문제**에 대한 **해결책**
  - **컨텍스트**: 패턴이 적용되는 상황 -> 반복적으로 일어날 수 있는 상황
  - **문제**: 컨텍스트 내에서 이루고자 하는 목적, 컨텍스트 내에서 생길 수 있는 제약조건도 문제에 포함 
  - **해결책**: 찾아내야 하는 것 -> 일련의 제약조건 내에서 목적을 달성할 수 있는 일반적인 디자인

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

## 안티 패턴
- 어떤 문제에 대한 나쁜 해결책에 이르는 길
- 좋은 해결책 같아 보이지만 결국 적용하고 나서 보면 좋지 않은 패턴
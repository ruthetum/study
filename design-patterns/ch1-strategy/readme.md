# Strategy pattern
- 알고리즘(행동)군을 정의하고 각각을 캡슐화하여 교환해서 사용할 수 있도록 만든다.
- 스트래티지를 활용하면 알고리즘을 사용하는 클라이언트와는 독립적으로 알고리즘을 변경할 수 있다.

## 방법
- 바뀌는 부분은 따로 뽑아서 캡슐화한다.
  - 그러면 나중에 바뀌지 않는 부분에는 영향을 미치지 않고 그 부분만 고치거나 확장할 수 있다.
- 상속보다는 구성을 활용한다.
- 구현보다는 인터페이스에 맞춰서 프로그래밍한다.

## Reference
- https://msyu1207.tistory.com/entry/1장-헤드퍼스트-디자인-패턴-전략-패턴
- https://refactoring.guru/design-patterns/strategy
- cache 정책을 strategy 패턴으로 : https://refactoring.guru/design-patterns/strategy/go/example
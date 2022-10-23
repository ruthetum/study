# Observer pattern
- 한 객체의 상태가 바귀면 그 객체에 의존하는 다른 객체들한테 연락이 가고 자동으로 내용이 갱신되는 방식으로 일대다(one to many) 의존성을 정의
- 일대다 관계는 주제와 옵저버에 의해 정의하고, 옵저버는 주제에 의존

## 방법
- 주제(subject)가 옵저버(observer)에 대해서 아는 것은 옵저버가 특정 인터페이스(observer의 interface)를 구현한다는 것
  - 옵저버의 구상 클래스가 무엇인지, 무엇을 하는지에 대해서 알 필요 없음
- 옵저버는 언제든지 새로 추가/삭제할 수 있어야 함
- 새로운 형식의 옵저버를 추가하려고 할 때도 주제를 전혀 변경할 필요 없어야 함
- 주제와 옵저버는 서로 독립적으로 재사용할 수 있어야 함
- 주제나 옵저버가 바뀌더라도 서로한테 영향을 미치지 않아야 함

## 구현 관련
- `before`의 경우 observer 구현체(display)에서 생성 시에 subject를 파라미터로 받고 있음
  - 이로 인해서 observer 구현체에 subject 패키지가 import
  - observer 구현체들을 display 패키지를 별도로 두지 않고, observer 패키지에서 위치시키는 경우 `import cycle not allowed` 발생
- `after`에서는 observer 구현체의 필드에 subject를 두지 않고, 클라이언트(메인)에서 register하는 방법으로 구현
  - 이로 인해서 observer 구현체에 subject 패키지가 import되지 않기 때문에 `import cycle not allowed`를 방지할 수 있음

## Reference
- https://msyu1207.tistory.com/entry/2%EC%9E%A5-%ED%97%A4%EB%93%9C%ED%8D%BC%EC%8A%A4%ED%8A%B8-%EB%94%94%EC%9E%90%EC%9D%B8-%ED%8C%A8%ED%84%B4-%EC%98%B5%EC%A0%80%EB%B2%84-%ED%8C%A8%ED%84%B4?category=1094201
- https://refactoring.guru/design-patterns/observer
- https://refactoring.guru/design-patterns/observer/go/example
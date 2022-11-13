# Template method pattern
- 메소드에서 알고리즘의 골격을 정의, 알고리즘의 여러 단계 중 일부를 서브 클래스에서 구현할 수 있음
- 알고리즘의 구조는 그대로 유지하면서 서브클래스에서 특정 단계를 재정의할 수 있음

## 헐리우드 원칙
- 저수준 구성요소에서 시스템을 접속을 할 수는 있지만, 언제 사용할지는 고수준 구성요소에서 결정
- 의존성 부패를 방지
  - 의존성 부패: 고수준 구성요소가 저수준 구성요소에 의존하고, 그 저수준 구성요소가 다시 고수준 구성요소에 의존하면서 의존성이 복잡하게 꼬여있는 상태

## Structure
![image](https://user-images.githubusercontent.com/22253556/71810981-f9b82f00-30b6-11ea-9c8a-47512a75e21d.png)
![image](https://refactoring.guru/images/patterns/diagrams/template-method/structure-2x.png?id=25082d6d6a76f51c6b64d8aeeaffdbb5)

## Reference
- https://refactoring.guru/ko/design-patterns/template-method
- https://refactoring.guru/ko/design-patterns/template-method/go/example
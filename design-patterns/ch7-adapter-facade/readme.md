# Adapter pattern
- 한 클래스의 인터페이스를 클라이언트에서 사용하고자하는 다른 인터페이스로 변환
- 어댑터를 이용하면 인퍼페이스 호환성 문제를 해결

## 객체 어댑터
- 구성(composition)을 사용

## 클래스 어댑터
- 상속을 사용
- 다중 상속이 이용한 경우 사용

## Structure
### 객체 어댑터
![image](https://t1.daumcdn.net/cfile/tistory/24231F4C575EACA210)
![image](https://refactoring.guru/images/patterns/diagrams/adapter/structure-object-adapter-2x.png)

### 클래스 어댑터
![iamge](https://t1.daumcdn.net/cfile/tistory/252CFB4F575EB62D0A)
![image](https://refactoring.guru/images/patterns/diagrams/adapter/structure-class-adapter-2x.png)

## Reference
- https://refactoring.guru/ko/design-patterns/adapter
- https://refactoring.guru/ko/design-patterns/adapter/go/example

# Facade pattern
- 어떤 서브시스템의 일련의 인터페이스에 통합된 인터페이스를 제공
- 파사드에서 고수준 인터페이스 정의해서 서브시스템을 쉽게 이용 가능

## 최소 지식 원칙
- 객체 사이의 상호작용은 될 수 있으면 아주 가까운 친구 사이에서만 허용하는 것이 좋다
  - 어떤 객체든 그 객체와 상호작용을 하는 클래스의 개수에 주의
  - 그런 객체들과 어떤 식으로 상호작용을 하는지에도 신경을 써야 함
- 여러 클래스가 복잡하게 얽혀있으면 시스템의 한 부분을 변경할 대 다른 부분까지 줄줄이 고쳐야 함

## Structure
![image](https://t1.daumcdn.net/cfile/tistory/2747C84E576006FE3F)
![image](https://refactoring.guru/images/patterns/diagrams/facade/structure-2x.png?id=528ca429555bce293b7c3bd90954e097)

## Reference
- https://refactoring.guru/ko/design-patterns/facade
- https://refactoring.guru/ko/design-patterns/facade/go/example
# Decorator pattern
- 객체에 추가적인 요건을 동적으로 첨가
- 데코레이터는 서브클래스를 만드는 것을 통해서 기능을 유연하게 확장할 수 있는 방법을 제공
- 콘크리트 컴포넌트와 데코레이터를 이용하여 구현

![image](https://blog.kakaocdn.net/dn/bnP6V5/btq1JVdeaQb/0DLKqgOGPfhb2qhfZmKWRk/img.png)


## 방법
1. 비즈니스 도메인이 그 위에 여러 선택적 레이어가 있는 기본 구성 요소(primary component)로 표시될 수 있는지 확인
2. 기본 구성 요소와 선택적 레이어(데코레이터) 모두에 공통적인 방법을 파악
   - 구성 요소 인터페이스(component interface)를 만들고 거기에서 해당 메서드를 선언
3. 구체적인 구성 요소 클래스를 만들고 기본 동작을 정의
4. 기본 데코레이터 클래스를 생성
   - 래핑된 개체에 대한 참조를 저장하기 위한 필드가 있어야 함
   - 필드는 데코레이터 뿐만 아니라 구체적인 구성 요소에 연결할 수 있도록 구성 요소 인터페이스 유형으로 선언되어야 
   - 기본 데코레이터는 모든 작업을 래핑된 개체에 위임
5. 모든 클래스가 구성 요소 인터페이스를 구현하는지 확인
6. 기본 데코레이터에서 확장하여 구체적인 데코레이터를 생성
   - 구체적인 데코레이터는 부모 메서드(항상 래핑된 개체에 위임) 호출 전후에 동작을 실행
7. 클라이언트 코드는 데코레이터를 만들고 클라이언트가 필요로 하는 방식으로 구성

## 구현 관련
```
.
├── component
│      ├── beverage.go     // Component interface
│      ├── dark_roast.go   // Concrete component
│      ├── decaf.go        // Concrete component
│      ├── espresso.go     // Concrete component
│      └── house_blend.go  // Concrete component
│
├── decorator
│      ├── milk.go         // Concrete decorator
│      ├── mocha.go        // Concrete decorator
│      ├── soy.go          // Concrete decorator
│      └── whip.go         // Concrete decorator
│
├── go.mod
└── main.go
```

## Reference
- https://refactoring.guru/ko/design-patterns/decorator
- https://refactoring.guru/ko/design-patterns/decorator/go/example
- https://github.com/alex-leonhardt/go-decorator-pattern
- https://levelup.gitconnected.com/the-decorator-pattern-in-go-66ed951b0f7c
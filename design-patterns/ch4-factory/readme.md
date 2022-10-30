# Factory-method pattern
- 객체 생성을 처리하는 클래스를 팩토리라고 부름
- 객체 생성을 서브클래스에 분리해서 위임(생성을 캡슐화)
- 팩토리 메서드는 템플릿 메서드의 특수화라고 생각
  - 동시에 대규모 템플릿 메서드의 한 단계의 역할을 팩토리 메서드가 할 수 있

![image](https://images.velog.io/images/ljo_0920/post/3f96dcce-7c72-4198-b1c1-13c818846b58/arch2.png)
<img width="954" alt="스크린샷 2022-10-30 오후 7 12 23" src="https://user-images.githubusercontent.com/59307414/198873337-e5d1f069-690f-4f37-9314-ae62e26d5acd.png">

## 구현 관련
```
.
├── pizza                           // Product
│     ├── chicago_chees_pizza.go
│     ├── chicago_veggie_pizza.go
│     ├── ny_cheese_pizza.go
│     ├── ny_veggie_pizza.go
│     └── pizza.go
│
├── pizzastore                      // Creator
│     ├── chicago_pizza_store.go
│     ├── ny_pizza_store.go
│     └── pizza_store.go
│
├── go.mod
└── main.go

```

# Abstract Factory pattern
- 인터페이스를 이용하여 서로 연관되거나 의존하는 객체를 구상 클래스로 지정하지 않고 생성

<img width="930" alt="스크린샷 2022-10-30 오후 7 17 09" src="https://user-images.githubusercontent.com/59307414/198873477-332fb47d-19f8-4a99-9fc4-f28bf3c48138.png">

## 구현 관련
```
.
├── cheese
│      ├── cheese.go
│      ├── mozzarella_cheese.go
│      └── parmesan_cheese.go
├── dough
│      ├── dough.go
│      ├── thick_dough.go
│      └── thin_dough.go
├── sauce
│      ├── marinara_sauce.go
│      ├── plum_tomato_sauce.go
│      └── sauce.go
│
├── pizza
│      ├── cheese_pizza.go
│      ├── pizza.go
│      └── veggie_pizza.go
│
├── pizzafactory
│      ├── chicago_pizza_ingredient_factory.go
│      ├── ny_pizza_ingredient_factory.go
│      └── pizza_ingredient_factory.go
│
├── pizzastore
│      ├── chicago_pizza_store.go
│      ├── ny_pizza_store.go
│      └── pizza_store.go
│
├── go.mod
└── main.go
```

# Factory method vs Abstract Factory
- 팩토리 메서드는 서브 클래스로 직접 생성
- 추상 팩토리는 구상 클래스를 지정하지 않고, 인터페이스로 전달
- 팩토리 비교: https://refactoring.guru/ko/design-patterns/factory-comparison


## Reference
- https://refactoring.guru/ko/design-patterns/factory-method
- https://refactoring.guru/ko/design-patterns/factory-method/go/example
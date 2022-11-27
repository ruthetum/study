# Compound pattern
- 서로 다른 패턴을 섞어 쓰는 방법을 의미
- GoF 디자인 패턴에는 따로 포함되지 않음
- 패턴을 잘 활용하기 위해 서로 다른 패턴을 섞어쓰는 디자인 방법이고, 패턴 몇 개를 결합해서 쓴다고 무조건 컴파운드 패턴이 되는 것은 아님
- 컴파운드 패턴이라 불릴 수 있으려면 여러 가지 문제를 해결하기 위한 용도로 쓰일 수 있는 일반적인 해결책 이어야 함
- 대표적인 예로 MVC, MVVM 패턴 등이 존재
- 책에서는 크게 두 가지에 대해서 언급
  - 오리 예제
  - MVC 패턴

## Duck exmaple
### QuackAble interface
- `QuackAble` interface를 통해 `duck` struct(class)를 구현
```go
package main

import "fmt"

type QuackAble interface {
	Quack()
}

type MallardDuck struct{}

func (d MallardDuck) Quack() {
	fmt.Println("Quack")
}

type RedheadDuck struct{}

func (d RedheadDuck) Quack() {
	fmt.Println("Quack")
}

type DuckCall struct{}

func (d DuckCall) Quack() {
	fmt.Println("Kwak")
}

type RubberDuck struct{}

func (d RubberDuck) Quack() {
	fmt.Println("Squack")
}

func main() {
	mallardDuck := MallardDuck{}
	redheadDuck := RedheadDuck{}
	duckCall := DuckCall{}
	rubberDuck := RubberDuck{}

	mallardDuck.Quack()
	redheadDuck.Quack()
	duckCall.Quack()
	rubberDuck.Quack()
}

[Output]
Quack
Quack
Kwak
Squack
```
- 종류가 다른 Duck 객체 구현 가능

### Adapter
- 거위가 추가되는 경우?
```go
type Goose struct{}

func (g Goose) Honk() {
	fmt.Println("Honk")
}
```
- 거위는 `Quack()`이 아닌 `Honk()`를 통해 소리를 냄
- 이 경우에 **어뎁터**를 이용해서 오리처럼 행동하게 할 수 있음
```go
type GooseAdapter struct {
	goose Goose
}

func (a GooseAdapter) Quack() {
	a.goose.Honk()
}
```

### Decorator
- 오리떼들이 낸 꽥소리(Quack() 호출)의 총 회수를 파악해야 하는 경우?
- 꽥소리를 낸 회수를 카운팅해주는 기능을 추가 -> **데코레이터**로 감싸준다.
```go
...

var numberOfQuack uint64 = 0

type QuackCounter struct {
    duck QuackAble
}

func (c *QuackCounter) Quack() {
    c.duck.Quack()
    atomic.AddUint64(&numberOfQuack, 1)
}

func main() {
    mallardDuck := &QuackCounter{MallardDuck{}}
    redheadDuck := &QuackCounter{RedheadDuck{}}
    duckCall := &QuackCounter{DuckCall{}}
    rubberDuck := &QuackCounter{RubberDuck{}}
    gooseDuck := GooseAdapter{Goose{}}
    
    mallardDuck.Quack()
    redheadDuck.Quack()
    duckCall.Quack()
    rubberDuck.Quack()
    gooseDuck.Quack()
    
    fmt.Println("Quack count:", numberOfQuack)
}

[Output]
Quack
Quack
Kwak
Squack
Honk
Quack count: 4
```
- 데코레이터를 쓸 때에는 객체들을 제대로 포장하지 않으면 원하는 행동을 추가할 수 없음
- 오리 객체를 성하는 작업을 한 군데에서 몰아서 하고, 데코레이터로 감싸는 부분은 빼내서 캡슐화 필요
- 오리를 한 군데에서 생산하기 위한 **팩토리** 필요

### Factory
- 오리를 생성하기 위한 추상 팩토리 생성
- 팩토리를 통해 오리 객체를 생성
```go
type IDuckFactory interface {
	CreateMallardDuck() QuackAble
	CreateRedheadDuck() QuackAble
	CreateDuckCall() QuackAble
	CreateRubberDuck() QuackAble
}

// 일반 duck 팩토리
type DuckFactory struct {}

func (f DuckFactory) CreateMallardDuck() QuackAble {
	return MallardDuck{}
}

func (f DuckFactory) CreateRedheadDuck() QuackAble {
	return RedheadDuck{}
}

func (f DuckFactory) CreateDuckCall() QuackAble {
	return DuckCall{}
}

func (f DuckFactory) CreateRubberDuck() QuackAble {
	return RedheadDuck{}
}

// 카운팅을 위한 duck 팩토리
type CountingDuckFactory struct {}

func (f CountingDuckFactory) CreateMallardDuck() QuackAble {
	return &QuackCounter{MallardDuck{}}
}

func (f CountingDuckFactory) CreateRedheadDuck() QuackAble {
	return &QuackCounter{RedheadDuck{}}
}

func (f CountingDuckFactory) CreateDuckCall() QuackAble {
	return &QuackCounter{DuckCall{}}
}
func (f CountingDuckFactory) CreateRubberDuck() QuackAble {
	return &QuackCounter{RubberDuck{}}
}

func main() {
    countingDuckFactory := CountingDuckFactory{}
    mallardDuck := countingDuckFactory.CreateMallardDuck()
    redheadDuck := countingDuckFactory.CreateRedheadDuck()
    duckCall := countingDuckFactory.CreateDuckCall()
    rubberDuck := countingDuckFactory.CreateRubberDuck()
	
    // ...
}
```

###  Composite
- 생성한 오리떼를 관리하는 기능을 추가해야 하는 경우?
- `QuackAble` 구현체들을 관리하기 위해 **컴포지트 패턴** 활용
  - 객체들로 구성된 컬렉션을 개별 객체와 똑같은 방법으로 다룰 수 있음
```go

type Flock struct {
	quackers []QuackAble
}

func NewFlock() *Flock {
	return &Flock{
		quackers: make([]QuackAble, 0),
	}
}

func (f *Flock) Add(q QuackAble) {
    f.quackers = append(f.quackers, q)
}

func (f *Flock) Quack() {
	// iterator 패턴 적용 가능
	for _, q := range f.quackers {
		q.Quack()
	}
}

func main() {
	
    // ...
	
    flockOfDucks := NewFlock()
    flockOfDucks.Add(redheadDuck)
    flockOfDucks.Add(duckCall)
    flockOfDucks.Add(rubberDuck)
    flockOfDucks.Add(gooseDuck)

    flockOfMallards := NewFlock()
    flockOfMallards.Add(mallardDuckOne)
    flockOfMallards.Add(mallardDuckTwo)
    flockOfMallards.Add(mallardDuckThree)
    flockOfMallards.Add(mallardDuckFour)

    flockOfDucks.Add(flockOfMallards)

    flockOfDucks.Quack()
    fmt.Println("Quack count:", numberOfQuack) // Output: 7

    flockOfMallards.Quack()
    fmt.Println("Quack count:", numberOfQuack) // Output: 11
}
```
- 컴포지트 패턴을 통해 오리떼(객체들의 집합)를 관리할 수 있음
  - 내부에서 **이터레이터 패턴**을 적용해서 컬렉션에 접근할 수도 있음

### Observer
- 개별 오리의 행동이나 상태를 확인하고 싶은 경우?
- **옵저버 패턴**을 이용해서 개별 객체의 행동을 확인
```go
type QuackObservable interface {
    GetName() string
    RegisterObserver(observer Observer)
    NotifyAll()
}

type Observable struct {
    Observers []Observer
    Duck      QuackObservable
}

func NewObservable(d QuackObservable) *Observable {
  return &Observable{
    Observers: make([]Observer, 0),
    Duck:      d,
  }
}

func (o *Observable) GetName() string {
    return o.Duck.GetName()
}

func (o *Observable) RegisterObserver(observer Observer) {
    o.Observers = append(o.Observers, observer)
}

func (o *Observable) NotifyAll() {
  for _, observer := range o.Observers {
    observer.Update(o.Duck)
  }
}

type Observer interface {
	Update(duck QuackObservable)
}

type Quackologist struct{}

func (l *Quackologist) Update(duck QuackObservable) {
	fmt.Println("Quackologist:", duck.GetName(), "just quacked")
}

func main() {
  // 오리 팩토리 생성
  countingDuckFactory := factory.CountingDuckFactory{}
  
  // 오리떼 생성
  mallardDuckOne := countingDuckFactory.CreateMallardDuck()
  mallardDuckTwo := countingDuckFactory.CreateMallardDuck()
  mallardDuckThree := countingDuckFactory.CreateMallardDuck()
  mallardDuckFour := countingDuckFactory.CreateMallardDuck()
  redheadDuck := countingDuckFactory.CreateRedheadDuck()
  duckCall := countingDuckFactory.CreateDuckCall()
  rubberDuck := countingDuckFactory.CreateRubberDuck()
  // 어뎁터 적용
  gooseDuck := adpater.GooseAdapter{Goose: goose.Goose{}}
  
  // 오리떼 집합 관리
  flockOfDucks := composite.NewFlock()
  flockOfDucks.Add(redheadDuck)
  flockOfDucks.Add(duckCall)
  flockOfDucks.Add(rubberDuck)
  flockOfDucks.Add(gooseDuck)
  
  flockOfMallards := composite.NewFlock()
  flockOfMallards.Add(mallardDuckOne)
  flockOfMallards.Add(mallardDuckTwo)
  flockOfMallards.Add(mallardDuckThree)
  flockOfMallards.Add(mallardDuckFour)
  
  flockOfDucks.Add(flockOfMallards)
  
  // 옵저버
  quackologist := &observer.Quackologist{}
  flockOfDucks.Register(quackologist)
  
  flockOfDucks.Quack()
  fmt.Println("Quack count:", decorator.NumberOfQuack)
  
  flockOfMallards.Quack()
  fmt.Println("Quack count:", decorator.NumberOfQuack)
}


[Output]
Quack
Quackologist: RedheadDuck just quacked
Kwak
Quackologist: DuckCall just quacked
Quack
Quackologist: RedheadDuck just quacked
Honk
Quack
Quackologist: MallardDuck just quacked
Quack
Quackologist: MallardDuck just quacked
Quack
Quackologist: MallardDuck just quacked
Quack
Quackologist: MallardDuck just quacked
Quack count: 7
Quack
Quackologist: MallardDuck just quacked
Quack
Quackologist: MallardDuck just quacked
Quack
Quackologist: MallardDuck just quacked
Quack
Quackologist: MallardDuck just quacked
Quack count: 11
```

## MVC Pattern

![image](https://sgc109.github.io/images/compound-pattern-feat-mvc/mvc-diagram.png)

- 모델 - 뷰 - 컨트롤러
- 예를 들어 MP3
  - 모델: MP3 내의 데이터, 애플리케이션 로직 등을 포함, 뷰에게 데이터의 상태가 변경되었음을 전달(상태 변화 통지)
  - 뷰: 디스플레이스 갱신
  - 컨트롤러: 사용자가 인터페이스를 조작할 때 해당 해동이 컨트롤러에 전달, 해당 행동에 의해 모델을 조작

### Model
- **옵저버 패턴** 활용
- Model의 상태가 변경 되었을 때 Controller 또는 View 에게 이 사실을 알리는데 사용

### View
- **컴포지트** 패턴 활용
- View를 구성하는 컴포넌트들은 계층 구조를 이룸 (ex. View, DOM, etc)

### Controller
- **스트래티지** 패턴 활용
- View 객체를 여러 스트래티지를 활용해서 사용
- View에서는 Model에서 어떤 내역이 처리되는지 알지 못함(View, Model 분리)


## JSP Model 2
![image](https://upload.wikimedia.org/wikipedia/commons/thumb/7/72/JSP_Model_2.svg/440px-JSP_Model_2.svg.png)
- JSP Model 2: MVC 패턴을 웹 애플리케이션에 맞는 형태로 적용
- Spring 과 같은 Web Framework 에서 사용하는 MVC 패턴은 JSP 를 사용하지 않는다고 하더라도 사실상 전통적인 MVC 패턴 보다는 이 JSP Model 2 에 해당
  - Spring MVC 에서 사용하는 패턴은 Servlet(DispatcherServlet)에서 HTTP 요청을 처리하는 것을 제외한 Controller 로직을 분리한 구조로, Front Controller 패턴
  - cf. https://docs.spring.io/spring-framework/docs/current/reference/html/web.html#mvc-servlet

![iamge](https://img1.daumcdn.net/thumb/R1280x0/?scode=mtistory2&fname=https%3A%2F%2Fblog.kakaocdn.net%2Fdn%2FlRW7F%2FbtrpyiCko1D%2FCXaGXmd3hnQImpr2GMWkL0%2Fimg.png)
![image](https://blog.kakaocdn.net/dn/Or4T1/btqFcNAEiAD/VLPsPQcUnUC8iWw8suH3Ek/img.png)

## Appendix
- MVC, MVP, MVVM: https://velog.io/@almondgood/15.-%EB%B3%B5%ED%95%A9-%ED%8C%A8%ED%84%B4
- 리액테에서 합성 컴포턴트를 만들 때 컴파운드 패턴이 활용
  - https://im-developer.tistory.com/238
  - https://www.patterns.dev/posts/compound-pattern/
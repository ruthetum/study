# Mediator pattern
- 한 집합에 속해있는 객체들의 상호작용을 캡슐화하여 새로운 객체로 정의

## Structure
![image](https://refactoring.guru/images/patterns/diagrams/mediator/structure-2x.png)

## Example
### Before
```go
type Alarm struct {}

func (a Alarm) OnEvent() {
	CheckCalendar()
	CheckSprinkler()
	StartCoffee()
	...
}

type CoffeePot struct {}

func (p CoffeePot) OnEvent() {
	CheckCalendar()
	CheckAlarm()
	...
}
...
```
### After
```go
type Mediator struct {
	alarm     Alarm
	pot       CoffeePot
	calendar  Calendar
	sprinkler Sprinkler
}

func (m Mediator) Notify(name string) {
    if name == "alarm" {
        m.Alarm()
    }
	...
}

func (m Mediator) Alarm() {
	m.calendar.CheckCalendar()
	m.sprinkler.CheckSprinkler()
	m.pot.StartCoffee()
}
```

- mediator에 각 컴포넌트들 간의 비즈니스 로직을 때려박는 패턴
- 컴포넌트들의 복잡성은 낮출 수 있지만 미디에이터가 복잡해지고, 재사용성이 떨어짐


## Reference
- https://refactoring.guru/ko/design-patterns/mediator
- https://ganghee-lee.tistory.com/8
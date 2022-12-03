# Bridge pattern
- 구현부에서 추상층을 분리하여 각자 독립적으로 변형하고 확장

## Structure
![image](https://refactoring.guru/images/patterns/diagrams/bridge/structure-ko-2x.png?id=e14227ad22486d0981c23b8709556383)

## Example
- 만능 리모컨 만들기
  - 리모컨과 TV가 존재
  - 리모컨 종류가 늘어나고, TV 종류가 늘어나면 조합 별로 구현체를 만들어야 하는가? -> 너무 코드가 많아짐
  - 추상화 계층(리모컨)과 구현 계층(TV)을 만든 후, 클래스 계층 간 연결(브릿지) 구성

```go
type RCARemoteControl struct {
	tv tv.TV
}

func NewRCARemoteControl(tv tv.TV) *RCARemoteControl {
	return &RCARemoteControl{tv: tv}
}

func (c *RCARemoteControl) On() {
	fmt.Println("RCA RemoteControl on")
	c.tv.On()
}

func (c *RCARemoteControl) Off() {
	fmt.Println("RCA RemoteControl off")
	c.tv.Off()
}

func (c *RCARemoteControl) SetChannel() {
	c.tv.TuneChannel()
}

type Sony struct {
}

func NewSony() Sony {
    return Sony{}
}

func (tv Sony) On() {
    fmt.Println("Sony TV on")
}

func (tv Sony) Off() {
    fmt.Println("Sony TV off")
}

func (tv Sony) TuneChannel() {
    fmt.Println("Sony TV tune channel")
}

func main() {
    sonyTV := tv.NewSony()

    rcaRemoteController = controller.NewRCARemoteControl(sonyTV)
    rcaRemoteController.SetChannel()
	
    ...
}
```

## Reference
- https://refactoring.guru/ko/design-patterns/bridge
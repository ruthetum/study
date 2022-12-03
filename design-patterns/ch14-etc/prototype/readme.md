# Prototype pattern
- 원본 객체를 복사하여 객체를 생성(클론)
- 클라이언트 코드에서 어떤 클래스의 인스턴스를 만드는 것인지 모르는 상태에서 새로운 인스턴스를 생성할 수 있음

## Structure
### 기초 구현
![image](https://refactoring.guru/images/patterns/diagrams/prototype/structure-2x.png?id=ba75079f42f08028ae4cdbda0cfecc26)

### 레지스트리 구현
![image](https://refactoring.guru/images/patterns/diagrams/prototype/structure-prototype-cache-2x.png?id=a1e4514bbcc9b10968b856f19b407105)

## Example
```go
type Manager struct {
	registry map[string]Monster
}

func NewManager() *Manager {
	return &Manager{registry: make(map[string]Monster)}
}

func (m *Manager) Register(name string, monster Monster) {
	m.registry[name] = monster
}

func (m *Manager) Create(prototype string) Monster {
	return m.registry[prototype].Clone()
}

type WellKnownMonster struct {
    name string
}

func NewWellKnownMonster(postfix string) WellKnownMonster {
    return WellKnownMonster{name: "wkm" + postfix}
}

func (m WellKnownMonster) Print() {
    fmt.Println(m.name)
}

func (m WellKnownMonster) Clone() Monster {
    return NewWellKnownMonster("_clone")
}

func main() {
    manager := monster.NewManager()
    
    wellKnownMonster := monster.NewWellKnownMonster("")
    dynamicPlayerMonster := monster.NewDynamicPlayerMonster("")
    
    manager.Register("wkm", wellKnownMonster)
    manager.Register("dpm", dynamicPlayerMonster)
    
    wellKnownMonster.Print() // wkm
    wellKnownMonsterClone := manager.Create("wkm")
    wellKnownMonsterClone.Print() // wkm_clone
}

```

## Reference
- https://refactoring.guru/ko/design-patterns/prototype
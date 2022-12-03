# Memento pattern
- 객체가 특정 상태로 다시 되돌아올 수 있도록 내부 상태를 실제화
    - 시스템에서 핵심적인 기능을 담당하는 객체의 중요한 상태를 보관
    - 핵심적인 객체의 캡슐화 유지
- 스냅샷 생성

## Structure
![image](https://refactoring.guru/images/patterns/diagrams/memento/structure1-2x.png)

## Example
```go
type Game struct {
	state string
}

func NewGame() *Game {
	return &Game{state: "A"}
}

func (g *Game) CreateSnapshot() Snapshot {
	return NewSnapshot(g.state)
}

func (g *Game) RestoreSnapshot(s Snapshot) {
	g.state = s.GetSavedState()
}


type Snapshot struct {
    state string
}

func NewSnapshot(state string) Snapshot {
    return Snapshot{state: state}
}

func (s *Snapshot) GetSavedState() string {
    return s.state
}

type SnapshotManager struct {
    snapshots []Snapshot
}

func NewSnapshotManager() *SnapshotManager {
    return &SnapshotManager{snapshots: make([]Snapshot, 0)}
}

func (sm *SnapshotManager) AddSnapshot(s Snapshot) {
    sm.snapshots = append(sm.snapshots, s)
}

func (sm *SnapshotManager) GetSnapshot(index int) Snapshot {
    return sm.snapshots[index]
}

func main() {
    snapshotManager := game.NewSnapshotManager()
    game := game.NewGame()
    
    fmt.Println(game.GetState())
    snapshotManager.AddSnapshot(game.CreateSnapshot())
    
    game.SetState("B")
    fmt.Println(game.GetState())
    snapshotManager.AddSnapshot(game.CreateSnapshot())
    
    game.SetState("C")
    fmt.Println(game.GetState())
    snapshotManager.AddSnapshot(game.CreateSnapshot())
    
    game.RestoreSnapshot(snapshotManager.GetSnapshot(1))
    fmt.Println(game.GetState())
}
```


## Reference
- https://refactoring.guru/ko/design-patterns/memento

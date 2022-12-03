package game

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

func (g *Game) GetState() string {
	return g.state
}

func (g *Game) SetState(state string) {
	g.state = state
}

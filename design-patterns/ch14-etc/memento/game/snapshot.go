package game

type Snapshot struct {
	state string
}

func NewSnapshot(state string) Snapshot {
	return Snapshot{state: state}
}

func (s *Snapshot) GetSavedState() string {
	return s.state
}

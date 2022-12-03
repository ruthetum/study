package monster

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

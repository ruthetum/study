package game

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

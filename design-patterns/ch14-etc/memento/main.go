package main

import (
	"fmt"
	"memento/game"
)

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

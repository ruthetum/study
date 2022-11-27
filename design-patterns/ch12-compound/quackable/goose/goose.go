package goose

import "fmt"

type Goose struct{}

func (g Goose) Honk() {
	fmt.Println("Honk")
}

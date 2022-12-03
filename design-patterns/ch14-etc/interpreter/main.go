package main

import (
	"fmt"
	"interpreter/logic"
)

func main() {
	logic.Values.Assign("A", true)  // A = true
	logic.Values.Assign("B", false) // B = false

	// A && B
	term1 := logic.NewAndLogic(
		logic.NewVariable("A"),
		logic.NewVariable("B"),
	).Evaluate()
	fmt.Println("A && B:", term1) // A && B: false

	// A || B
	term2 := logic.NewOrLogic(
		logic.NewVariable("A"),
		logic.NewVariable("B"),
	).Evaluate()
	fmt.Println("A || B:", term2) // A || B: true
}

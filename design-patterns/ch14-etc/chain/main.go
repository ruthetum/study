package main

import "chain/handler"

func main() {

	spamHandler := handler.NewSpamHandler()
	fanHandler := handler.NewFanHandler()
	complainHandler := handler.NewComplainHandler()
	locHandler := handler.NewLocHandler()

	spamHandler.SetNext(fanHandler)
	fanHandler.SetNext(complainHandler)
	complainHandler.SetNext(locHandler)

	test := []string{"spam", "loc", "fan", "complain"}
	for _, t := range test {
		spamHandler.Execute(t)
	}
}

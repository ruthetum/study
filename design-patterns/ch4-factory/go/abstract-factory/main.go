package main

import "abstract-factory/pizzastore"

func main() {
	nyStore := pizzastore.NewNYPizzaStore()
	chicagoStore := pizzastore.NewChicagoPizzaStore()

	nyPizza, _ := nyStore.OrderPizza("cheese")
	nyPizza.Info()

	chicagoPizza, _ := chicagoStore.OrderPizza("veggie")
	chicagoPizza.Info()
}

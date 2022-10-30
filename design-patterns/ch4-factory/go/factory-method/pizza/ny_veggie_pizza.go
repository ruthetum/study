package pizza

type NYVeggiePizza struct {
	Pizza
}

func CreateNYVeggiePizza() IPizza {
	toppings := []string{
		"잘게 썬 레지아노 치즈", "마늘", "양파", "버섯", "고추",
	}
	return &NYVeggiePizza{
		Pizza: Pizza{
			name:     "뉴욕 야채 피자",
			dough:    "씬 크러스트 도우",
			sauce:    "마리나라 소스",
			toppings: toppings,
		},
	}
}

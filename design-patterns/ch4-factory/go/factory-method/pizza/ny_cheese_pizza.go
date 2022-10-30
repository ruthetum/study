package pizza

type NYCheesePizza struct {
	Pizza
}

func CreateNYCheesePizza() IPizza {
	toppings := []string{
		"잘게 썬 레지아노 치즈",
	}
	return &NYCheesePizza{
		Pizza: Pizza{
			name:     "뉴욕 치즈 피자",
			dough:    "씬 크러스트 도우",
			sauce:    "마리나라 소스",
			toppings: toppings,
		},
	}
}

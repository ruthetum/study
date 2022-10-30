package pizza

import "fmt"

type ChicagoVeggiePizza struct {
	Pizza
}

func CreateChicagoVeggiePizza() IPizza {
	toppings := []string{
		"모짜렐라 치즈", "마늘", "양파", "버섯", "고추",
	}
	return &ChicagoVeggiePizza{
		Pizza: Pizza{
			name:     "시카고 치즈 피자",
			dough:    "두꺼운 도우",
			sauce:    "플럼 토마토 소스",
			toppings: toppings,
		},
	}
}

func (p *ChicagoVeggiePizza) Cut() {
	fmt.Println("네모난 모양으로 자르는 중")
}

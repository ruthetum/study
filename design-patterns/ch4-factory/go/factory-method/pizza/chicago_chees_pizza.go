package pizza

import "fmt"

type ChicagoCheesePizza struct {
	Pizza
}

func CreateChicagoCheesePizza() IPizza {
	toppings := []string{
		"모짜렐라 치즈", "마늘", "양파", "버섯", "고추",
	}
	return &ChicagoCheesePizza{
		Pizza: Pizza{
			name:     "시키고 치즈 피자",
			dough:    "두꺼운 도우",
			sauce:    "플럼 토마토 소스",
			toppings: toppings,
		},
	}
}

func (p *ChicagoCheesePizza) Cut() {
	fmt.Println("네모난 모양으로 자르는 중")
}

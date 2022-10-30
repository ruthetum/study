package pizza

import "fmt"

type IPizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
	GetName() string
	SetName(name string)
	Info()
}

type Pizza struct {
	name   string
	dough  string
	sauce  string
	cheese string
}

func (p *Pizza) SetName(name string) {
	p.name = name
}

func (p *Pizza) Prepare() {
	fmt.Println("준비 중", p.name)
	fmt.Println("도우 전달...")
	fmt.Println("소스 추가...")
	fmt.Println("치즈 추가...")
}

func (p *Pizza) Bake() {
	fmt.Println("굽는 중")
}

func (p *Pizza) Cut() {
	fmt.Println("자르는 중")
}

func (p *Pizza) Box() {
	fmt.Println("포장 중")
}

func (p *Pizza) GetName() string {
	return p.name
}

func (p *Pizza) Info() {
	fmt.Println("이름:", p.name)
	fmt.Println("도우:", p.dough)
	fmt.Println("소스:", p.sauce)
	fmt.Println("치즈:", p.cheese)
	fmt.Println()
}

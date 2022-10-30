package main

import (
	"decorator/component"
	"decorator/decorator"
	"fmt"
)

func main() {
	espresso := component.CreateEspresso()
	fmt.Println(fmt.Sprintf("%v: $%v", espresso.GetDescription(), espresso.Cost()))
	// output: 에스프레소[TALL]: $1.99

	darkRoast := component.CreateDarkRoast()
	darkRoast.SetSize(component.VentiSize)
	darkRoastWithMocha := decorator.AddMocha(darkRoast)
	darkRoastWithMochaTwice := decorator.AddMocha(darkRoastWithMocha)
	fmt.Println(fmt.Sprintf("%v: $%v", darkRoastWithMochaTwice.GetDescription(), darkRoastWithMochaTwice.Cost()))
	// output: 다크 로스트 커피[VENTI], 모카, 모카: $1.59

	houseBlend := component.CreateHouseBlend()
	houseBlend.SetSize(component.GrandeSize)
	hbWithMilk := decorator.AddMilk(houseBlend)
	hbWithMilkSoy := decorator.AddSoy(hbWithMilk)
	hbWithMilkSoyWhip := decorator.AddWhip(hbWithMilkSoy)
	fmt.Println(fmt.Sprintf("%v: $%v", hbWithMilkSoyWhip.GetDescription(), hbWithMilkSoyWhip.Cost()))
	// output: 하우스 블렌드 커피[GRANDE], 우유, 두유, 휘핑: $1.39
}

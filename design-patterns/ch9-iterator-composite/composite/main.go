package main

import "composite/menu"

func main() {
	pancakeHouse := menu.NewMenu("팬케이크 하우스 메뉴", "아침 메뉴")
	diner := menu.NewMenu("객체마을 식당 메뉴", "점심 메뉴")
	cafe := menu.NewMenu("카페 메뉴", "저녁 메뉴")
	all := menu.NewMenu("전체 메뉴", "전체 메뉴")

	all.Add(pancakeHouse)
	all.Add(diner)
	all.Add(cafe)

	pancake := menu.NewItem("K&B 팬케이크 세트", "스크램블 에그와 토스트가 곁들여진 팬케이크", 2.99, true)
	blt := menu.NewItem("BLT", "통밀 위에 베이컨, 상추, 토마토를 얹은 메뉴", 2.99, false)
	waffle := menu.NewItem("와플", "와플, 취향에 따라 블루베리나 딸기를 얹을 수 있습니다.", 3.59, true)

	pancakeHouse.Add(pancake)
	diner.Add(blt)
	cafe.Add(waffle)

	waitress := menu.NewWaitress(all)
	waitress.Print()
}

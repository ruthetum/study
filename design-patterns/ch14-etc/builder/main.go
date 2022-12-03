package main

import (
	"builder/fcm"
	"fmt"
)

const (
	token = "1q2w3e4r"
	aos   = "aos"
	ios   = "ios"
)

func main() {

	iosBuilder := fcm.NewIOSBuilder()
	iosBuilder.SetToken(token)
	// 하위 데이터도 빌더 패턴 적용 가능
	iosBuilder.SetIOS(fcm.IOSData{
		APN:      "APN",
		Content:  "hello",
		Redirect: "design/pattern/14",
	})
	iosMessage := iosBuilder.Build()
	fmt.Printf("IOS Message: %v", iosMessage)

	builder := fcm.NewBuilder()
	// 체이닝 메서드도 적용 가능
	aosMessage := builder.
		SetPlatform(aos).
		SetToken(token).
		SetAOS(fcm.AOSData{
			Data: "hello",
			Path: "design/pattern/14",
		}).
		Build()
	fmt.Printf("AOS Message: %v", aosMessage)
}

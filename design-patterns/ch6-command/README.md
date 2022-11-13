# 커맨드 패턴 (Command Pattern)

## Prologue
![image](https://oopy.lazyrockets.com/api/v2/notion/image?src=https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2F6ce7d4f6-cf81-4e0e-a534-e9088ca1104e%2FScreen_Shot_2021-01-12_at_16.03.18.png&blockId=8ed9d03a-f3eb-4297-9a77-1139afad9d87)
- 프로그래밍이 가능한 리모컨에 7개의 슬롯이 있고, 각 슬롯에 필요한 프로그램을 연결
- 각 프로그램을 통해 특정 기능을 on/off 할 수 있음

### Issue
![image](https://blog.kakaocdn.net/dn/bm3r8n/btrG25TCBiY/KaiiPX312oGUzr8osgpkH1/img.png)
- 프로그램 별로 on/off 기능을 추가하는 경우 모든 클래스에 대해 코드들을 추가해줘야 함
  - 프로그램의 종류가 많아질수록 각 메서드를 각각 정의 및 추해줘야 하는 문제 발생
- 해당 클래스들을 관리하는 컨트롤러(리모컨)가 존재하는 경우 다양한 객체와 의존성을 깊게 갖기 때문에 변경에 어려움 발생

## Command Pattern
- 하나의 객체를 통해 여러 객체들에 명령해야할 때 사용되는 패턴

- **실행될 기능을 캡슐화**
  - 커맨드 객체는 명령을 받는 객체들에 대한 의존성 감소
  - 실행될 기능의 변경에도 호출자 클래스를 수정 없이 그대로 사용

- 캡슐화된 method 호출을 로그 기록용으로 저장하거나 취소 기능을 구현하고 재사용 가능

### Flow

고객 → 웨이트리스에게 주문 → 주문서 작성 → 주방장에게 주문 전달

![image](https://oopy.lazyrockets.com/api/v2/notion/image?src=https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2F328063da-393c-4324-a8a7-38aa7a0a0d46%2FScreen_Shot_2021-01-12_at_16.05.21.png&blockId=1ba33d99-5b57-4c4a-8918-9e3a6aa4bd27)

- 웨이터 입장 
  - 어떤 주문인지 조차도 몰라도 된다. 
  - 누가 식사를 준비할지 몰라도 된다.
- 주방장 입장 
  - 주문서를 읽고 필요한 메뉴만 준비하면 된다.
  - 누가 주문했는지 알 필요가 없다

![image](https://oopy.lazyrockets.com/api/v2/notion/image?src=https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2F7a1f0e55-68c7-492f-b783-f158ad989def%2FScreen_Shot_2021-01-12_at_16.10.21.png&blockId=1ec3b0e9-5e64-48ac-9fbd-3a07b4e36e33)

- Invoker에서 단순히 `execute()`를 호출
  - `makeBurger()`, `makeShake()`와 같은 특정 행동은 Receiver에 정의된 대로 호출

### Architecture
![image](https://user-images.githubusercontent.com/59307414/200165270-192b63e7-6c6a-405d-8e28-58488c375130.png)
![image](https://t1.daumcdn.net/cfile/tistory/255222435744588802)

**Command**
- 실행될 기능에 대한 인터페이스
- 실행될 기능을 execute 메서드로 선언함

**ConcreteCommand**
- 실제로 실행되는 기능을 구현
- 즉, Command라는 인터페이스를 구현함

**Invoker**
- 기능의 실행을 요청하는 호출자 클래스

**Receiver**
- ConcreteCommand에서 execute 메서드를 구현할 때 필요한 클래스
- 즉, ConcreteCommand의 기능을 실행하기 위해 사용하는 수신자 클래스

### Example
**Structure**
```
.
├── command
│      ├── command.go         // Command interface 
│      ├── light_off.go       // Concrete command
│      ├──  light_on.go       // Concrete command
│      ├── stereo_off.go      // Concrete command
│      └── stereo_on.go       // Concrete command
│
├── receiver
│      ├── interface
│      │      ├── light.go    // Receiver interface
│      │      └── stereo.go   // Receiver interface
│      │ 
│      ├── kitchen.go         // Concrete receiver
│      ├── living_room.go     // Concrete receiver
│      └── stereo_player.go   // Concrete receiver
│
├── invoker
│      └── remote_control.go  // Invoker
│
└── main.go
```
**Main**
```go
package main

func main() {
	// 리모컨 객체 생성 - invoker(호출자)
	remote := invoker.NewRemoteControl()

	// 커맨드 기능을 실행하기 위한 객체 생성 - receiver(수신자)
	livingRoom := receiver.NewLivingRoom()
	// 실제로 실행되는 기능을 구현하는 객체 생성 - command
	livingRoomLightOn := command.NewLightOnCommand(livingRoom)
	livingRoomLightOff := command.NewLightOffCommand(livingRoom)

	kitchen := receiver.NewKitchen()
	kitchenLightOn := command.NewLightOnCommand(kitchen)
	kitchenLightOff := command.NewLightOffCommand(kitchen)

	stereoPlayer := receiver.NewStereoPlayer()
	stereoPlayerOn := command.NewStereoOnWithCDCommand(stereoPlayer)
	stereoPlayerOff := command.NewStereoOffWithCDCommand(stereoPlayer)

	// 호출자(invoker)에 커맨드에 맵핑
	_ = remote.SetOnCommand(0, livingRoomLightOn)
	_ = remote.SetOffCommand(0, livingRoomLightOff)

	_ = remote.SetOnCommand(1, kitchenLightOn)
	_ = remote.SetOffCommand(1, kitchenLightOff)

	_ = remote.SetOnCommand(2, stereoPlayerOn)
	_ = remote.SetOffCommand(2, stereoPlayerOff)

	// 호출자(invoker)에서 원하는 커맨드 실행
	_ = remote.OnButtonWasPushed(0)
	_ = remote.OnButtonWasPushed(2)
	_ = remote.UndoButtonWasPushed()
}
```

### Command pattern vs Strategy Pattern
- 특정 행동을 캡슐화한다는 점에서 command pattern 은 strategy pattern 과 유사하다.
- 사용 목적에 따라 command pattern을 적용해야 할지, strategy pattern을 적용해야할지를 결정할 수 있다.
- strategy pattern의 경우 특정 행위를 **어떻게 구현**하고, 수정하느냐에 집중을 하고, command pattern은 **무엇을 실행**하느냐에 집중을 한다. 실제로 어떻게 구현하는지에 대한 내용은 외부에서 작성하게 된다.
- 디테일한 내용은 아래 링크에 정리가 잘 되어 있습니다.
  - https://tecoble.techcourse.co.kr/post/2021-10-04-strategy-command-pattern/

## Reference
- 커맨드 패턴 예제
  - https://refactoring.guru/ko/design-patterns/command

- 커맨드 패턴 설명
  - https://gmlwjd9405.github.io/2018/07/07/command-pattern.html
  - https://kotlinworld.com/370
  - https://blog.yevgnenll.me/posts/what-is-command-pattern

- 커맨트 패턴 vs 전략 패턴
  - https://tecoble.techcourse.co.kr/post/2021-10-04-strategy-command-pattern/
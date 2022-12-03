# Builder pattern
- 구현부에서 추상층을 분리하여 각자 독립적으로 변형하고 확장

## Structure
![image](https://refactoring.guru/images/patterns/diagrams/builder/structure-2x.png)

## Example
- `@Builder`로 자연스럽게 많이 사용해보셨을 듯 합니다.
  - 혹은 FCM 메시지 만들 때 간혹 사용했던 것으로 기억
  - aos, ios 플랫폼 별로 담아줘야 할 내용이 달라서 빌더 패턴 이용해서 message form을 만들었던 기억
- 책에서는 계획표 생성인데 그냥 FCM 만드는 예제 (실제로는 아니고 간단히)

```go
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
}
```

## Reference
- https://refactoring.guru/ko/design-patterns/builder
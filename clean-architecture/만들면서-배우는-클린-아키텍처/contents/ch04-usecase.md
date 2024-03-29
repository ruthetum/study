# ch 04. 유스케이스 구현하기

## 유스케이스 (Usecase)

일반적으로 유스케이스는 아래의 단계를 따른다.

1. 입력을 받는다.

2. 비즈니스 규칙을 검증한다

3. 모델 상태를 조작한다.

4. 출력을 반환한다.

### 입력을 받는다
유스케이스는 인커밍 어뎁터로부터 입력을 받는다.

이 단계를 '입력 유효성 검증'으로 명명하지 않는 이유는 유스케이스 코드는 도메인 로직에만 신경써야 하기 때문에 입력 유효성 검증으로 오염되면 안 된다.

그래서 입력 유효성 검증은 다른 곳에서 처리한다.

### 비즈니스 규칙을 검증한다
입력의 유효성에 대해서는 검증하지 않지만 비즈니스 규칙(business rule)에 대해서는 검증할 책임이 있다.

그리고 도메인 엔티티와 이 책임을 공유한다.

### 모델 상태를 조작한다
비즈니스 규칙을 충족하면 유스케이스는 입력을 기반으로 어떤 방법으로든 모델의 상태를 변경한다.

일반적으로 도메인 객체의 상태를 바꾸고, 영속성 어뎁터를 통해 구현된 포트로 이 상태를 전달해서 저장한다.

또는 또 다른 아웃고잉 어뎁터를 호출할 수도 있다.

### 출력을 반환한다
마지막 단계에서는 아웃고잉 어뎁터에서 온 출력값을, 유스케이스를 호출한 어뎁터로 반환할 출력 객체로 변환한다.

## 코드 구현
<img width="601" alt="스크린샷" src="https://github.com/ruthetum/study/assets/59307414/92bb2dc5-5700-46a9-b0ee-5839a208ae4e">

> 하나의 서비스가 하나의 유스케이스를 구현하고, 도메인 모델을 변경하고, 변경한 상태를 저앟기 위해 아웃 고잉 포트를 호출

```java
@Service
@RequiredArgsConstructor
@Transactional(readOnly = true)
public class SendMoneyService implements SendMoneyUseCase {

    private final LoadAccountPort loadAccountPort;
    private final AccountLock accountLock;
    private final UpdateAccountStatePort updateAccountStatePort;

    @Override
    public boolean sendMoney(SendMoneyCommand command) {
        // TODO: 비즈니스 규칙 검증
        // TODO: 모델 상태 조작
        // TODO: 출력 값 반환
    }
}
```

서비스(`SendMoneyService`)는 인커밍 포트 인터페이스인 `SendMoneyUseCase`를 구현한다.

계좌를 불러오기 위해서는 아웃고잉 포트 인터페이스인 `LoadAccountPort`를 호출한다.

그리고 데이터베이스의 계좌 상태를 업데이트하기 위해 `UpdateAccountStatePort`를 호출한다.

## 입력 유효성 검증

호출하는 어뎁터가 유스케이스에 입력을 전달하기 전에 입력 유효성을 검증하면 유효성 검증 을 각 어댑터에서 전부 구현해야 한다. 그럼 그 과정에서 실수할 수도 있고, 유효성 검증을 해야 한다는 사실을 잊어버리게 될 수도 있다.

애플리케이션 계층에서 입력 유효성을 검증해야 하는 이유는, 그렇게 하지 않을 경우 애플리케이션 코어의 바깥쪽으로부터 유효하지 않은 입력값을 받게 되고, 모델의 상태를 해칠 수 있기 때문이다.

> 입력 유효성 검증을 유스케이스 클래스가 아니라면 도대체 어디에서 진행해야 할까?

입력 모델(input model)이 이 문제를 다루도록 해보자. ‘송금하기’ 유스케이스에서 입력 모델은 예제 코드에서 본 `SendMoneyCommand` 클래스다. 정확히는 생성자 내에서 입력 유효성을 검증한다.

## 유스케이스마다 다른 입력 모델
각 유스케이스 전용 입력 모델은 유스케이스를 훨씬 명확하게 만들고 다른 유스케이스와의 결합도를 제거해서 불필요한 부수효과가 발생하지 않게 한다.

물론 비용이 안 드는 것은 아니다. 들어오는 데이터를 각 유스케이스에 해당하는 입력 모델에 매핑해야 하기 때문이다.

## 비즈니스 규칙 검증하기
입력 유효성 검증은 유스케이스 로직의 일부가 아닌 반면, 비즈니스 규칙 검증은 분명히 유스케이스 로직의 일부다.

그렇다면 언제 입력 유효성을 검증하고 언제 비즈니스 규칙을 검증해야 할까?

둘 사이의 아주 실용적인 구분점은 비즈니스 규칙을 검증하는 것은 도메인 모델의 현재 상태에 접근해야 하는 반면, 입력 유효성 검증은 그럴 필요가 없다는 것이다.

입력 유효성을 검증하는 일은 @NotNull 애너테이션을 붙인 것처럼 선언적으로 구현할 수 있지만 비즈니스 규칙을 검증하는 일은 조금 더 맥락이 필요하다.

- 입력 유효성 검증: 구문상(syntactical)의 유효성을 검증
  - e.g. 송금되는 금액은 0보다 커야 한다

- 비즈니스 규칙 검증: 유스케이스 맥락 속에서 의미적(semantical)인 유효성을 검증
  - e.g. 출금 계좌는 초과 출금되어서는 안 된다
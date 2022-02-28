# Kafka vs RabbitMQ vs Redis Queue
## 메시징 플랫폼
- 흔히 메시징 플랫폼이라고 불리는 플랫폼들은 크게 두 가지로 나눌 수 있다.

    1. 메시지 브로커 (Message Broker)
    2. 이벤트 브로커 (Event Broker)

- 메시지 브로커는 이벤트 브로커 역할을 할 수 없지만, 이벤트 브로커는 메시지 브로커 역할을 할 수 있다.

## 메시지 브로커(Message Broker)
- 많은 기업들에서 <b>대규모 메시지 기반 미들웨어 아키텍처</b>에서 사용되어짐
- 메시지 브로커에 존재하는 큐에 데이터를 보내고 받는 프로듀서와 컨슈머를 통해 메세지를 통신하고 네트워크를 맺는 용도로 사용
- 메시지를 받아서 적절히 처리하고 나면 <b>즉시 또는 짧은 시간 내에 삭제</b>되는 구조

## 이벤트 브로커(Event Broker)
- 이벤트 또는 메시지라고 불리는 레코드를 딱 하나만 보관, 인덱스를 통해 개별 엑세스를 관리
- 업무상 <b>필요한 시간동안 데이터를 보관</b> 가능
- 서비스에서 나오는 이벤트를 마치 데이터베이스에 저장하듯이 이벤트 브로커의 큐에 저장
- 이를 통해 딱 한 번 일어난 이벤트를 <b>단일 진실 공급원</b>으로 사용할 수 있음
- 장애가 발생했을 때 <b>장애가 발생한 지점부터 재처리 가능</b>
- 많은 양의 실시간 스트림 데이터를 효과적으로 처리할 수 있음
- 이벤트 브로커를 통해 이벤트 기반 마이크로서비스 아키텍처를 설계 가능

|메시지 브로커|이벤트 브로커|
|---|---|
|![image](https://user-images.githubusercontent.com/59307414/155964463-e5fcea5b-ed92-4157-8f63-fc2c32ea7db5.png)<br/>![image](https://user-images.githubusercontent.com/59307414/155964533-8cd21e34-920c-45d5-a50e-aad88fd83924.png)|![image](https://user-images.githubusercontent.com/59307414/155964552-ec572210-4717-4841-b9eb-e9560fd6a500.png)<br/>![image](https://user-images.githubusercontent.com/59307414/155964584-53e4bdd9-05fb-4792-b4cf-b4aa63acdf4c.png)|
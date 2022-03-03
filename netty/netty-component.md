## Netty 
- Netty는 프로토콜 서버 및 클라이언트와 같은 네트워크 애플리케이션의 빠르고 쉬운 개발을 가능하게 하는 NIO 클라이언트 서버 프레임워크
- 기존의 소켓 프로그래밍은 클라이언트가 접속하게 되면 스레드를 할당해야 하는데(1:1관계), 정말 많은 클라이언트가 접속을 하게 될 경우 그 숫자만큼 스레드를 생성해야 해서 리소스의 낭비로 이루어지고, 문맥 교환과 관련된 문제와 입력이나 출력 데이터에 관련한 무한 대기 현상이 발생하는 문제가 있다.
- 이러한 네트워크 문제 때문에 개발된 방법이 자바의 NIO 방식(Non-Blocking Input Ouput)이다.
    - <b>비동기 네트워크 통신</b>
- Tomcat서버가 10,000건의 커넥션을 처리한다면, netty는 100,000건 이상의 커넥션을 처리할 수 있는 장점이 있다.

### Netty 컴포넌트
- 핵심 컴포넌트로는 Channel, CallBack, Future, Event와 Handler, Event Loop와 PipeLine이 있다.

#### Channel
- 하나 이상의 입출력 작업을 수행할 수 있는 하드웨어 장치, 파일, 네트워크 소켓이나 프로그램 컴포넌트와 같은 Open된 Connection 을 의미

#### CallBack
- 다른 메서드로 자신에 대한 참조를 제공하는 메서드
- 이벤트를 처리할 때 Netty 내부적으로 콜백을 이용하는데, 이때 ChannelHandler 인터페이스를 통해 이벤트를 처리

#### Future
- 작업이 완료가 될 경우 애플리케이션에 알림
- Future 객체는 비동기 작업의 결과를 담는 Plachloder의 역할을 함
- 이때 ChannelFuture 인터페이스를 이용해 결과값을 활용

#### Event, Handler
- Netty는 작업 상태의 변화를 알리기 위해 이벤트를 이용하고, 발생한 이벤트를 기준으로 Handler를 통해 트리거

#### Event Loop
- 새로운 event를 반복적으로 확인
![eventloop](https://img1.daumcdn.net/thumb/R1280x0/?scode=mtistory2&fname=https%3A%2F%2Fblog.kakaocdn.net%2Fdn%2FyEhK2%2FbtqHBRuUD2o%2FijCfGliB3cD7LcwZe6Tz01%2Fimg.png)

#### PipeLine
- Event loop에서 event를 받아 handler에 전달하는 역할
- ![pipeline](https://img1.daumcdn.net/thumb/R1280x0/?scode=mtistory2&fname=https%3A%2F%2Fblog.kakaocdn.net%2Fdn%2FEN6fG%2FbtqHLfuMOba%2FRY7O3U5BYp5Z9hHACgIqxK%2Fimg.png)

#### Inbound event
- 이벤트 루프가 발생 시킨 이벤트 (소켓연결, 데이터수신 등)를 작성한 inbound event handler에게 전달

#### Outbound event
- 사용자가  요청한 동작(쓰기, 읽기 일시중단 등)을 작성한 outbound event handler에게 전달
- 최종적으로 이벤트 루프에 전달되어 I/O가 수행되도록 함

### Flow
- ![flow](https://img1.daumcdn.net/thumb/R1280x0/?scode=mtistory2&fname=https%3A%2F%2Fblog.kakaocdn.net%2Fdn%2FPWASk%2FbtqHJpEmBF9%2FR3O3z2HcrDZhEDtZZQazG0%2Fimg.png)


### Reference
- https://netty.io/
- https://hbase.tistory.com/116
- https://narup.tistory.com/118

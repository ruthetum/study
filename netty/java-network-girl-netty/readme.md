# [자바 네트워크 소녀 Netty](http://www.kyobobook.co.kr/product/detailViewKor.laf?ejkGb=KOR&mallGb=KOR&barcode=9788968482243&orderClick=LAG&Kc=)

## ch1. 네티 맛보기
- [기본 에코 서버 만들기](./netty-example/src/main/java/ch1/echo/)
    - 텔넷 접속으로 `putty` 활용 (connection type : `Raw`)

- 데이터 이동의 방향성
    - 네티는 이벤트를 인바운드(Inbound) 이벤트와 아웃바운드(Outbound) 이벤트로 구분한 추상화 모델 제공

        ![image](https://user-images.githubusercontent.com/59307414/156346568-3286825f-b641-4624-8937-1e4924ed6f6a.png)

---

## ch2. 네티의 주요 특징

> 네티는 <b>비동기 이벤트 기반</b> 네트워크 애플리케이션 프레임워크로써 유지보수를 고려한 고성능 <b>프로토콜 서버</b>와 클라이언트를 <b>빠르게 개발할 수 있다.

- [블로킹 서버와 논블로킹 서버](./netty-example/src/main/java/ch2/)
    - `ServerSocket` vs `ServerSocketChannel`
    - `serverSocketChannel.configureBlocking(false)` : ServerSocketChannel blocking 설정 dafault 값은 true, non-blocking으로 쓰기 위해서는 configureBlocking(false) 설정

### 이벤트 기반 프로그래밍
- 이벤트를 먼저 정의해두고 발생한 이벤트에 따라서 코드가 실행되도록 프로그램을 작성
- ex. Non-Blocking Socket의 `Selector`를 사용한 I/O 이벤트 감지

---

## ch3. 부트스트랩
- Bootstrap은 네트워크 애플리케이션의 동작 방식에 대한 설정을 담당
    - `Bootstrap` : 클라이언트 애플리케이션 담당
    - `ServerBootstrap` : 서버 애플리케이션 담당

### `group` : 이벤트 루프 설정
- 데이터 송수신 처리를 위한 이벤트 루프를 설정

```java
// 이벤트 루프 설정 API
@SuppressWarnings("unchecked")
public B group(EventLoopGroup group) {
    if (group == null) {
        throw new NullPointerException("group");
    }
    if (group != null) {
        throw new IllegalStateException("group set already");
    }
    // 하나의 이벤트 루프만 설정
    this.group = group;
    return (B) this;
}
```

### `channel` : 소켓 입출력 모드 설정
- 소켓의 입출력 모드를 설정
- 부트스트랩의 channel 메서드에 드록된 소켓 채널 생성 클래스가 소켓 채널을 설정

|클래스|내용|
|---|---|
|`LocalServerChannel.class`|하나의 자바 가상머신에서 가사 통신을 위한 서버 소켓 채널을 생성하는 클래스|
|`OioServerSocketChannel.class`|블로킹 모드의 서버 소켓 채널을 생성하는 클래스|
|`NioServerSocketChannel.class`|논블로킹 모드의 서버 소켓 채널을 생성하는 클래스|
|`EpollServerSocketChannel.class`|리눅스 커널의 epoll 입출력 모드를 지원하는 서버 소켓 클래스(windows에서 실행 시 에러)|
|`OioSctpServerSocketChannel.class`|SCTP 전송 계층을 사용하는 블로킹 모드의 서버 소켓 채널|
|`NioSctpServerSocketChannel.class`|SCTP 전송 계층을 사용하는 논블로킹 모드의 서버 소켓 채널|
|`NioUdtByteAcceptorChannel.class`|UDT 프로토콜을 지원하는 논블로킹 모드의 서버 소켓 채널을 생성하는 클래스, 내부적으로 스트림 데이터를 처리하도록 구현되어 있음|
|`NioUdtMessageAcceptorChannel.class`|UDT 프로토콜을 지원하는 논블로킹 모드의 서버 소켓 채널을 생성하는 클래스, 내부적으로 데이터그램 패킷을 처리하도록 구현되어 있음|

> [SCTP(Stream Control Transmission Protocol)](https://ko.wikipedia.org/wiki/%EC%8A%A4%ED%8A%B8%EB%A6%BC_%EC%A0%9C%EC%96%B4_%EC%A0%84%EC%86%A1_%ED%94%84%EB%A1%9C%ED%86%A0%EC%BD%9C)

- 스트림 제어 전송 프로토콜
- 전송 계층의 프로토콜로 UDP 메시지 스트리밍 특성과 TCP의 연결형 및 신뢰성 제공 특성을 조합한 프로토콜
- TCP처럼 연결지향적 프로토콜이며 혼잡 제어를 통해 신뢰성 있는 순차적 메시지 전송을 보장

- TCP와 SCTP handshake

    ![tcp-sctp-handshake](https://user-images.githubusercontent.com/59307414/158057493-9f990e65-98ff-4206-bca1-6b91490df3aa.png)


- TCP와 SCTP 종료

    ![tcp-sctp-close](https://user-images.githubusercontent.com/59307414/158057428-56714f22-b6b0-43a4-a93d-e737569ebc6c.png)

    - SCTP에서 반 닫힘(Half Closed) 상태를 지원하지 않음
    - TCP 연결에서 반 닫힘 상태는 TCP 연결 해제를 요청하는 측의 `FIN`패킷에 대한 결과로 상대방의 `ACK`와 `FIN`패킷을 받은 상태를 의미하고, 이 상태에서 커널의 포트 상태가 `TIME_WAIT`으로 변경
    - SCTP에서는 `TIME_WAIT` 상태가 존재하지 않음

> UDT(UDP-based Data Transfer)
- 애플리케이션 계층의 UDP 기반의 프로토콜
- 신뢰성 있는 데이터 전송 프로토콜로 애플리케이션 계층에 속하기 때문에 특정 하드웨어나 플랫폼에 종속적이지 않음
- 기존 UDP와 다르게 여결 유지를 지원하며 양뱡향 통신을 지원

### channelFactory - 소켓 입출력 모드 설정
- channel 메서드와 동일하게 소켓의 입출력 모드를 설정하는 API

### handler - 서버 소켓 채널의 이벤트 핸들러 설정
- 서버 소켓 채널의 이벤트를 처리할 핸들러 설정 API

### childHandler - 소켓 채널의 데이터 가공 핸들러 설정
- 클라이언트 소켓 채널로 송수신되는 데이터를 가공하는 데이터 핸들러 설정 API

---

## ch4. 채널 파이프라인과 코텍
- `채널 파이프라인` : 채널에서 발생한 이벤트가 이동하는 통로
- 이 통로를 통해서 이동하는 이벤트를 처리하는 클래스가 `이벤트 핸들러`
- 이 이벤트 핸들러를 상속받아서 구현한 구현체들이 `코덱`
    - 자주 사용하는 이벤트 핸들러를 미리 구현해둔 묶음은 `io.netty.handler.codec` 패키지에 존재

### 채널 파이프라인
- 파이프라인 구조

    ![channel-pipline](https://user-images.githubusercontent.com/59307414/158058547-d62a15f0-7244-484c-aca8-7afeb50a9aeb.png)

    - 네티의 이벤트 흐름은 전기의 흐름과 유사
    - 발전소(채널)에서 발생한 전기(이벤트)가 전선(채널 파이프라인)을 타고 가정으로 이동
    - 가정에서는 콘센트/멀티탭(이벤트 핸들러)를 통해 가전제품(이벤트 처리 코드)을 연결해놓고 기능이 작동

```java
public static void main(String[] args) throws Exception {
    EventLoopGroup bossGroup = new NioEventLoopGroup(1);
    EventLoopGroup workerGroup = new NioEventLoopGroup();
    try {
        ServerBootstrap b = new ServerBootstrap();
        b.group(bossGroup, workerGroup)
                .channel(NioServerSocketChannel.class)
                .childHandler(new ChannelInitializer<SocketChannel>() { // ①
                    @Override
                    public void initChannel(SocketChannel ch) { // ②
                        ChannelPipeline p = ch.pipeline(); // ③
                        p.addLast(new EchoServerHandler()); // ④
                    }
                });

        ChannelFuture f = b.bind(8888).sync();

        f.channel().closeFuture().sync();
    }
    finally {
        workerGroup.shutdownGracefully();
        bossGroup.shutdownGracefully();
    }
}
```

- ① : `childHandler` 메서드를 통해서 연결된 클라이언트 소켓 채널이 사용할 채널 파이프라인 설정
- ② : 클라이언트 소켓 채널이 생성될 때 자동으로 호출(`initChannel` 메서드), 파이프라인 설정
- ③ : `initChannel` 메서드의 인자로 입력된 소켓 채널(연결된 클라이언트 소켓 채널)에 설정된 채널 파이프라인을 가져옴
- ④ : `add` 메서드를 통해 이벤트 핸들러 등록


### 이벤트 핸들러  
- 네티는 소켓 채너에서 발생하는 이벤트를 인바운드(inbound) 이벤트와 아웃바운드(outbound) 이벤트로 추상화

#### 인바운드(inbound) 이벤트
- 발생 순서

    ![inbound-event](https://user-images.githubusercontent.com/59307414/158059146-4941356c-966b-48b7-a6c1-028819479773.png)

- `channelRegistered`
    - 서버에서는 아래의 상황에서 발생
        - 서버 소켓 채널을 생성할 때
        - 새로운 클라이언트가 서버에 접속하여 클라이언트 소켓 채널이 생성될 때
    - 클라이언트에서는 서버 접속을 위한 connect 메서드를 수행할 때 이벤트 발생

- `channelActive`
    - `channelRegistered` 이벤트 이후에 발생
    - 채널이 생성되고 이벤트 루프에 등록된 이후에 네티 API를 사용하여 채널 입출력을 수행할 상태가 되었음을 알려줌
    - 아래의 상황에서 사용하기 적합함
        - 서버에 연결된 클라이언트의 연결 개수를 셀 때
        - 서버에 연결된 클라이언트에게 최초 연결에 대한 메세지를 전송할 때
        - 클라이언트가 연결된 서버에 최초로 메세지를 전송할 때
        - 클라이턴트가 서버에 연결된 상태에 대한 작업이 필요할 때
    - 서버 또는 클라이언트가 상대방에 연결한 직후 최초 한 번 수행할 작업을 처리하기에 적합

- `channelRead`
    - 데이터 수신

- `channelReadComplete`
    - 데이터 수신 완료
    
    > `channelRead` vs  `channelReadComplete`
    > - 클라이언트가 서버로 'A', 'B', 'C'라는 데이터를 순차적으로 전송한 경우
    > - 서버에서는 'A', 'B', 'C' 각각의 데이터를 전송받았을 때 `channelRead` 이벤트 발생
    > - 이후에 ByteBuf 인터페이스 구현체를 통해 입력받은 객체의 데이터가 'ABC'인 경우 `channelReadComplete` 이벤트 발생
    > `channelRead`이벤트는 채널에 데이터가 있을 때 발생하고, `channelReadComplete` 이벤트는 채널의 데이터를 다 읽어서 더 이상 데이터가 없을 때 발생
    
- `channelInactive`
    - 채널이 비활성화되었을 때 발생
    - 이벤트가 발생한 이후에는 채널에 대한 입출력 작업을 수행할 수 없음

- `channelUnregisterd`
    - 채널이 이벤트 루프에서 제거되었을 때 발생
    - 이벤트를 수신한 이후에는 채널에서 발생한 이벤트를 처리할 수 없음

#### 아웃바운드(outbound) 이벤트
- 소켓 채널에서 발생하는 이벤트 중 네티 사용자(개발자)가 요청한 동작에 해당하는 이벤트를 말함
    - 연결 요청
    - 데이터 전송
    - 소켓 닫기
- `ChannelOutboundHandler` 인터페이스로 제공
    - 모든 `ChannelOutboundHandler` 이벤트는 `ChannelHandlerContext` 객체를 인수로 받음

- `bind`
    - 서버 소켓 채널이 클라이언트의 연결을 대기하는 IP와 Port가 설정되었을 때 발생

- `connect`
    - 클라이언트 소켓 채널이 서버에 연결되었을 때 발생

- `disconnect`
    - 클라이언트 소켓 채널이 끊어졌을 때 발생

- `close`
    - 클라이언트 소켓 채널의 연결이 닫혔을 때 발생

- `write`
    - 소켓 채널에 데이터가 기록되었을 때 발생
    - 소켓 채널에 기록된 데이터 버퍼가 인수로 입력

- `flush`
    - 소켓 채널에 대한 flush 메서드가 호출되었을 때 발생
    - 별도의 인수 없음

### 코덱
> `송신데이터` → `인코딩` → `소켓 채널` → `디코딩` → `수신 데이터`
- `ChannelOutboundHandler` : 인코더 역할
- `ChannelInboundHandler` : 디코더 역할

---

## ch5. 이벤트 모델
- 이벤트 루프 기반 프레임워크
    - JS : Node.js, ...
    - Java : Vert.x, Netty, ...

### 이벤트 루프(Event Loop)
- 이벤트 기반 애플리케이션이 이벤트를 처리하는 방법은 크게 두 가지

1. 이벤트 리스너와 이벤트 처리 스레드 활용
    - 이벤트를 처리하는 로직을 가진 이벤트 메서드를 대상 객체의 이벤트 리스너로 등록
    - 객체에 이벤트가 발생했을 때 이벤트 처리 스레드에서 등록된 메서드를 수행

2. 이벤트 큐 활용
    - 이벤트 큐에 이벤트를 등록하고 이벤트 루프가 큐에 접근하여 처리

### 단일 스레드와 다중 스레드 이벤트 루프
|구분|장점|단점|
|---|---|---|
|단일 스레드|- 단순하고 예측 가능한 동작<br>- 이벤트가 발생한 순서대로 수행|- 다중 코어 CPU를 효율적으로 사용하지 못함<br>- node.js의 경우 다중 인스턴스를 실행해서 보완하기도 함|
|다중 스레드|- 다중 코어 CPU를 효율적으로 사용<br>- 전체 처리 시간 단축|- 구현이 복잡<br>- 스레드 경합 발생<br>- 이벤트의 발생 순서와 실행 순서의 불일치|

- 애플리케이션을 다중 스레드로 구현하면 전체 처리 시간을 단축할 수 있음
- 하지만 다중 스레드의 장점을 얻기 위해 스레드의 개수를 너무 많이 설정하거나 제한하지 않으면 과도한 garbage collection이 발생하거나 Out of Memory 에러가 발생할 수 있음
- 다중 스레드 아키텍처는 자원을 효율적으로 사용하지만 context switching 비용과 스레드 경합이라는 단점도 존재
- 다중 스레드 구현 시 초기에 스레드가 증가할 때는 처리량이 증가하지만 최대 성능을 기점으로 처리량이 줄어듦
    - 최대 성능에 도달, 시스템의 CPU 사용률이 100%에 근접하게 되는 순간 스레드의 개수가 더 늘어나면 스레드 경합으로 이해 CPU 사용률이 100%인 상태에서 오히려 처리량은 감소
- 따라서 사용하는 시스템에 적정한 수치로 스레드 개수를 설정해야 함
    - 스레드 개수의 적정 수치를 찾기 위해서 애플리케이션 부하 테스트 도구 사용 : `JMeter`, `nGrinder`, ...

### 네티 이벤트 루프
- 네티는 단일/다중 스레드 이벤트 루프 모두 사용 가능
- 다중 스레드 이벤트 루프

    ![image](https://user-images.githubusercontent.com/59307414/158062006-10bf0aa2-b101-464c-8f3e-e94d568b9c50.png)

    - 객체가 Event 1부터 Event 4까지 이벤트를 발생한 상황
    - 이벤트 루프 스레드에서 Event 1과 Event 2를 처리하고, 먼저 수행이 완료된 이벤트 루프 스레드가 Event 3를 처리하게 됨
    - 이벤트 루프 스레드 2가 먼저 수행이 완료되면 Event 3는 이벤트 루프 스레드 2에서 작업이 수행
    - 만약 Evnet 3가 먼저 작업이 완료된다면 이벤트 발생 순서와 다르게 이벤트 처리 순서는 Event 2, 3, 1, 4 순으로 처리되면서 발생 순서와 실행 순서가 일치하지 않게 됨

- 네티의 이벤트 루프와 채널의 구조

    - 네티는 다중 스레드 이벤트 루프를 사용함에도 불구하고 이벤트 발생 순서와 실행 순서를 일치시킬 수 있음

    ![image](https://user-images.githubusercontent.com/59307414/158062359-63579fe1-773f-4441-a95f-c8be13eb9dcd.png)

    - 아래 3가지 특징을 통해 네티의 이벤트 루프는 이벤트 발생 순서와 실행 순서를 일치시킴

        > - 네티의 이벤트는 채널에서 발생함
        > - 이벤트 루프 객체는 이벤트 큐를 가짐
        > - 네티의 채널은 하나의 이벤트 루프에 등록됨
    
    - 이벤트 루프들이 이벤트 큐를 공유하는 경우 발생 순서와 처리 순서의 불일치가 발생
    - 네티는 이벤트 큐를 이벤트 루프 스레드의 내부에 둠으로써 원인을 제거

---

## Reference
- [자바 네트워크 소녀 Netty 소스 코드 예제](https://github.com/krisjey/netty.book.kor)
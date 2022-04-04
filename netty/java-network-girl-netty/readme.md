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

### 네티 비동기 I/O 처리
- 네티는 비동기 호출을 위한 두 가지 패턴을 제공
    - 1. 리액터 패턴의 구현체인 이벤트 핸들러
    - 2. 퓨처(Future) 패턴
- 퓨처 패턴은 미래에 완료될 작업을 등록하고, 처리 결과를 확인하는 객체를 통해서 작업의 완료를 확인
- 퓨처 패턴은 메서드를 호출하는 즉시 퓨처 객체를 돌려준다.
> Ex. 빵집에 빵을 예약하고 정해진 시간에 받으러 가는 경우
> 1. 빵집에서 빵을 주문하고 계산을 하면, 빵의 주문서를 받는다.
> 2. 빵이 완성되고 정해진 시간에 빵집을 찾아가면 주문서와 빵을 교환한다.

```java
public class SpecialCake {
    public static void main(String[] args) {
        Bakery bakery = new Bakery();

        // 빵을 주문하고 주문서를 받는다.
        Future future = bakery.orderCake();

        // 다른 일을 하다가
        doSomething();

        // 빵이 완성되었는지 확인한다.
        if (future.isDone()) {
            Cake cake = future.getCake();
        } else {
            while(future.isDone() != true) {
                // 다른 일을 한다.
                doSomething();
            }
            Cake cake = future.getCake();
        }
    }
}
```

#### Netty에서 퓨처 패턴
```java
public class EchoServer {
    public static void main(String[] args) throws Exception {
        EventLoopGroup bossGroup = new NioEventLoopGroup(1);
        EventLoopGroup workerGroup = new NioEventLoopGroup();
        try {
            ServerBootstrap b = new ServerBootstrap();
            b.group(bossGroup, workerGroup)
                    .channel(NioServerSocketChannel.class)
                    .childHandler(new ChannelInitializer<SocketChannel>() {
                        @Override
                        public void initChannel(SocketChannel ch) {
                            ChannelPipeline p = ch.pipeline();
                            p.addLast(new EchoServerHandler());
                        }
                    });

            ChannelFuture future = b.bind(8888).sync();     // ①
            future.channel().closeFuture().sync();          // ②
        }
        finally {
            workerGroup.shutdownGracefully();
            bossGroup.shutdownGracefully();
        }
    }
}
```
- ①
    - 8888번 포트를 사용하도록 바인드하는 비동기 bind 메서드를 호출
    - bind 메서드는 포트 바인딩이 완료되기 전에 ChannelFuture 객체를 반환
    - sync 메서드는 주어진 ChannelFuture 객체의 작업이 완료될 때까지 블로킹
    - 따라서 bind 메서드의 처리가 완료될 때 sync 메서드도 같이 완료
- ②
    - future 객체를 통해서 채널을 얻어 오고,
    - 바인드가 완료된 서버 채널의 CloseFuture 객체를 반환
    - CloseFuture 객체는 채널의 연결이 종료될 때 연결 종료를 이벤트를 받음

---

## ch6. 바이트 버퍼
- 자바 바이트 버퍼 vs 네티 바이트 버퍼

### 자바 NIO 바이트 버퍼
- ByteBuffer, CharBuffer, IntBuffer, ShortBuffer, ...
- 바이트 버퍼 클래스는 내부의 배열 상태를 관리하는 세 가지 속성을 가지고 있음

#### 바이트 버퍼 클래스의 속성 세 가지
1. `capacity`
- 버퍼에 저장할 수 있는 데이터의 최대 크기로 한 번 정하면 변경이 불가능
- 버퍼를 생성할 때 생성자의 인수로 인력한 값

2. `position`
- 읽기 또는 쓰기가 작업 중인 위치, 버퍼 객체가 생성될 때 0으로 초기화
- 데이터 입력(put 메서드)이나 데이터 조회(get 메서드) 호출 시 자동으로 증가

3. `limit`
- 읽고 쓸 수 있는 버퍼 공간의 최대치
- limit 메서드로 값을 조절할 수 있으나 capacity 값보다 크게 설정할 수 없음

#### 자바 바이트 버퍼 생성 메서드 세 가지
1. `allocate` : 힙 버퍼 생성
- JVM 힙 영역에 바이트 버퍼를 생성
- 메서드의 인수는 생성할 바이트 버퍼의 크기이며 capacity의 값으로 설정
- 생성되는 바이트 버퍼의 값은 모두 0으로 초기화

2. `allocateDirect` : 다이렉트 버퍼 생성
- JVM 힙 영역이 아닌 운영체제의 커널 영역에 바이트 버퍼를 생성
- allocateDirect 메서드는 ByteBuffer 추상 클래스만 사용 가능 -> 다이렉트 버퍼는 ByteBuffer로만 생성 가능
- 메서드 인수는 allocate와 마찬가지로 생성할 바이트 버퍼의 크기이며 capacity의 값으로 설정
- 생성되는 바이트 버퍼의 값은 모두 0으로 초기화

3. `wrap`
- 입력된 바이트 배열을 사용해서 바이트 버퍼를 생성
- 입력에 사용된 바이트 배열이 변경되면 wrap 메서드를 사용해서 생성한 바이트 버퍼의 내용도 변경

#### 자바 바이트 버퍼 사용법
```java
public void ByteBufferFuction1() {
    ByteBuffer buffer = ByteBuffer.allocate(11);

    buffer.put((byte) 1);
    buffer.put((byte) 2);
    buffer.put((byte) 3);


    // put 메서드를 사용하면서 position이 증가(+1)돼서 3이 출력되지 않음
    System.out.println(buffer.get());       // 0
    System.out.println(buffer.position())   // 4
    
    // rewind 메서드를 통해 position을 0으로 변경해줘야 됨
    buffer.rewind();
    System.out.println(buffer.get()); // 1
}
```
- get, put 메서드 모두 호출하게 되면 postion이 증가
- rewind 메서드를 사용하면 position을 0으로 변경

```java
public void ByteBufferFuction2() {
    byte[] temp = {1, 2, 3, 4, 5, 0, 0, 0, 0, 0, 0};

    // wrap 메서드로 버퍼 생성
    ByteBuffer buffer = ByteBuffer.wrap(temp);

    System.out.println(buffer.position());   // 0
    System.out.println(buffer.limit());      // 11

    System.out.println(buffer.get());        // 1
    System.out.println(buffer.get());        // 2
    System.out.println(buffer.get());        // 3
    System.out.println(buffer.position());   // 현재 position : 3

    buffer.flip();
    System.out.println(buffer.position());   // 0
    // flip 메서드로 인해 limit이 flip 메서드 직전의 position 값으로 변경
    System.out.println(buffer.limit());      // 3
}
```
- flip 메서드는 get, put 메서드가 호출된 이후의 position 정보를 저장


#### 자바 바이트 버퍼 정리
- 자바 바이트 버퍼는 사용할 때 읽기와 쓰기를 분리해서 생각해야 되고, 멀티 스레드 환경에서 바이트 버퍼를 공유하지 않아야 됨
- 네티는 이런 자바 바이트 버퍼의 문제를 해결하기 위해 읽기를 위한 인덱스와 쓰기를 위한 인덱스를 구분해서 제공

### 네티 바이트 버퍼
- 읽기 인덱스와 쓰기 인덱스 구분
- flip 메서드없이 읽기/쓰기 가능
- 가변 바이트 버퍼
- 바이트 버퍼 풀
- 복합 버퍼
- 자바 바이트 버퍼와 네티 바이트 버퍼 상호 변환

#### 네티 바이트 버퍼 생성 방법
- 네티 바이트 버퍼는 자바 바이트 버퍼와 달리 프레임워크 레벨의 바이트 버퍼 풀을 제공
- 이 바이트 버퍼 풀을 이용해서 바이트 버퍼를 재사용
- 네티 바이트 버퍼를 생성할 때는 두 가지를 선택해야 됨
    - 풀링 여부
    - 다이렉트 버퍼 여부

#### 네티 바이트 버퍼의 종류와 생성 방법 네 가지
|버퍼 종류|풀링 사용|풀링 사용 X|
|:---:|:---:|:---:|
|힙 버퍼|`PooledHeapByteBuf`|`UnpooledHeapByteBuf`|
|다이렉트 버퍼|`PooledDirectByteBuf`|`UnpooledDirectByteBuf`|

|버퍼 생성 방법|풀링 사용|풀링 사용 X|
|:---:|:---:|:---:|
|힙 버퍼|`ByteBufAllocator.DEFAULT.heapBuffer()`|`Unpooled.buffer()`|
|다이렉트 버퍼|`ByteBufAllocator.DEFAULT.directBuffer()`|`Unpooled.directBuffer()`|
- 풀링을 사용하는 힙, 다이렉트 버퍼의 경우 `ByteBufAllocator` 하위 추상 구현체인 `PooledBufAllocator` 클래스로 생성

### 네티 바이트 버퍼 사용법
#### 읽기/쓰기

```java
public void NettyByteBufferFuction1() {
    // 네티 바이트 버퍼 생성 방법 4가지
    // 풀링 O, 힙 버퍼
    ByteBuf pooledHeapByteBuf = PooledBufAllocator.DEFAULT.heapBuffer(11);
    // 풀링 O, 다이렉트 버퍼
    ByteBuf pooledDirectByteBuf = PooledBufAllocator.DEFAULT.heapBuffer(11);
    // 풀링 X, 힙 버퍼
    ByteBuf unpooledHeapByteBuf = Unpooled.buffer(11);
    // 풀링 X, 다이렉트 버퍼
    ByteBuf unpooledDirectByteBuf = Unpooled.directBuffer(11);


    ByteBuf buf = PooledBufAllocator.DEFAULT.heapBuffer(11);

    // 정수 65537 작성 -> 4바이트 작성
    buf.writeInt(65537);

    // 읽어들일 수 있는 바이트 : 4
    System.out.println(buf.readableBytes());    // 4
    // 기록할 수 있는 바이트 : 7
    System.out.println(buf.writableBytes());    // 7 

    // 2 바이트 읽기 : 1
    // 65537 = 0x10001 -> 4바이트 패딩 : 0x00010001
    System.out.println(buf.readShort());        // 1

    // 읽어들일 수 있는 바이트는 4바이트에서 2바이트를 읽어서 2
    System.out.println(buf.readableBytes());    // 2
    // 기록할 수 있는 바이트는 그대로 7
    System.out.println(buf.writableBytes());    // 7

    // 남은 데이터 있는지 확인 : 2바이트 남음
    System.out.println(buf.isReadable());       // true

    // 버퍼 초기화
    buf.clear();

    // 읽어들일 수 있는 바이트 : 0
    System.out.println(buf.readableBytes());    // 0
    // 기록할 수 있는 바이트 : 11
    System.out.println(buf.writableBytes());    // 11 
}
```

#### 가변 크기 버퍼
- 자바 바이트 버퍼는 버퍼를 생성할 때 크기를 지정해야 하고, 한 번 생성된 바이트 버퍼의 크기를 변경할 수 없음
- 네티 바이트 버퍼는 생성된 바이트 버퍼의 크기를 동적으로 변경할 수 있음

```java
public void NettyByteBufferFuction2() {
    ByteBuf buf = PooledBufAllocator.DEFAULT.heapBuffer(11);

    String source = "hello world";

    // 버퍼에 11바이트 저장
    buf.writeBytes(source.getBytes());
    
    // 읽어들일 수 있는 바이트 : 11
    System.out.println(buf.readableBytes());    // 11
    // 기록할 수 있는 바이트 : 0
    System.out.println(buf.writableBytes());    // 0 

    // 버퍼에 저장된 문자열 확인
    System.out.println(buf.toString(Charset.defaultCharset())); // "hello world"


    // 버퍼 크기 감소
    buf.capacity(6);
    // 버퍼에 저장된 문자열 확인, capacity가 변경되면서 문자열이 잘려짐
    System.out.println(buf.toString(Charset.defaultCharset())); // "hello "

    // 버퍼 크기 증가
    buf.capacity(13);
    // 버퍼에 추가 저장
    buf.writeBytes("world".getBytes());
    // 버퍼에 저장된 문자열 확인, capacity가 변경되면서 추가 저장 가능
    System.out.println(buf.toString(Charset.defaultCharset())); // "hello world"

    // capacity 확인 : 13
    System.out.println(buf.capacity());    // 13
    // 읽어들일 수 있는 바이트 : 11
    System.out.println(buf.readableBytes());    // 11
    // 기록할 수 있는 바이트 : 2
    System.out.println(buf.writableBytes());    // 2 
}
```

#### 바이트 버퍼 풀링
- 바이트 버퍼 풀을 사용하면 버퍼를 할당하고 해제할 때 일어나는 가비지 컬렉션의 횟수를 감소할 수 있음

#### 바이트 버퍼 상호 변환
```java
public void NettyByteBufferFuction3() {
    String source = "hello world";

    ByteBuf buf = Unpooled.buffer(11);
    buf.writeBytes(source.getBytes());

    // 네티 바이트 버퍼를 자바 NIO 버퍼로 변환
    ByteBuffer nioByteBuffer = buf.nioBuffer();
    // 출력 확인
    System.out.println(new String(
        nioByteBuffer.array(),
        nioByteBuffer.arrayOffset(),
        nioByteBuffer.remaining()));        // "hello world"


    // 자바 NIO 버퍼를 네티 바이트 버퍼로 변환
    ByteBuffer byteBuffer = ByteBuffer.wrap(source.getBytes());
    ByteBuf nettyBuffer = Unpooled.wrappedBuffer(byteBuffer);
    // 출력 확인
    System.out.println(nettyBuffer.toString(Charset.defaultCharset())); // "hello world"
}
```

#### 채널과 바이트 버퍼 풀
- channelRead 메서드의 인수로 사용되는 바이트 버퍼는 네티 바이트 버퍼
- channelRead 메서드가 실행된 이후의 네티 바이트 버퍼는 바이트 버퍼 풀로 돌아감

```java
public class EchoServerHandler extends ChannelInboundHandlerAdapter {

    @Override
    public void channelRead(ChannelHandlerContext ctx, Object msg) {
        
        ByteBuf readMessage = (ByteBuf) msg;
        System.out.println(readMessage.toString(Charset.defaultCharset()));

        // ChannelHandlerContext를 통해서 네티 프레임워크에 초기화된 ByteBufAllocator 참조
        // ByteBufAllocator는 바이트 버퍼 풀을 관리하는 인터페이스
        // (설정에 따라 힙 또는 다이렉트 버퍼 풀 생성, 기본적으로 다이렉트 버퍼 풀 생성)
        ByteBufAllocator byteBufAllocator = ctx.alloc();

        // ByteBufAllocator의 buffer 메서드를 사용하여 생성된 바이트 버퍼는 ByteBufAllocator의 풀에서 관리
        // (release 메서드를 호출하면 버퍼 풀로 돌아감)
        ByteBuf newBuffer = byteBufAllocator.buffer();


        // newBuffer 사용


        // write 메서드의 인수로 버퍼가 입력되면 데이터를 채널에 기록하고 버퍼 풀로 돌아감
        ctx.write(msg);
    }

    @Override
    public void exceptionCaught(ChannelHandlerContext ctx, Throwable cause) {
        cause.printStackTrace();
        ctx.close();
    }
}
```

### 정리
- 네티로 애플리케이션을 작성할 때 반드시 네티 바이트 버퍼를 사용해야 하는 것은 아님
- 하지만 더 나은 성능을 제공하기 위해 네티 바이트 버퍼를 사용하는 것이 이득
- 자바 바이트 버퍼를 사용할 경우 반드시 호출해야 하는 flip 메서드를 호출하지 않아도 되기 때문에 애플리케이션 버그 발생률을 많이 낮출 수 있음
- 네티 바이트 버퍼 풀을 사용하여 가비지 컬렉션 빈도를 낮추며 더 빠르고 안정적인 애플리케이션을 개발할 수 있음

---

## Reference
- [자바 네트워크 소녀 Netty 소스 코드 예제](https://github.com/krisjey/netty.book.kor)
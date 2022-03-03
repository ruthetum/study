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

### `channel` : 소케 입출력 모드 설정

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

> SCTP(Stream Control Transmission Protocol)
> - 스트림 제어 전송 프로토콜
> - 전송 계층의 프로토콜로 UDP 메시지 스트리밍 특성과 TCP의 연결형 및 신뢰성 제공 특성을 조합한 프로토콜
> - TCP처럼 연결지향적 프로토콜이며[2] 혼잡 제어를 통해 신뢰성 있는 순차적 메시지 전송을 보장
> - [WIKI](https://ko.wikipedia.org/wiki/%EC%8A%A4%ED%8A%B8%EB%A6%BC_%EC%A0%9C%EC%96%B4_%EC%A0%84%EC%86%A1_%ED%94%84%EB%A1%9C%ED%86%A0%EC%BD%9C)


## Reference
- [자바 네트워크 소녀 Netty 소스 코드 예제](https://github.com/krisjey/netty.book.kor)
# Netty
- Netty는 프로토콜 서버 및 클라이언트와 같은 네트워크 애플리케이션의 빠르고 쉬운 개발을 가능하게 하는 NIO 클라이언트 서버 프레임워크
- 기존의 소켓 프로그래밍은 클라이언트가 접속하게 되면 스레드를 할당해야 하는데(1:1관계), 정말 많은 클라이언트가 접속을 하게 될 경우 그 숫자만큼 스레드를 생성해야 해서 리소스의 낭비로 이루어지고, 문맥 교환과 관련된 문제와 입력이나 출력 데이터에 관련한 무한 대기 현상이 발생하는 문제가 있다.
- 이러한 네트워크 문제 때문에 개발된 방법이 자바의 NIO 방식(Non-Blocking Input Ouput)이다.
    - <b>비동기 네트워크 통신</b>
- Tomcat서버가 10,000건의 커넥션을 처리한다면, netty는 100,000건 이상의 커넥션을 처리할 수 있는 장점이 있다.

---

## Content
- [자바 네트워크 소녀 Netty 정리 및 실습](./java-network-girl-netty/)
- [Netty 컴포넌트 정리](./netty-component.md)

---

## Reference
- https://netty.io/
- https://github.com/netty/netty
- [자바 네트워크 소녀 Netty](http://www.kyobobook.co.kr/product/detailViewKor.laf?ejkGb=KOR&mallGb=KOR&barcode=9788968482243&orderClick=LAG&Kc=)
- SpringBoot + Netty TCP 소켓 서버 : https://i-hope9.github.io/2020/12/14/SpringBoot-Netty-2-SocketServer.html
- **a tour of netty** 👍🏻: https://medium.com/geekculture/a-tour-of-netty-5020ecee5494

---

### Spring
- WebFlux Docs : https://docs.spring.io/spring-framework/docs/current/reference/html/web-reactive.html#webflux
    - WebFlux Default Decoder `DecoderHttpMessageReader` : https://docs.spring.io/spring-framework/docs/current/javadoc-api/org/springframework/http/codec/DecoderHttpMessageReader.html
- Reactor Netty Configuration : https://www.baeldung.com/spring-boot-reactor-netty
- Concurrency in Spring WebFlux : https://www.baeldung.com/spring-webflux-concurrency
- How to build TCP server with SpringBoot : https://programmer.help/blogs/spring-boot-build-tcp-server.html

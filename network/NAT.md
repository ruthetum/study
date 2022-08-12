# NAT (Network Address Translation)

- 사설 IP를 공인 IP로 변경할 때 필요한 주소 변환 서비스

![image](https://user-images.githubusercontent.com/59307414/184296268-8e606d51-657a-40de-b8b5-de6bf17e90f8.png)

- 내부 네트워크에 위치한 호스트들의 사설 IP와 포트 번호에 대한 정보를 가지고 있음
- 외부로 나갈 때의 동일한 공인 IP와 각기 다른 포트 번호를 가지고 있음
- 목적지 주소의 공인 IP와 서비스에 사용된 동일한 서비스 포트 번호를 가지고 있음

### Flow

> 위 테이블에서 첫 번째 호스트(192.168.1.1:9688)이 외부 호스트(68.1.31.1:23)에 접속할 때 시나리오를 가정
> 
1. 사설 IP(192.168.1.1)를 사용하는 호스트가 68.1.31.1에 접속을 원한다.
2. 라우터는 NAT를 사용해서 공인 IP(101.89.101.12)와 포트 번호 8801을 할당한다.
3. 요청은 인터넷을 통해 외부 호스트(68.1.31.1:23)에 도달한다.
4. 외부 호스트는 요청의 클라이언트 IP를 공인 IP(101.89.101.12), 포트 8801로 인식하고 응답한다.
5. 외부 호스트의 응답이 라우팅에 도착한다.  라우터는 포트 번호 8801을 응답 패킷에서 인식한다.
6. 라우터는 사설 IP(192.168.1.1:9688)의 호스트에게 할당된 포트임을 NAT 테이블에서 찾아 응답을 돌려준다.

### 라우터 A에 연결된 디바이스 B, 디바이스 C가 각각 호스트 D에 접속할 때 호스트 D는 디바이스 B와 C를 어떻게 구분할 수 있을까?

→ 동일한 공인 IP를 갖지만 각기 다른 포트 번호를 가지고 있기 때문에 이를 통해 디바이스를 구분할 수 있다.

## NAT 사용 목적

1. **공인 IP 주소 절약**
    - 공인 IP 주소는 한정적이기 때문에 모든 디바이스가 고유한 공인 IP을 할당한다면 자원이 고갈됨
    - 따라서 사설 IP를 통해 자원을 절약
2. 보안
    - 공개된 주소는 외부에서 공격 당할 확률이 존재
    - 따라서 내부망과 공개망 사이에 방화벽을 운영하여 외부 공격으로부터 내부망을 보호

## Static NAT (1:1 NAT)

- IP 주소들을 1:1 매핑으로 변환해주는 간단한 NAT 유형
- 호환되지 않는 주소 체계를 가지고 있는 두 개의 IP 네트워크가 서로 통신해야할 때 사용될 수 있음
- Static NAT 명령어를 입력하자마자 NAT 변환 테이블에 변환 정보가 입력

![image](https://user-images.githubusercontent.com/59307414/184296310-9b122afe-1a24-4f25-86ec-c23261e37242.png)

## Dynamic NAT

- Static NAT처럼 NAT 라우터가 사설 IP와 공인 IP 간의 매핑 정보를 생성
- IP 패킷이 네트워크에서 나가거나 들어올 때 패킷 헤더의 IP 주소를 변환
- 이 때 사설 IP와 공인 IP를 매핑하는 것을 정적인 정보에 기반하지 않고, 동적으로 수행
    - 사용 가능한 공인 IP 주소들의 pool을 만들어 두고, 어떤 사설 IP 주소들이 NAT에 의해 주소 변환이 되어야 하는지 결정하기 위해 매핑 기준을 정의
    - Dynamic NAT에서는 주소 변환이 필요한 트래픽을 라우터가 받기 전까지 NAT 테이블에 변환 정보가 존재하지 않음

![image](https://user-images.githubusercontent.com/59307414/184296779-57f23a9c-9d9e-4cee-aace-1e3efef2d2ad.png)

## Cone NAT vs Symmetric NAT

- 공유기에 사용되는 NAT 방식에는 크게 Cone NAT과 Symmetric NAT가 있다.

### Cone NAT

- 특정 PC의 내부의 IP:Port가 외부 IP:Port로 변환될 때 Destination에 상관 없이 외부 Port가 항상 일정
- Cone NAT에는 다시 Full cone NAT과 (Address) Restricted cone NAT과 Port Restricted cone NAT의 3 종류가 존재

**Normal NAT (Full Cone NAT)**

> A full cone NAT is one where all requests from the same internal IP address and port are mapped to the same external IP address and port. Furthermore, any external host can send a packet to the internal host, by sending a packet to the mapped external address.
> 
- 동일한 내부 IP 주소 및 포트의 모든 요청이 동일한 외부 IP 주소 및 포트에 매핑
- PC1에서 공유기를 거쳐 IP1와 Port1로 패킷이 나가면, 외부에서 IP1와 Port1로 들어오는 모든 패킷은 PC1로 전달

![image](https://user-images.githubusercontent.com/59307414/184296451-f8590935-1848-4049-8e3a-ccdfd94fb704.png)

---

**Restricted Cone NAT (Address Restricted Cone NAT)**

> A restricted cone NAT is one where all requests from the same internal IP address and port are mapped to the same external IP address and port. Unlike a full cone NAT, an external host (with IP address X) can send a packet to the internal host only if the internal host had previously sent a packet to IP address X.
> 
- 외부 호스트(IP 주소 X)는 내부 호스트가 이전에 IP 주소 X로 패킷을 보낸 경우에만 내부 호스트로 패킷을 보낼 수 있음
- PC1에서 외부 서버로 패킷을 보낼 때 서버 주소와 포트는 IP2와 Port2인 경우 서버에서 PC1으로 보내는 패킷의 Source는 IP2이다. 서버에서 보내는 Source 주소 IP2의 패킷이 공유기에 들어올 때 PC1으로 전달된다. 포트는 아무 포트가 되어도 무방하고 Port2일 필요가 없다.

![image](https://user-images.githubusercontent.com/59307414/184296486-4139c527-2c51-49e5-aece-23bfccff1c88.png)

---

**Port Restricted Cone NAT**

> A port restricted cone NAT is like a restricted cone NAT, but the restriction includes port numbers. Specifically, an external host can send a packet, with source IP address X and source port P, to the internal host only if the internal host had previously sent a packet to IP address X and port P.
> 
- Address Restricted Cone NAT와 유사하지만 포트번호까지도 제한함
- 외부 호스트는 내부 호스트가 이전에 IP 주소 X와 포트 P로 패킷을 보낸 경우에만 소스 IP 주소 X와 소스 포트 P를 가진 패킷을 내부 호스트로 보낼 수 있음

![image](https://user-images.githubusercontent.com/59307414/184296531-6bb5c3ea-f502-428c-b6d8-868b83dd78e9.png)

---

### Symmetric NAT

- 특정 PC의 내부의 IP:Port가 외부 IP:Port로 변환될 때 Destination에 따라 다른 외부 Port가 사용

**Symmetric NAT**

> A symmetric NAT is one where all requests from the same internal IP address and port, to a specific destination IP address and port, are mapped to the same external IP address and port. If the same host sends a packet with the same source address and port, but to a different destination, a different mapping is used. Furthermore, only the external host that receives a packet can send a UDP packet back to the internal host.
> 
- 동일한 내부 IP 주소 및 포트에서 특정 대상의 IP 주소 및 포트에 대한 모든 요청이 동일한 외부 IP 주소 및 포트에 매핑
    - 패킷을 보내는 외부 서버마다 다른 NAT 매핑을 사용
- 동일한 호스트가 동일한 소스 주소 및 포트로 다른 대상으로 패킷을 보내는 경우 다른 매핑이 사용
- PC에서 패킷을 특정 서버로 보내면 그 서버에서 보낸 패킷만 PC로 전달

![image](https://user-images.githubusercontent.com/59307414/184296587-7221f43f-b776-4c71-a49c-13b9bd228053.png)

---

### Symmetric NAT를 사용하는 경우 공인 IP 주소를 발견해도 연결이 불가능한 이유

- Cone 방식은 Router와 매핑이 이루어지면 통신할 때 계속해서 한번 매핑된 정보를 쓴다.
- 하지만 Symmetric 방식은 목적지(외부 호스트의 IP:Port)를 기반으로 매핑 정보를 다르게 한다.
- 즉, 호스트가 통신하는 대상에 따라 각각 다른 매핑 정보를 가져가는 것이고, 동일한 목적지가 아니라면 매핑 정보가 다르다.
- 따라서 TURN 서버를 이용해서 우회한다.

![image](https://user-images.githubusercontent.com/59307414/184296625-bac0b63a-7805-4ab8-9089-7cfed103e34f.png)

---

### cf 1.  IP Masquerade (MASQ)

- NAT 안에 속해 있는 기술
    
    > 예를 들어 클라이언트 2명 A(192.168.0.1), B(192.168.0.2)가 하나의 라우터에 있는 경우 클라이언트들은 공인 IP를 가지고 있지 않기에 외부 통신을 할 수 없다.
    > 
- 이런 문제를 보안하기 위해 IP 마스커레이드가 생긴 것인데 사설 대역 IP를 가진 클라이언트들이 패킷을 전송하면 라우터에서 목적지에 모두 자신이 보낸 것처럼하여 패킷을 전송한다.
    - 클라이언트 A가 패킷을 전송하면 발신자가 192.168.0.1에서 라우터를 거치면서 공인 IP로 변환이 되서 나가게 된다.
    - 수신측에서는 클라이언트의 사설 IP가 아닌 공인 IP가 보낸 것으로 받게 된다.
- IP Masquerade는 주소 변환뿐만 아니라 포트 번호까지 포워딩한다.
    - 클라이언트가 외부로 패킷을 보낼 때 라우터에서 공인 IP에 무작위로 포트번호를 붙여서 보내준다.
    - 외부에서 응답 패킷을 보낼때 전송받은 포트번호로 패킷을 보낼 수 있도록 하고 해당 응답패킷을 받은 라우터는 포트 번호를 이용해 클라이언트에게 패킷을 전달한다.

### cf 2. PAT (Port Address Translation)

- Dynamic NAT의 한 종류로, 공인 IP 주소 1개에 사설 IP 주소 여러개를 매핑하는 것
- 변환된 IP 주소로는 사내망 호스트들을 구분할 수 없기 때문에 포트번호를 부여하여 구분한다. 대부분의 홈 네트워크는 PAT를 사용하고 있다.
- ISP(Internet Service Provider)는 홈 네트워크 라우터에 하나의 공인 IP 주소를 할당한다.
- 컴퓨터 A가 인터넷에 로그온하면 라우터는 사설 주소에 포트 번호를 붙여서 유일한 주소를 할당한다.
- 그 다음에 컴퓨터 B가 인터넷에 로그온하면 컴퓨터 A에게 할당한 똑같은 사설 주소에 다른 포트 번호를 붙여서 유일한 주소를 또 만들어서 할당

### Reference

- [https://en.wikipedia.org/wiki/Network_address_translation](https://en.wikipedia.org/wiki/Network_address_translation)
- [https://docs.microsoft.com/ko-kr/azure/rtos/netx-duo/netx-duo-nat/chapter1](https://docs.microsoft.com/ko-kr/azure/rtos/netx-duo/netx-duo-nat/chapter1)
- [https://brunch.co.kr/@sangjinkang/61](https://brunch.co.kr/@sangjinkang/61)
- [http://www.ibiblio.org/pub/linux/docs/howto/other-formats/html_single/IP-Masquerade-HOWTO.html#IPMASQ-INTRO1.1](http://www.ibiblio.org/pub/linux/docs/howto/other-formats/html_single/IP-Masquerade-HOWTO.html#IPMASQ-INTRO1.1)
- [https://nsinc.tistory.com/100#:~:text=IP Masquerade](https://nsinc.tistory.com/100#:~:text=IP%20Masquerade)
- [https://lascrea.tistory.com/105](https://lascrea.tistory.com/105)
- [https://hyolo.tistory.com/47](https://hyolo.tistory.com/47)
- [https://better-together.tistory.com/124](https://better-together.tistory.com/124)
- [https://dh2i.com/kbs/kbs-2961448-understanding-different-nat-types-and-hole-punching/#:~:text=A symmetric NAT is one,a different mapping is used](https://dh2i.com/kbs/kbs-2961448-understanding-different-nat-types-and-hole-punching/#:~:text=A%20symmetric%20NAT%20is%20one,a%20different%20mapping%20is%20used).
- [https://tomatohj.tistory.com/42](https://tomatohj.tistory.com/42)
- [https://lovejaco.github.io/posts/webrtc-connectivity-and-nat-traversal/](https://lovejaco.github.io/posts/webrtc-connectivity-and-nat-traversal/)
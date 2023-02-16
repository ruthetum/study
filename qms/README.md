# QMS (Queue Management System)
대기열 시스템, 접속자 대기 시스템

- 처리율 제한 장치, 트래픽 유량 제어와 유사
- 스로틀링(Throttling)을 통해 차단만 하는 것이 아니라 입장까지 제어

## 대기열 시스템?
> 단말이 직접 복수의 서버 중 서버 부하를 고려하여 하나의서버를선택하고,
>
> 선택된 서버를 통해 접속시간 순으로 자신의 접속대기자에 대한 접속처리를 수행하되,
>
> 서버마다 서로 다른 수의 접속대기 이탈자가 발생하는 환경에서도
> 
> 모든 서버에 접속요청한 접속대기자들에 대한 접속처리를 보다 공정하게 수행할 수 있도록 해주는 기술

흔히 생각되는 티켓팅, 이벤트 사이트 혹은 수강신청 사이트 같이 순간적으로 트래픽이 몰릴 때 해당 트래픽을 제어하는 시스템

### 대기열 시스템 배경
- 한 서버에 들어올 수 있는 동접 수의 제한이 있다.
- 서버를 늘리는 것은 시간과 자원이 필요하다.
- 필연적으로 서버에 접속할 수 없는 사용자가 발생한다.
- **넘치는 사용자를 단순히 튕겨내는 것보다는 뭔가 알려주면서, 사용자가 접속할 수 있을 때까지 기다리게 해야 한다.**

### 대기열 시스템 목표
- 기존 시스템에 부하를 주지 않고 사용자들을 기다릴 수 있게 해야함
- 기존 시스템보다 안정적이고 빨라야 함
- 오래 기다린 사람이 가급적이면 먼저 접속해야 함 (선입선출)
- 기다리는 동안 사용자에게 유용한 정보를 제공해야 함 (대기자 수, 예상 대기 시간)

## 대기열 시스템 적용 및 구현 사례
### 1. 처리율 제한 장치의 설계
- [가상 면접 사례로 본 처리율 제한 장치의 설계](./처리율-제한-장치-설계.md)

### 2. NDC2019 - 실버바인 대기열 서버 설계 리뷰
- [실버바인 대기열 서버 설계 리뷰](./실버바인-대기열-서버-설계-리뷰.md)

### 3. 지마켓 대기열 시스템
- [지마켓 대기열 시스템 리뷰](./지마켓-대기열-시스템.md)

## 프로토타입
- [프로토타입](./prototype)

## Reference
- [NDC2019 - 실버바인 대기열 서버 설계 리뷰](http://ndcreplay.nexon.com/NDC2019/sessions/NDC2019_0069.html)
  - 프레젠테이션 자료: http://ndc.vod.nexoncdn.co.kr/NDC2019/slides/NDC2019_0069/index.html
- [G마켓 대기열 시스템](https://dev.gmarket.com/46)
- [대기열 시스템 구현하기 with Spring, Redis, WebSocket](https://dev-jj.tistory.com/entry/%ED%94%84%EB%A1%9C%EB%AA%A8%EC%85%98%EC%9D%84-%EB%8C%80%EB%B9%84%ED%95%9C-%EB%8C%80%EA%B8%B0%EC%97%B4-%EC%8B%9C%EC%8A%A4%ED%85%9C-%EA%B5%AC%EC%84%B1%ED%95%98%EA%B8%B0-Redis-WebSocket-Spring?category=828965)
- [웹 서비스 대기열 서버 구현을 위한 고찰](https://moonsiri.tistory.com/156)
- [데브와이 - 대기열 시스템 솔루션 판매](https://devy.kr/)
  - [신세계 적용 사례](https://ssgmsp.com/news/2020-07-16-News-INC-Devy-kr/)
- [How to Control The Flow of Waiting Lines Effectively](https://blog.timify.com/control-waiting-lines-with-queue-management-systems/)
- [넷퍼넬](https://netfunnel.io/)
  - 대학교에서 수강신청 사이트 트래픽 제어하는데 많이 사용 (ex. 성균관대학교, 고려대학교)
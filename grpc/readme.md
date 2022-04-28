# gRPC
- 분산된 이기종 애플리케이션을 연결, 호출, 운영, 디버깅할 수 있는 프로세스 간 통신 기술

- **프로토콜 버퍼**를 IDL(Interface Definition Language)로 사용해서 서비스 인터페이스를 정의

> 프로토콜 버퍼
>
> : 구조화된 데이터를 직렬화할 수 있는 오픈소스, 데이터 저장용이나 통신용 프로그램에 주로 사용되며, 프로토버프(protobuf)로 불림
> - Google Protocl Buffer : https://developers.google.com/protocol-buffers
> - Google Protocl Buffer - Java tutorial : https://developers.google.com/protocol-buffers/docs/javatutorial
> - Protocl Buffer Github : https://github.com/protocolbuffers
> - 프로토콜 버퍼? : https://bcho.tistory.com/1182
> - Java로 프로토콜 버퍼 구현해보기 : https://velog.io/@jakeseo_me/Java%EB%A1%9C-Google-Protocol-Buffer-%EA%B5%AC%ED%98%84%ED%95%B4%EB%B3%B4%EA%B8%B0

<br>

## Pros and Cons
<details>
<summary>장점</summary>
<div markdown="1">

### 프로세스 간 통신 효율성
- gRPC는 JSON이나 XML 같은 텍스트 형식을 사용하는 대신 프로토콜 버퍼 기반 바이너리 프로토콜을 사용하여 gRPC 서비스 및 클라이언트와 통신
    - 클라이언트에서 서버로 메시지를 전송할 때 REST API를 사용한다면 1) 바이너리 콘텐츠를 만들고, 2) 바이너리 콘텐츠를 텍스트 형식으로 변환하여 전송, 3) 서버에서 텍스트 형식의 콘텐츠를 다시 바이너리 콘텐츠로 변환
- 추가로 HTTP/2 위에 프로토콜 버퍼로 구현되기 때문에 통신 속도 향상

### 간단하고 명확한 서비스 인터페이스와 스키마
- gRPC는 애플리케이션 개발용 계약 유선(Contract-First) 접근 방식을 권장
    - 서비스 인터페이스를 먼저 정의하고 나중에 구현 세부 사항을 작업

- 정의된 서비스 인터페이스를 바탕으로 개발하기 때문에 일관되고 안정적인 확장 가능한 애플리케이션 개발 환경 제공

### 엄격한 타입 점검 형식
- gRPC 서비스를 정의하고자 프로토콜 버퍼를 사용하기 때문에 gRPC 서비스 계약은 애플리케이션 통신에 사용할 데이터 타입을 명확하게 정의

### 폴리글랏
- 특정 언어에 구애 받지 않고 사용 가능

### 이중 스트리밍(Duplex Streaming)
- gRPC는 클라이언트나 서버 측 스트리밍을 기본적으로 지원하며 서비스 정의 자체에 포함되기 때문에 스트리밍 서비스나 스트리밍 클라이언트 개발에 용이

- 요청-응답 스타일의 메시징 방식이나 클라이언트 및 서버 측 스트리밍을 구축할 때 RESTful 메시징 스타일에 비해 좋음

### 유용한 내장 기능 지원
- 인증, 암호화, 복원력(데드라인, 타임아웃), 메타데이터 교환, 압축, 로드밸런싱, 서비스 검색 등과 같은 기능을 지원

---

</div>
</details>

<details>
<summary>단점</summary>
<div markdown="1">

### 외부 서비스 부적합
- 인터넷을 통해 애플리케이션이나 서비스를 외부 클라이언트에 제공하려는 경우 gRPC가 새롭기 때문에 적합하지 않을 수 있음

- 계약 기반이면서 엄격한 타입을 정의하기 때문에 외부 당사자에게 노출되는 서비스의 유연성을 방해할 수 있으며, 사용자는 비교적 적은 제어권을 갖음

> 이에 대한 해결책으로 gRPC Gateway가 설계됨

### 서비스 정의의 급격한 변경에 따른 개발 프로세스 복잡성
- 미리 정의된 인터페이스에서 기능을 지원하기 때문에 서비스 정의가 급격히 변경되면 클라이언트와 서비 코드 모두 다시 생성해야 함

- 이 과정이 CI 프로세스에도 통합돼야 하기 때문에 전체 개발 수명 주기를 복잡하게 만들 수 있음

> 하지만 대부분의 gRPC 서비스 정의 변경은 서비스 계약을 위반하지 않게 수용될 수 있고, 주요 변경 사항이 없는 한 다른 버전의 프로토를 사용하여 클라이언트 및 서버으와 문제없이 상호 운영할 수 있음
>
> 즉, 대부분의 경우 코드 재생성이 필요하지 않음 

### 상대적으로 적은 생태계
- REST나 HTTP 프로토콜에 비해 상대적으로 생태계 크기가 적음

- 브라우저와 모바일 애플리케이션에서 gRPC의 지원도 아직 초기 단계이기 때문에 모든 프로세스 간 통신에 필수적으로 사용해야 하는 기술은 아님

- 비즈니스 유스 케이스나 요구 사항에 따라 적절한 메시징 프로토콜을 사용해야 함

---

</div>
</details>

<br>

## Settings
- *Go* 언어 설치  
    최신 버전의 Go 언어는 [Go 공식 사이트](https://golang.org/dl/)에서 다운로드 받아 설치합니다. 설치에 대한 설명은 공식 문서([https://golang.org/doc/install](https://golang.org/doc/install))를 참조하세요.
 
- *Java* 설치  
    예제들은 JDK 1.8 기준으로 작성 및 테스트되었으며 [Java 공식 사이트](https://www.java.com/en/download/)에서 다운로드 및 설치합니다.

- *Gradle* 설치  
    Java 예제를 실행하기 위해서는 [Gradle](https://gradle.org/) 설치가 필요합니다.

- *protoc* 설치  
    gRPC는 Protocal buffer 기반으로 [protoc](https://developers.google.com/protocol-buffers/docs/downloads) 설치가 필요합니다.

<br>

## Books (gRPC 시작에서 운영까지)
- 2장. [gRPC 시작](./books/chapter2)
- 3장. [gRPC 통신 패턴](./books/chapter3)
- 4장. [gRPC: 동작 원리](./books/chapter4)
- 5장. [gRPC: 고급 기능](./books/chapter5)
- 6장. [보안 적용 gRPC](./books/chapter6)
- 7장. [서비스 수준 gRPC 실행](./books/chapter7)
- 8장. [gRPC 생태계](./books/chapter8)

<br>

## Reference
- [gRPC 시작에서 운영까지](http://www.acornpub.co.kr/book/grpc)
- https://github.com/switchover/grpc-up-and-running
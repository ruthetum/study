# Apache vs Nginx

> Web Server
> - HTTP를 통해 웹 브라우저에서 요청하는 HTML 문서나 오브젝트(이미지 파일 등)을 전송해주는 서비스 프로그램을 말한다.
> - Apache, Nginx, etc. 

## Overview
- Apache와 Nginx는 웹 서버를 운영할 수 있는 오픈 소스 서버 기술을 제공

- Apache는 오랫동안 이용되어 왔고, 최근에는 Nginx가 점유가 높아짐

    - 최근 점유율

        ![possesion](https://w3techs.com/diagram/history_overview/web_server/ms/m)

- 웹 서버의 선택은 인터넷 트래픽에 영향을 주기 때문에 상황에 알맞는 선택이 필요

- 두 기술은 작동 방식이 다르기 때문에 각각의 작동 방식 및 장단점을 알고 사용하자

<br>

|Apache|Nginx|
|:---:|:---:|
|<img src="https://www.shost.vn/wp-content/uploads/2022/01/apache-webserver.png" width="300" height="75">|<img src="https://blog.kakaocdn.net/dn/bbdjJa/btroxBROmfE/2zNE70TvmQ7Lnzh6U9BvMk/img.jpg" width="300" height="75">|

### Apache
- Apache HTTP는 1995년부터 많이 사용되어진 웹 서버로 HTTP 표준을 준수하도록 구축된 오픈 소스이며, 고성능 웹 서버

### Nginx
- Apache의 C10K 문제점 해결을 위해 만들어진 Event-Driven 구조의 웹 서버
    - C10K : 일만개의 클라이언트 문제, 한 시스템에 동시 접속자수가 1만명이 넘어갈 때 효율적 방안


## Compairson
### 설계 구조
- 클라이언트의 요청을 처리하고 응답하는 방식의 차이

#### Apache
- 프로세스 기반 접근 방식으로 하나의 스레드가 하나의 요청을 처리하는 구조
- 매 요청마다 스레드를 생성 및 할당해야 하기 때문에 리소스를 많이 잡아먹음

![thread programming](https://img1.daumcdn.net/thumb/R1280x0/?scode=mtistory2&fname=https%3A%2F%2Fblog.kakaocdn.net%2Fdn%2FcFcSPQ%2FbtrcB4KXbHM%2Fg4FqyuDKYEHLIxozKYHDn0%2Fimg.png)

#### Nginx
- 이벤트 중심 접근 방식으로 하나의 스레드 내에서 여러 요청을 처리하는 구조
    - 비동기 Event-Driven 구조 : Event Handler에서 비동기 방식으로 먼저 처리되는 요청을 진행 
- 코어 모듈이 Apache보다 적은 리소스로도 많은 트래픽을 효율적으로 처리 가능

![event driven programming](https://img1.daumcdn.net/thumb/R1280x0/?scode=mtistory2&fname=https%3A%2F%2Fblog.kakaocdn.net%2Fdn%2FbAbGyw%2FbtrcGUADagZ%2FEvavl3JpiSZbPq9T3Hzm01%2Fimg.png)

#### cf.
- 적은 스레드가 사용되면 CPU 소모가 적고, context switching 비용이 감소

- Spring에서 Sevlet 기반 MVC(멀티 스레드)와 WebFlux(이벤트 루프) 비교할 때와 유사

- 대용량 트래픽을 처리하기 위해 가벼움과 높은 성능을 목표로 한다면 Nginx가 적합

<br>

### 성능 차이
#### 정적 컨텐츠
- Apache는 서버 컴퓨터의 디스크에 저장하는 파일 기반 방법을 사용하여 정적 콘텐츠를 처리
- Nginx는 설계 아키텍처 구조상 Nginx가 적은 비용으로 효율적으로 제공
- Nginx는 최대 1,000개의 동시 연결을 실행하는 벤치마크 테스트에 따르면 Apache보다 2.5배 더 빠른 성능을 발휘

#### 동적 컨텐츠
- 두 웹 서버의 성능이 비슷함
- Apache는 외부 구성 요소에 의존할 필요 없이 웹 서버 자체 내에서 동적 컨텐츠를 처리할 수 있음
- Nginx는 동적 컨텐츠를 웹 서버 내에서 처리하지 않지만 SCGI 핸들러와 FastCGI 모듈을 사용해서 동적 컨텐츠 제공할 수 있음

<br>

### OS 지원 여부
#### Apache
- Linux 및 BSD를 포함한 모든 Unix 계열 OS 지원
- Windows 모두 지원

#### Nginx
- 거의 모든 Unix 계열 OS 지원
- Windows는 부분적으로 지원

<br>

### 중앙 집중/분산 구조
#### Apache
- .htaccess 파일을 통해 디렉토리 별로 추가 구성을 허용
- 이로 인해 권한이 없는 사용자가 웹 사이트의 특정 측면을 제어할 수 있음

#### Nginx
- 추가 구성을 허용하지 않음
- 권한이 없는 사용자가 웹 사이트의 특정 측면을 제어할 수 없지만 추가 구성을 제공하지 않음으로써 성능 향상
- 디렉토리 구성을 허용하지 않음으로 .htaccess 파일을 검색하고 사용자가 만든 요구 사항을 해석할 필요 없기 때문에 Apache보다 빠르게 요청을 처리할 수 있음
 
### 요청 처리 및 해석하는 방법의 차이
#### Apache
- 요청을 해석하기 위해 파일 시스템 위치 전달
- URI 위치를 사용하지만 일반적으로 더 추상적인 디렉토리 구조를 사용

#### Nginx
- 요청을 해석하기 위해 URI를 전달
- URI로 전달함으로써 웹 서버뿐만 아니라 프록시 서버, 로드 밸런서 및 HTTP 캐시로 쉽게 동작 가능
- 서버에서 클라이언트로 데이터가 전송되는 속도가 Apache보다 더 빠름

<br>
 
### 기능 모듈의 차이
#### Apache
- 동적으로 로드 가능한 다양한 60개의 공식 모듈을 제공
모든 모듈을 가지고 서버가 실행되지만 실제 사용되는 모듈은 소수임 = 무거움

#### Nginx
- 타사 플러그인 과정으로 선택되고 컴파일되기 때문에 동적으로 모듈을 로드할 수 없음
- 사용하려는 기능만 선택해서 서버를 실행 = 가벼움

<br>

### 유연성
#### Apache
- 동적 모듈과 로딩을 지원함

#### Nginx
- 아직까지는 동적 모듈과 로딩을 지원하지 않음

### 보안
- 두 웹 서버 모두 C언어 기반으로 확장된 보안을 제공
- 하지만 Nginx가 코드가 더 작기 때문에 미래 지향적인 보안 관점에서 장점을 가짐(비슷하지만 Nginx가 조금 더 안전한 것으로 간주)

<br>

![compairson](https://serverguy.com/wp-content/uploads/2020/10/Apache-Vs-NGINX-Infographic_Edit-2_New-SG-Logo.png)

### Summary

<table>
    <thead>
        <tr>
            <th>항목</th>
            <th>Apache</th>
            <th>Nginx</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>작동 방식</td>
            <td>프로세스 기반 접근 방식<br>(멀티 스레드 or 멀티 프로세스)</td>
            <td>이벤트 중심 접근 방식<br>(단일 스레드)</td>
        </tr>
        <tr>
            <td>성능 차이<br>(정적 컨텐츠)</td>
            <td colspan="2">Nginx가 효율적</td>
        </tr>
        <tr>
            <td>성능 차이<br>(동적 컨텐츠)</td>
            <td colspan="2">비슷함<br>(다만 Nginx는 내부에서 처리하지 못 하고 핸들러를 통해 처리)</td>
        </tr>
        <tr>
            <td>OS 지원 범위</td>
            <td colspan="2">거의 비슷하지만 조금 다름<br>(Apache가 좀 더 넓게 지원)</td>
        </tr>
        <tr>
            <td>요청 해석</td>
            <td>파일 디렉토리 구조 접근</td>
            <td>URI 전달</td>
        </tr>
        <tr>
            <td>무게</td>
            <td>사용하지 않는 모듈도 포함하기 때문에 무거움</td>
            <td>가벼움</td>
        </tr>
    </tbody>
</table>

- Apache는 오랫동안 이용되어 왔기 때문에 신뢰할 수 있고 문제가 발생했을 때 참고할 자료가 Nginx에 비해 많다. 추가로 다양한 동적 모듈을 로드하는데 이점이 있다.

- 하지만 다량의 트래픽이 발생하는 경우에는 Nginx가 안정성과 속도면에서 확실히 효율적이다.

- 추가로 두 웹 서버를 함께 사용해도 된다. Apache 앞단에 Nginx를 프록시 서버로 활용할 수 있다.

    ![both](https://img1.daumcdn.net/thumb/R1280x0/?scode=mtistory2&fname=https%3A%2F%2Fblog.kakaocdn.net%2Fdn%2FR0kWS%2FbtrcEKSVioY%2FQ4Z1g17XQiYA5nVKYCDzF1%2Fimg.png)


## Reference
- https://webinstory.tistory.com/entry/Apache-vs-Nginx-%EB%B9%84%EA%B5%90

- https://sorjfkrh5078.tistory.com/289

- https://velog.io/@ksso730/Nginx-Apache-%EB%B9%84%EA%B5%90

- https://w3techs.com/technologies/history_overview/web_server/ms

- https://serverguy.com/comparison/apache-vs-nginx/

- https://themeisle.com/blog/nginx-vs-apache/
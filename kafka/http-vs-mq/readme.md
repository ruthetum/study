# HTTP vs MQ
## Overview
- [지난 인턴십에서 진행했던 프로젝트](https://github.com/stove-smooth/sgs-smooth)에서 마이크로서비스 간 요청을 MQ를 이용하여 처리하지 못 했다.
- 변명이라면 변명이지만 MSA 및 Spring Cloud에 대한 개념이 조금도 없는 상황에서 진행했던 프로젝트였고, 기본적인 구조 이해 및 설계에 정신이 없었다.
- 추가적으로 WebRTC를 바탕으로 한 시그널링 서버+미디어 서버 구축 부분에 공부 및 구현을 신경쓰다보니 MQ를 공부하고 서비스간 요청에 대해 효율적으로 처리하기보다는 Feign Client를 이용한 단순하게 API 처리를 하게 되었다.
- 당시 적용하지 못 했던 부분을 이번 실험을 통해 공부하고 적용해보면서 성능을 비교할 계획이다. 

---

# Subject
## 상황
- 실험은 간단하게 2가지 상황을 비교할 계획이다.

- 첫 번째 상황은 단순하게 두 개의 서비스에 각각 HTTP 요청을 보내는 상황이고,

    ![image](https://user-images.githubusercontent.com/59307414/155879986-45ea3a60-e9cd-4c16-9500-2d2ab21f5c29.png)

- 두 번째 상황은 kafka에 topic을 publish하고, 해당 토픽을 각 서비스에 subscribe하여 처리하는 상황이다.

    ![image](https://user-images.githubusercontent.com/59307414/155879990-dfb19455-8a50-4c65-b0a5-177d64fce8ae.png)

- 이후 각 상황에 대해서 nGrinder를 이용하여 성능을 비교할 계획이다.

---

# Experiment
## 환경
- Service
    - 모두 Local 환경에서 동작
    - `source`
        - HTTP 요청
        - Producer(publish)
        - port : 8080 실행
    - `target`
        - HTTP 응답
        - Consumer(subscribe)
        - port : 8081 / 8082 실행

- Kafka
    - AWS EC2 t2.micro(CPU 1, 메모리 1GiB)
    - heap size 옵션 : `KAFKA_HEAP_OPTS="-Xmx400m -Xms400m"`

- Test tool
    - nGrinder

## 준비
- [EC2에 Zookeeper, Kafka 설치 및 실행](../usage-install-ec2.md)

    ![image](https://user-images.githubusercontent.com/59307414/155881543-dd325967-2e17-4447-a7cb-36880588bdc7.png)

- [source](./source/), [target](./target/) application 작성

## 결과
### vusers : 10, run count : 5000
<b>HTTP<b>
![image](https://user-images.githubusercontent.com/59307414/155887762-514c1d68-9097-479b-94e9-d2a180587ed9.png)

<b>Kafka<b>
![image](https://user-images.githubusercontent.com/59307414/155887840-22bb132e-5ef6-4654-ae13-f05d282a3d23.png)

---

### vusers : 10, run count : 10000
<b>HTTP<b>
![image](https://user-images.githubusercontent.com/59307414/155888124-37735557-5906-4e8a-9f4d-a3430f59593e.png)

<b>Kafka<b>
![image](https://user-images.githubusercontent.com/59307414/155887973-051edd74-3616-441b-921e-9efd27bff475.png)

---

## Conclusion
- 기본 설정으로 단순 비교해봤을 때 kafka를 활용해서 처리하는 경우가 앞도적으로 성능이 좋다.
- 역시 대용량 서비스에서 서비스간 요청 혹은 파이프라인 처리 시에는 kafka와 같은 MQ를 사용해서 효율적으로 처리하자. 

---
## Reference
- AWS에 카프카 클러스터 설치하기 : https://blog.voidmainvoid.net/325
- EC2 카프카 세팅(heap memory) : https://seulcode.tistory.com/539
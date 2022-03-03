# Spring MVC vs WebFlux

## 시나리오
![image](https://user-images.githubusercontent.com/59307414/156568821-e19254ce-a762-4338-b8e7-e8daf1dcc5f6.png)

- `Order Service`에 요청을 보내면, `Order Service`는 `Item Service`에 Redis에 저장된 값을 감소시키고 감소시킨 값을 반환

- 같은 시나리오에서 아래의 상황을 비교
    1) Spring MVC로만 구성했을 때
    2) Spring WebFlux로만 구성했을 때
    3) Spring WebFlux + Redis-Reactive

> Redis에 값을 저장하는 이유
> - 관계형 데이터베이스에 값을 저장하는 경우 DB connection 과정에서 결국 병목이 발생하기 때문에 MVC와 WebFlux의 성능을 비교하는 것이 무의미해질 수 있음

---

## 준비
- [Spring MVC 기반 소스 코드]
- [Spring WebFlux 기반 소스 코드] + Redis Reactive

## 실험
> Vusers: 10, Run Count: 10000

### MVC + Redis
![image](https://user-images.githubusercontent.com/59307414/156555317-d5b0b0a3-8339-4d77-a493-5795006e31c3.png)

### WebFlux + Redis
![image](https://user-images.githubusercontent.com/59307414/156555293-8f1194f2-560b-4287-a379-eb002b525e54.png)

### WebFlux + Redis-Reactive
![image](https://user-images.githubusercontent.com/59307414/156565946-ef055241-5351-41c5-8b97-800ea35b5077.png)

---

## Conclusion

|구분|TPS|Peak TPS|Mean Test Time(ms)|Run Time|
|---|---|---|---|---|
|MVC + Redis|618.8|788|13.31|00:02:52|
|WebFlux + Redis|952.7|1261|8.00|00:01:51|
|WebFlux + Redis-Reactive|971.7|1271|7.60|00:01:50|

- WebFlux가 MVC보다 동일한 리소스의 상황에서 효율적으로 처리하는 것을 확인할 수 있음
- Redis-reactive를 활용했을 때 그렇지 않을 때보다 일정 부분 성능이 개선됨을 확인할 수 있음
    - 비동기 + 비동기
    - <i>실험을 잘못한건지 dynamic한 차이를 보이지는 못 함. 이후 공부 필요</i>
- 서비스 예제를 `order`, `item`으로 잡아서 상황이 애매할 수 있지만, <b>만약 서비스 제공 기능이 현재 게임의 실시간 랭킹을 보여주는 기능이고, 플레이어의 활동에 따라 점수가 산정돼서 실시간 랭킹을 제공해야 한다면 WebFlux를 이용했을 때 효과적으로 서비스를 제공할 수 있을 것으로 생각됨</b>


## Reference
- NHN Forward - 내가 만든 WebFlux가 느렸던 이유
    - 동영상: https://www.youtube.com/watch?v=I0zMm6wIbRI
    - PDF: https://rlxuc0ppd.toastcdn.net/presentation/%5BNHN%20FORWARD%202020%5D%EB%82%B4%EA%B0%80%20%EB%A7%8C%EB%93%A0%20WebFlux%EA%B0%80%20%EB%8A%90%EB%A0%B8%EB%8D%98%20%EC%9D%B4%EC%9C%A0.pdf
    - map() vs flatMap() : https://madplay.github.io/post/difference-between-map-and-flatmap-methods-in-java
- https://taes-k.github.io/2019/06/28/spring-msa-6/
- https://tech.kakao.com/2018/05/29/reactor-programming/
- https://hantsy.github.io/spring-reactive-sample/data/data-redis.html
- https://umbum.dev/1045
- https://heekim0719.tistory.com/387
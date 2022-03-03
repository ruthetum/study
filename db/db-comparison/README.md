# DB 성능 비교 테스트
- 상황에 맞는 데이터베이스를 선택하기 위해서 각 데이터베이스의 성능 측정
- 기본적으로 관계형 데이터베이스를 활용할 때 성능이 떨어지는 것을 알고있지만 직접 테스트하고 비교해보기 위해 진행
    - 다수의 생성, 조회 작업이 진행될 때 트랜잭션으로 인해 성능이 떨어질 것으로 예상

## 종류
1. MySQL
2. Mongo DB
3. Redis

## 실험 1
> Insert 50000건 (Vuesrs: 10, Run count: 5000)

![image](https://user-images.githubusercontent.com/59307414/156512498-9b5f5683-6bf9-4a61-81a0-f837d979bb1b.png)

![image](https://user-images.githubusercontent.com/59307414/156512441-ba6c61ed-ef33-4343-b0d4-f8ee6ac8e3c5.png)

![image](https://user-images.githubusercontent.com/59307414/156514922-2730868e-7f4b-4172-81ad-0a714b085582.png)

---

|구분|TPS|Peak TPS|Mean Test Time(ms)|Run Time|
|---|---|---|---|---|
|MySQL|459.8|562|20.78|00:01:56|
|Mongo|1075.7|1663|6.73|00:00:54|
|Redis|1180|1937|5.57|00:00:50|

- 단순 삽입 작업에서 Mongo와 Redis의 성능이 우월한 것을 확인할 수 있음
- 다만 현재 실험에서는 50000건을 저장하는 것만을 테스트했지만 다량의 데이터, 영속적인 보관을 고려했을 때 다수의 삽입 및 저장에는 Mongo DB를 선택하는 것이 좋을 것으로 생각

## 실험 2
> Select 50000건 (Vuesrs: 10, Run count: 5000)
> - 특정 아이디를 기준으로 단건 조회

![image](https://user-images.githubusercontent.com/59307414/156516735-5e4057b2-538f-4031-bb63-258fa02eddde.png)

![image](https://user-images.githubusercontent.com/59307414/156516803-a522d8ba-1550-4ee8-83c3-6f12191ce812.png)

![image](https://user-images.githubusercontent.com/59307414/156516859-b868d699-057a-46af-a2b0-f841b8b27558.png)

---

|구분|TPS|Peak TPS|Mean Test Time(ms)|Run Time|
|---|---|---|---|---|
|MySQL|881.2|1196|8.22|00:01:05|
|Mongo|1125|1602|6.65|00:00:51|
|Redis|1377|2040|5.20|00:00:42|

- 역시 조회 시에는 Redis의 성능이 우월하다
- 캐시를 괜히 사용하는 게 아니다. 따라서 순수하게 DB에만 저장하고 조회를 하기보다는 적절하게 Cache를 사용하자

## 실험 3
> Select 50000건, Insert 50000건 (Vuesrs: 10, Run count: 5000)
> - 번갈아가면서 실행

![image](https://user-images.githubusercontent.com/59307414/156519812-7c0b472d-1150-4ac8-bbfc-0dae4509b7c6.png)

![image](https://user-images.githubusercontent.com/59307414/156519772-02aeb892-55d1-4a55-876f-1fec368b96ce.png)

![image](https://user-images.githubusercontent.com/59307414/156519698-ff82da01-f687-4c0e-b414-ef5dafc5a327.png)

---

|구분|TPS|Peak TPS|Mean Test Time(ms)|Run Time|
|---|---|---|---|---|
|MySQL|394.1|449|24.42|00:02:14|
|Mongo|709.7|945|11.27|00:01:17|
|Redis|916.2|1259|8.35|00:01:01|

- 읽기/쓰기 작업을 동시에 진행하다보니 관계형 데이터베이스(MySQL)에서는 성능이 매우 떨어지는 것을 확인할 수 있음
- 테이블간의 결합이나 정해진 관계가 구체적이지 않은 경우, 다수의 삽입/조회 작업이 있는 경우 비정형 데이터베이스를 사용하는 것이 좋음
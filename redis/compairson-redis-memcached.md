# Redis vs Memcached

## 인메모리 데이터 저장소
- 인메모리란 컴퓨터의 메인 메모리 RAM에 데이터를 올려서 속도를 높여서 사용하는 방법이다.
- SSD,HDD 같은 저장공간에서 데이터를 가져오는 것보다 RAM에 올려진 데이터를 가져오는데 걸리는 속도가 수 백배(HDD 기준) 이상 빠르다. 따라서 인메모리 데이터 저장소를 사용할 때 속도면에서 효과를 볼 수 있다.
- 하지만 용량적으로 크기가 작기 때문에 메인 데이터베이스로 활용하기에는 무리가 있다.
- 대표적으로 Redis, Memcached 등이 있다.

## Redis vs Memcached
- Redis와 Memcached는 많이 사용되는 오픈 소스이자 인메모리 데이터 저장소이다.
- 두 가지 모두 사용하기 쉽고 고성능을 자랑한다.
- 하지만 엔진으로 사용할 때 차이를 반드시 고려해야한다.
- Memcached는 명료하고 단순함을 위하여 개발된 반면, Redis는 다양한 용도에 효과적으로 사용할 수 있도록 많은 특징을 가지고 개발되었다.

### 공통점
1. 1ms 이하의 응답 대기시간
- 1ms 이하의 응답시간을 제공
- 데이터를 메모리에 저장하기 때문에 디스크 기반의 데이터베이스보다 빠르게 데이터를 읽을 수 있음

2. 개발의 용이성
- 문법적으로 사용하기 쉽고 개발코드의 양 또한 적음

3. 데이터 파티셔닝
- 데이터를 여러 노드에 분산하여 저장시킬 수 있음
- 수요가 증가할 때 더 많은 데이터를 효과적으로 처리하기 위하여 스케일 아웃이 가능

4. 다양한 프로그래밍 언어 지원
- Java, Python, C, C++, C#, JavaScript 등 다양한 언어를 지원

### Redis
#### 장점
- [다양한 데이터 구조를 지원](https://github.com/ruthetum/study/tree/main/redis#2-redis-%EB%8D%B0%EC%9D%B4%ED%84%B0-%ED%83%80%EC%9E%85-%ED%99%9C%EC%9A%A9%ED%95%98%EA%B8%B0)

- 스냅샷 지원
    - 특정시점에 데이터를 디스크에 저장하여 파일 보관이 가능
    - 장애 상황 발생 시 복구에 사용할 수 있음

- 복제
    - Master-Slave 구조로 여러 개의 복제본을 생성할 수 있음
    - 데이터베이스 읽기를 확장할 수 있기 때문에 고가용성 클러스터를 제공

- 트랜잭션

- Pub/Sub Messaging
    - Publish(발행)과 Subscribe(구독)방식의 메시지를 검색 가능
    - 높은 성능을 요구하는 채팅, 실시간 스트리밍, SNS 피드 그리고 서버 상호통신에 사용할 수 있음

#### 단점
- **싱글 스레드** 작동으로 인한 속도 차이
    - 싱글 스레드로 작동하기 때문에 한 번에 1개의 명령어만 실행할 수 있음
    - `keys *`와 같이 모든 키를 조회하거나, `flushall`와 같이 모든 데이터를 삭제하는 명령어를 사용할 때 Memcached는 1ms 정도 소요되지만, Redis의 경우 100만건의 데이터 기준 1초로 엄청난 속도 차이가 발생
    - **해결책**
        - `keys *` → `scan`으로 대체
        - Hash나 Sorted Set 등의 자료구조를 활용

- 메모리를 2배로 사용
    - Redis는 싱글 스레드로 작동하기 때문에 스냅샷을 뜰 때 자식 프로세스를 하나 만들낸 후 새로 변경된 메모리 페이지를 복사해서 사용함
    - Redis는 copy-on-write 방식을 사용하고 있지만 보통 사용할 때 데이터 변경이 빈번하기 때문에 실제 메모리 양만큼의 메모리를 자식 프로세스가 복사하게 됨
    - 그래서 실제로 필요한 메모리 양보다 더 많은 메모리를 사용하게 됨

- Redis는 메모리를 직접 처리할 수 없기 때문에 메모리 파편화가 발생하기 쉬움
    - **해결책**
        - 다양한 크기의 데이터 사용을 줄이고 유사한 크기의 데이터를 사용하여 파편화를 줄일 수 있음

### Memcached
#### 장점
- 멀티스레드를 지원하기 때문에 멀티 프로세스를 사용할 수 있음
    - 따라서 스케일 업을 통해 더욱 많은 작업을 처리할 수 있음

- Redis는 트래픽이 몰리면, 응답속도가 불안정하다고 한다. 반면 트래픽이 몰려도 Memcached의 응답 속도는 안정적인 편
    - cf. http://preview.hanbit.co.kr/2647/sample_ebook.pdf

- 메모리 파편화가 Redis에 비해 덜함
    - 내부적으로 slab 할당자를 사용
    - cf. https://www.slideshare.net/charsyam2/cache-governancepub

- Redis에 비하면 메타 데이터를 적게 사용하기 때문에 메모리 사용량이 상대적으로 적음

#### 단점
- 지원하는 타입이 다양하지 않음

## Reference
- https://aws.amazon.com/ko/elasticache/redis-vs-memcached/
- https://chrisjune-13837.medium.com/redis-vs-memcached-10e796ddd717
- https://americanopeople.tistory.com/148
- https://zangzangs.tistory.com/72
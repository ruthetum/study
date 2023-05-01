# EC2 환경에서 Kafka 설치하기
1. 인스턴스 생성
- 인스턴스 : Amazom Linux 2

2. AWS EC2 inbound 방화벽 설정
- 9092 port 개방

3. Java, Kafka 설치
- Install openjdk
    ```
    sudo yum install -y java-1.8.0-openjdk-devel.x86_64
    ```

- Download kafka
    ```
    wget http://mirror.navercorp.com/apache/kafka/3.1.0/kafka_2.13-3.1.0.tgz

    tar xvf kafka_2.13-3.1.0.tgz
    ```

4. EC2, Kafka 설정
- Kafka 실행 최소 Heap size 설정 제거
    ```
    export KAFKA_HEAP_OPTS="-Xmx400m -Xms400m"
    ```
- Kafka 2.5.0은 1G의 Heap memory가 default
- 테스트용 ec2인 t2.micro에 실행하기 위해 heap size 환경변수 선언

    > cf. 링크드인에서 테스트한 최적의 Java option 추천값
    ```
    -Xmx6g -Xms6g -XX:MetaspaceSize=96m -XX:+UseG1GC
    -XX:MaxGCPauseMillis=20 -XX:InitiatingHeapOccupancyPercent=35 -XX:G1HeapRegionSize=16M -XX:MinMetaspaceFreeRatio=50 -XX:MaxMetaspaceFreeRatio=80
    ```


5. Kafka 서버 설정 정보 수정
    ```
    cd kafka_2.13-3.1.0/
    vi config/server.properties
    ```

    ```
    listeners=PLAINTEXT://:9092
    advertised.listeners=PLAINTEXT://{aws ec2 public ip}:9092
    // advertised.listeners=PLAINTEXT://3.39.64.136:9092
    ```
    > - broker.id : 정수로 된 브로커 번호. 클러스터 내 고유번호로 지정
    > - listeners : kafka 통신에 사용되는 host:port
    advertised.listeners : Kafka client가 접속할 host:port
    > - log.dirs : 메시지를 저장할 디스크 디렉토리. 세그먼트가 저장됨
    > - log.segment.bytes : 메시지가 저장되는 파일의 크기 단위
    > - log.retention.ms : 메시지를 얼마나 보존할지 지정. 닫힌 세그먼트를 처리
    > - zookeeper.connect : 브로커의 메타데이터를 저장하는 주키퍼의 위치
    > - auto.create.topics.enable : 자동으로 토픽이 생성여부
    > - num.partitions : 자동생성된 토픽의 default partition 개수
    > - message.max.bytes : kafka broker에 쓰려는 메시지 최대 크기

6. Zookeeper 실행, Kafka 실행
- Zookeeper 실행
    ```
    bin/zookeeper-server-start.sh -daemon config/zookeeper.properties
    ```

- Kafka 실행
    ```
    bin/kafka-server-start.sh -daemon config/server.properties
    ```

7. 실행 확인
- `jps` 명령어로 확인

    ![image](https://user-images.githubusercontent.com/59307414/155881543-dd325967-2e17-4447-a7cb-36880588bdc7.png)
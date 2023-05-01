# 카프카 Local 환경에서 작동하기 (Windows)
## 설치
- Download URL : https://kafka.apache.org/downloads
    ![image](https://user-images.githubusercontent.com/59307414/155874677-a2a952de-9d05-473f-89e4-7b4964d17c9e.png)

    ![image](https://user-images.githubusercontent.com/59307414/155874689-926e8222-69d1-481d-91a0-e16ec07ef70a.png)

## 실행
- kafka에서 Broker는 kafka의 서버를 뜻하며 동일 노드 내에서 여러개의 Broker를 띄울 수 있음
- 이렇게 분산되어서 여러개의 Broker가 띄워져 있으면 이 분산 Message Queue를 관리해주는 역할을 하는것이 Zookeeper
- kafka 서버를 띄우기 앞서 Zookeeper를 반드시 띄워야 함

### Zookeeper 실행
- `.\kafka\bin\windows\zookeeper-server-start.bat kafka\config\zookeeper.properties`
- USAGE : `.\zookeeper-server-start.bat zookeeper.properties`

    ![image](https://user-images.githubusercontent.com/59307414/155874857-59662a66-e312-4ae1-ac27-aba7df7729de.png)

### Kafka 실행
- `.\kafka\bin\windows\kafka-server-start.bat kafka\config\server.properties`
- USAGE : `.\kafka-server-start.bat server.properties`

    ![image](https://user-images.githubusercontent.com/59307414/155874960-9595b1b6-437a-43a9-a616-22674492919e.png)

## 실행확인
- kafka 기본 port : 9092
- zookeeper 기본 port : 2181
    
    ![image](https://user-images.githubusercontent.com/59307414/155875029-95b4abdc-6342-483f-8f06-b54aed00882f.png)

---

## 테스트

### 토픽(topic) 생성 (broker)
- `.\kafka\bin\windows\kafka-topics.bat --create --bootstrap-server localhost:9092 --replication-factor 1 -partitions 1 --topic heedong`

    ![image](https://user-images.githubusercontent.com/59307414/155875314-aaf69e9c-ddbc-4483-b1ee-0bea75047639.png)

- `replication` : 1
- `partition` : 1
- `topic` : heedong
- 토픽 생성 관련된 예전 자료를 보면 `--bootstrap-server localhost:9092` 대신 `--zookeeper localhost:2181` 옵션을 주는 경우가 있는데 해당 옵션은 deprecated됨
    - https://stackoverflow.com/questions/53428903/zookeeper-is-not-a-recognized-option-when-executing-kafka-console-consumer-sh

### 생성한 topic에 message 보내기 (producer)
- `.\kafka\bin\windows\kafka-console-producer.bat --broker-list localhost:9092 --topic heedong`

    ![image](https://user-images.githubusercontent.com/59307414/155877035-f1610f54-8b8f-452d-b2e6-e858b41687c4.png)

### 토픽에 있는 메세지 가져오기 (consumer)
- `.\kafka\bin\windows\kafka-console-consumer.bat --bootstrap-server localhost:9092 --topic heedong --from-beginning`

    ![image](https://user-images.githubusercontent.com/59307414/155877059-a91ffad8-62cd-4861-bc67-2029483963a8.png)

### 토픽(topic 삭제)
- `.\kafka\bin\windows\kafka-topics.bat --bootstrap-server localhost:9092 --delete --topic heedong`

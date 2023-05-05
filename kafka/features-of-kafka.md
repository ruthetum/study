## kafka 특징: 카프카가 빅데이터 파이프라인으로 적합한 이유
### 1. 높은 처리량
카프카는 프로듀서가 브로커로 데이터를 보낼 때, 컨슈머가 브로커로부터 데이터를 받을 때 모두 묶어서 전송

동일한 양의 데이터를 보낼 때 네트워크 통신 횟수를 최소한으로 줄이기 때문에 동일 시간 내에 더 많은 데이터를 전송할 수 있음

많은 양의 데이터를 묶음 단위로 처리하는 배치로 빠르게 처리할 수 있기 때문에 대용량 실시간 로그 데이터를 처리하는데 적합

파티션 단위를 통해 동일 목적의 데이터를 여러 파티션에 분배하고 병렬 처리할 수 있음

파티션 개수만큼 컨슈머 개수를 늘려서 동일 시간당 데이터 처리량을 증가

### 2. 확장성
데이터 파이프라인에서 데이터를 모을 때 데이터가 얼마나 들어올지는 예측하기 어려움

하루에 1000건 가량 들어오는 로그 데이터도 예상치 못한 이벤트로 인해 100만 건 이상 데이터가 들어올 수 있음

카프카는 가변적인 환경에서 안정적으로 확장 가능하도록 설계

데이터가 적을 때는 카프카 클러스터의 브로커를 최소한의 개수로 운영하다가 데이터가 많아지면 클러스터의 브로커 개수를 늘려서 스케일 아웃할 수 있음, 반대로 처리해야 하는 데이터의 개수가 적어지는 경우 브로커 개수를 줄여서 스케일 인 가능

가변적인 환경에서 확장성있게 운영 가능

### 3. 영속성
영속성이란 데이터를 생성한 프로그램이 종료되더라도 사라지지 않은 데이터의 특성을 의미

카프카는 다른 메시징 플랫폼과 다르게 전송받은 데이터를 메모리에 저장하지 않고 파일 시스템에 저장

보편적으로 파일 시스템에 데이터를 적재하고 사용하는 것은 느리다고 생각할 수 있지만, 카프카는 운영체제 레벨에서 파일 시스템을 최대한 활용하는 방법을 적용

운영체제에서는 파일 I/O 성능 향상을 위해 페이지 캐시(page cache) 영역을 메모리에 따로 생성해서 사용

페이지 캐시 메모리 영역을 사용하여 한 번 읽은 파일 내용은 메모리에 저장했다가 다시 사용하는 방식이기 때문에 카프카가 파일 시스템에 데이터를 저장, 전송하더라도 처리량이 높음

디스크 기반의 파일 스템을 활용하기 때문에 브로커 애플리케이션에 장애가 발생해서 갑자기 종료되더라도 프로세스를 재시작하여 안전하게 데이터를 처리할 수 있음

### 4. 고가용성
3개 이상의 서버들로 운영되는 카프카 클러스터는 일부 서버에 장애가 발생하더라도 무중단으로 안전하고 지속적으로 데이터를 처리할 수 있음

클러스터로 이루어진 카프카는 데이터 복제(replication)를 통해 고가용성 특징을 지님

프로듀서로 전송받은 데이터를 여러 브로커 중 1대의 브로커에만 저장하는 것이 아니라 또 다른 브로커에도 저장

한 브로커에 장애가 발생하도라도 복제된 데이터가 나머지 브로커에 저장되어 있으므로 저장된 데이터를 기준으로 지속적으로 데이터 처리 가능

서버를 직접 운영하는 온프레미스(on-premise) 환경의 서버 랙 또는 퍼블릭 클라우드(public cloud)의 리전 단위 장애에도 데이터를 안전하게 복제할 수 있는 브로커 옵션 존재
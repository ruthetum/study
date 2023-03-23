# 대기열 시스템 프로토타입
Golang(+Echo Framework), Redis를 이용해 대기열 시스템 프로토타입을 작성합니다.

## 구성
### [대기열 서버](queue)
게임(소켓) 서버의 동시 접속자 수를 제한합니다.

대기열 서버는 아래의 기능을 수행합니다.

|          기능          | 구현                                         |
|:--------------------:|--------------------------------------------| 
|        대기표 발권        | Redis에 대기 번호 값 증가                          |
|   대기표 폴링 및 입장권 발급    | Redis 진입 번호 조회 후 대기 번호가 진입 번호 미만인 경우 대기표 발권 |
|  게임 서버 피드백 엔드포인트 제공  | Redis에 진입 번호 값 증가                          |
|  진입 번호 조정 엔드포인트 제공   | 게임 서버 피드백 외에 수동으로 Redis에 진입 번호 값 증가                             |
| 대기자 수(+ 예상 대기 시간) 제공 | Redis에 저장된 값 비교 (대기 번호 - 진입 번호)            |
|        알림 기능         | 트래픽이 몰리는 경우 로그 출력                          |

### [게임 서버(소켓 서버)](game)
게임(소켓) 서버는 대기열 서버에서 발급한 입장권을 통해 접속을 허용합니다.

게임(소켓) 서버는 아래의 기능을 수행합니다.

|          기능          | 구현                                                             |
|:--------------------:|----------------------------------------------------------------| 
|     입장권 검증 및 로그인     | 입장권 유효성 검사, 유저 식별, 입장 순번 캐싱                                    |
|      대기열 서버 피드백      | 동시 접속 인원에 따라 대기열 서버에 진입 번호 값 증가 요청<br/>더 이상 접속자가 없는 경우 초기화 피드백 |

### [부하 테스트기](stress)
부하 테스트기는 클라이언트의 롤을 담당하여 대기열 서버에 폴링 및 게임 서버에 접속합니다.

## 실행
```shell
// redis 세팅
docker-compose up -d

// 대기열 서버 실행
go run ./queue/main.go

// 게임 서버 실행
go run ./game/main.go

// 부하 테스트기 실행
go run ./stress/main.go
```
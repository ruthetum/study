# Chapter05. 트랜잭션과 잠금
- MySQL의 동시성에 영향을 미치는 요소: 잠금, 트랜잭션(트랜잭션 격리 수준)

- 트랜잭션: 작업의 완전성을 보장해주는 것
- 잠금: 트랜잭션과 비슷한 개념 같지만, 동시성을 제어하기 위한 기능

| 트랜잭션                 | 잠금              |
|----------------------|-----------------|
| 데이터의 정합성을 보장하기 위한 기능 | 동시성을 제어하기 위한 기능 |

- MySQL에서 사용되는 잠금은 크게 [MySQL 엔진 레벨](#MySQL-엔진의-잠금)과 [스토리지 엔진 레벨](#InnoDB-스토리지-엔진-잠금)로 나눌 수 있음
  
<br/>

# 트랜잭션
- 작업의 완전성을 보장해주기 위한, 데이터의 정합성을 보장하기 위한 기능
- InnoDB는 트랜잭션을 지원하지만, MyISAM과 Memory는 지원하지 않음
  - 트랜잭션때문에 속도가 떨어져서 MyISAM, Memory를 고려할 수 있지만 오히려 트랜잭션이 없을 때 고려해야할 문제가 더 많이 발생할 수 있음

## InnoDB vs MyISAM
```mysql
# 두 종류의 테이블 생성 후 레코드 생성
CREATE TABLE tab_myisam (fdpk INT NOT NULL, PRIMARY KEY (fdpk)) ENGINE=MYISAM;
INSERT INTO tab_myisam (fdpk) VALUES (3);

CREATE TABLE tab_innodb (fdpk INT NOT NULL, PRIMARY KEY (fdpk)) ENGINE=INNODB;
INSERT INTO tab_innodb (fdpk) VALUES (3);

# AUTO-COMMIT 활성화
SET autocommit=ON;

# 레코드 추가
INSERT INTO tab_myisam (fdpk) VALUES (1), (2), (3); # PK 중복으로 인한 에러 발생
INSERT INTO tab_innodb (fdpk) VALUES (1), (2), (3); # PK 중복으로 인한 에러 발생

# 조회
SELECT COUNT(*) FROM tab_myisam; # output: 3 / 1, 2는 추가됨
SELECT COUNT(*) FROM tab_innodb; # output: 1 / 트랜잭션에 의해 해당 쿼리가 수행되지 않음
```

## 트랜잭션 주의사항
- 트랜잭션은 최소한의 코드에만 적용하는 것이 좋음
- 단건 작업에 대한 트랜잭션의 범위가 커질 수록 커넥션을 오래 유지하는 문제 발생
- 외부 기능(ex. FTP, 메일, ...)이 트랜잭션에 묶여져 있다면 외부 기능에 문제가 발생했을 떄 DB까지 위험해질 수 있음

<br/>

# MySQL 엔진의 잠금
- MySQL에서 사용되는 잠금은 크게 스토리지 엔진 레벨과 MySQL 엔진 레벨로 나눌 수 있음
- MySQL 엔진 레벨의 잠금은 모든 스토리지 엔진에 영향을 미치지만, 스토리지 엔진 레벨의 잠금은 스토리지 엔진 간 상호 영향을 미치지 않음
- MySQL 엔진에서는 테이블 데이터 동기화를 위한 테이블 락 외에도 메타데이터 락, 네임드 락 기능도 제공
  - 메타데이터 락(Metadata Lock): 테이블의 구조를 잠금
  - 네임드 락(Named Lock): 사용자의 필요에 맞게 사용할 수 있는 잠금

## 글로벌 락(GLOBAL LOCK)
- MySQL에서 제공하는 잠금 가운데 가장 범위가 큰 잠금으로, `FLUSH TABLES WITH READ LOCK` 명령으로 획득할 수 있음
- 한 세션에서 글로벌 락을 획득하면 다른 세션에서 SELECT를 제외한 대부분의 DDL, DML 문장을 실행하는 경우 글로벌 락이 해제될 때까지 대기

### 백업 락
- 8.0부터 InnoDB가 기본 스토리지 엔진으로 채택되면서 글로벌 락보다 가벼운 락이 필요했고, 이를 지원하기 위한 백업 락이 도입
  - cf. Xtrabackup, Enterprise Backup
- 특정 세션에서 백업 락을 획득하면 모든 세션에서 아래와 같은 테이블의 스키마나 사용자 인증 관련 정보를 변경할 수 없음
  - 데이터베이스 및 테이블 등 모든 객체 생성 및 변경, 삭제
  - REPAIT TABLE과 OPTIMIZE TABLE 명령
  - 사용자 관리 및 비밀번호 변경
- 다만 백업 락은 일반적인 테이블의 데이터 변경은 허용

## 테이블 락(Table Lock)
- 개별 테이블 단위로 설정하는 잠금으로, `LOCK TABLES table name [READ | WRITE ]` 명령으로 권한 획득
  - `UNLOCK TABLES`로 잠금 해제
- 특별한 상황이 아니면 애플리케이션에서 거의 사용할 필요 없음
  - 테이블을 잠그는 작업은 글로벌 락과 동일하게 온라인 작업에 영향을 끼침

### 묵시적인 테이블 락
- MyISAM이나 Memory 테이블에서 데이터를 변경하는 쿼리를 실행하면 발생
  - MySQL 서버가 데이터가 변경되는 테이블에 잠금을 설정하고, 데이터를 변경한 후, 즉시 잠금을 해제하는 형태로 사용
  - 쿼리가 실행되는 동안 자동으로 획득됐다가, 쿼리 실행 완료 후 해제
- InnoDB 테이블의 경우 스토리지 엔진 차원에서 레코드 기반 잠금을 제공하기 때문에 단순 데이터 변경 쿼리로 인해 묵시적인 테이블 락이 설정되지는 않음
  - 정확히히 테이블 락이 설정되지만 대부분의 DML 쿼리에서는 무시되고, DDL에 한해 영향을 미침

## 네임드 락(Named Lock)
- `GET_LOCK()`을 이용해서 임의의 문자열에 대해 잠금을 설정할 수 있음
  - ex. `SELECT GET_LOCK('mylock'`, 2)
- 사용자가 지정한 문자열에 대해 획득하고, 해제하는 잠금으로 자주 사용되지는 않음
  - 예를 들어 DB 1대에 5대의 웹 서버가 접속하는 경우 웹 서버별 어떤 정보를 동기화해야 하는 경우처럼 여러 클라이언트가 상호 동기화를 처리해야하는 경우 네임드 락을 활용

## 메타데이터 락(Metadata Lock)
- 데이터베이스 객체(테이블, 뷰, ...)의 이름이나 구조를 변경하는 경우에 획득하는 잠금
- 명시적으로 획득하거나 해제할 수 있는 잠금은 아니고, 자동으로 획득/해제되는 잠금

<br/>

# InnoDB 스토리지 엔진 잠금
- MySQL 엔진과 별개로 스토리지 엔진 내부에서 레코드 기반 잠금 방식을 탑재
  - 이로 인해 동시성 처리를 제공
  - 하지만 이원화된 잠금 처리때문에 과거에는 스토리지 엔진에서 사용되는 잠금 정보를 MySQL 명령을 통해 접근하기에 까다로움
- 최근 버전에서는 InnoDB의 트랜잭션, 잠금, 잠금 대기 중인 트랜잭션 목록을 조회할 수 있음
  - `information_schema` 데이터베이스에 존재하는 `INNODB_TRX`, `INNODB_LOCKS`, `INNODB_LOCK_WAITS`를 조인해서 확인
- InnoDB의 잠금에 대한 모니터링도 강화되면서 Performance Schema`를 이용해서 스토리지 엔진 내부 잠금(세마포어)에 대한 모니터링도 가능

## InnoDB 스토리지 엔진의 잠금
- 레코드 기반의 잠금 기능 제공
  - 작은 공간으로 관리되기 때문에 레코드 락이 페이지 락, 테이브 락으로 레벨업되는 경우는 없음
- 일반 DBMS와는 다르게 InnoDB 엔진에서는 레코드 락뿐 아니라 레코드와 레코드 사이의 간격을 잠그는 갭(GAP) 락이 존재

![image](https://user-images.githubusercontent.com/59307414/210166791-2cc00689-6c04-49e6-8853-1aa050a2fd79.png)

### 레코드 락
- InnoDB 엔진에서 레코드 락은 레코드 자체가 아닌 인덱스의 레코드를 잠금
- 인덱스가 하나도 없는 테이블이더라도 내부적으로 자동 생성된 클러스터 인덱스를 이용하여 잠금 설정
  - 레코드 자체를 잠그냐, 인덱스를 잠그냐는 많은 차이가 발생

#### 레코드 잠금 vs 인덱스 잠금
- ㅇㅇ

### 갭 락
- 레코드 자체가 아닌 레코드와 바로 인접한 레코드 사이의 간격만을 잠그는 것을 의미
- 이를 통해 레코드와 레코드 사이의 간격에 새로운 레코드가 생성(INSERT)되는 것을 제어
- 갭 락은 그 자체보다 넥스트 키 락의 일부분으로 자주 사용

### 넥스트 키 락
- 레코드 락과 갭 락을 합쳐 놓은 형태의 잠금
- `innodb_locks_unsafe_for_binlog` 시스템 변수가 비활성화(0으로 설정)되면 변경을 위해 검색하는 레코드에 넥스트 키 락 방식으로 잠금 설정
- 바이너리 로그에 기록되는 쿼리가 레플리카 서버에서 실행될 때 소스 서버에서 만들어 낸 결과와 동일한 결과를 만들어내도록 보장하는 것이 목적
- 하지만 의외로 넥스트 키 락과 갭 락으로 이한 데드락이 발생하거나 다른 트랜잭션을 기다리게 하는 일이 자주 발생
  - 가능하다면 바이너리 로그 포맷을 ROW 형태로 바꿔서 넥스트 키 락이나 갭 락을 줄이는 것이 좋음
  - 5.5 버전까지는 많이 사용되지는 않았지만, 5.7 / 8.0 버전으로 업그레이드되면서 ROW 포맷의 바이너리 로그에 대한 안정성이 높아짐
  - 8.0에서는 ROW 포맷의 바이너리 로그가 기본 설정

### 자동 증가 락
- `AUTO_INCREMENT`와 관련, 중복되지 않고 저장된 순서대로 증가하는 일련번호를 제공하기 위해 자동 증가 락(Auto increment lock)이라고 하는 테이블 수준의 잠금 지원
- 자동 증가 락은 INSERT, REPLACE 쿼리 같이 새로운 레코드를 저장하는 쿼리에서만 필요
  - UPDATE, DELETE 등에서는 걸리지 않음

#### 자등 증가 락 작동 방식
| `innodb_autoinc_lock_mode=0`                       | `innodb_autoinc_lock_mode=1`                                                                                                                                                                                                                         |`innodb_autoinc_lock_mode=2`|
|----------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---|
| MySQL 5.0과 동일한 잠금 방식<br/><br/>모든 Insert 문장은 자등 증가 락을 사용 | 레코드를 처리할 때 MySQL 서버가 Insert되는 레코드의 수를 정확히 예측할 수 있을 떄는 자동 증가 락을 사용하지 않고, 그보다 가볍고 빠른 래치(mutex)를 이용하여 처리<br/><br/>래치는 자동 증가 락과 달리 아주 짧은 시간동안만 잠금을 걸고, 필요한 증가 값을 반환받으면 즉시 잠금 해제<br/><br/>예측할 수 없을 때에는 자동 증가 락 사용<br/><br/> 연속 모드(consecutive mode)라고도 함  | 절대 자동 증가 락을 사용하지 않고 경량화된 래치만을 사용<br/><br/> 단건 생성이어도 연속된 자동 증가 값을 보장하지 않음<br/><br/>인터리빙 모드(interleaved mode)라고 함<br/><br/>동시 처리 성능은 높아짐<br/><br/>자동 증가 기능에 대해서는 유니크한 값이 생성된다는 것은 보장<br/><br/>STATEMENT 포맷의 바이너리 로그를 사용하는 복제의 경우 소스 서버와 레플리카 서버의 자동 증가 값이 달라질 수 있기 때문에 주의해야 함

> 5.7까지는 `innodb_autoinc_lock_mode`의 기본값이 1이었지만, 8.0 버전부터는 기본값이 2로 변경
> 
> 8.0부터는 바이너리 로그 포맷이 STATEMENT가 아니라 ROW 포맷이 기본값이 되었기 때문
> 
> 만약 8.0에서 STATEMENT 포맷의 바이너리 로그를 사용한다면 `innodb_autoinc_lock_mode`를 1로 설정해서 사용할 것을 권장

## 인덱스와 잠금
- InnoDB는 레코드를 잠그는 것이 아니라 인덱스를 잠그는 방식으로 처리
- 변경해야 할 레코드를 찾기 위해 검색한 인덱스의 레코드를 모두 락을 걸어야 함

```mysql
# 테이블
CREATE TABLE employees (
    emp_no INT NOT NULL, PRIMARY KEY (emp_no),
    birth_date DATE,
    first_name VARCHAR(14),
    last_name VARCHAR(16),
    gender ENUM('M', 'F'),
    hire_date DATE
) ENGINE=INNODB;

# first_name index 설정
ALTER TABLE employees ADD INDEX first_name_index(first_name);

# first_name이 Georgi 레코드 253개가 있고, first_name이 Georgi면서 last_name이 Klassen인 레코드가 1개가 있음
SELECT COUNT(*) FROM employees WHERE first_name = 'Georgi'; # output: 253
SELECT COUNT(*) FROM employees WHERE first_name = 'Georgi' AND last_name = 'Klassen'; # output: 1

# first_name이 Georgi면서 last_name이 Klassen인 사원의 입사일(hire_date)을 오늘로 변경
UPDATE employees SET hire_date = NOW() WHERE first_name = 'Georgi' AND last_name = 'Klassen';
```

- 위 예제에서 단건 업데이트지만 first_name에 인덱스가 설정되어 있기 때문에 253건의 레코드가 모두 잠김
- 테이블에 인덱스가 하나도 없다면 테이블을 풀스캔하면서 모든 레코드를 잠금
- 그래서 인덱스를 설정하는 게 만능이 아니기 때문에 InnoDB에서는 인덱스 설계가 중요한 이유 중 하나

## 레코드 수준의 잠금 확인 및 해제
- `information_schema` 데이터베이스에 존재하는 `INNODB_TRX`, `INNODB_LOCKS`, `INNODB_LOCK_WAITS`를 조인해서 확인
- `SHOW PROCESSLIST` 명령어를 통해 프로세스 목록 및 스레드를 조회할 수도 있음
- `KILL` 명령어를 통해 스레드 삭제
  
<br/>

# MySQL의 격리 수준

| Isolation level  | DIRTY READ | NON-REPEATABLE READ | PHANTOM READ |
|:----------------:|------|-----------------|--------------|
| READ UNCOMMITTED | **발생**   | **발생**              | **발생**           |
|  READ COMMITTED  |      | **발생**              | **발생**           |
| REPEATABLE READ  |      |                 | 발생(InnoDB는 없음) |
|   SERIALIZABLE   |      |                 |              |

### 부정합 문제
#### Dirty read
- 트랜잭션에서 처리한 작업이 완료되지 않았음에도 불구하고 다른 트랜잭션에서 볼 수 있게 되는 현상
- 데이터가 나타났다가 사라졌다하는 현상을 초래

#### Non-repeatable read
- 하나의 트랜잭션 내에서 동일한 SELECT 쿼리를 실행했을 때 항상 같은 결과를 보장해야 한다는 REPEATABLE READ 정합성에 어긋나는 현상

#### Phantom read
- SELECT ... FOR UPDATE 쿼리와 같은 쓰기 잠금을 거는 경우 다른 트랜잭션에서 수행한 변경 작업에 의해 레코드가 보였다가 안 보였다가 하는 현상

## 격리 수준
### READ UNCOMMITTED
- 한 트랜잭션의 변경된 내용을 커밋이나 롤백과 상관 없이 다른 트랜잭션에서 읽을 수 있는 격리 수준
- 모든 부정합 문제 발생

### READ COMMITTED
- COMMIT이 완료된 데이터만 조회 가능한 격리 수준 (undo 영역 활용)
- 더티 리드 해결
- 오라클에서 기본 설정으로 주로 사용

### REPEATABLE READ
- 트랜잭션이 시작되기 전에 커밋된 내용에 관해서만 조회할 수 있는 격리 수준 (트랜잭션 번호를 기준으로 활용)
- NON-REPEATABLE-READ 해결
- PHANTOM READ 발생 (InnoDB에서는 PHANTOM READ 해결)
  - ex. `FOR UPDATE`
  > FOR UPDATE
  > 
  > SELECT FOR UPDATE쿼리는 가정 먼저 LOCK을 획득한 SESSION의 SELECT 된 ROW들이 UPDATE 쿼리후 COMMIT 되기 이전까지 다른 SESSION들은 해당 ROW들을 수정하지 못하도록 하는 기능
- MySQL에서 기본 설정으로 주로 사용
  



> **InnoDB 스토리지 엔진에서 PHANTOM READ 해결**
> - InnoDB 스토리지 엔진은 레코드 락과 갭 락을 합친 넥스트 키 락을 사용
> 
> 테이블에 c1 = 13 , c1 = 17 인 두 레코드가 있다고 가정할 때 `SELECT c1 FROM t WHERE c1 BETWEEN 10 AND 20 FOR UPDATE` 쿼리를 수행하면,
> 
> 10 <= c1 <= 12, 14 <= c1 <= 16, 18 <= c1 <= 20 인 영역은 전부 갭 락에 의해 락이 걸려서 해당 영역에 레코드를 삽입할 수 없음
>
> 또한 c1 = 13, c1 = 17인 영역도 레코드 락에 의해 해당 영역에 레코드를 삽입할 수 없다. 참고로 INSERT 외에 UPDATE, DELETE 쿼리도 마찬가지이다.
>
> 이러한 방식으로 InnoDB 스토리지 엔진은 넥스트 키 락을 이용하여 PHANTOM READ 문제를 해결

### SERIALIZABLE
- 한 트랜잭션을 다른 트랜잭션으로부터 완전히 분리하는 격리 수준
- 모든 부정합 문제 해결

> ref.
> 
> 격리수준
> - https://zzang9ha.tistory.com/381
> - https://steady-coding.tistory.com/562
> - https://tecoble.techcourse.co.kr/post/2022-11-07-mysql-isolation/
> 
> FOR UPDATE 구문
> - https://jinhokwon.github.io/mysql/mysql-select-for-update/
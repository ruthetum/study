# Isolation Level 변경에 따른 성능 확인
## Isolation Level
[여기서 확인](../real-mysql/ch05-transaction-lock.md#MySQL의-격리-수준)

## Isolation Level 확인하기
```mysql
SELECT @@GLOBAL.transaction_isolation, @@SESSION.transaction_isolation;
# or
SELECT @@GLOBAL.tx_isolation, @@SESSION.tx_isolation;
```

## Isolation Level 설정하기
1. 서버 띄울 때 설정
- mysql.conf or mysql.ini 등의 설정 파일에 명시
```mysql
[mysqld]
transaction-isolation = REPEATABLE READ
transaction-read-only = OFF
```

2. 문장으로 설정하기
```mysql
SET SESSION TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;
SET SESSION TRANSACTION ISOLATION LEVEL READ COMMITTED;
SET SESSION TRANSACTION ISOLATION LEVEL REPEATABLE READ;
SET SESSION TRANSACTION ISOLATION LEVEL SERIALIZABLE;
```

3. 변수로 설정하기
```mysql
UPDATE @@SESSION.transaction_isolation = 'read-uncommitted';
UPDATE @@SESSION.transaction_isolation = 'read-committed';
UPDATE @@SESSION.transaction_isolation = 'repeatable-read';
UPDATE @@SESSION.transaction_isolation = 'serializable';
```
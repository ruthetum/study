# 11. 쿼리 작성 및 최적화
- SQL은 어떠한(what) 데이터를 요청하기 위한 언어이지, 어떻게(how) 데이터를 읽을지를 표현하는 언어는 아님
  - 따라서 쿼리가 빠르게 수행되게 하려면 데이터베이스 서버에서 쿼리가 어떻게 요청을 처리할지 예측할 수 있어야 함
- 일반적으로 애플리케이션 코드를 튜닝해서 성능을 2배 이끌어내는 것은 쉽지 않음
  - 하지만 DBMS에서 몇 십배, 몇 백배 성능 향상이 이뤄지는 것은 흔한 일

## 11.1 쿼리 작성과 연관된 시스템 변수
### 11.1.1 SQL 모드
- `sql_mode`라는 시스템 설정에는 여러 개의 값을 설정할 수 있음
  - 이를 통해 SQL 작성 결과에 영향을 줄 수 있음
- `STRICT_ALL_TABLES` : 일반적인 MySQL에서는 저장하려는 값의 길이가 칼럼 길이를 넘어가더라도 경고만 발생시킨 후 초과한 부분을 자르고 저장한다. 이 옵션을 주면 경고가 아닌 에러를 발생시켜 칼럼 길이를 넘는 데이터를 저장하는 것을 막을 수 있다.
- `STRICT_TRANS_TABLES` : MySQL 서버는 칼럼 타입과 호환되지 않는 값을 저장하려고 할 때 최대한 변환해서 저장하려고 하나, 이 옵션을 주면 강제 변환하지 않고 에러를 발생시킨다.
- `TRADITIONAL` : 위 두 방법보다 조금 더 엄격한 방법으로 ANSI 표준 모드로 동작하도록 한다.
- `ONLY_FULL_GROUP_BY` : MySQL 쿼리에서는 GROUP BY절에 포함되지 않은 칼럼이라도 집합 함수 없이 SELECT나 HAVING절에 이용할 수 있도록 되어 있다. 사실 SQL 표준과는 거리가 있는 방식인데, 이 옵션은 이를 방지하도록 해준다.
- `IGNORE_SPACE` : 프로시저나 함수명과 괄호 사이에 공백이 들어가 있어도 무시해준다.
- `ANSI` : 위의 여러가지 옵션을 조합해서 MySQL 서버가 최대한 SQL 표준에 맞게 동작하도록 한다.

### 11.1.2 영문 대소문자 구분
- MySQL은 설치된 운영체제에 따라 대소문자를 구분함
  - 유닉스 계열 : 대소문자 구분
  - 윈도우 계열 : 대소문작 구분 X
- `lower_case_table_names` 시스템 변수를 통해 설정 가능
  - 0 : 대소문자 구분
  - 1 : 테이블 이름을 모두 소문자로 저장
  - 2 : windows, macOS에서만 가능(저장은 대소문자 구분, 쿼리는 구분 X)

### 11.1.3 MySQL 예약어
- 예약어로 설정된 키워드를 사용하려면 역따옴표(`)나 쌍따옴표(")로 감싸야 함

## 11.2 매뉴얼의 SQL 문법 표기를 읽는 방법
- `[ ]`(대괄호): 해당 키워드 또는 표현식이 선택 사항
- `|`(파이프): 해당 키워드 또는 표현식 중 단 하나만 선택
- `{ }`(중괄호): 해당 키워드 또는 표현식 중 반드시 하나 이상 선택
- `...`: 명시된 키워드 또는 표현식이 반복될 수 있음

## 11.3 MySQL 연산자와 내장 함수
- 다른 DBMS에서 사용되는 기본적인 연산자는 MySQL에서 비슷하게 사용
  - MySQL에서만 사용되는 연산자나 표기법도 존재
- ANSI 표준 형태가 아닌 연산자도 존재하는데 가독성이 좋지 않으므로 되도록이며 ANSI 표준 형태의 연산자 사용을 권장

### 11.3.1 리터럴 표기법 문자열
- 문자열
- 숫자
- 날짜
- 불리언

### 11.3.2 MySQL 연산자
- 동등 비교 (`=`, `<=>`)
  - `<=>`: NULL 값에 대한 비교 수행
- 부정 비교 (`!=`, `<>`)
- NOT 연산자 (`!`)
- AND 연산자 (`&&`)와 OR 연산자 (`||`)
- 나누기 연산자 (`/`, `DIV`)와 나머지 연산자 (`%`, `MOD`)
- REGEXP 연산자
- LIKE 연산자
  - `REGEXP`는 인덱스를 전혀 사용하지 못하지만, `LIKE`는 인덱스를 이용해서 처리할 수 있음
- BETWEEN 연산자
  - `BETWEEN`은 간혹 `IN`연산자와 비슷한 처리를 하는 것 같지만 실제로는 인덱스를 선형적으로 검색(BETWEEN)하냐, 동등 비교(Equal)를 여러 번 수행하느냐의 차이기 때문에 IN 연산자를 통한 처리가 효율적
- IN 연산자

### 11.3.3 MySQL 내장 함수
- NULL 값 비교 및 대체 함수 (`ISNULL`, `IFNULL`)
- 현재 시각 조회 (`NOW`, `SYSDATE`)
  - SYSDATE 함수의 경우 래플리카에서 동일하게 작동하지 않고, 비교 시에 인덱스를 효율적으로 사용하지 못하므로 NOW 함수를 사용하는 것을 권장
- 날짜와 시간의 포맷 (`DATE_FORMAT`, `STR_TO_DATE`)
- 날짜와 시간의 연산 (`DATE_ADD`, `DATE_SUB`)
- 타임스탬프 연산 (`UNIX_TIMESTAMP`, `FROM_UNIXTIME`)
- 문자열 처리 (`RPAD`, `LPAD` / `RTRIM`, `LTRIM`, `TRIM`)
- 문자열 결합 (`CONCAT`)
- GROUP BY 문자열 결합 (`GROUP_CONCAT`)
- 값의 비교와 대체 (`CASE WHEN ...THEN ... END`)
- 타입의 변환 (`CAST`, `CONVERT`)
- 이진값과 16진수 문자열 변환 (`HEX`, `UNHEX`)
- 암호화 및 해시함수 (`MD5`, `SHA`, `SHA2`)
  - MD5, SHA 함수의 결과를 UNHEX 함수를 이용해 이진값으로 변환하여 저장하면 저장공간을 효율적으로 사용
  - 다시 사람이 읽을 수 있는 형태로 되돌릴 때는 HEX 함수를 이용
- 처리 대기 (`SLEEP`)
  - 개발이나 디버깅 용도로 활용 시에 유용
- 벤치마크 (`BENCHMARK`)
  - 쿼리의 성능을 측정할 때 사용
- IP 주소 변환 (`INET_ATON`, `INET_NTOA`)
  - 보통 IP 주소를 VARCHAR(15) 타입에 '.'으로 구분해서 저장, 이 경우 저장공간이 크게 필요
  - INET_ATON은 IPv4 주소(문자열)를 정수형으로 변환, INET_NTOA는 정수형을 IPv4 주소(문자열)로 변환
  - INET6_ATON 함수를 이용하면 IPv6 주소(+IPv4)도 변환 가능
- JSON 포맷 (`JSON_PRETTY`)
- JSON 필드 크기 (`JSON_STORAGE_SIZE`)
- JSON 필드 추출 (`JSON_EXTRACT`)
- JSON 오브젝트 포함 여부 (`JSON_CONTAINS`)
- JSON 오브젝트 생성 (`JSON_OBJECT`)
- JSON 컬럼으로 집계 (`JSON_OBJECTAGG`, `JSON_ARRAYAGG`)
- JSON 데이터를 테이블로 변환(`JSON_TABLE`)

## 11.4 SELECT
### SELECT 절의 처리 순서
- 일반적으로 아래 순서를 따름
  ```
  (SELECT) -> (FROM) -> (WHERE) -> (GROUP BY) -> (DISTINCT) -> (HAVING) -> (ORDER BY) -> (LIMIT)
  ```

- GROUP BY절이 없이 ORDER BY만 사용된 쿼리에서는 다음 순서를 따르기도 함
  - 드라이빙 테이블 읽어서 ORDER BY 적용
  - 드리븐 테이블 WHERE 적용 및 조인 실행
  - LIMIT 적용

### WHERE 절의 인덱스 사용
- 인덱스를 사용하기 위한 기본적인 규칙은 **인덱스된 칼럼의 값 자체를 변경하지 않고 그대로 사용**

- WHERE 조건이 인덱스를 사용할 수 있는 기준은 5장에서 보았듯이 범위 제한 조건과 체크 조건으로 나뉘는데, 범위 제한 조건으로 동등 비교 조건이나 IN으로 구성된 조건이 인덱스를 구성하는 칼럼과 순서대로 얼마나 일치하는가에 따라 달라진다.
결국 WHERE 절에서는 각 조건이 명시된 순서는 중요하지 않고, 해당 칼럼에 대한 조건이 있는지 없는지가 중요함

- 각 조건이 OR로 연결되면 비교해야 할 레코드가 더 늘어나기 때문에 풀 테이블 스캔을 사용하거나, 인덱스를 각각의 칼럼에 맞추어 사용하더라도 index_merge 방식으로 접근해야 해서 효율이 떨어질 수 있음
  - 인덱스를 효율적으로 이용할 수 있도록 OR 대신 AND로 조건을 주는 것이 좋음

### GROUP BY 절의 인덱스 사용
- GROUP BY 절의 각 칼럼은 비교 연산자를 가지지 않으므로 WHERE 절처럼 범위 제한 조건이나 체크 조건 등을 고려할 필요는 없음

- 사용 조건은 다음과 같음
  - GROUP BY 절에 명시된 칼럼이 인덱스 칼럼의 순서와 위치가 같아야 함
  - 인덱스의 앞쪽부터 빠지는 칼럼 없이 사용해야 하며, 인덱스에 존재하지 않는 칼럼으로 GROUP BY를 사용해서도 안됨

- 여기서 조금 더 주의할 점은 WHERE 조건 절에 앞쪽, 예를 들어 COL1, COL2 가 동등 비교 조건으로 사용된다면, GROUP BY 절에는 COL1, COL2 가 빠지고 COL3 부터 사용해도 인덱스를 사용할 수 있음

  ```sql
  # 다음 두 쿼리는 같은 결과를 가져옴
  ... WHERE COL1 = '상수' ... GROUP BY COL2, COL3
  ... WHERE COL1 = '상수' ... GROUP BY COL1, COL2, COL3
  ```

### ORDER BY 절의 인덱스 사용
- 기본적으로 ORDER BY 는 GROUP BY 절의 사용 조건과 거의 비슷함
  - 여기에 조건이 더 하나 추가되는데 정렬되는 각 칼럼의 오름차순 혹은 내림차순 옵션이 인덱스와 같거나 정반대의 경우에만 사용할 수 있음
- MySQL의 인덱스는 오름차순으로 정렬되어 있기 때문에 모든 ORDER BY의 조건이 오름차순 혹은 내림차순 한 방향으로만 걸려있어야 함

### WHERE 조건과 ORDER BY(또는 GROUP BY) 절의 인덱스 사용
- 다음 3가지 중 하나를 사용함
  - WHERE 절과 ORDER BY 절이 동시에 같은 인덱스를 사용
    - WHERE, ORDER BY 절의 대상 칼럼이 모두 하나의 인덱스에 연속해서 포함되어 있을 때 사용 가능
    - 가장 빠른 성능
    
  - WHERE 절만 인덱스 이용
    - 인덱스를 통해 검색 후 Filesort를 통해 정렬을 수행
    - WHERE 절의 조건에 일치하는 레코드 건수가 많지 않을 때 사용
  
  - ORDER BY 절만 인덱스 이용
    - ORDER BY 절의 순서대로 레코드를 읽으면서 WHERE 절을 하나씩 적용하는 방식
    - 아주 많은 레코드를 조회해서 정렬해야 할 경우 이 방법을 사용
  
### GROUP BY 절과 ORDER BY 절의 인덱스 사용
- GROUP BY 절의 칼럼과 ORDER BY 절의 칼럼이 순서와 내용이 모두 같아야 함 
- **둘 중 하나라도 인덱스를 이용할 수 없을 때에는 둘다 인덱스를 사용하지 못함**

### WHERE, ORDER BY, GROUP BY 절의 인덱스 사용
1. WHERE 절이 인덱스를 사용할 수 있는가?
2. GROUP BY 절이 인덱스를 사용할 수 있는가?
3. GROUP BY 절과 ORDER BY 절이 동시에 인덱스를 사용할 수 있는가?

### DISTINCT
- DISTINCT에서 주의할 점은 **DISTINCT는 SELECT되는 레코드를 유니크하게 가져오는 것이지 칼럼 하나를 유니크하게 가져오는 것이 아니라는 것**
- 즉 다음 두 쿼리는 똑같이 (COL1 + COL2) 의 유니크 값을 가져옴. 괄호는 적용되지 않음

  ```sql
  SELECT DISTINCT COL1, COL2 FROM test_table;
  SELECT DISTINCT(COL1), COL2 FROM test_table;
  ```

- 집합 함수 내에서 사용된 DISTINCT는 조금 다름
  - 집합 함수의 인자 칼럼들 중에서 중복을 제거하고 남은 값들을 가져옴

### JOIN
- 보통의 인덱스 레인지 스캔으로 레코드를 읽는 작업은 다음 순서를 따름

1. 인덱스 탐색 : 인덱스에서 조건을 만족하는 값이 저장된 위치를 찾음
2. 인덱스 스캔 : 1번에서 찾은 위치에서 필요한 만큼 인덱스를 읽음
3. 2번에서 읽는 인덱스 키와 레코드 주소로 최종 레코드를 읽음

- 조인 시 드라이빙 테이블을 읽을 때는 인덱스 탐색을 한 번만 수행하고, 그 이후로는 스캔만 수행하면 되지만, 드리븐 테이블에서는 탐색과 스캔 작업을 드라이빙 테이블에서 읽은 레코드 건수만큼 반복
  - 옵티마이저는 항상 드라이빙 테이블이 아니라 드리븐 테이블을 최적으로 읽을 수 있게 실행 계획을 수립

- 따라서 조인하는 테이블 칼럼 중 한 쪽 테이블 칼럼에만 인덱스가 있을 경우 **옵티마이저는 인덱스가 있는 테이블을 드리븐 테이블로 선택할 확률이 높음**
  - 드라이빙 테이블을 풀 테이블 스캔 하더라도 동일함
  
- 추가로 몇 가지 조인에 대해 알아두어야 하는 점은 아래와 같음

  - OUTER JOIN 에서 OUTER 로 조인되는 테이블 칼럼에 대한 조건은 WHERE 절이 아니라 모두 ON 절에 명시해야 함
    - 그렇지 않으면 옵티마이저는 INNER JOIN 과 같은 방법으로 처리
  
  - 페이징 처리 시 COUNT 쿼리를 날릴 때 데이터를 가져오던 쿼리문에서 SELECT 절만 바꿔서 가져오는 실수를 주의해야 함
    - OUTER JOIN 시 조인 전과 후의 레코드 건수의 차이가 없다면 당연히 조인을 제거하고 COUNT 쿼리를 날리는 것이 좋음

  - NOT EXISTS나 NOT IN(subquery)는 상당히 비효율적이므로 가능하다면 OUTER JOIN을 이용한 ANTI JOIN으로 처리
    - OUTER JOIN 을 하면서 필요 없는 쪽의 테이블 칼럼을 IS NULL 로 걸러서 가져오는 방법
  
  - INNER JOIN과 OUTER JOIN이 가져와야 하는 레코드 건수가 같다면 둘의 성능 차이는 거의 없음

  - FOREIGN KEY는 조인과 아무런 연관이 없음
    - FOREIGN KEY를 생성하는 주 목적은 데이터의 무결성을 보장하기 위함
    - 아무런 상관이 없고 값만 같은 칼럼으로도 조인 가능

  - 보통 드라이빙 테이블부터 읽은 순으로 조인의 결과 순서가 보장되지만, 조인 버퍼를 사용한 조인은 정렬이 보장되지 않음


### GROUP BY
- 쿼리에 GROUP BY가 사용되면 GROUP BY 절에 사용되지 않은 칼럼은 반드시 집합 함수로 감싸서 사용해야 함
- 다른 DBMS와 달리 MySQL의 GROUP BY는 정렬을 기본으로 수행
- MySQL의 GROUP BY가 불필요한 정렬을 수행하지 않게 하려면 ORDER BY NULL 키워드를 사용해야 함

### ORDER BY
- 어떤 DBMS도 ORDER BY절이 명시되지 않은 쿼리는 어떠한 정렬도 보장하지 않음
  - 인덱스를 사용한 SELECT 절이라고 ORDER BY절을 사용하지 않아도 되는 것은 아님
  - 정렬이 필요한 곳에는 꼭 ORDER BY를 명시해야 함

- 인덱스를 사용하지 못하는 정렬 작업은 실행 계획에 "Using filesort" 코멘트가 표시되는데, 이는 디스크 파일을 이용했다는 뜻은 아니고 그냥 퀵 소트 정렬 알고리즘을 사용했다 정도의 의미로만 이해
  - 실제로 메모리만 이용했는지 디스크의 파일까지 이용했는지는 알 수 없음

- 여러 개의 칼럼을 조합해서 정렬할 때 각 칼럼의 정렬 순서가 오름차순, 내림차순이 혼용되면 인덱스를 이용할 수 없음
  - 전부 다 오름차순이거나, 전부 다 내림차순이어야 인덱스를 사용해서 정렬할 수 있음

### 서브 쿼리
- MySQL 서버는 서브 쿼리를 최적으로 실행하지 못할 때가 많음
- 최대한 서브 쿼리를 지양하고, JOIN으로 해결하거나 두 번의 쿼리로 나눠서 실행하는 것이 훨씬 좋음

- IN(subquery) 같은 경우는 옵티마이저가 EXISTS (subquery)형태로 변환하여 실행

### 집합 연산
- 집합 연산자에는 보통 UNION, INTERSECT, 그리고 MINUS가 있는데, MySQL은 UNION 기능만 제공
- INTERSECT와 MINUS는 JOIN을 활용하면 충분히 같은 결과를 가져올 수 있음
  - UNION 은 기본적으로 두 집합을 합치는 것이기 때문에 DISTINCT 조건이 기본적으로 붙음
    - 두 집합 간에 중복된 결과과 나오지 않는다는 것이 보장된다면, 중복제거 작업을 하지 않는 UNION ALL을 쓰는 것이 좋음

### SELECT INTO OUTFILE
- SELECT 쿼리의 결과를 화면으로 출력하는 것이 아니라 파일로 저장할 수 있음
- 테이블 단위로 데이터를 덤프받아서 적재하거나, 엑셀 혹은 다른 DBMS로 옮길 때 csv 형태로 파일을 뽑아서 유용하게 사용할 수 있음
- 파일 덤핑 시 주의할 점은 아래와 같음
  - SELECT 결과는 MySQL 클라이언트가 아니라 MySQL 서버가 기동 중인 장비의 디스크로 저장
  - 파일과 파일이 저장되는 디렉터리는 MySQL 서버를 기동 중인 OS의 계정이 쓰기 권한을 가지고 있어야 함
  - 이미 동일 디렉터리, 동일 이름의 파일이 있다면 덮어쓰는 것이 아니라 에러가 발생 

- 다음 쿼리와 같이 OUTFILE 옵션 뒤에는 결과를 저장할 파일 경로와 이름을 적고, FIELDS 옵션에는 각 칼럼의 구분자를, LINES 옵션에는 각 레코드의 구분자를 명시
  ```sql
  SELECT col1, col2, col3  
  INTO OUTFILE '/tmp/result.csv'  
  FIELDS TERMINATED BY ','  
  LINES TERMINATED BY ' \\n'  
  FROM my_table  
  WHERE col1 BETWEEN 100 AND 200;
  ```

> ref. https://wbluke.tistory.com/30

## 11.5 INSERT
- 일반적인 웹 서비스에는 단건 처리 혹은 소량의 레코드를 삽입하기 때문에 성능에 대해 크게 고려할 부분이 많지 않음
- 오히려 다수의 INSERT가 발생하는 경우 INSERT 문장보다는 테이블의 구조가 성능에 더 큰 영향을 미침
  - 대부분의 경우 INSERT와 SELECT 성능을 동시에 빠르게 만들 수 있는 테이블 구조는 없음
  - 따라서 어느 정도 타협하면서 테이블 구조를 설계해야 함
- 아래 내용에서는 INSERT 문장의 주의사항 및 테이블 용도에 따라 테이블 구조를 선택하는 방법에 대해 이야기함

### 11.5.1 고급 

- INSERT에서 사용할 수 있는 유용한 옵션은 아래와 같음
  - INSERT IGNORE
  - INSERT ON DUPLICATE KEY UPDATE
- 두 옵션 모두 유니크 인덱스나 PK에 대해 중복 레코드를 어떻게 처리할지 결정
  - INSERT IGNORE는 추가로 INSERT 문장에 대한 에러 핸들링 기능도 포함

#### INSERT IGNORE
- IGNORE 옵션은 저장하는 레코드의 PK나 유니크 인덱스 컬럼의 값이 아래와 같은 경우 무시하고 다음 레코드를 처리
  - 이미 테이블에 존재하는 레코드와 중복되는 경우
  - 저장하는 레코드의 컬럼이 테이블의 컬럼과 호환되지 않는 경우
- 여러 레코드를 하나의 INSERT 문장으로 처리하는 경우 유용함

  ```sql
  INSERT IGNORE INTO salaries (emp_no, salary, from_date, to_date) VALUES 
      (10001, 60117, '1986-06-26', '1987-06-26'),
      (10001, 62102, '1987-06-26', '1988-06-25'),
      (10001, 66074, '1988-06-25', '1989-06-25'),
      (10001, 66596, '1989-06-25', '1990-06-25'),
      (10001, 66961, '1990-06-25', '1991-06-25');
      
  INSERT IGNORE INTO salaries
      SELECT emp_no, (salary+100), '2020-01-01', '2022-01-01'
      FROM salaries WHERE to_date>='2020-01-01';
  ```

  - 에러가 발생하는 경우 경고 수준의 메시지로 바꾸고, 나머지 레코드의 INSERT를 계속 진행

- 프로그램 코드에서 중복을 무시하기 위해 INSERT IGNORE 옵션을 사용한다면 데이터 중복 이외의 에러가 발생할 여지가 없는지 확인할 필요가 있음
  - 제대로 검증되지 않은 INSERT IGNORE 문장은 의도하지 않은 에러도 무시

#### INSERT ON DUPLICATE KEY UPDATE
- INSERT IGNORE 문자은 중복이나 에러 발생 건에 대해서 모두 무시하지만 INSERT ... ON DUPLICATE KEY UPDATE 문장은 PK나 유니크 인덱스의 중복이 발생하면 UPDATE 문장 수행
- MySQL 서버의 REPLACE 문장도 INSERT ... ON DUPLICATE KEY UPDATE 문장과 비슷한 역할을 수행
  - 하지만 내부적으로 REPLACE 문장은 DELETE와 INSERT 조합으로 작동
  - INSERT ... ON DUPLICATE KEY UPDATE 문장은 UPDATE하는 방식으로 작동

  ```sql
  INSERT INTO dailty_statistic (target_date, stat_name, stat_value)
      VALUES ('DATE(NOW())', 'VISIT', 1)
      ON DUPLICATE KEY UPDATE stat_value=stat_value+1;
  ```

### 11.5.2 LOAD DATA 명령 주의 사항
- LOAD DATA INFILE은 SELECT INTO OUTFILE 쿼리에 대응하는 적재 기능의 쿼리다. CSV 파일 포맷 또는 일정 규칙을 지닌 구분자로 구분된 데이터 파일을 읽어 MySQL 서버의 테이블로 저장
  - ref. https://dev.mysql.com/doc/refman/8.0/en/load-data.html
  ```sql
  LOAD DATA INFILE '/tmp/employees.csv'
  IGNORE INTO TABLE employees
  FIELDS
      TERMINATED BY ','
      OPTIONALLY ENCLOSED BY '"' ESCAPED BY '"'
  LINES
      TERMINATED BY '\n'
      STARTING BY ''
  (emp_no, birth_date, first_name, last_name, gender, hire_date);
  ```

- 일반적으로 RDBMS에서 데이터를 빠르게 적재할 수 있는 방법으로 LOAD DATA 명령이 자주 소개
  - MySQL 서버의 LOAD DATA 명령도 내부적으로 MySQL 엔진과 스토리지 엔진의 호출 횟수를 최소화하고, 스토리지 엔진에 직접 데이터를 적재하기 때문에 일반적인 INSERT 보다 빠름

- 다만 아래의 단점 존재
  - 단일 스레드로 실행
    - 테이블에 여러 인덱스가 있다면 LOAD DATA 문장이 레코드를 INSERT하고, 인덱스에도 키 값을 INSERT해야 함
    - 테이블에 레코드가 INSERT되면 될수록 테이블과 인덱스의 크기가 커지게 되지만, 단일 스레드로 작업이 실행되기 때문에 시간이 지날수록 INSERT 속도가 떨어짐
  - 단일 트랜잭션으로 실행
  - 단일 트랜잭션으로 처리도기 때문에 LOAD DATA 문장이 시작한 시점부터 언두 로그가 삭제되지 못하고 유지돼야 함
    - 언두 로그를 디스크로 기록하는 부하와 함께, 언두 로그가 많이 쌓이면 레코드를 읽는 쿼리들이 필요한 레코드를 찾는데 더 많은 오버헤드를 발생

- LOAD DATA 명령으로 적재하는 데이터가 많지 않다면 위 두 가지 사항은 큰 문제가 되지 않음
  - 하지만 데이터가 크다면 실행 시간이 길이지고, 다른 트랜잭션 쿼리들의 성능에 영향을 미칠 수 있음

- 가능하다면 LOAD DATA 문장으로 잭재할 데이터 파일을 하나보다는 여러 개 파일로 나누어서 실행
  - 여럭 개의 트래잭션으로 나누어 실행

### 11.5.3 성능을 위한 테이블 구조
#### 대량 INSERT 성능
- 수백, 수천 건의 레코드를 INSERT한다면 해당 레코드의 PK 값 기준으로 미리 정렬해서 INSERT 문장을 구성하는 것이 좋음
- InnoDB 엔진은 INSERT할 때마다 PK를 검색해서 레코드가 저장될 위치를 찾기 때문에 PK로 정렬되지 않았다면 INSERT를 할 때마다 B-Tree의 랜덤한 위치의 페이지를 메모리로 읽어야 함 

#### PK 선정
- 앞서 언급된 것처럼 INSERT 성능을 결정하는 부분은 PK
- InnoDB 엔진을 사용하는 테이블의 PK는 클러스터링 키
  - 이는 세컨더리 인덱스를 이용하는 쿼리보다 PK를 이용하는 쿼리의 성능이 훨씬 빨라지는 효과를 냄
- 하지만 단순히 INSERT 성능만을 위해 PK를 설계해서는 안 됨
  - 결국 SELECT와 INSERT의 성능은 대립되고, 해당 테이블 및 서비스에 맞는 PK를 잘 설계해야 됨
  - 물론 둘다 만족할 수 있는 PK가 있다면 상관없음
- 보통의 애플리케이션은 쓰기 작업보다 읽기 작업이 압도적으로 많음
  - SELECT가 거의 실행되지 않고, INSERT가 많이 실행되는 테이불이면 테이블의 PK를 단조 증가/감소하는 패턴의 값을 선택하는 것이 좋음 (e.g. 로그)
    - 추가적으로 SELECT가 많지 않고, INSERT가 많다면 인덱스의 개수를 최소화하는 것이 좋음
      - 인덱스가 많을수록 쓰기 성능이 떨어지기 때문
  - 반대로 SELECT가 많이 실행되고, INSERT가 거의 실행되지 않는 테이블이라면 PK를 SELECT 쿼리를 빠르게 만드는 방향으로 구성해야 함

#### Auto-increment 컬럼
- INSERT에 최적화된 테이블을 생성하기 위해서 아래 두 가지 요소를 고려
  - 단조 증가 또는 단조 감소되는 값으로 PK 선정
  - 세컨더리 인덱스 최소화

- Auto-increment 값을 PK로 해서 테이블을 생성하는 것은 MySQL 서버에서 가장 빠른 INSERT를 보장하는 방법

- MySQL 서버에서 자동 증가 값의 채번을 위해서는 잠금이 필요, 이를 AUTO-INC 잠금이라고 하고, innodb_autoinc_lock_mode 시스템 변수를 통해 잠금 방식 설정 가능
  - innodb_autoinc_lock_mode = 0: 항상 AUTO-INC 잠금을 걸고 한 번에 1씩만 증가된 값을 가져옴
    - 5.1버전까지의 방식으로 하위호환을 위해 유지된 설정이고, 현재 사용할 필요 없음
  - innodb_autoinc_lock_mode = 0 = 1 (5.7 default): 단순히 레코드 한 건씩 INSERT하는 쿼리에서는 AUTO-INC 잠금을 사용하지 않고 뮤텍스를 이용해 더 가볍고 빠르게 처리
    - 하나의 INSERT 문장으로 여러 레코드를 INSERT하거나 LOAD DATA 명령으로 INSERT하는 쿼리에서는 AUTO-INC 잠금을 걸고 , 필요한만큼의 자동 증가값을 한 번에 가져와서 사용
  - innodb_autoinc_lock_mode = 2(8.0 default): LOAD DATA나 BULK INSERT를 포함한 INSERT 계열의 문장을 실행할 때 AUTO-INC 잠금을 사용하지 않음
    - 자동 증가 값을 적당히 미리 할당받아서 처리 (따라서 가장 빠름)
    - 채번된 번호는 단조 증가하는 유니크한 번호까지만 보장하고, INSERT 순서와 채번된 번호의 연속성은 보장하지 않음
      - 하나의 INSERT 문장으로 저장된 레코드라도 중간에 번호가 띄엄띄엄 발급될 수 있음
    - 쿼리 기반의 복제(Statement Based Replication)를 사용하는 MySQL에서느 소스서버와 래플리카 서버의 자동 증가 값이 동기화되지 못할 수도 있음

- MySQL 8.0 버전부터는 복제의 바이너리 로그 포맷 기본 값이 STATEMENT에서 ROW로 변경됐기 때문에 기본값이 2로 변경
  - MySQL 서버의 버전과 관계없이 복제를 STATEMENT 바이너리 로그 포맷으로 사용 중이면 innodb_autoinc_lock_mode 값을 1로 설정해야 함

- 자동 증가 값이 반드시 연속적이어야 하는 경우에도 1로 설정해야 함
  - 하지만 시간이 조금씩 지나면 연속된 값에 빈 공간이 생길 수 있기 때문에 자동 증가 값이 반드시 연속적이이어야 한다는 요건에 집착하면 안됨

> ### MySQL binlog format for replication
> - Statement : 명령문 기반의 로깅 방식
> - Row : 행 기반의 데이터 로깅 방식
> - Mixed : Statement와 Row의 장점을 혼합한 로깅 방식
> 
> #### Ref.
> - https://omty.tistory.com/63
> - http://channy.creation.net/project/dev.kthcorp.com/2011/09/16/mysql-replication-binlog-format-mixed-vs-row/index.html
> - https://hyunki1019.tistory.com/105

## 11.6 UPDATE와 DELETE
- 보통은 하나의 테이블에서 단건 또는 여러 건의 레코드를 변경 또는 삭제하기 위해 사용
- MySQL 서버에서는 여러 테이블을 조인해서 한 개 이상 테이블의 레코드를 변경하거나 삭제하는 기능을 제공
  - JOIN UPDATE, JOIN DELETE

### 11.6.1 UPDATE ... ORDER BY ... LIMIT n
- 일반적으로 WHERE 절에 일치하는 모든 레코드를 업데이트
- ORDER BY 절과 LIMIT 절을 동시에 사용하면 특정 컬럼으로 정렬해서 상위 몇 건만 업데이트하는 것도 가능
  - 한 번에 너무 많은 레코드를 업데이트하면 MySQL 서버에 과부하를 유발하거나 다른 커넥션의 처리를 방해할 수 있음
- 바이너리 로그 포맷이 STATEMENT 로 설정되어 있는 경우 ORDER BY ... LIMIT이 포함된 UPDATE 나 DELETE 를 실행할 때 주의해야 함
  - ROW 일 때는 문제되지 않음
  - STATEMENT 로 설정되어 있는 겅우 ORDER BY에 의해 정렬되더라도 중복된 값의 순서가 소스 서버와 래플리카 서버에서 달라질 수 있기 때문
    - PK로 정렬하면 문제가 없음

### 11.6.2 JOIN UPDATE
- 두 개 이상의 테이블을 조인해서 조인된 결과 레코드를 변경 및 삭제하는 쿼리
- 조인된 테이블 중에서 특정 테이블의 컬럼값을 다른 테이블의 컬럼에 업데이트 해야할 때 주로 사용
  - 또는 다른 테이블의 컬럼 값을 참조하지 않더라도 조인되는 양쪽 테이블에 공통으로 존재하는 레코드만 찾아서 업데이트하는 용도로 사용할 수도 있음
- JOIN UPDATE로 조인되는 모든 테이블에 대해 읽기 참조만 되는 테이블은 읽기 잠금이 걸리고, 컬럼이 변경되는 테이블은 쓰기 잠금이 걸림
  - 웹 서비스같은 OLTP 환경에서는 데드락을 유발할 가능성이 있기 때문에 빈번한 사용은 권장하지 않음
  - 배치 프로그램이나 통계용 UPDATE 문에서 유용하게 사용할 수 있음
- 추가적으로 테이블의 조인 순서에 따라 UPDATE 문의 성능이 달라질 수 있기 때문에 실행 계획을 먼저 확인해보는 것을 권장
- JOIN UPDATE 문에서는 GROUP BY 나 ORDER BY 절을 사용할 수 없음
  - 문법적으로 지원하지 않기 때문에 서브쿼리를 이용해서 처리
- JOIN 순서에 따라 성능이 달라지기 때문에 조인 순서를 잘 설정해야 함
  - STRAIGHT_JOIN 이나 JOIN_ORDER 옵티마이저 힌트를 사용해서 순서 지정

### 11.6.3 여러 레코드 UPDATE
- 일반적으로 하나의 UPDATE 문장으로 여러 개의 레코드를 업데이트하는 경우 아래와 같이 레코드들들 동일한 값으로만 업데이트할 수 있음
  ```sql
  UPDATE deparments SET emp_count = 10;
  UPDATE deparments SET emp_count = emp_count + 10;
  ```
  
- 8.0 버전부터는 레코드 생성 문법을 이용해 레코드 별로 서로 다른 값을 업데이트할 수 있음

  ```sql
  # 2건의 레코드((1,1), (2,4))를 가지는 임시 테이블 new_user_level 생성
  # new_user_level 임시 테이블과 user_level 테이블을 조인해서 user_lv 칼럼을 업데이트
  UPDATE user_level ul
      INNER JOIN(VALUES ROW(1,1),
                        ROW(2,4)) new_user_level (user_id, user_lv)
                              ON ul.user_id = new_user_level.user_id
      SET ul.user_lv = ul.user_lv + new_user_level.user_lv;
  ```
  
  - VALUES ROW(...), ROW(...) 문법을 사용하면 SQL 문장 내에서 임시 테이블을 생성하는 효과를 냄

### 11.6.4 JOIN DELETE
- JOIN DELETE 문을 사용하려면 단일 테이블의 DELETE 문과는 조금 다른 문법으로 쿼리를 작성
  
  ```sql
  DELETE e
  FROM employees e, dept_emp de, departments d
  WHERE e.emp_no = de.emp_no AND de.dept_no = d.dept_no AND d.dept_name = 'd001';
  ```
  
  - 일반적으로 DELETE 문은 DELETE FROM table ... 문법을 사용하지만, JOIN DELETE 문은 DELETE와 FROM 절 사이에 삭제할 테이블을 명시

  ```sql
  DELETE e, de, d
  FROM employees e, dept_emp de, departments d
  WHERE e.emp_no = de.emp_no AND de.dept_no = d.dept_no AND d.dept_name = 'd001';
  ```

  - 하나의 테이블에서만 레코드를 삭제할 수 있을 뿐만 아니라 여러 테이블에서 동시에 레코드를 삭제할 수 있음

- JOIN DELETE 역시 JOIN UPDATE 처럼 SELECT 쿼리로 변환하여 실행 계획을 확인할 수 있음
  - STRAIGHT_JOIN 이나 JOIN_ORDER 옵티마이저 힌트를 사용해서 순서 지정

## 11.7 스키마 조작(DDL)

### 11.7.1 온라인 DDL

### 11.7.2 데이터베이스 변경

### 11.7.3 테이블 스페이스 변경

### 11.7.4 테이블 변경

### 11.7.5 칼럼 변경

### 11.7.6 인덱스 변경

### 11.7.7 테이블 변경 묶음 실행

### 11.7.8 프로세스 조회 및 강제 종료

### 11.7.9 활성 트랜잭션 조회

## 11.8 쿼리 성능 테스트

### 11.8.1 쿼리의 성능에 영향을 미치는 요소
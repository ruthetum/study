# Replication lag
> Replication lag occurs when the slaves (or secondaries) cannot keep up with the updates occuring on the master (or primary).
> 
> Unapplied changes accumulate in the slaves’ relay logs and the version of the database on the slaves becomes increasingly different from that of the master.

복제 지연은 슬레이브 인스턴스가 마스터 인스턴스에서 발생하는 업데이트를 따라잡을 수 없을 때 발생

적용되지 않은 변경 사항이 슬레이브의 relay 로그에 남으면서, 슬레이브 인스턴스의 데이터와 마스터 인스턴스의 데이터가 차이가 점점 커짐

## 원인
- Long Query
- Write 쿼리량 증가
- Lock 이슈
- 슬레이브 인스턴스의 로드 증가

## 지표
> https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-metrics.html

`ReplicaLag`을 이용해서 확인 가능

### 함께 확인하면 좋은 지표
- lock wait
  - [aurora mysql wait events](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Managing.Tuning.wait-events.html)
  - https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/ams-waits.row-lock-wait.html
  - `synch/cond/innodb/row_lock_wait`
- history list length
  - https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/proactive-insights.history-list.html
  - `trx_rseg_history_len`
- temporary table
  - https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/proactive-insights.temp-tables.html
  - `Created_tmp_disk_tables`, `Created_tmp_tables`

## 해결책
- 쿼리 튜닝
  - Long Query 개선
- Read Replica 추가
  - 단순 slave instance 부하로 인한 경우 replica 조정
- transaction or lock 이슈 개선
  - innodb_flush_log_at_trx_commit 조정 (데이터 특성에 따라)
  - innodb_lock_wait_timeout 조정
  - read 요청 시 해당 세션에 대한 isolation level 조정
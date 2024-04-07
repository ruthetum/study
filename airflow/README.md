# Airflow
> - https://airflow.apache.org
> - https://github.com/wilump-labs/airflow-in-actions

## Airflow 소개
- 파이썬을 이용해 워크플로우를 만들고 관리할 수 있는 오픈소스 기반 워크플로우 관리 도구

## Airflow 특징
- 파이썬을 제작된 도구이며 사용자가 워크플로우 생성 시 파이썬으로 구현
- 하나의 워크플로우는 DAG(Directed Acyclic Graph)이라 부르며 DAG 안에는 1개 이상의 Task가 존재
- Task간 선후행 연결이 가능하되 순환되지 않고 방향성을 가짐
- Cron 기반의 스케줄링
- 모니터링 및 실패 작업에 대한 재실행 기능이 간편

## Airflow 단점
- 실시간 워크플로우 관리에 적합하지 않음(최소 분 단위 실행)
- 워크플로우(DAG) 개수가 많아질 경우 모니터링이 쉽지 않음
- 워크플로우 GUI 환경에서 만들지 않기에 파이썬에 익숙하지 않다면 다루기 쉽지 않음
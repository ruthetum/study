# Emerging Architectures for Modern Data Infrastructure
> https://a16z.com/emerging-architectures-for-modern-data-infrastructure/

## 데이터 인프라의 목적
1. 데이터 기반의 결정 (비즈니스 리더들의 의사결정을 도움)
2. 데이터 기반의 제품 강화 (서비스/제품을 데이터의 도움을 받아 향상)

### 데이터 인프라 기초
#### Production Systems
- 각각의 도구에 데이터가 저장되기 때문에 각각의 분석 도구가 필요함
  - ERP, CRM, Database(MySQL, PostgreSQL, Oracle, etc)
- Normalized Schema -> many, small tables

#### Data Warehouse
- 통합된 분석 보고서 작성을 위해 다양한 소스로 부터 데이터를 저장
- Dimensional Schema -> fewer, simpler tables

#### ETL
- 데이터를 production system에서 data warehouse로 옮기는 과정
- Extract(추출) -> Transform(변환) -> Load(적재)

#### ELT
- ETL은 추출과 변환을 자동화하기 어려운 단점 존재, 이를 해결하기 ELT 등장
  - 데이터를 추출한 후 일단 적재, 그 후 변환
- Extract(추출) -> Load(적재) -> Transform(변환) 

## 통합 데이터 인프라 아키텍처

![image](https://github.com/ruthetum/study/assets/59307414/aca822de-3cda-4496-af82-a84615196f18)

![image](https://github.com/ruthetum/study/assets/59307414/e2c9f226-3909-4dba-9b3d-e72fca1e2566)


## Reference
- https://a16z.com/emerging-architectures-for-modern-data-infrastructure/
- https://www.youtube.com/watch?v=g_c742vW8dQ&list=PLL-_zEJctPoJ92HmbGxFv1Pv_ugsggGD2&index=1
- https://esevan.tistory.com/13#recentComments

--

https://www.youtube.com/watch?v=F-KDcb3FBuw&list=WL&index=21
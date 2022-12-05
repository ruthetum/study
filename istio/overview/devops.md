# DevOps
## 등장
- 인프라/운영팀(안정성 강조)과 개발팀(신속성 강조)의 목표가 다르다보니 충돌하는 영역 발생
- MSA 흐름에서 와서 개발팀과 운영팀의 사이클에 깊게 관여하는 DevOps 문화 대두

## 개발 문화
- 애자일 개발문화가 Ops까지 연결되는 지점
- 개발팀과 운영팀의 경계를 허물고 디벨롭하는 사이클
- 개발 라이프사이클의 개선
    - 지속적 통합(CI)과 지속적 배포(CD)를 중심으로 운영/피드백 과정까지 참여

# SRE(Site Reliability Engineering)
- 구글의 DevOps 구현
- 사이트 안정성에 무게중심
- 가용성에 대한 목표 정의
- 장애 발생에 대한 대응/분석/부검

## SRE 목표
### Metrics & Monitoring
- SLI(Service Level Indicator) 지표 정의
- SLO(Service Level Objective): SLI 지표들의 Goal

```
Ex.
- SLI: 메인 서비시의 응답 시간
- SLO: 우리 서비스의 매주 99% 해당 서비스의 응답시간 100ms 이하여야 한다.
```

### Capacity Planning
- 시스템 자원의 확보
- 자원 활용의 효율성 / 소프트웨어 성능 / 시스템 안정성

### Change Management
- 애자일 문화의 보편화
- 시스템 장애 원인의 대부분은 변경 작업 때문
    - 언제든 롤백할 수 있게
- 점진적인 배포와 변경 / 빠른 원인 분석
- 회복성(Resilience)

### Emergency Response
- 장애 처리에 대해 빠른 반응
- 사람의 복구가 아닌 시스템에 의한 복구
- 장애 분석 및 재발 방지에 대한 플랜

### Culture
- Toil Management (Operation 관리)
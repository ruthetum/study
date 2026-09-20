# Herdr 활용 가이드 — 주요 기능·장단점·사용법·팁

작성일: 2026-09-20  
대상 제품: [Herdr](https://herdr.dev/), [herdrdev/herdr](https://github.com/herdrdev/herdr)  
기준: 현재 공식 문서와 제공된 Herdr 관련 영상. 설치·성능 벤치마크는 수행하지 않았다.

## 1. 어떤 도구인가

Herdr는 기존 터미널에서 실행하는 에이전트 중심 멀티플렉서다. 여러 터미널을 묶고 그 안의 에이전트 상태를 보여주며, 백그라운드 서버가 실행 프로세스를 유지한다. Ghostty·iTerm·Windows Terminal 같은 터미널 앱 안에서 사용하는 계층이다.

Rust 기반이며 Apache-2.0 라이선스다. 기존 Claude Code·Codex 등 CLI를 실행하므로 에이전트 자체의 기능·인증·모델 이용 한도는 별도로 적용된다. [공식 저장소](https://github.com/herdrdev/herdr)

**추천 대상:** 기존 IDE를 유지하면서 여러 CLI 에이전트를 운영하려는 개발자, 원격 서버 작업이 많은 사람, 터미널 창을 수시로 닫거나 연결을 바꾸는 사람.

**처음부터 필요하지 않은 경우:** 한 프로젝트에서 에이전트 하나만 돌리는 경우. 도구보다 작업 분리와 검증 습관부터 갖추는 편이 효과적일 수 있다.

## 2. 화면 구조와 상태

| 단위 | 의미 | 구성 예시 |
|---|---|---|
| Session | 별도의 지속 런타임 공간 | 회사 작업 / 개인 작업 |
| Workspace | 프로젝트·작업 묶음 | backend-api / mobile-app |
| Tab | 목적별 화면 구성 | implementation / review / server |
| Pane | 실제 터미널 하나 | Claude / Codex / 테스트 명령 |
| Agent | 페인 안에서 감지한 코딩 에이전트 | 구현 담당 / 검토 담당 |

처음에는 기본 Session 하나와 프로젝트별 Workspace만 사용해도 충분하다. Session은 완전히 분리된 런타임이 필요할 때 추가한다. [공식 개념](https://herdr.dev/docs/concepts/)

| 상태 | 의미 | 권장 대응 |
|---|---|---|
| `working` | 작업 중 | 기다리거나 필요한 경우 진행 확인 |
| `blocked` | 질문·권한·결정 입력 필요 | 해당 화면을 읽고 응답 |
| `done` | 입력 가능한 완료 상태, 아직 확인되지 않음 | 결과와 검증 내용을 읽음 |
| `idle` | 확인된 대기·완료 상태 | 다음 작업 가능 여부 확인 |
| `unknown` | 상태를 확실히 분류하지 못함 | 터미널 내용 직접 확인 |

**`done`은 테스트 통과나 요구사항 충족의 인증이 아니다.** 상태는 실행·입력 준비 상태이며, 사람이 봤는지에 따라 `done`과 `idle`이 달라질 수 있다. [공식 상태 정의](https://herdr.dev/docs/concepts/)

## 3. 주요 기능과 장단점

| 기능·장점 | 활용 가치 |
|---|---|
| 에이전트 목록·상태 표시 | 여러 창을 열어 읽지 않고 개입할 작업을 찾음 |
| 백그라운드 서버 | 클라이언트 분리 중에도 작업을 유지 |
| 마우스와 prefix 단축키 | 클릭으로 시작하고 자주 쓰는 조작만 단축키로 전환 |
| CLI·socket API·스킬 | 다른 에이전트 생성·입력·상태 관찰·결과 회수 |
| 원격 머신 연결 | 로컬과 SSH 머신의 작업을 관리 |
| 플러그인 확장 | 파일 보기·보조 워크플로우 등을 추가 |
| 기존 도구와 조합 | IntelliJ 등 익숙한 IDE를 계속 사용 가능 |

기능 근거: [공식 소개](https://herdr.dev/), [공식 저장소](https://github.com/herdrdev/herdr). 마지막 항목은 도입 방식에 관한 제안이다.

| 단점·한계 | 대응 |
|---|---|
| 파일 편집·브라우저·협업 규칙을 하나의 앱으로 모두 해결하는 방식은 아님 | 필요한 IDE·브라우저만 병행 |
| 여러 페인 생성만으로 파일이 격리되지 않음 | 수정 작업은 별도 worktree와 담당 범위로 분리 |
| 상태 감지가 UI 변경에 영향을 받을 수 있음 | 새 에이전트 버전 도입 때 대기·완료 표시 확인 |
| 직접 구성한 협업 흐름의 책임이 남음 | 의존성·타임아웃·재시도·결과 형식을 지시 |
| 플러그인이 늘면 유지보수도 증가 | 실제 불편을 해결하는 확장부터 하나씩 추가 |
| 실행 호스트가 꺼지면 계산이 계속되지 않음 | 야간 작업은 켜져 있는 호스트에서 실행 |

제공 영상은 가벼움과 관리 편의를 높이 평가하지만 같은 환경에서의 Orca 대비 RAM·CPU 수치는 제시하지 않는다. [비교 영상](https://www.youtube.com/watch?v=VgngB-ZF8ac), [Herdr 가이드 영상](https://www.youtube.com/watch?v=rSHizW6dI3c)

## 4. 설치와 첫 실행

### 4.1 macOS

Homebrew를 사용한다면:

```bash
brew install herdr
```

공식 직접 설치 방식(macOS·Linux):

```bash
curl -fsSL https://herdr.dev/install.sh | sh
```

Windows도 stable 바이너리를 제공한다. 플랫폼별 설치 명령과 ZIP 설치 방법은 [공식 설치 문서](https://herdr.dev/docs/install/)를 따른다. 직접 설치와 패키지 관리자 설치를 중복하지 않는 편이 관리하기 쉽다.

### 4.2 프로젝트에서 실행

```bash
cd /path/to/project
herdr
```

Herdr 화면의 터미널에서 기존에 사용하던 에이전트를 시작한다.

```bash
claude
```

다른 페인에서는 `codex` 등 필요한 CLI를 실행한다. 에이전트의 설치·인증은 각각 먼저 완료되어 있어야 한다.

초기 화면 구성 제안:

| Tab | Pane 구성 |
|---|---|
| implementation | 주 구현 에이전트 |
| review | 읽기 중심 검토 에이전트 |
| server | 개발 서버와 로그 |

### 4.3 작업을 유지한 채 나오기

`Ctrl+B`를 누르고 손을 뗀 다음 `q`를 누르면 detach한다. 나중에 `herdr`를 실행해 재접속한다. **페인 닫기나 서버 종료와 혼동하지 않는다.** [공식 키보드 안내](https://herdr.dev/docs/keyboard/)

## 5. 자주 쓰는 단축키

기본 prefix는 `Ctrl+B`다. 아래 `prefix → c`는 `Ctrl+B`를 누른 뒤 놓고 `c`를 누르라는 뜻이다.

| 작업 | 기본 키 |
|---|---|
| 단축키 도움말 | `prefix → ?` |
| 새 Tab | `prefix → c` |
| 오른쪽 분할 | `prefix → v` |
| 아래쪽 분할 | `prefix → -` |
| Pane 이동 | `prefix → h / j / k / l` |
| 현재 Pane 확대·복원 | `prefix → z` |
| Workspace 탐색 | `prefix → w` |
| 다음 / 이전 Tab | `prefix → n / p` |
| 클라이언트 분리 | `prefix → q` |
| Pane 닫기 | `prefix → x` — 실행 중 작업을 끝내려는 경우에만 사용 |

마우스로 탭·페인을 선택하고 우클릭 메뉴와 분할 경계 드래그를 사용할 수 있다. 처음부터 단축키를 모두 외울 필요는 없다. [공식 키보드 문서](https://herdr.dev/docs/keyboard/)

## 6. 스킬과 Integration의 차이

| 구분 | 역할 |
|---|---|
| Herdr 스킬 | 에이전트에게 Herdr를 조작하는 방법을 알려줌 |
| Agent integration | 지원 에이전트의 세션 식별·상태 관련 신호를 연동 |
| 플러그인 | 사용자 기능이나 워크플로우 확장 |

스킬 설치:

```bash
npx skills add herdrdev/herdr --skill herdr -g
```

`-g`는 전역 설치다. 현재 프로젝트에만 설치하려면 생략한다. 설치된 바이너리에 맞는 스킬을 읽으려면:

```bash
herdr --skill
```

에이전트는 Herdr가 관리하는 페인 안에서 실행해야 한다. 공식 스킬은 `HERDR_ENV=1`을 확인한다. **밖에서 값을 억지로 설정해 내부 세션인 것처럼 만들라는 뜻이 아니다.** [공식 스킬 가이드](https://herdr.dev/docs/agent-skill/)

사용 중인 integration 확인과 설치 예:

```bash
herdr integration status
herdr integration install claude
herdr integration install codex
```

Integration이 모든 에이전트에서 동일한 상태 감지 역할을 하는 것은 아니다. Claude Code·Codex 등은 현재 문서상 화면 규칙이 상태 판단을 담당하고, integration은 세션 복원을 위한 식별을 제공한다. [Integrations](https://herdr.dev/docs/integrations/), [Agents](https://herdr.dev/docs/agents/)

## 7. 에이전트 협업 사용법

### 7.1 “오케스트레이션이 없다”의 정확한 해석

Herdr에는 에이전트 시작·지시·출력 읽기·상태 대기 도구가 있다. 따라서 총괄 에이전트가 다른 페인을 만들고 협업하는 것이 가능하다. 다만 **업무를 어떤 순서로 나누고 무엇을 성공으로 볼지는 사용자가 설계**한다. [공식 자동화 가이드](https://herdr.dev/docs/agent-automation/)

### 7.2 자연어로 시작하는 권장 프롬프트

```text
Herdr 스킬을 사용해주세요. 현재 Herdr 안에서 실행 중인지 확인해주세요.

이 프로젝트의 PR 후보 변경을 검토하려고 합니다.
- 현재 에이전트는 총괄 역할입니다.
- 별도 Pane에 Codex 검토 담당을 실행해주세요.
- 검토 담당은 코드 변경 없이 버그·회귀·누락된 검증만 보고합니다.
- 검토 대상은 현재 작업 브랜치와 기준 브랜치의 diff입니다.
- 결과는 작업 전용 Markdown 파일에 저장하고 경로를 알려주세요.
- 사람이 직접 보는 Pane의 입력·포커스를 임의로 빼앗지 마세요.
- 막히면 원인을 보고하고, 무제한 재시도하지 마세요.
- 마지막에 심각도 / 근거 / 재현 방법 / 권장 조치로 정리해주세요.
```

이는 작성자의 예시다. 총괄이 검토 의견을 그대로 적용하기보다 실제 코드와 테스트로 재확인하도록 설계한다.

### 7.3 CLI로 검토 에이전트 실행하기

아래 예시는 **Herdr 내부의 셸**, 설치된 Codex CLI, `jq`가 있다는 가정이다. `reviewer`라는 이름이 기존 실행 에이전트와 겹치지 않아야 한다. 복사 전에 현재 버전의 `herdr --skill`과 도움말을 확인한다. 이 보고서에서 실행한 명령은 아니다.

```bash
split_result=$(herdr pane split --current --direction right --no-focus)
review_pane=$(printf '%s\n' "$split_result" | jq -r '.result.pane.pane_id')

herdr agent start reviewer --kind codex --pane "$review_pane"
herdr agent prompt reviewer "Review the current diff without editing files. Report bugs and missing verification." --wait --timeout 120000
herdr agent read reviewer --source recent-unwrapped --lines 120
```

명령·응답 필드는 [공식 자동화 예시](https://herdr.dev/docs/agent-automation/)를 바탕으로 단순화했다. Pane ID는 응답에서 받아 사용한다. `w1:p2`처럼 임의로 예상하지 않는다.

### 7.4 자동화에서 중요한 실패 처리

- `agent start`는 기존 셸 페인에서 실행한다. 에디터·다른 에이전트가 점유한 페인에 덮어 실행하지 않는다.
- `agent prompt --wait`가 반환되어도 결과 파일과 검증을 확인한다.
- 시간 초과는 입력이 전송되지 않았다는 뜻이 아니다. 재전송 전에 출력부터 읽는다.
- `blocked` 상태에는 질문·승인 화면이 있을 수 있다. 내용을 확인한 뒤 의도적으로 응답한다.
- 긴 결과는 터미널 스크롤백만 읽기보다 파일로 저장하도록 요청한다.

명령 동작 근거: [CLI reference](https://herdr.dev/docs/cli-reference/). 재시도·결과 검증 절차는 이를 반영한 운영 제안이다.

## 8. 지속성과 원격 사용

| 상황 | 실행 중 프로세스 | 복원·운영 방식 |
|---|---|---|
| 클라이언트 detach | 유지 | 재접속해서 계속 사용 |
| SSH 연결 단절 | 원격 서버가 살아 있으면 유지 | 다시 연결 |
| Herdr 서버 재시작 | 기존 프로세스 종료 | 레이아웃 및 지원되는 세션 복원 경로 사용 |
| 컴퓨터 재부팅 | 기존 프로세스 종료 | 원래 프로세스의 생존과 대화 재개는 별개 |
| 실행 머신 절전·전원 차단 | 정상적인 지속 실행 불가 | 켜져 있는 실행 호스트 필요 |

화면 기록 복원은 별도 설정이며, 지원 에이전트의 대화 재개도 integration 조건을 만족해야 한다. 임의의 개발 서버·테스트 프로세스가 재부팅 뒤 이어지는 것은 아니다. [공식 세션 복원](https://herdr.dev/docs/session-state/)

SSH 설정에 `workbox`라는 호스트가 등록되어 있다는 가정의 원격 접속 예:

```bash
herdr --remote workbox
```

현재는 저장한 원격 머신을 관리하는 기능도 제공한다. 로컬과 원격의 파일 위치·CLI·인증·개발 의존성은 각각 확인한다. 화면에서 다른 머신을 선택했다고 기존 페인의 명령 대상까지 자동으로 바뀐다고 가정하지 않는다. [머신 연결](https://herdr.dev/docs/connecting-machines/), [CLI reference](https://herdr.dev/docs/cli-reference/)

**야간 작업 제안:** 목표·산출물·종료 기준·재시도 상한을 먼저 적고 실행한다. 아침 보고서에는 실제 구현, 분석만 완료, 미검증, 사용자 결정 필요를 나누도록 한다. [실전 영상](https://www.youtube.com/watch?v=U9AConAimzM)의 성능 분석·버그 수정·UX 제안은 산출물 종류를 구분해야 한다는 좋은 사례다.

## 9. 플러그인과 모바일

Herdr는 플러그인으로 파일 뷰어·보조 UI·알림 등을 확장할 수 있다. 다만 커뮤니티 확장을 공식 기본 기능과 구분한다. 영상에서 소개한 파일 뷰어의 정확한 패키지는 자막만으로 확정하지 않아 설치 명령을 임의로 제시하지 않았다.

현재 공개 저장소에는 `herdr-sidebar`, `herdr-browser`, 모바일·원격 제어 확장 등이 등록되어 있다. 목록이 있다는 사실만으로 유지보수·안정성·모바일 알림 품질이 검증된 것은 아니다. 필요한 기능이 생길 때 [공식 사이트의 Plugins](https://herdr.dev/)와 [플러그인 저장소 목록](https://github.com/topics/herdr-plugin)에서 해당 프로젝트의 README를 확인한다.

**선택 팁:** 모바일 사용까지 바로 시작하려면 별도의 SSH 클라이언트나 확장의 연결 방식·알림 경로를 검토해야 한다. Orca companion과 같은 경험이 Herdr 기본 설치에 포함된다고 가정하지 않는다.

## 10. Kotlin/Spring·RN 적용 예시

아래 구성은 실무 제안이다.

### Kotlin/Spring

| 작업 | 역할과 검증 |
|---|---|
| API 구현 | 별도 worktree에서 구현·단위 테스트 |
| 요구사항 검토 | nullable·empty·실패·기준시각·백필 정책 검토 |
| 데이터 처리 검토 | 중복 이벤트·재시도·동시성 검토 |
| 통합 | 한 담당자가 변경 병합 후 관련 모듈 검증 |

Herdr는 작업 상태 관리에 집중하고 IntelliJ는 코드 탐색·디버깅에 활용한다. 테스트를 여러 개 동시에 돌릴 경우 공용 DB·Redis·Kafka 자원이 간섭하지 않는지 확인한다.

### RN

| 작업 | 역할과 검증 |
|---|---|
| 기능 구현 | 작업별 코드 변경, 수정 파일 소유권 지정 |
| 로컬 데이터 | 마이그레이션·업그레이드 경로·기존 데이터 보존 검토 |
| UI 검토 | 구현된 화면과 요구사항 비교 |
| 플랫폼 검증 | iOS·Android별 실행 결과와 미검증 영역 기록 |

Pane는 에이전트 화면을 나눌 뿐이다. Metro 포트, 네이티브 빌드, 시뮬레이터와 worktree가 어떤 조합인지 이름으로 표시해두면 혼동을 줄일 수 있다.

## 11. 문제 해결과 유지보수

| 증상 | 확인·대응 |
|---|---|
| `herdr`를 찾지 못함 | 설치 경로·PATH 확인 후 터미널 재시작 |
| 에이전트가 목록에 없음 | 해당 페인의 실제 전경 프로세스와 지원 여부 확인 |
| `idle`인데 질문이 보임 | 감지 규칙이 새 화면을 모를 수 있으므로 실제 화면 기준으로 판단 |
| 상태가 이상함 | `herdr agent explain <target>`으로 판단 근거 확인 |
| 재접속했는데 셸만 나타남 | 서버 재시작 여부와 native session integration 확인 |
| 협업 에이전트가 멈춤 | 질문·권한·한도·오류 내용을 읽고 타임아웃 처리 |
| 결과가 잘림 | 작업별 Markdown 보고서를 요청 |
| 새 CLI 명령이 안 됨 | 설치된 클라이언트와 실행 서버의 버전 차이 확인 |

특히 새로운 질문 UI는 `blocked` 대신 `idle`로 표시될 수 있다고 공식 문서가 설명한다. [상태 감지·진단](https://herdr.dev/docs/agents/)

업데이트는 설치 방식에 맞춘다.

```bash
# Homebrew 설치인 경우
brew upgrade herdr

# 공식 직접 설치인 경우
herdr update
```

현재 업데이트는 호환되는 기존 서버를 유지할 수 있다. 새 서버 기능 반영에 재시작이 필요할 수 있으므로 장기 실행 작업이 끝난 뒤 전환하는 편이 좋다. [공식 설치·업데이트](https://herdr.dev/docs/install/)

## 12. 처음 일주일의 운영 기준

1. 에이전트 2개로 시작하고 프로젝트·작업 이름을 명확히 한다.
2. 구현과 읽기 전용 검토부터 분리한다.
3. 수정 작업이 늘어날 때 worktree와 공용 자원을 구분한다.
4. 상태 배지 다음에 실제 산출물·검증 결과를 확인한다.
5. 재시도·토큰 사용보다 사람의 검토가 밀리면 병렬도를 낮춘다.
6. 필요한 플러그인만 하나씩 추가한다.

성공 기준은 띄운 에이전트 수가 아니라 **기다리는 작업을 빨리 발견하고, 검증 가능한 결과를 안정적으로 회수하는 것**이다.

# Chain of responsibility pattern
- 역할 사슬 또는 책임 연쇄 패턴이라고 불림
- 요청을 받는 객체를 연쇄적으로 묶어 요청을 처리하는 객체를 만날 때까지 객체 Chain을 따라 요청을 전달

## Structure
![image](https://refactoring.guru/images/patterns/diagrams/chain-of-responsibility/structure-2x.png)

## Example
- 책에서는 메일 분류기
  1. 메일 수신
  2. 스펨 메일 감지
  3. 팬 메일 감지
  4. 힝의 메일 감지
  5. 신규 설치 요청 메일 감지

```go
type SpamHandler struct {
	target string
	next   Handler
}

func NewSpamHandler() *SpamHandler {
	return &SpamHandler{target: "spam"}
}

func (h *SpamHandler) Execute(mailType string) {
	if mailType == h.target {
		fmt.Println("process spam mail")
		return
	}
	if h.next != nil {
		h.next.Execute(mailType)
	}
}

func (h *SpamHandler) SetNext(handler Handler) {
	h.next = handler
}
```

## Reference
- https://refactoring.guru/ko/design-patterns/chain-of-responsibility
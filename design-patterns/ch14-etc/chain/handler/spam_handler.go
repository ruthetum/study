package handler

import "fmt"

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

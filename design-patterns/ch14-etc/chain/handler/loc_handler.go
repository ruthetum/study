package handler

import "fmt"

type LocHandler struct {
	target string
	next   Handler
}

func NewLocHandler() *LocHandler {
	return &LocHandler{target: "loc"}
}

func (h *LocHandler) Execute(mailType string) {
	if mailType == h.target {
		fmt.Println("process loc mail")
		return
	}
	if h.next != nil {
		h.next.Execute(mailType)
	}
}

func (h *LocHandler) SetNext(handler Handler) {
	h.next = handler
}

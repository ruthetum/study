package handler

import "fmt"

type FanHandler struct {
	target string
	next   Handler
}

func NewFanHandler() *FanHandler {
	return &FanHandler{target: "fan"}
}

func (h *FanHandler) Execute(mailType string) {
	if mailType == h.target {
		fmt.Println("process fan mail")
		return
	}
	if h.next != nil {
		h.next.Execute(mailType)
	}
}

func (h *FanHandler) SetNext(handler Handler) {
	h.next = handler
}

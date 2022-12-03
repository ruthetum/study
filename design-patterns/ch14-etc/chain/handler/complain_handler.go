package handler

import "fmt"

type ComplainHandler struct {
	target string
	next   Handler
}

func NewComplainHandler() *ComplainHandler {
	return &ComplainHandler{target: "complain"}
}

func (h *ComplainHandler) Execute(mailType string) {
	if mailType == h.target {
		fmt.Println("process complain mail")
		return
	}
	if h.next != nil {
		h.next.Execute(mailType)
	}
}

func (h *ComplainHandler) SetNext(handler Handler) {
	h.next = handler
}

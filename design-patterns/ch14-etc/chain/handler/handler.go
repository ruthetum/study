package handler

type Handler interface {
	Execute(mailType string)
	SetNext(handler Handler)
}

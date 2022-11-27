package server

type Server interface {
	HandleRequest(string, string) (int, string)
}

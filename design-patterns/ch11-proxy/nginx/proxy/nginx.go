package proxy

import "nginx/server/http"

type Nginx struct {
	server            *http.Server
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func NewNginx() *Nginx {
	return &Nginx{
		server:            http.NewServer(),
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func (n *Nginx) HandleRequest(url, method string) (int, string) {
	allowed := n.CheckRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	return n.server.HandleRequest(url, method)
}

func (n *Nginx) CheckRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}
	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}

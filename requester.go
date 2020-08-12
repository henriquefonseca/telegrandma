package telegrandma

import "net/http"

// requester is an interface for http clients
type requester interface {
	Get(url string, headers map[string]string) (*http.Response, error)
}

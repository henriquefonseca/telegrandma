package telegrandma

import "net/http"

// Requester is an interface for http clients
type Requester interface {
	Get(url string, headers map[string]string) (*http.Response, error)
}

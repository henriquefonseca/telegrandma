package telegrandma

import "net/http"

// interface for http clients
type Requester interface {
	Get(url string, headers map[string]string) (*http.Response, error)
}

package telegrandma

import "net/http"

type Requester interface {
	Get(url string, headers map[string]string) (*http.Response, error)
}

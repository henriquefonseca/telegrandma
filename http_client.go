package telegrandma

import (
	"log"
	"net/http"
)

// Http client struct for requests
type HttpClient struct{}

// Get is a function for making requests using GET Http method.
//
// It receives an url and headers and creates a new GET request, returning *http.Response.
//
// If headers is null, headers default will be applied
func (hc *HttpClient) Get(url string, headers map[string]string) (*http.Response, error) {
	log.Printf("Requesting url %v\n", url)

	t := &http.Transport{}
	client := &http.Client{Transport: t}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error on requesting: %v", err)
	}

	if headers == nil {
		headers = newDefaultHeaders().Headers
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	return client.Do(req)
}

type defaultHeaders struct {
	Headers map[string]string
}

func newDefaultHeaders() defaultHeaders {
	headers := map[string]string{
		"User-Agent":      `GrandmaClient`,
		"Accept":          `text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3`,
		"Accept-Encoding": `gzip, deflate, br`,
		"Accept-Language": `pt-BR,pt;q=0.9,en-US;q=0.8,en;q=0.7`,
	}

	return defaultHeaders{Headers: headers}
}

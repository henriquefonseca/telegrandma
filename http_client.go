package telegrandma

import (
	"log"
	"net/http"
)

type HttpClient struct{}

func (hc *HttpClient) Get(url string, headers map[string]string) (*http.Response, error) {
	log.Printf("Requesting url %v\n", url)

	t := &http.Transport{}
	client := &http.Client{Transport: t}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error on requesting: %v", err)
	}

	if headers == nil {
		headers = NewDefaultHeaders().Headers
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	return client.Do(req)
}

type DefaultHeaders struct {
	Headers map[string]string
}

func NewDefaultHeaders() DefaultHeaders {
	headers := map[string]string{
		"User-Agent":      `GrandmaClient`,
		"Accept":          `text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3`,
		"Accept-Encoding": `gzip, deflate, br`,
		"Accept-Language": `pt-BR,pt;q=0.9,en-US;q=0.8,en;q=0.7`,
	}

	return DefaultHeaders{Headers: headers}
}

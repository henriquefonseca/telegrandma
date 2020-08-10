package telegrandma

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// HttpClientMock: Http client mock for tests
type HttpClientMock struct {
	ExpectedUrl            string
	ResponseBody           string
	ResponseHttpStatusCode int
}

// Get: is a function for mocking GET requests.
//
// It receives an url and headers and compares url with the string received by SetExpectedUrl function.
func (hcMock *HttpClientMock) Get(url string, headers map[string]string) (*http.Response, error) {
	log.Printf("Mocking request to url %v\n", url)

	if hcMock.ExpectedUrl != "" && hcMock.ExpectedUrl != url {
		return &http.Response{
				StatusCode: 404,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{}`))),
			},
			fmt.Errorf("Expected [%s] url but received [%s]", hcMock.ExpectedUrl, url)
	}

	if hcMock.ResponseHttpStatusCode == 0 {
		hcMock.SetResponseHttpStatusCode(200)
	}

	if hcMock.ResponseBody == "" {
		hcMock.SetResponseBody(`{}`)
	}

	r := ioutil.NopCloser(bytes.NewReader([]byte(hcMock.ResponseBody)))

	return &http.Response{
		StatusCode: hcMock.ResponseHttpStatusCode,
		Body:       r,
	}, nil
}

// SetExpectedUrl: sets the url HttpClientMock expects to receive
func (hcMock *HttpClientMock) SetExpectedUrl(expectedUrl string) {
	hcMock.ExpectedUrl = expectedUrl
}

// SetResponseHttpStatusCode: sets the http status code for response
func (hcMock *HttpClientMock) SetResponseHttpStatusCode(responseHttpStatusCode int) {
	hcMock.ResponseHttpStatusCode = responseHttpStatusCode
}

// SetResponseBody sets: the response body
func (hcMock *HttpClientMock) SetResponseBody(responseBody string) {
	hcMock.ResponseBody = responseBody
}

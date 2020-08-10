package telegrandma

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// HTTPClientMock is the Http client mock for tests
type HTTPClientMock struct {
	ExpectedURL            string
	ResponseBody           string
	ResponseHTTPStatusCode int
}

// Get is a function for mocking GET requests.
//
// It receives an url and headers and compares url with the string received by SetExpectedURL function.
func (hcMock *HTTPClientMock) Get(url string, headers map[string]string) (*http.Response, error) {
	log.Printf("Mocking request to url %v\n", url)

	if hcMock.ExpectedURL != "" && hcMock.ExpectedURL != url {
		return &http.Response{
				StatusCode: 404,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{}`))),
			},
			fmt.Errorf("Expected [%s] url but received [%s]", hcMock.ExpectedURL, url)
	}

	if hcMock.ResponseHTTPStatusCode == 0 {
		hcMock.SetResponseHTTPStatusCode(200)
	}

	if hcMock.ResponseBody == "" {
		hcMock.SetResponseBody(`{}`)
	}

	r := ioutil.NopCloser(bytes.NewReader([]byte(hcMock.ResponseBody)))

	return &http.Response{
		StatusCode: hcMock.ResponseHTTPStatusCode,
		Body:       r,
	}, nil
}

// SetExpectedURL sets the url HTTPClientMock expects to receive
func (hcMock *HTTPClientMock) SetExpectedURL(expectedURL string) {
	hcMock.ExpectedURL = expectedURL
}

// SetResponseHTTPStatusCode sets the http status code for response
func (hcMock *HTTPClientMock) SetResponseHTTPStatusCode(responseHTTPStatusCode int) {
	hcMock.ResponseHTTPStatusCode = responseHTTPStatusCode
}

// SetResponseBody sets the response body
func (hcMock *HTTPClientMock) SetResponseBody(responseBody string) {
	hcMock.ResponseBody = responseBody
}

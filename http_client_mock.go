package telegrandma

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// hTTPClientMock is the Http client mock for tests
type hTTPClientMock struct {
	ExpectedURL            string
	ResponseBody           string
	ResponseHTTPStatusCode int
}

// Get is a function for mocking GET requests.
//
// It receives an url and headers and compares url with the string received by setExpectedURL function.
func (hcMock *hTTPClientMock) Get(url string, headers map[string]string) (*http.Response, error) {
	log.Printf("Mocking request to url %v\n", url)

	if hcMock.ExpectedURL != "" && hcMock.ExpectedURL != url {
		return &http.Response{
				StatusCode: 404,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{}`))),
			},
			fmt.Errorf("(hTTPClientMock): expected url [%s] but got [%s]", hcMock.ExpectedURL, url)
	}

	if hcMock.ResponseHTTPStatusCode == 0 {
		hcMock.setResponseHTTPStatusCode(200)
	}

	if hcMock.ResponseBody == "" {
		hcMock.setResponseBody(`{}`)
	}

	r := ioutil.NopCloser(bytes.NewReader([]byte(hcMock.ResponseBody)))

	return &http.Response{
		StatusCode: hcMock.ResponseHTTPStatusCode,
		Body:       r,
	}, nil
}

// setExpectedURL sets the url hTTPClientMock expects to receive
func (hcMock *hTTPClientMock) setExpectedURL(expectedURL string) {
	hcMock.ExpectedURL = expectedURL
}

// setResponseHTTPStatusCode sets the http status code for response
func (hcMock *hTTPClientMock) setResponseHTTPStatusCode(responseHTTPStatusCode int) {
	hcMock.ResponseHTTPStatusCode = responseHTTPStatusCode
}

// setResponseBody sets the response body
func (hcMock *hTTPClientMock) setResponseBody(responseBody string) {
	hcMock.ResponseBody = responseBody
}

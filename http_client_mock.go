package telegrandma

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type HttpClientMock struct {
	ExpectedUrl            string
	ResponseBody           string
	ResponseHttpStatusCode int
}

func (hcMock *HttpClientMock) Get(url string, headers map[string]string) (*http.Response, error) {
	log.Printf("Mocking request to url %v\n", url)

	if hcMock.ExpectedUrl != "" && hcMock.ExpectedUrl != url {
		return &http.Response{
				StatusCode: 404,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{}`))),
			},
			errors.New(fmt.Sprintf("Expected [%s] url but received [%s]", hcMock.ExpectedUrl, url))
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

func (hcMock *HttpClientMock) SetExpectedUrl(expectedUrl string) {
	hcMock.ExpectedUrl = expectedUrl
}

func (hcMock *HttpClientMock) SetResponseHttpStatusCode(responseHttpStatusCode int) {
	hcMock.ResponseHttpStatusCode = responseHttpStatusCode
}

func (hcMock *HttpClientMock) SetResponseBody(responseBody string) {
	hcMock.ResponseBody = responseBody
}

package client

import (
	"io"
	"net/http"
)

type authTransport struct {
	apiToken string
	base     http.RoundTripper
}

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req2 := req.Clone(req.Context())
	req2.Header.Set("Authorization", "Bearer "+t.apiToken)
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("Accept", "application/json")
	return t.base.RoundTrip(req2)
}

func (ac *ArtifactsMMOClient) CreateRequest(method string, url string, body io.Reader) *http.Request {
	request, err := http.NewRequest(method, ac.baseUrl+url, body)
	if err != nil {
		panic(err)
	}
	return request
}

func (ac *ArtifactsMMOClient) CloseBody(response *http.Response) {
	err := response.Body.Close()
	if err != nil {
		panic(err)
	}
}

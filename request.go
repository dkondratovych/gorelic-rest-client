package newrelic

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type NewRelicRequest struct {
	Config *RequestConfig
}

type Params map[string]interface{}

type RequestConfig struct {
	Method string
	URL    *url.URL
	Header http.Header
	Body   *strings.Reader
}

func (r *NewRelicRequest) NewRequest() (*http.Request, error) {
	request, err := http.NewRequest(
		r.Config.Method,
		r.Config.URL.String(),
		ioutil.NopCloser(r.Config.Body),
	)

	if err != nil {
		return nil, err
	}

	request.Header = r.Config.Header

	return request, nil
}

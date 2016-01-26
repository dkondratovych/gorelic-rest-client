package newrelic

import (
	"net/url"

	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	dataStoreMemcacheApiPath string = "/v2/applications/%d/metrics/data.%s"
)

type DataStoreMemcache struct {
	NewRelicRequest
}

type DataStoreMemcacheData struct {
	MetricData struct {
		From    time.Time `json:"from"`
		To      time.Time `json:"to"`
		Metrics []struct {
			Name       string `json:"name"`
			Timeslices []struct {
				From   time.Time `json:"from"`
				To     time.Time `json:"to"`
				Values struct {
					CallCount           int     `json:"call_count"`
					AverageResponseTime float32 `json:"average_response_time"`
				} `json:"values"`
			} `json:"timeslices"`
		} `json:"metrics"`
	} `json:"metric_data"`
}

func (a *DataStoreMemcacheData) CallCount() int {
	return a.MetricData.Metrics[0].Timeslices[0].Values.CallCount
}

func (a *DataStoreMemcacheData) AverageResponseTime() float32 {
	return a.MetricData.Metrics[0].Timeslices[0].Values.AverageResponseTime
}

type DataStoreMemcacheConfig struct {
	Params Params
}

func (ac DataStoreMemcacheConfig) Init() (*RequestConfig, error) {
	// Full fill url template
	applicationId, _ := ac.Params["ApplicationId"].(int)
	responseFormat, _ := ac.Params["ResponseFormat"].(string)
	apiKey, _ := ac.Params["ApiKey"].(string)

	path := fmt.Sprintf(browserNetworkRespTimeApiPath, applicationId, responseFormat)

	// Full fill body template
	startDate, _ := ac.Params["StartDate"].(string)
	endDate, _ := ac.Params["EndDate"].(string)

	body := url.Values{}
	body.Add("names[]", "Datastore/Memcached/allWeb")
	body.Add("values[]", "average_response_time")
	body.Add("values[]", "call_count")

	body.Add("from", startDate)
	body.Add("to", endDate)
	body.Add("summarize", "true")

	header := http.Header{}
	header.Add("X-Api-Key", apiKey)
	header.Add("Content-Type", "application/x-www-form-urlencoded")
	header.Add("Content-Length", strconv.Itoa(len(body.Encode())))

	requestConfig := &RequestConfig{
		Method: "POST",
		URL: &url.URL{
			Scheme: "https",
			Host:   "api.newrelic.com",
			Path:   path,
		},
		Header: header,
		Body:   strings.NewReader(body.Encode()),
	}

	return requestConfig, nil
}

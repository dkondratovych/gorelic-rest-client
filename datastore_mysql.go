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
	dataStoreMysqlApiPath string = "/v2/applications/%d/metrics/data.%s"
)

type DataStoreMysql struct {
	NewRelicRequest
}

type DataStoreMysqlData struct {
	MetricData struct {
		From    time.Time `json:"from"`
		To      time.Time `json:"to"`
		Metrics []struct {
			Name       string `json:"name"`
			Timeslices []struct {
				From   time.Time `json:"from"`
				To     time.Time `json:"to"`
				Values struct {
					AverageResponseTime float32 `json:"average_response_time"`
					CallCount           int     `json:"call_count"`
				} `json:"values"`
			} `json:"timeslices"`
		} `json:"metrics"`
	} `json:"metric_data"`
}

func (a *DataStoreMysqlData) CallCount() int {
	return a.MetricData.Metrics[0].Timeslices[0].Values.CallCount
}

func (a *DataStoreMysqlData) AverageResponseTime() float32 {
	return a.MetricData.Metrics[0].Timeslices[0].Values.AverageResponseTime
}

type DatastoreMysqlConfig struct {
	Params Params
}

func (ac DatastoreMysqlConfig) Init() (*RequestConfig, error) {
	// Full fill url template
	applicationId, _ := ac.Params["ApplicationId"].(int)
	responseFormat, _ := ac.Params["ResponseFormat"].(string)
	apiKey, _ := ac.Params["ApiKey"].(string)

	path := fmt.Sprintf(browserNetworkRespTimeApiPath, applicationId, responseFormat)

	// Full fill body template
	startDate, _ := ac.Params["StartDate"].(string)
	endDate, _ := ac.Params["EndDate"].(string)

	body := url.Values{}
	body.Add("names[]", "Datastore/MySQL/allWeb")
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

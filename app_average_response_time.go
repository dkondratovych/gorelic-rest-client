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
	averageRespTimeApiPath string = "/v2/applications/%d/metrics/data.%s"
)

type AppAverageResponseTime struct {
	NewRelicRequest
}

type AppAverageResponseTimeData struct {
	MetricData struct {
		From    time.Time `json:"from"`
		To      time.Time `json:"to"`
		Metrics []struct {
			Name       string `json:"name"`
			Timeslices []struct {
				From   time.Time `json:"from"`
				To     time.Time `json:"to"`
				Values struct {
					AverageCallTime float32 `json:"average_call_time"`
					CallCount       int     `json:"call_count"`
				} `json:"values"`
			} `json:"timeslices"`
		} `json:"metrics"`
	} `json:"metric_data"`
}

func (a *AppAverageResponseTimeData) AverageCallTime() float32 {
	return a.MetricData.Metrics[0].Timeslices[0].Values.AverageCallTime
}

func (a *AppAverageResponseTimeData) CallCount() int {
	return a.MetricData.Metrics[0].Timeslices[0].Values.CallCount
}

type AppAverageResponseTimeConfig struct {
	Params Params
}

func (ac AppAverageResponseTimeConfig) Init() (*RequestConfig, error) {
	// Full fill url template
	applicationId, _ := ac.Params["ApplicationId"].(int)
	responseFormat, _ := ac.Params["ResponseFormat"].(string)
	apiKey, _ := ac.Params["ApiKey"].(string)

	path := fmt.Sprintf(averageRespTimeApiPath, applicationId, responseFormat)

	// Full fill body template
	startDate, _ := ac.Params["StartDate"].(string)
	endDate, _ := ac.Params["EndDate"].(string)

	body := url.Values{}
	body.Add("names[]", "HttpDispatcher")
	body.Add("values[]", "average_call_time")
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

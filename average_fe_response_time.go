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
	averageFeRespTimeApiPath string = "/v2/applications/%d/metrics/data.%s"
)

type AverageFeResponseTime struct {
	NewRelicRequest
}

type AverageFeResponseTimeData struct {
	MetricData struct {
		From    time.Time `json:"from"`
		To      time.Time `json:"to"`
		Metrics []struct {
			Name       string `json:"name"`
			Timeslices []struct {
				From   time.Time `json:"from"`
				To     time.Time `json:"to"`
				Values struct {
					AverageFeResponseTime float32 `json:"average_fe_response_time"`
				} `json:"values"`
			} `json:"timeslices"`
		} `json:"metrics"`
	} `json:"metric_data"`
}

func (a *AverageFeResponseTimeData) AverageCallTime() float32 {
	return a.MetricData.Metrics[0].Timeslices[0].Values.AverageFeResponseTime
}

type AverageFeResponseTimeConfig struct {
	Params Params
}

func (ac AverageFeResponseTimeConfig) Init() (*RequestConfig, error) {
	// Full fill url template
	applicationId, _ := ac.Params["ApplicationId"].(int)
	responseFormat, _ := ac.Params["ResponseFormat"].(string)
	apiKey, _ := ac.Params["ApiKey"].(string)

	path := fmt.Sprintf(averageFeRespTimeApiPath, applicationId, responseFormat)

	// Full fill body template
	startDate, _ := ac.Params["StartDate"].(string)
	endDate, _ := ac.Params["EndDate"].(string)

	body := url.Values{}
	body.Add("names[]", "EndUser")
	body.Add("values[]", "average_fe_response_time")

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

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
	averageDomContentLoadTimeApiPath string = "/v2/applications/%d/metrics/data.%s"
)

type AverageDomContentLoadTime struct {
	NewRelicRequest
}

type AverageDomContentLoadTimeData struct {
	MetricData struct {
		From    time.Time `json:"from"`
		To      time.Time `json:"to"`
		Metrics []struct {
			Name       string `json:"name"`
			Timeslices []struct {
				From   time.Time `json:"from"`
				To     time.Time `json:"to"`
				Values struct {
					AverageDomContentLoadTime float32 `json:"average_dom_content_load_time"`
				} `json:"values"`
			} `json:"timeslices"`
		} `json:"metrics"`
	} `json:"metric_data"`
}

func (a *AverageDomContentLoadTimeData) AverageDomContentLoadTime() float32 {
	return a.MetricData.Metrics[0].Timeslices[0].Values.AverageDomContentLoadTime
}

type AverageDomContentLoadTimeConfig struct {
	Params Params
}

func (ac AverageDomContentLoadTimeConfig) Init() (*RequestConfig, error) {
	// Full fill url template
	applicationId, _ := ac.Params["ApplicationId"].(int)
	responseFormat, _ := ac.Params["ResponseFormat"].(string)
	apiKey, _ := ac.Params["ApiKey"].(string)

	path := fmt.Sprintf(averageRespTimeApiPath, applicationId, responseFormat)

	// Full fill body template
	startDate, _ := ac.Params["StartDate"].(string)
	endDate, _ := ac.Params["EndDate"].(string)

	body := url.Values{}
	body.Add("names[]", "EndUser/RB")
	body.Add("values[]", "average_dom_content_load_time")

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

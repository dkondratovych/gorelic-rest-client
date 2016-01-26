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
	webApplicationTimeApiPath string = "/v2/applications/%d/metrics/data.%s"
)

type WebApplicationTime struct {
	NewRelicRequest
}

type WebApplicationTimeData struct {
	MetricData struct {
		From    time.Time `json:"from"`
		To      time.Time `json:"to"`
		Metrics []struct {
			Name       string `json:"name"`
			Timeslices []struct {
				From   time.Time `json:"from"`
				To     time.Time `json:"to"`
				Values struct {
					TotalAppTime float32 `json:"total_app_time"`
					CallCount    int     `json:"call_count"`
				} `json:"values"`
			} `json:"timeslices"`
		} `json:"metrics"`
	} `json:"metric_data"`
}

func (a *WebApplicationTimeData) TotalAppTime() float32 {
	return a.MetricData.Metrics[0].Timeslices[0].Values.TotalAppTime
}

func (a *WebApplicationTimeData) CallCount() int {
	return a.MetricData.Metrics[0].Timeslices[0].Values.CallCount
}

// Web application = EndUser:total_app_time / EndUser:call_count
func (a *WebApplicationTimeData) WebApplicationTime() float32 {
	if a.CallCount() == 0 {
		return float32(0)
	} else {
		// We have to return here in ms, that's why we multiply on 1000
		return float32(a.TotalAppTime()) / float32(a.CallCount()) * float32(1000)
	}
}

type WebApplicationTimeConfig struct {
	Params Params
}

/**
The Web application time is the time spent in the application code.
To calculate this value, use this equation:
Web application = EndUser:total_app_time / EndUser:call_count
*/
func (ac WebApplicationTimeConfig) Init() (*RequestConfig, error) {
	// Full fill url template
	applicationId, _ := ac.Params["ApplicationId"].(int)
	responseFormat, _ := ac.Params["ResponseFormat"].(string)
	apiKey, _ := ac.Params["ApiKey"].(string)

	path := fmt.Sprintf(webApplicationTimeApiPath, applicationId, responseFormat)

	// Full fill body template
	startDate, _ := ac.Params["StartDate"].(string)
	endDate, _ := ac.Params["EndDate"].(string)

	body := url.Values{}
	body.Add("names[]", "EndUser")
	body.Add("values[]", "total_app_time")

	body.Add("names[]", "EndUser")
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

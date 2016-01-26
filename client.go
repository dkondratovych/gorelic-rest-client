package newrelic

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var errorMap = map[int]string{
	400: "Bad request, missing parametr",
	401: "Invalid request, API key required",
	403: "New Relic API access has not been enabled",
	500: "A server error occured, please contact New Relic support",
}

type NewRelicClient struct {
	client *http.Client
}

type NewRelicError struct {
	Error struct {
		Title string `json:"title"`
	} `json:"error"`
}

func (n NewRelicError) Title() string {
	return n.Error.Title
}

func (c *NewRelicClient) ExternalAllWeb(params Params) (*ExternalAllWebData, *http.Response, error) {
	config, err := (&ExternalAllWebConfig{Params: params}).Init()
	if err != nil {
		return nil, nil, err
	}

	srtRequest, err := (&ExternalAllWeb{NewRelicRequest{Config: config}}).NewRequest()
	if err != nil {
		return nil, nil, err
	}

	dataStoreMySqlData := &ExternalAllWebData{}
	httpResp, err := c.do(srtRequest, dataStoreMySqlData)

	if err != nil {
		return nil, httpResp, err
	}

	return dataStoreMySqlData, httpResp, nil
}

func (c *NewRelicClient) DataStoreMemcached(params Params) (*DataStoreMemcacheData, *http.Response, error) {
	config, err := (&DataStoreMemcacheConfig{Params: params}).Init()
	if err != nil {
		return nil, nil, err
	}

	srtRequest, err := (&DataStoreMemcache{NewRelicRequest{Config: config}}).NewRequest()
	if err != nil {
		return nil, nil, err
	}

	dataStoreMySqlData := &DataStoreMemcacheData{}
	httpResp, err := c.do(srtRequest, dataStoreMySqlData)

	if err != nil {
		return nil, httpResp, err
	}

	return dataStoreMySqlData, httpResp, nil
}

func (c *NewRelicClient) DataStoreMysql(params Params) (*DataStoreMysqlData, *http.Response, error) {
	config, err := (&DatastoreMysqlConfig{Params: params}).Init()
	if err != nil {
		return nil, nil, err
	}

	srtRequest, err := (&DataStoreMysql{NewRelicRequest{Config: config}}).NewRequest()
	if err != nil {
		return nil, nil, err
	}

	dataStoreMySqlData := &DataStoreMysqlData{}
	httpResp, err := c.do(srtRequest, dataStoreMySqlData)

	if err != nil {
		return nil, httpResp, err
	}

	return dataStoreMySqlData, httpResp, nil
}

func (c *NewRelicClient) HttpDispatcher(params Params) (*NewRelicHttpDispatcherData, *http.Response, error) {
	config, err := (&NetworkDispatcherConfig{Params: params}).Init()
	if err != nil {
		return nil, nil, err
	}

	srtRequest, err := (&NewRelicHttpDispatcher{NewRelicRequest{Config: config}}).NewRequest()
	if err != nil {
		return nil, nil, err
	}

	newRelicHttpDispatcher := &NewRelicHttpDispatcherData{}
	httpResp, err := c.do(srtRequest, newRelicHttpDispatcher)

	if err != nil {
		return nil, httpResp, err
	}

	return newRelicHttpDispatcher, httpResp, nil
}

func (c *NewRelicClient) AppAverageResponseTime(params Params) (*AppAverageResponseTimeData, *http.Response, error) {

	config, err := (&AppAverageResponseTimeConfig{Params: params}).Init()
	if err != nil {
		return nil, nil, err
	}

	srtRequest, err := (&AppAverageResponseTime{NewRelicRequest{Config: config}}).NewRequest()
	if err != nil {
		return nil, nil, err
	}

	appAverageResponseTime := &AppAverageResponseTimeData{}
	httpResp, err := c.do(srtRequest, appAverageResponseTime)

	if err != nil {
		return nil, httpResp, err
	}

	return appAverageResponseTime, httpResp, nil
}

func (c *NewRelicClient) BrowserNetworkTime(params Params) (*BrowserNetworkTimeData, *http.Response, error) {
	config, err := (&BrowserNetworkTimeConfig{Params: params}).Init()
	if err != nil {
		return nil, nil, err
	}

	srtRequest, err := (&BrowserNetworkTime{NewRelicRequest{Config: config}}).NewRequest()
	if err != nil {
		return nil, nil, err
	}

	bowserNetworkTime := &BrowserNetworkTimeData{}
	httpResp, err := c.do(srtRequest, bowserNetworkTime)

	if err != nil {
		return nil, httpResp, err
	}

	return bowserNetworkTime, httpResp, nil
}

func (c *NewRelicClient) AverageDomContentLoad(params Params) (*AverageDomContentLoadTimeData, *http.Response, error) {
	config, err := (&AverageDomContentLoadTimeConfig{Params: params}).Init()
	if err != nil {
		return nil, nil, err
	}

	srtRequest, err := (&AverageDomContentLoadTime{NewRelicRequest{Config: config}}).NewRequest()
	if err != nil {
		return nil, nil, err
	}

	averageDomContentLoadData := &AverageDomContentLoadTimeData{}
	httpResp, err := c.do(srtRequest, averageDomContentLoadData)

	if err != nil {
		return nil, httpResp, err
	}

	return averageDomContentLoadData, httpResp, nil
}

func (c *NewRelicClient) AverageFeResponseTime(params Params) (*AverageFeResponseTimeData, *http.Response, error) {
	config, err := (&AverageFeResponseTimeConfig{Params: params}).Init()
	if err != nil {
		return nil, nil, err
	}

	srtRequest, err := (&AverageFeResponseTime{NewRelicRequest{Config: config}}).NewRequest()
	if err != nil {
		return nil, nil, err
	}

	averageFeResponseTimeData := &AverageFeResponseTimeData{}
	httpResp, err := c.do(srtRequest, averageFeResponseTimeData)

	if err != nil {
		return nil, httpResp, err
	}

	return averageFeResponseTimeData, httpResp, nil
}

func (c *NewRelicClient) WebApplicationTime(params Params) (*WebApplicationTimeData, *http.Response, error) {
	config, err := (&WebApplicationTimeConfig{Params: params}).Init()
	if err != nil {
		return nil, nil, err
	}

	srtRequest, err := (&WebApplicationTime{NewRelicRequest{Config: config}}).NewRequest()
	if err != nil {
		return nil, nil, err
	}

	webApplicationTimeData := &WebApplicationTimeData{}
	httpResp, err := c.do(srtRequest, webApplicationTimeData)

	if err != nil {
		return nil, httpResp, err
	}

	return webApplicationTimeData, httpResp, nil
}

func (c *NewRelicClient) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)

	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	err = c.checkResponse(resp)

	if err != nil {
		return resp, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, v)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *NewRelicClient) checkResponse(resp *http.Response) error {

	switch resp.StatusCode {
	case http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError, http.StatusBadRequest, http.StatusNotFound:

		respBody, err := ioutil.ReadAll(resp.Body)

		n := &NewRelicError{}
		err = json.Unmarshal(respBody, n)

		// If response body is empty let's return default error message
		if err != nil {
			return &ResponseError{
				Response:     resp,
				ErrorCode:    resp.StatusCode,
				ErrorMessage: errorMap[resp.StatusCode],
			}
		}

		return &ResponseError{
			Response:     resp,
			ErrorCode:    resp.StatusCode,
			ErrorMessage: n.Title(),
		}

	default:
		return nil

	}

	return nil
}

type INewRelicClient interface {
	// @TODO
	// 1) add all package interface methods here
	// 2) Add interface for each request type, we need to check ofr interfce, but not for *Average bla bla bla

}

func NewRelicRestClient() *NewRelicClient {
	return &NewRelicClient{
		client: &http.Client{},
	}
}

package clients

import (
	"email-service/config"
	"email-service/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client interface {
	do(string, string, map[string]string) (*http.Response, error)
}

type Ergast struct {
	config *config.ErgastClient
	httpClient *http.Client
}

type F1API struct {
	config *config.F1APIClient
	httpClient *http.Client
}

func NewF1APIClient(cfg *config.F1APIClient, timeout time.Duration) *F1API {
	return &F1API{cfg, &http.Client{Timeout: timeout}}
}

// do is an API call wrapper returning a http.Response
func (fa *F1API) do(method string, endpoint string, params map[string]string) (*http.Response, error) {
	// Create new request
	baseURL := fmt.Sprintf("%s%s", fa.config.BaseURL, endpoint)
	req, err := http.NewRequest(method, baseURL, nil)
	if err != nil {
		return nil, err
	}

	// Add headers to the request
	req.Header.Add("x-rapidapi-host", fa.config.Host)
	req.Header.Add("x-rapidapi-key", fa.config.APIKey)

	// Iterate through params and add to query
	q := req.URL.Query()
	for key, val := range params {
		q.Set(key, val)
	}
	// Add query to request
	req.URL.RawQuery = q.Encode()

	return fa.httpClient.Do(req)
}


func (fa *F1API) GetCurrentEvents() (CurrentEvents *models.CurrentEvents, err error) {
	// Create params
	params := map[string]string{
		"date": "2021-12-04",
		"timezone": fa.config.Timezone,
	}

	// Make the request
	res, err := fa.do(http.MethodGet, fa.config.EventEndpoint, params)
	if err != nil {
		return
	}
	defer res.Body.Close()

	// Read the body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the json into a CurrentEvents
	var CurrentEvent *models.CurrentEvents
	if err = json.Unmarshal(body, &CurrentEvent); err != nil {
		return
	}
	return
}


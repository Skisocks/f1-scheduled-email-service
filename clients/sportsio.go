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

type CurrentEventGetter interface {
	GetEventsResponse() (*models.EventsResponse, error)
}

type sportsIO struct {
	config     *config.SportsIO
	httpClient *http.Client
}

func NewSportsIOClient(cfg *config.SportsIO, timeout time.Duration) *sportsIO {
	return &sportsIO{
		config:     cfg,
		httpClient: &http.Client{Timeout: timeout}}
}

// do is an API call wrapper returning a http.Response
func (sio *sportsIO) do(method string, endpoint string, params map[string]string) (*http.Response, error) {
	// Create new request
	baseURL := fmt.Sprintf("%s%s", sio.config.BaseURL, endpoint)
	req, err := http.NewRequest(method, baseURL, nil)
	if err != nil {
		// Todo: handle error
		return nil, err
	}

	// Add headers to the request
	req.Header.Add("x-rapidapi-host", sio.config.Host)
	req.Header.Add("x-rapidapi-key", sio.config.APIKey)

	// Iterate through params and add to query
	q := req.URL.Query()
	for key, val := range params {
		q.Set(key, val)
	}
	// Add query to request
	req.URL.RawQuery = q.Encode()

	return sio.httpClient.Do(req)
}

func (sio *sportsIO) GetEventsResponse() (*models.EventsResponse, error) {
	// Create params
	params := map[string]string{
		"date":     "2021-12-04",
		"timezone": sio.config.Timezone,
	}

	// Make the request
	res, err := sio.do(http.MethodGet, sio.config.EventEndpoint, params)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Read the body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// Todo: handle error
		return nil, err
	}

	// Unmarshal the json into a EventsResponse
	var CurrentEvent *models.EventsResponse
	if err = json.Unmarshal(body, &CurrentEvent); err != nil {
		// Todo: handle error
		return nil, err
	}
	return CurrentEvent, nil
}

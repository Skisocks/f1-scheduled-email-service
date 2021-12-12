package clients

import (
	"email-service/config"
	"email-service/models"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"time"
)

type CurrentEventGetter interface {
	GetEventsResponse() *models.EventsResponse
}

type sportsIO struct {
	logger     *zap.Logger
	config     *config.SportsIO
	httpClient *http.Client
}

func NewSportsIOClient(
	logger *zap.Logger,
	cfg *config.SportsIO,
) *sportsIO {
	return &sportsIO{
		logger:     logger,
		config:     cfg,
		httpClient: &http.Client{Timeout: time.Second * cfg.Timeout},
	}
}

// do is an API call wrapper returning a http.Response
func (sio *sportsIO) do(method string, endpoint string, params map[string]string) (*http.Response, error) {
	// Create new request
	baseURL := fmt.Sprintf("%s%s", sio.config.BaseURL, endpoint)
	req, err := http.NewRequest(method, baseURL, nil)
	if err != nil {
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

	response, err := sio.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	err = response.Body.Close()
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (sio *sportsIO) GetEventsResponse() *models.EventsResponse {
	// Create params
	params := map[string]string{
		"date":     time.Now().Format("2006-01-02"),
		"timezone": sio.config.Timezone,
	}

	// Make the request
	res, err := sio.do(http.MethodGet, sio.config.EventEndpoint, params)
	if err != nil {
		sio.logger.Error(fmt.Sprintf("failed to make request to Ergast API: %s", err))
	}

	// Check for non 2xx code
	if res.StatusCode != http.StatusOK {
		sio.logger.Error(fmt.Sprintf("non-OK HTTP status code: %d", res.StatusCode))
	}

	// Read the body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		sio.logger.Error(fmt.Sprintf("failed to read the response body: %s", err))
	}

	// Unmarshal the json into a EventsResponse
	var CurrentEvent *models.EventsResponse
	if err = json.Unmarshal(body, &CurrentEvent); err != nil {
		sio.logger.Error(fmt.Sprintf("failed to unmarshall the response: %s", err))
	}
	return CurrentEvent
}

package clients

import (
	"email-service/config"
	"fmt"
	"net/http"
	"time"
)

type StandingsGetter interface {
	GetDriverStandings()
	GetConstructorStandings()
}

type ergast struct {
	config     *config.Ergast
	httpClient *http.Client
}

func NewErgastClient(cfg *config.Ergast, timeout time.Duration) *ergast {
	return &ergast{
		config:     cfg,
		httpClient: &http.Client{Timeout: timeout},
	}
}

func (er *ergast) Do(method string, endpoint string, params map[string]string) (*http.Response, error) {
	// Create new request
	baseURL := fmt.Sprintf("%s/%s%s", er.config.BaseURL, er.config.Season, endpoint)
	req, err := http.NewRequest(method, baseURL, nil)
	if err != nil {
		// Todo: handle error
		return nil, err
	}

	// Iterate through params and add to query
	q := req.URL.Query()
	for key, val := range params {
		q.Set(key, val)
	}
	// Add query to request
	req.URL.RawQuery = q.Encode()

	return er.httpClient.Do(req)
}

func (er *ergast) GetDriverStandings() {

}

func (er *ergast) GetConstructorStandings() {

}

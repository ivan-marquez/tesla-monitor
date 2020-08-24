// Package tesla implements functions to consume Tesla Powerwall API
package tesla

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// GetAggregates retrieves aggregates for solar, site and battery status
func (t *Tessel) GetAggregates(p *Powerwall) (*Powerwall, error) {
	endpoint := fmt.Sprintf("%s/aggregates", t.url)
	req, _ := http.NewRequest("GET", endpoint, nil)
	timeoutCtx, cancelFunc := context.WithTimeout(req.Context(), 5*time.Millisecond)

	defer cancelFunc()

	req = req.WithContext(timeoutCtx)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return p, fmt.Errorf("Error getting aggregates: %v", err)
	}

	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&p.Aggregates); err != nil {
		return p, err
	}

	return p, nil
}

// GetBatteryPercentage retrieves the battery percentage
func (t *Tessel) GetBatteryPercentage(p *Powerwall) (*Powerwall, error) {
	endpoint := fmt.Sprintf("%s/status", t.url)
	req, _ := http.NewRequest("GET", endpoint, nil)
	timeoutCtx, cancelFunc := context.WithTimeout(req.Context(), 5*time.Millisecond)

	defer cancelFunc()

	req = req.WithContext(timeoutCtx)
	res, err := http.Get(endpoint)
	if err != nil {
		return p, err
	}

	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&p.BatteryPercentage); err != nil {
		return p, fmt.Errorf("Error decoding battery percentage Json: %v", err)
	}

	return p, nil
}

// New function package constructor
func New(URL string) *Tessel {
	return &Tessel{
		url: URL,
	}
}

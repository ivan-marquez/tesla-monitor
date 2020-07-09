// Package iot implements functions to consume Tesla Powerwall API
package iot

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Tessel type represents IoT device
type Tessel struct {
	url string // Tessel URL
}

// Powerwall type with solar system data
type Powerwall struct {
	Aggregates        map[string]interface{} // Solar system aggregates
	BatteryPercentage map[string]interface{} // Powerwall battery percentage
}

// GetAggregates retrieves aggregates for solar, site and battery status
func (t *Tessel) GetAggregates(p *Powerwall) (*Powerwall, error) {
	endpoint := fmt.Sprintf("%s/aggregates", t.url)

	res, err := http.Get(endpoint)
	if err != nil {
		return p, fmt.Errorf("Error getting aggregates: %v", err)
	}

	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&p.Aggregates); err != nil {
		return p, fmt.Errorf("Error decoding aggregates Json: %v", err)
	}

	return p, nil
}

// GetBatteryPercentage retrieves the battery percentage
func (t *Tessel) GetBatteryPercentage(p *Powerwall) (*Powerwall, error) {
	endpoint := fmt.Sprintf("%s/status", t.url)

	res, err := http.Get(endpoint)
	if err != nil {
		return p, fmt.Errorf("Error getting battery percentage: %v", err)
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

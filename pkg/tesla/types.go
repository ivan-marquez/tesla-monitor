package tesla

import (
	"time"
)

type props struct {
	LastCommunicationTime time.Time `json:"last_communication_time"`
	InstantPower          float32   `json:"instant_power"`
	InstantReactivePower  float32   `json:"instant_reactive_power"`
	InstantApparentPower  float32   `json:"instant_apparent_power"`
	Frequency             float32   `json:"frequency"`
	EnergyExported        float32   `json:"energy_exported"`
	EnergyImported        float32   `json:"energy_imported"`
	InstantAverageVoltage float32   `json:"instant_average_voltage"`
	InstantTotalCurrent   float32   `json:"instant_total_current"`
	IACurrent             float32   `json:"i_a_current"`
	IBCurrent             float32   `json:"i_b_current"`
	ICCurrent             float32   `json:"i_c_current"`
	Timeout               int32     `json:"timeout"`
}

// BatteryStatus represents battery percentage
type BatteryStatus struct {
	Percentage int `json:"percentage"`
}

// Tesla type represent API aggregates
type Tesla struct {
	Site    props `json:"site"`
	Battery props `json:"battery"`
	Load    props `json:"load"`
	Solar   props `json:"solar"`
}

// Tessel type represents IoT device
type Tessel struct {
	url string // Tessel URL
}

// Powerwall type with solar system data
type Powerwall struct {
	Aggregates        Tesla         // Solar system aggregates
	BatteryPercentage BatteryStatus // Powerwall battery percentage
}

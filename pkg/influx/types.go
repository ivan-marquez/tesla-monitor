package influx

import "time"

type props struct {
	LastCommunicationTime time.Time
	InstantPower          float32
	InstantReactivePower  float32
	InstantApparentPower  float32
	Frequency             float32
	EnergyExported        float32
	EnergyImported        float32
	InstantAverageVoltage float32
	InstantTotalCurrent   float32
	IACurrent             float32
	IBCurrent             float32
	ICCurrent             float32
	Timeout               int32
}

// BatteryStatus represents battery percentage
type BatteryStatus struct {
	Percentage int
}

// Tesla type represent API aggregates
type Tesla struct {
	Site    props
	Battery props
	Load    props
	Solar   props
}

// Powerwall type with solar system data
type Powerwall struct {
	Aggregates        Tesla         // Solar system aggregates
	BatteryPercentage BatteryStatus // Powerwall battery percentage
}

package influx

import (
	"context"
	"fmt"
	"strconv"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go"
)

type Influx struct {
	influxdb2.Client
}

func parsePoint(pw *Powerwall) (pt map[string]string) {

	pt["site_last_communication_time"] = pw.Aggregates.Site.LastCommunicationTime.String()
	pt["site_instant_power"] = fmt.Sprintf("%f", pw.Aggregates.Site.InstantPower)
	pt["site_instant_reactive_power"] = fmt.Sprintf("%f", pw.Aggregates.Site.InstantReactivePower)
	pt["site_instant_apparent_power"] = fmt.Sprintf("%f", pw.Aggregates.Site.InstantApparentPower)
	pt["site_frequency"] = fmt.Sprintf("%f", pw.Aggregates.Site.Frequency)
	pt["site_energy_exported"] = fmt.Sprintf("%f", pw.Aggregates.Site.EnergyExported)
	pt["site_energy_imported"] = fmt.Sprintf("%f", pw.Aggregates.Site.EnergyImported)
	pt["site_instant_average_voltage"] = fmt.Sprintf("%f", pw.Aggregates.Site.InstantAverageVoltage)
	pt["site_instant_total_current"] = fmt.Sprintf("%f", pw.Aggregates.Site.InstantTotalCurrent)
	pt["battery_last_communication_time"] = pw.Aggregates.Battery.LastCommunicationTime.String()
	pt["battery_instant_power"] = fmt.Sprintf("%f", pw.Aggregates.Battery.InstantPower)
	pt["battery_instant_reactive_power"] = fmt.Sprintf("%f", pw.Aggregates.Battery.InstantReactivePower)
	pt["battery_instant_apparent_power"] = fmt.Sprintf("%f", pw.Aggregates.Battery.InstantApparentPower)
	pt["battery_frequency"] = fmt.Sprintf("%f", pw.Aggregates.Battery.Frequency)
	pt["battery_energy_exported"] = fmt.Sprintf("%f", pw.Aggregates.Battery.EnergyExported)
	pt["battery_energy_imported"] = fmt.Sprintf("%f", pw.Aggregates.Battery.EnergyImported)
	pt["battery_instant_average_voltage"] = fmt.Sprintf("%f", pw.Aggregates.Battery.InstantAverageVoltage)
	pt["battery_instant_total_current"] = fmt.Sprintf("%f", pw.Aggregates.Battery.InstantTotalCurrent)
	pt["battery_percentage"] = strconv.Itoa(pw.BatteryPercentage.Percentage)
	pt["load_last_communication_time"]
	pt["load_instant_power"]
	pt["load_instant_reactive_power"]
	pt["load_instant_apparent_power"]
	pt["load_frequency"]
	pt["load_energy_exported"]
	pt["load_energy_imported"]
	pt["load_instant_average_voltage"]
	pt["load_instant_total_current"]
	pt["solar_last_communication_time"]
	pt["solar_instant_power"]
	pt["solar_instant_reactive_power"]
	pt["solar_instant_apparent_power"]
	pt["solar_frequency"]
	pt["solar_energy_exported"]
	pt["solar_energy_imported"]
	pt["solar_instant_average_voltage"]
	pt["solar_instant_total_current"]

	return
}

// IngestData inserts data points to InfluxDB
func (ifx *Influx) IngestData(pw *Powerwall) {
	// Get the blocking write client
	writeAPI := ifx.WriteAPIBlocking("", "powerwall")

	// TODO: how to write Json as a new point
	// create point using full params constructor
	p := influxdb2.NewPoint("stat",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"avg": 24.5, "max": 45},
		time.Now())
	// Write data
	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		fmt.Printf("Write error: %s\n", err.Error())
	}
}

// New function package constructor
func New(dbURL, user, password string) *Influx {
	return &Influx{
		influxdb2.NewClient(dbURL, fmt.Sprintf("%s:%s", user, password)),
	}
}

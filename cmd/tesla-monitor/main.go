// tesla-adx package retrieves data from Tesla Powerwall through the Tessel board
// and ingests data to Azure Data Explorer through a Cron job.
package main

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Logging setup
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Env variables
	// tesselURL := os.Getenv("TESSEL_URL")
	// AzClientSecret := os.Getenv("AZ_CLIENT_SECRET")
	// AzClientID := os.Getenv("AZ_CLIENT_ID")
	// AzTenantID := os.Getenv("AZ_TENANT_ID")
	// AzClusterURL := os.Getenv("AZ_CLUSTER_URL")

	// var (
	// 	powerwall iot.Powerwall
	// 	wg        sync.WaitGroup
	// )

	// tessel := iot.New(tesselURL)
	// log.Info("Setting up Kusto client...")
	// adx, err := adx.New(AzClusterURL, AzClientID, AzTenantID, AzClientSecret)
	// log.Info("Azure authentication successful")
	// c := cron.New()

	// // cron job to ingest data to ADX
	// err = c.AddFunc("@every 5s", func() {
	// 	wg.Add(2)

	// 	log.Info("Retrieving Powerwall data...")
	// 	go func() {
	// 		tessel.GetAggregates(&powerwall)
	// 		wg.Done()
	// 	}()

	// 	go func() {
	// 		tessel.GetBatteryPercentage(&powerwall)
	// 		wg.Done()
	// 	}()

	// 	wg.Wait()

	// 	// merge properties into one to send powerwall.Aggregates only
	// 	powerwall.Aggregates["battery"].(map[string]interface{})["percentage"] = powerwall.BatteryPercentage["percentage"]

	// 	payload, err := json.Marshal(powerwall.Aggregates)
	// 	if err != nil {
	// 		// TODO: improve error handling
	// 		log.Errorf("Error encoding to Json: %v", err)
	// 	}

	// 	log.Info("Ingesting payload to Kusto...")
	// 	err = adx.IngestData(payload)
	// 	if err != nil {
	// 		log.Error(err)
	// 		goto localStore
	// 	}

	// 	log.Info("Data ingestion successful")

	// localStore:
	// 	log.Info("Storing payload locally...")
	// 	// TODO: send payload to local store in case ingestion to ADX fails
	// })

	// if err != nil {
	// 	log.Error(err)
	// }

	// c.Start()
	// select {}

	// influx.Test()
}

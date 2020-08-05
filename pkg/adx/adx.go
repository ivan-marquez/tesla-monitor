// Package adx implements functions to create an Azure Data Explorer client and data ingestion
package adx

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-kusto-go/kusto"
	"github.com/Azure/azure-kusto-go/kusto/ingest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

// ADX type with Azure Data Explorer Auth credentials
type ADX struct {
	url          string
	clientID     string
	tenantID     string
	clientSecret string
	kustoClient  *kusto.Client
}

// IngestData inserts passed payload to Kusto cluster
func (adx *ADX) IngestData(payload []byte) error {
	const db = "SolarSystem"
	const table = "SystemStatus"
	const ingestionMapping = "SystemStatusMapping"

	in, err := ingest.New(adx.kustoClient, db, table)
	if err != nil {
		return fmt.Errorf("Error setting up ingestion: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := in.Stream(ctx, payload, ingest.JSON, ingestionMapping); err != nil {
		return fmt.Errorf("Error ingesting data: %v", err)
	}

	return nil
}

// New function package constructor
func New(url, clientID, tenantID, clientSecret string) (*ADX, error) {
	authorizer := kusto.Authorization{
		Config: auth.NewClientCredentialsConfig(clientID, clientSecret, tenantID),
	}

	kc, err := kusto.New(url, authorizer)
	if err != nil {
		return nil, fmt.Errorf("Error authorizing with Azure: %v", err)
	}

	return &ADX{
		url,
		clientID,
		tenantID,
		clientSecret,
		kc,
	}, nil
}

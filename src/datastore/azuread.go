package datastore

import (
	"context"
	"errors"
	"os"

	"google.golang.org/appengine"

	"google.golang.org/appengine/datastore"
)

type (
	AzureCredentials struct {
		ClientID     string
		ClientSecret string
	}

	azureClientID struct {
		ClientID string `datastore:"CLIENT_ID"`
	}

	azureClientSecret struct {
		ClientSecret string `datastore:"CLIENT_SECRET"`
	}
)

const (
	// clientIDName is the name of the client id in the env and
	// google cloud datastore
	clientIDName = "AZURE_AD_CLIENT_ID"

	// clientSecretName is the name of the client secret in the env and
	// google cloud datastore
	clientSecretName = "AZURE_AD_CLIENT_SECRET"
)

// findAzureCredentialsFromEnv retrieves the credentials from the environment,
// if one or more environment variable is missing an error is returned
func findAzureCredentialsFromEnv(tokenName string) (string, error) {
	token := os.Getenv(tokenName)
	if token == "" {
		return "", errors.New("API token variable is missing")
	}
	return token, nil
}

// findAzureIDFromDatastore retrieves azure client ID
func findAzureIDFromDatastore(ctx context.Context) (string, error) {
	var info azureClientID
	q := datastore.NewQuery(clientIDName).Limit(1)
	iterator := q.Run(ctx)

	_, err := iterator.Next(&info)

	if err != nil {
		return "", err
	}
	return info.ClientID, nil
}

// findAzureSecretFromDatastore retrieves azure client ID
func findAzureSecretFromDatastore(ctx context.Context) (string, error) {
	var info azureClientSecret
	q := datastore.NewQuery(clientSecretName).Limit(1)
	iterator := q.Run(ctx)

	_, err := iterator.Next(&info)

	if err != nil {
		return "", err
	}
	return info.ClientSecret, nil
}

// FindAzureCredentialsInformation retrieves information about azure ad creds
// from either the local environment or the Google Cloud Datastore,
// depending if we are running the service in dev or production
func FindAzureCredentialsInformation(ctx context.Context) (AzureCredentials, error) {
	var creds AzureCredentials
	var err error

	if appengine.IsDevAppServer() {
		creds.ClientID, err = findTokenFromEnv(clientIDName)
		if err != nil {
			return creds, err
		}
		creds.ClientSecret, err = findTokenFromEnv(clientSecretName)
	} else {
		creds.ClientID, err = findAzureIDFromDatastore(ctx)
		if err != nil {
			return creds, err
		}
		creds.ClientSecret, err = findAzureSecretFromDatastore(ctx)
	}
	return creds, err
}

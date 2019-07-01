package datastore

import (
	"context"
	"errors"
	"os"

	"google.golang.org/appengine"

	"google.golang.org/appengine/datastore"
)

const (
	// AssignmentsTokenName is the name of the token in the env and
	// google cloud datastore for the assignments service/api
	AssignmentsTokenName = "API_ASSIGNMENTS_TOKEN"

	// TasksTokenName is the name of the token in the env and
	// google cloud datastore for the tasks service/api
	TasksTokenName = "API_TASKS_TOKEN"
)

// TokenAPI is the data structure fitting google cloud datastore
type TokenAPI struct {
	Token string `datastore:"TOKEN_VALUE"`
}

// findTokenFromEnv retrieves the API's token from the environment,
// if one or more environment variable is missing an error is returned
func findTokenFromEnv(tokenName string) (string, error) {
	token := os.Getenv(tokenName)
	if token == "" {
		return "", errors.New("API token variable is missing")
	}
	return token, nil
}

// findTokenFromDatastore retrieves the API's token
func findTokenFromDatastore(ctx context.Context, tokenName string) (string, error) {
	var info TokenAPI
	q := datastore.NewQuery(tokenName).Limit(1)
	iterator := q.Run(ctx)

	_, err := iterator.Next(&info)

	if err != nil {
		return "", err
	}
	return info.Token, nil
}

// FindAPITokenInformation retrieves information about the API token
// from either the local environment or the Google Cloud Datastore,
// depending if we are running the service in dev or production
func FindAPITokenInformation(ctx context.Context, tokenName string) (string, error) {
	var token string
	var err error

	if appengine.IsDevAppServer() {
		token, err = findTokenFromEnv(tokenName)
	} else {
		token, err = findTokenFromDatastore(ctx, tokenName)
	}
	return token, err
}

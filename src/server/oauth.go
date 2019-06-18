package server

import (
	"context"
	"net/http"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/azuread"
	"google.golang.org/appengine"

	"github.com/pedafy/pedafy/src/datastore"
)

func (s *Server) initOauth(ctx context.Context) error {
	creds, err := datastore.FindAzureCredentialsInformation(ctx)
	if err != nil {
		return err
	}
	s.oAuthID = creds.ClientID
	s.oAuthSecret = creds.ClientSecret
	s.setOauthProvider()
	return nil
}

func (s *Server) setOauthProvider() {

	// TODO: simply the code here

	if appengine.IsDevAppServer() {
		goth.UseProviders(
			azuread.New(s.oAuthID, s.oAuthSecret, "http://localhost:9000/auth/azuread/callback", nil, "User.Read"),
		)
	} else {
		goth.UseProviders(
			azuread.New(s.oAuthID, s.oAuthSecret, "https://pedafy.com/auth/azuread/callback", nil, "User.Read"),
		)
	}
}

func (s *Server) loginOauthHandler(w http.ResponseWriter, r *http.Request) {

}

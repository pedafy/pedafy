package server

import (
	"context"
	"net/http"

	"github.com/pedafy/pedafy/src/server/user"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/azuread"
	"google.golang.org/appengine"
)

func (s *Server) initOauth(ctx context.Context) error {
	// creds, _ := datastore.FindAzureCredentialsInformation(ctx)
	// // if err != nil {
	// // 	return err
	// // }
	// s.oAuthID = creds.ClientID
	// s.oAuthSecret = creds.ClientSecret
	s.oAuthID = "3405088c-5068-4118-b567-eab3450c6779"
	s.oAuthSecret = "IvV9CyFSYrnQj/gFNf+HChubNT2XcFbIBs1Jhs2XNB0="
	s.setOauthProvider()
	return nil
}

func (s *Server) setOauthProvider() {

	// TODO: simplify the code here

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
	authUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}
	user.NewUser(w, r, authUser)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

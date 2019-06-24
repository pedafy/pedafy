package server

import (
	"log"
	"net/http"
	"os"

	"github.com/pedafy/pedafy/src/server/user"

	"github.com/markbates/goth/gothic"

	"github.com/pedafy/pedafy/src/template"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

// registerHandlers will add all the handlers to the http router
func (s *Server) registerHandlers() {
	r := mux.NewRouter()

	r.Methods(http.MethodGet).PathPrefix("/public/").HandlerFunc(s.serveStatic)

	r.Methods(http.MethodGet).Path("/_ah/start").HandlerFunc(s.startupHandler)

	r.Methods(http.MethodGet).Path("/").HandlerFunc(s.homeHandler)
	r.Methods(http.MethodGet).Path("/login").HandlerFunc(s.loginHandler)
	r.Methods(http.MethodGet).Path("/logout").HandlerFunc(s.logoutHandler)

	// OAuth
	r.HandleFunc("/auth/{provider}/callback", s.loginOauthHandler).Methods(http.MethodGet)
	r.HandleFunc("/auth/{provider}", gothic.BeginAuthHandler).Methods(http.MethodGet)

	r.Methods(http.MethodGet).Path("/tig").HandlerFunc(s.tigHomeHandler)
	r.Methods(http.MethodGet).Path("/tig/{id:[0-9]+}").HandlerFunc(s.tigHandler)
	r.Methods(http.MethodGet).Path("/tig/modify/{id:[0-9]+}").HandlerFunc(s.modifyTigHandler)
	r.Methods(http.MethodGet).Path("/tig/review/{id:[0-9]+}").HandlerFunc(s.reviewTigHandler)
	r.Methods(http.MethodGet).Path("/tig/new").HandlerFunc(s.newTigHandler)

	r.Methods(http.MethodGet).Path("/task").HandlerFunc(s.taskHomeHandler)
	r.Methods(http.MethodGet).Path("/task/{id:[0-9]+}").HandlerFunc(s.taskHandler)
	r.Methods(http.MethodGet).Path("/task/modify/{id:[0-9]+}").HandlerFunc(s.modifyTaskHandler)
	r.Methods(http.MethodGet).Path("/task/new").HandlerFunc(s.newTaskHandler)
	r.Methods(http.MethodPost).Path("/task/modify/{id:[0-9]+}").HandlerFunc(s.modifyTaskHandlerAPI)
	r.Methods(http.MethodPost).Path("/task/new").HandlerFunc(s.newTaskHandlerAPI)

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))
}

func (s *Server) homeHandler(w http.ResponseWriter, r *http.Request) {
	user, loggedIn := user.GetUser(r)

	p := template.NewPage("Pedafy - Home", loggedIn == nil, user, nil)
	if err := template.RenderTemplate(w, p, "home.gohtml"); err != nil {
		log.Fatal(err.Error())
	}
}

func (s *Server) loginHandler(w http.ResponseWriter, r *http.Request) {
	user, loggedIn := user.GetUser(r)

	p := template.NewPage("Pedafy - Login", loggedIn == nil, user, nil)
	if err := template.RenderTemplate(w, p, "login.gohtml"); err != nil {
		log.Fatal(err.Error())
	}
}

func (s *Server) logoutHandler(w http.ResponseWriter, r *http.Request) {
	user.LogoutUser(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *Server) startupHandler(w http.ResponseWriter, r *http.Request) {
	if s.isTokenSet() == false {
		ctx := appengine.NewContext(r)
		if err := s.fetchTokenAPI(ctx); err != nil {
			log.Fatal(err.Error())
		}
		if err := s.initOauth(ctx); err != nil {
			log.Fatal(err.Error())
		}
		user.Init()
	}
}

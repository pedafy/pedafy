package server

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

// registerHandlers will add all the handlers to the http router
func (s *Server) registerHandlers() {
	r := mux.NewRouter()

	r.Methods(http.MethodGet).Path("/_ah/start").HandlerFunc(s.startupHandler)

	r.Methods(http.MethodGet).Path("/").HandlerFunc(s.homeHandler)
	r.Methods(http.MethodGet).Path("/login").HandlerFunc(s.loginHandler)

	r.Methods(http.MethodGet).Path("/tig").HandlerFunc(s.tigHomeHandler)
	r.Methods(http.MethodGet).Path("/tig/{id:[0-9]+}").HandlerFunc(s.tigHandler)
	r.Methods(http.MethodGet).Path("/tig/modify/{id:[0-9]+}").HandlerFunc(s.modifyTigHandler)
	r.Methods(http.MethodGet).Path("/tig/review/{id:[0-9]+}").HandlerFunc(s.reviewTigHandler)
	r.Methods(http.MethodGet).Path("/tig/new").HandlerFunc(s.newTigHandler)

	r.Methods(http.MethodGet).Path("/task").HandlerFunc(s.taskHomeHandler)
	r.Methods(http.MethodGet).Path("/task/{id:[0-9]+}").HandlerFunc(s.taskHandler)
	r.Methods(http.MethodGet).Path("/task/modify/{id:[0-9]+}").HandlerFunc(s.modifyTaskHandler)
	r.Methods(http.MethodGet).Path("/task/new").HandlerFunc(s.newTaskHandler)

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))
}

func (s *Server) homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func (s *Server) loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func (s *Server) startupHandler(w http.ResponseWriter, r *http.Request) {
	if s.isTokenSet() == false {
		ctx := appengine.NewContext(r)
		err := s.fetchTokenAPI(ctx)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
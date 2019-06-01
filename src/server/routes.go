package server

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func (s *Server) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

// RegisterHandlers will add all the handlers to the http router
func (s *Server) RegisterHandlers() {
	r := mux.NewRouter()

	r.Methods(http.MethodGet).Path("/").HandlerFunc(s.home)
	r.Methods(http.MethodGet).Path("/tig").HandlerFunc(s.home)
	r.Methods(http.MethodGet).Path("/tig/{id:[0-9]+}").HandlerFunc(s.home)
	r.Methods(http.MethodGet).Path("/tig/modify/{id:[0-9]+}").HandlerFunc(s.home)
	r.Methods(http.MethodGet).Path("/tig/review/{id:[0-9]+}").HandlerFunc(s.home)
	r.Methods(http.MethodGet).Path("/tig/new").HandlerFunc(s.home)

	r.Methods(http.MethodGet).Path("/task").HandlerFunc(s.home)
	r.Methods(http.MethodGet).Path("/task/{id:[0-9]+}").HandlerFunc(s.home)
	r.Methods(http.MethodGet).Path("/task/modify/{id:[0-9]+}").HandlerFunc(s.home)
	r.Methods(http.MethodGet).Path("/task/new").HandlerFunc(s.home)

	r.Methods(http.MethodGet).Path("/login").HandlerFunc(s.home)
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))
}
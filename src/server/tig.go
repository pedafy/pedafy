package server

import "net/http"

func (s *Server) tigHomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func (s *Server) tigHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func (s *Server) modifyTigHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func (s *Server) reviewTigHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func (s *Server) newTigHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}




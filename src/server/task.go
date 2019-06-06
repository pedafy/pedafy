package server

import "net/http"

func (s *Server) taskHomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func (s *Server) taskHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func (s *Server) modifyTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func (s *Server) newTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}



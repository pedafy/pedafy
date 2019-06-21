package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pedafy/pedafy/src/server/user"
	"github.com/pedafy/pedafy/src/template"
)

type taskPageInfo struct {
	Tasks  []Task
	Status []Status
}

func (s *Server) taskHomeHandler(w http.ResponseWriter, r *http.Request) {
	user, loggedIn := user.GetUser(r)

	ts, err := s.taskGetAll()
	if err != nil {
		log.Println(err.Error())
		ts = []Task{}
	}

	sts, err := s.taskStatusGetAll()
	if err != nil {
		log.Println(err.Error())
		sts = []Status{}
	}

	data := taskPageInfo{
		Tasks:  ts,
		Status: sts,
	}

	p := template.NewPage("Pedafy - Tasks", loggedIn == nil, user, data)
	if err := template.RenderTemplate(w, p, "tasks.gohtml"); err != nil {
		log.Fatal(err.Error())
	}
}

func (s *Server) taskHandler(w http.ResponseWriter, r *http.Request) {
	user, loggedIn := user.GetUser(r)

	vars := mux.Vars(r)
	ids := vars["id"]
	taskID, _ := strconv.Atoi(ids)

	ts, err := s.taskGetOne(taskID)
	if err != nil {
		log.Println(err.Error())
		ts = []Task{}
	}

	sts, err := s.taskStatusGetAll()
	if err != nil {
		log.Println(err.Error())
		sts = []Status{}
	}

	data := taskPageInfo{
		Tasks:  ts,
		Status: sts,
	}

	p := template.NewPage(fmt.Sprintf("Pedafy - Task #%d", taskID), loggedIn == nil, user, data)
	if err := template.RenderTemplate(w, p, "one_task.gohtml"); err != nil {
		log.Fatal(err.Error())
	}
}

func (s *Server) modifyTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func (s *Server) newTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func (s *Server) modifyTaskHandlerAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func (s *Server) newTaskHandlerAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

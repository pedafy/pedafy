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

type assignmentPageInfo struct {
	Assignments []Assignment
	Tasks       []Task
	Status      []StatusAssignment
}

func (s *Server) tigHomeHandler(w http.ResponseWriter, r *http.Request) {
	user, loggedIn := user.GetUser(r)

	var data assignmentPageInfo
	var templateName string

	if true {
		// Student user

		// use real student ID
		as, err := s.assignmentGetByAssignedOne(1)
		if err != nil {
			log.Println(err.Error())
			as = []Assignment{}
		}

		sts, err := s.asignmentsStatusGetAll()
		if err != nil {
			log.Println(err.Error())
			sts = []StatusAssignment{}
		}

		var ts []Task

		for _, a := range as {
			t, err := s.taskGetOne(a.TaskID)
			ts = append(ts, t[0])
			if err != nil {
				log.Println(err.Error())
			}
		}

		data = assignmentPageInfo{
			Assignments: as,
			Status:      sts,
			Tasks:       ts,
		}
		templateName = "my_assignments"

	} else {
		// Helper and Admin user
	}
	p := template.NewPage("Pedafy - Assignments", loggedIn == nil, user, data)
	if err := template.RenderTemplate(w, p, templateName+".gohtml"); err != nil {
		log.Fatal(err.Error())
	}
}

func (s *Server) tigHandler(w http.ResponseWriter, r *http.Request) {
	user, loggedIn := user.GetUser(r)

	vars := mux.Vars(r)
	ids := vars["id"]
	assignmentID, _ := strconv.Atoi(ids)

	as, err := s.assignmentGetOne(assignmentID)
	if err != nil {
		log.Println(err.Error())
		as = []Assignment{}
	}

	sts, err := s.asignmentsStatusGetAll()
	if err != nil {
		log.Println(err.Error())
		sts = []StatusAssignment{}
	}
	assignmentStatus := []StatusAssignment{{0, ""}}
	for _, s := range sts {
		if s.ID == as[0].StatusID {
			assignmentStatus[0] = s
			break
		}
	}

	ts, err := s.taskGetOne(as[0].TaskID)
	if err != nil {
		log.Println(err.Error())
		ts = []Task{}
	}

	data := assignmentPageInfo{
		Assignments: as,
		Status:      assignmentStatus,
		Tasks:       ts,
	}

	p := template.NewPage(fmt.Sprintf("Pedafy - Assignment #%d", assignmentID), loggedIn == nil, user, data)
	if err := template.RenderTemplate(w, p, "one_assignment.gohtml"); err != nil {
		log.Fatal(err.Error())
	}
}

func (s *Server) modifyTigHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func (s *Server) modifyTigHandlerAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func (s *Server) reviewTigHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func (s *Server) reviewTigHandlerAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func (s *Server) newTigHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func (s *Server) newTigHandlerAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/pedafy/pedafy/src/server/user"
	"github.com/pedafy/pedafy/src/template"
)

type Date struct {
	Month string
	Year  string
	Day   string
}

type assignmentPageInfo struct {
	Assignments []Assignment
	Tasks       []Task
	Status      []StatusAssignment

	DueDate        Date
	AccomplishDate Date
	Accomplished   bool
}

func (s *Server) tigHomeHandler(w http.ResponseWriter, r *http.Request) {
	user, loggedIn := user.GetUser(r)

	if loggedIn != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	var data assignmentPageInfo
	var templateName string

	if user.Login == "florent.poinsard@epitech.eu" {
		// Student users

		// use real student ID
		as, err := s.assignmentGetByAssignedOne(user.Login)
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

	} else if user.Login == "florent1.poinsard@epitech.eu" {
		// Helper & Admin users

		as, err := s.assignmentsGetAll()
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
		templateName = "all_assignments"

	}
	p := template.NewPage("Pedafy - Assignments", loggedIn == nil, user, data)
	if err := template.RenderTemplate(w, p, templateName+".gohtml"); err != nil {
		log.Fatal(err.Error())
	}
}

func (s *Server) tigHandler(w http.ResponseWriter, r *http.Request) {
	user, loggedIn := user.GetUser(r)

	if loggedIn != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
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
	user, loggedIn := user.GetUser(r)

	if loggedIn != nil || user.Login != "florent1.poinsard@epitech.eu" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	vars := mux.Vars(r)
	ids := vars["id"]
	aID, _ := strconv.Atoi(ids)

	as, err := s.assignmentGetOne(aID)
	if err != nil {
		log.Println(err.Error())
		as = []Assignment{}
	}

	sts, err := s.asignmentsStatusGetAll()
	if err != nil {
		log.Println(err.Error())
		sts = []StatusAssignment{}
	}

	ts, err := s.taskGetAll()
	if err != nil {
		log.Println(err.Error())
		ts = []Task{}
	}

	y, m, d := as[0].DueDate.Date()
	dueDate := Date{
		Month: strconv.Itoa(int(m) - 1),
		Year:  strconv.Itoa(y),
		Day:   strconv.Itoa(d),
	}

	data := assignmentPageInfo{
		Assignments: as,
		Status:      sts,
		Tasks:       ts,
		DueDate:     dueDate,
	}

	if as[0].CompletionDate != nil {
		completionDate := Date{}
		y, m, d = as[0].CompletionDate.Date()
		completionDate = Date{
			Month: strconv.Itoa(int(m) - 1),
			Year:  strconv.Itoa(y),
			Day:   strconv.Itoa(d),
		}
		data.AccomplishDate = completionDate
		data.Accomplished = true
	} else {
		data.Accomplished = false
	}

	p := template.NewPage("Pedafy - Modify assignment", loggedIn == nil, user, data)
	if err := template.RenderTemplate(w, p, "modify_assignment.gohtml"); err != nil {
		log.Fatal(err.Error())
	}
}

func (s *Server) modifyTigHandlerAPI(w http.ResponseWriter, r *http.Request) {
	user, loggedIn := user.GetUser(r)

	if loggedIn != nil || user.Login != "florent1.poinsard@epitech.eu" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	vars := mux.Vars(r)
	ids := vars["id"]
	tigID, _ := strconv.Atoi(ids)

	dueDate := r.FormValue("due_date")
	dueDateTime, err := time.Parse("Jan 02, 2006", dueDate)
	if err != nil {
		log.Fatal(err.Error())
	}

	accomplishDate := r.FormValue("accomplishement_date")
	accomplishDateTime := time.Time{}
	if accomplishDate != "" {
		accomplishDateTime, err = time.Parse("Jan 02, 2006", accomplishDate)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	statusID, err := strconv.Atoi(r.FormValue("status_id"))
	if err != nil {
		log.Fatal(err.Error())
	}

	taskID, err := strconv.Atoi(r.FormValue("task_id"))
	if err != nil {
		log.Fatal(err.Error())
	}
	// TODO: fix the creator ID
	newTig, err := s.assignmentModify(tigID, user.Login, r.FormValue("assigned"), statusID, taskID, dueDateTime, accomplishDateTime, r.FormValue("title"), r.FormValue("description"))
	if err != nil {
		log.Fatal(err.Error())
	}
	http.Redirect(w, r, fmt.Sprintf("/tig/%d", newTig.ID), http.StatusSeeOther)
}

func (s *Server) newTigHandler(w http.ResponseWriter, r *http.Request) {
	user, loggedIn := user.GetUser(r)

	if loggedIn != nil || user.Login != "florent1.poinsard@epitech.eu" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	as, err := s.assignmentsGetAll()
	if err != nil {
		log.Println(err.Error())
		as = []Assignment{}
	}

	sts, err := s.asignmentsStatusGetAll()
	if err != nil {
		log.Println(err.Error())
		sts = []StatusAssignment{}
	}

	ts, err := s.taskGetAll()
	if err != nil {
		log.Println(err.Error())
		ts = []Task{}
	}

	data := assignmentPageInfo{
		Assignments: as,
		Status:      sts,
		Tasks:       ts,
	}

	p := template.NewPage("Pedafy - New tig", loggedIn == nil, user, data)
	if err := template.RenderTemplate(w, p, "new_assignment.gohtml"); err != nil {
		log.Fatal(err.Error())
	}
}

func (s *Server) newTigHandlerAPI(w http.ResponseWriter, r *http.Request) {
	user, loggedIn := user.GetUser(r)

	if loggedIn != nil || user.Login != "florent1.poinsard@epitech.eu" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	dueDate := r.FormValue("due_date")
	dueDateTime, err := time.Parse("Jan 02, 2006", dueDate)
	if err != nil {
		log.Fatal(err.Error())
	}

	statusID, err := strconv.Atoi(r.FormValue("status_id"))
	if err != nil {
		log.Fatal(err.Error())
	}
	taskID, err := strconv.Atoi(r.FormValue("task_id"))
	if err != nil {
		log.Fatal(err.Error())
	}
	// TODO: fix the creator ID
	newTig, err := s.assignmentNew(user.Login, r.FormValue("assigned"), statusID, taskID, dueDateTime, r.FormValue("title"), r.FormValue("description"))
	if err != nil {
		log.Fatal(err.Error())
	}
	http.Redirect(w, r, fmt.Sprintf("/tig/%d", newTig.ID), http.StatusSeeOther)
}

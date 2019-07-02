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

	if loggedIn != nil || user.Login != emailPedago {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
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

	if loggedIn != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
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
	user, loggedIn := user.GetUser(r)

	if loggedIn != nil || user.Login != emailPedago {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
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

	p := template.NewPage("Pedafy - Modify task", loggedIn == nil, user, data)
	if err := template.RenderTemplate(w, p, "modify_task.gohtml"); err != nil {
		log.Fatal(err.Error())
	}
}

func (s *Server) modifyTaskHandlerAPI(w http.ResponseWriter, r *http.Request) {
	user, loggedIn := user.GetUser(r)

	if loggedIn != nil || user.Login != emailPedago {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	vars := mux.Vars(r)
	ids := vars["id"]
	taskID, _ := strconv.Atoi(ids)

	statusID, err := strconv.Atoi(r.FormValue("status"))
	if err != nil {
		log.Fatal(err.Error())
	}
	// TODO: fix the creator ID
	newTask, err := s.taskModify(taskID, user.Login, statusID, r.FormValue("title"), r.FormValue("description"))
	if err != nil {
		log.Fatal(err.Error())
	}
	http.Redirect(w, r, fmt.Sprintf("/task/%d", newTask.ID), http.StatusSeeOther)
}

func (s *Server) newTaskHandler(w http.ResponseWriter, r *http.Request) {
	user, loggedIn := user.GetUser(r)
	if loggedIn != nil || user.Login != emailPedago {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	sts, err := s.taskStatusGetAll()
	if err != nil {
		log.Println(err.Error())
		sts = []Status{}
	}

	data := taskPageInfo{
		Status: sts,
	}

	p := template.NewPage("Pedafy - New task", loggedIn == nil, user, data)
	if err := template.RenderTemplate(w, p, "new_task.gohtml"); err != nil {
		log.Fatal(err.Error())
	}
}

func (s *Server) newTaskHandlerAPI(w http.ResponseWriter, r *http.Request) {
	user, loggedIn := user.GetUser(r)
	if loggedIn != nil || user.Login != emailPedago {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	statusID, err := strconv.Atoi(r.FormValue("status"))
	if err != nil {
		log.Fatal(err.Error())
	}
	// TODO: fix the creator ID
	newTask, err := s.taskNew(user.Login, statusID, r.FormValue("title"), r.FormValue("description"))
	if err != nil {
		log.Fatal(err.Error())
	}
	http.Redirect(w, r, fmt.Sprintf("/task/%d", newTask.ID), http.StatusSeeOther)
}

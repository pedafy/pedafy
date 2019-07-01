package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type (
	Assignment struct {
		ID             int        `json:"id"`
		CreatorID      int        `json:"creator_id"`
		AssignedID     int        `json:"assigned_id"`
		StatusID       int        `json:"status_id"`
		TaskID         int        `json:"task_id"`
		CreatedAt      *time.Time `json:"created_at"`
		LastEdit       *time.Time `json:"last_edit"`
		DueDate        *time.Time `json:"due_date"`
		CompletionDate *time.Time `json:"completion_date"`
		Title          string     `json:"title"`
		Description    string     `json:"description"`
	}

	AssignmentData struct {
		Data Assignment `json:"data"`
	}

	AssignmentArrayData struct {
		Data []Assignment `json:"data"`
	}

	StatusAssignment struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	StatusAssignmentArrayData struct {
		Data []StatusAssignment `json:"data"`
	}
)

func (s *Server) assignmentsGetAll() ([]Assignment, error) {
	var data AssignmentArrayData

	client := &http.Client{}
	req, _ := http.NewRequest("GET", s.EndpointServices[ServiceAssignments]+"/assignments?sort=new", nil)
	req.Header.Set("Authorization", s.TokenAPI[ServiceAssignments])
	res, err := client.Do(req)
	if err != nil {
		return data.Data, err
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(responseBody, &data)
	return data.Data, err
}

func (s *Server) assignmentGetOne(assignmentID int) ([]Assignment, error) {
	var data AssignmentArrayData

	client := &http.Client{}
	req, _ := http.NewRequest("GET", s.EndpointServices[ServiceAssignments]+fmt.Sprintf("/assignments/id/%d", assignmentID), nil)
	req.Header.Set("Authorization", s.TokenAPI[ServiceAssignments])
	res, err := client.Do(req)
	if err != nil {
		return data.Data, err
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(responseBody, &data)
	return data.Data, err
}

func (s *Server) assignmentGetByAssignedOne(assignedID int) ([]Assignment, error) {
	var data AssignmentArrayData

	client := &http.Client{}
	req, _ := http.NewRequest("GET", s.EndpointServices[ServiceAssignments]+fmt.Sprintf("/assignments/assigned_id/%d", assignedID), nil)
	req.Header.Set("Authorization", s.TokenAPI[ServiceAssignments])
	res, err := client.Do(req)
	if err != nil {
		return data.Data, err
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(responseBody, &data)
	return data.Data, err
}

func (s *Server) assignmentNew(creatorID, assignedID, statusID, taskID int, dueDate *time.Time, title, description string) (Assignment, error) {
	var data AssignmentData

	client := &http.Client{}

	a := Assignment{
		AssignedID:  assignedID,
		TaskID:      taskID,
		CreatorID:   creatorID,
		StatusID:    statusID,
		Title:       title,
		Description: description,
		DueDate:     dueDate,
	}

	aJSON, _ := json.Marshal(a)

	req, _ := http.NewRequest("PUT", s.EndpointServices[ServiceAssignments]+"/assignment", bytes.NewBuffer(aJSON))
	req.Header.Set("Authorization", s.TokenAPI[ServiceAssignments])
	res, err := client.Do(req)
	if err != nil {
		return data.Data, err
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(responseBody, &data)
	return data.Data, err
}

func (s *Server) assignmentModify(assignmentID, creatorID, assignedID, statusID, taskID int, dueDate, completionDate time.Time, title, description string) (Assignment, error) {
	var data AssignmentData

	client := &http.Client{}

	a := Assignment{
		AssignedID:     assignedID,
		TaskID:         taskID,
		CreatorID:      creatorID,
		StatusID:       statusID,
		Title:          title,
		Description:    description,
		DueDate:        &dueDate,
		CompletionDate: &completionDate,
	}

	aJSON, _ := json.Marshal(a)

	req, _ := http.NewRequest("POST", s.EndpointServices[ServiceAssignments]+fmt.Sprintf("/assignment/%d", assignmentID), bytes.NewBuffer(aJSON))
	req.Header.Set("Authorization", s.TokenAPI[ServiceAssignments])
	res, err := client.Do(req)
	if err != nil {
		return data.Data, err
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(responseBody, &data)
	return data.Data, err
}

func (s *Server) assignmentReview(assignmentID, creatorID, assignedID, taskID int, dueDate, completionDate *time.Time, title, description string) (Assignment, error) {
	var data AssignmentData

	client := &http.Client{}

	var statusID int
	sts, err := s.asignmentsStatusGetAll()

	for _, s := range sts {
		if s.Name == "programmed" {
			statusID = s.ID
			break
		}
	}

	a := Assignment{
		AssignedID:     assignedID,
		TaskID:         taskID,
		CreatorID:      creatorID,
		StatusID:       statusID,
		Title:          title,
		Description:    description,
		DueDate:        dueDate,
		CompletionDate: completionDate,
	}

	aJSON, _ := json.Marshal(a)

	req, _ := http.NewRequest("POST", s.EndpointServices[ServiceAssignments]+fmt.Sprintf("/assignment/review/%d", assignmentID), bytes.NewBuffer(aJSON))
	req.Header.Set("Authorization", s.TokenAPI[ServiceAssignments])
	res, err := client.Do(req)
	if err != nil {
		return data.Data, err
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(responseBody, &data)
	return data.Data, err
}

func (s *Server) asignmentsStatusGetAll() ([]StatusAssignment, error) {
	var data StatusAssignmentArrayData

	client := &http.Client{}
	req, _ := http.NewRequest("GET", s.EndpointServices[ServiceAssignments]+"/status", nil)
	req.Header.Set("Authorization", s.TokenAPI[ServiceAssignments])
	res, err := client.Do(req)
	if err != nil {
		return data.Data, err
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(responseBody, &data)
	return data.Data, err
}

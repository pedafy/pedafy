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
	Task struct {
		ID          int        `json:"id"`
		CreatorID   string     `json:"creator_id"`
		StatusID    int        `json:"status_id"`
		Title       string     `json:"title"`
		Description string     `json:"description"`
		CreatedAt   *time.Time `json:"created_at"`
		LastEdit    *time.Time `json:"last_edit"`
	}

	TaskData struct {
		Data Task `json:"data"`
	}

	TaskArrayData struct {
		Data []Task `json:"data"`
	}

	Status struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	StatusArrayData struct {
		Data []Status `json:"data"`
	}
)

func (s *Server) taskGetAll() ([]Task, error) {
	var data TaskArrayData

	client := &http.Client{}
	req, _ := http.NewRequest("GET", s.EndpointServices[ServiceTasks]+"/tasks?sort=new", nil)
	req.Header.Set("Authorization", s.TokenAPI[ServiceTasks])
	res, err := client.Do(req)
	if err != nil {
		return data.Data, err
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(responseBody, &data)
	return data.Data, err
}

// DEPRECATED
//
// func (s *Server) taskGetUserTask(userID int) ([]Task, error) {
// 	var data TaskArrayData

// 	client := &http.Client{}
// 	req, _ := http.NewRequest("GET", s.EndpointServices[ServiceTasks]+fmt.Sprintf("/tasks/user_id/%d", userID), nil)
// 	req.Header.Set("Authorization", s.TokenAPI[ServiceTasks])
// 	res, err := client.Do(req)
// 	if err != nil {
// 		return data.Data, err
// 	}
// 	defer res.Body.Close()

// 	responseBody, err := ioutil.ReadAll(res.Body)

// 	err = json.Unmarshal(responseBody, &data)
// 	return data.Data, err
// }

func (s *Server) taskGetOne(taskID int) ([]Task, error) {
	var data TaskArrayData

	client := &http.Client{}
	req, _ := http.NewRequest("GET", s.EndpointServices[ServiceTasks]+fmt.Sprintf("/tasks/id/%d", taskID), nil)
	req.Header.Set("Authorization", s.TokenAPI[ServiceTasks])
	res, err := client.Do(req)
	if err != nil {
		return data.Data, err
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(responseBody, &data)
	return data.Data, err
}

func (s *Server) taskNew(creatorID string, statusID int, title, description string) (Task, error) {
	var data TaskData

	client := &http.Client{}

	t := Task{
		CreatorID:   creatorID,
		StatusID:    statusID,
		Title:       title,
		Description: description,
	}

	tJSON, _ := json.Marshal(t)

	req, _ := http.NewRequest("PUT", s.EndpointServices[ServiceTasks]+"/task", bytes.NewBuffer(tJSON))
	req.Header.Set("Authorization", s.TokenAPI[ServiceTasks])
	res, err := client.Do(req)
	if err != nil {
		return data.Data, err
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(responseBody, &data)
	return data.Data, err
}

func (s *Server) taskModify(taskID int, creatorID string, statusID int, title, description string) (Task, error) {
	var data TaskData

	client := &http.Client{}

	t := Task{
		CreatorID:   creatorID,
		StatusID:    statusID,
		Title:       title,
		Description: description,
	}

	tJSON, _ := json.Marshal(t)

	req, _ := http.NewRequest("POST", s.EndpointServices[ServiceTasks]+fmt.Sprintf("/task/%d", taskID), bytes.NewBuffer(tJSON))
	req.Header.Set("Authorization", s.TokenAPI[ServiceTasks])
	res, err := client.Do(req)
	if err != nil {
		return data.Data, err
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(responseBody, &data)
	return data.Data, err
}

func (s *Server) taskStatusGetAll() ([]Status, error) {
	var data StatusArrayData

	client := &http.Client{}
	req, _ := http.NewRequest("GET", s.EndpointServices[ServiceTasks]+"/status", nil)
	req.Header.Set("Authorization", s.TokenAPI[ServiceTasks])
	res, err := client.Do(req)
	if err != nil {
		return data.Data, err
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(responseBody, &data)
	return data.Data, err
}

package server

import (
	"context"
	"log"

	"google.golang.org/appengine"

	"github.com/pedafy/pedafy/src/template"

	"github.com/pedafy/pedafy/src/datastore"
)

type serviceName = string

const (
	ServiceAssignments = "ASSIGNMENTS_SERVICE"
	ServiceTasks       = "TASKS_SERVICE"
	ServiceUsers       = "USERS_SERVICE"
)

// Server main structure
type Server struct {
	EndpointServices map[serviceName]string
	TokenAPI         map[serviceName]string

	// Oauth
	oAuthID     string
	oAuthSecret string
}

// Start the server by setting all its data
func (s *Server) Start() {

	s.registerHandlers()
	s.initEndpoint()
	err := template.Init("../../public/template")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (s *Server) initEndpoint() {
	s.EndpointServices = make(map[serviceName]string)
	if appengine.IsDevAppServer() {
		s.EndpointServices[ServiceAssignments] = "http://localhost:9001"
		s.EndpointServices[ServiceTasks] = "http://localhost:9002"
	} else {
		s.EndpointServices[ServiceAssignments] = "https://api.pedafy.com/tig/v1"
		s.EndpointServices[ServiceTasks] = "https://api.pedafy.com/task/v1"
	}
}

// fetchTokenAPI will request all the API token that are needed
// for API calls
func (s *Server) fetchTokenAPI(ctx context.Context) error {
	var err error
	s.TokenAPI = make(map[serviceName]string)
	s.TokenAPI[ServiceAssignments], err = datastore.FindAPITokenInformation(ctx, datastore.AssignmentsTokenName)
	if err != nil {
		return err
	}
	s.TokenAPI[ServiceTasks], err = datastore.FindAPITokenInformation(ctx, datastore.TasksTokenName)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) isTokenSet() bool {
	return s.TokenAPI[ServiceAssignments] != "" && s.TokenAPI[ServiceTasks] != ""
}

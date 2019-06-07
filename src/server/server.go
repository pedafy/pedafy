package server

import (
	"context"

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
}

func (s *Server) Start() {
	s.registerHandlers()
	template.Init("../../public/template")
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
